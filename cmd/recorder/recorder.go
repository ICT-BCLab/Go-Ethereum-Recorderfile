package recorder

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbConnected = false  // 数据库是否已连接
	db          *gorm.DB // gorm数据库连接对象
	allModels   []interface{}
	// registerInfo: key: modelName; value: register point info
	registerInfo = make(map[string]string)
)

// RegisterModel: the argument model must be an address of an object
func RegisterModel(model interface{}) {
	if dbConnected {
		panic("'RegisterModel' must be called before 'Start'")
	}
	modelType := reflect.TypeOf(model)
	if modelType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("%s should be a pointer", modelType.Name()))
	}
	modelName := modelType.Elem().Name()
	if caller, ok := registerInfo[modelName]; ok {
		panic(fmt.Sprintf("%q has been registered at %s", modelName, caller))
	}
	allModels = append(allModels, model)
	caller := getCaller()
	registerInfo[modelName] = caller
}

/*
Start: initialize package recorder

	@dns is the address of db
	if dns contains "sqlite", we treat it as a sqlite db address, and will init *gorm.DB with sqlite;
	else, we treat dns as a mysql db address
	mysql dns format："user:passwd@tcp(ip:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	eg:
	dns = "chainmaker:chainmaker@tcp(127.0.0.1:3306)/chainmaker_recorder?charset=utf8mb4&parseTime=True&loc=Local"

	@port is the backdoor port for dynamically update configuration
*/
func Start(dns string, port uint16) error {
	if dbConnected {
		panic("db has been connected")
	}
	var err error
	if strings.Contains(dns, "sqlite") {
		// use sqlite, we don't have to start a mysql server,
		// sqlite will create a db file locally
		db, err = gorm.Open(sqlite.Open(dns), &gorm.Config{}) // &gorm.Config{}表示使用默认配置
	} else {
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	}
	if err != nil {
		return fmt.Errorf("connect to dns(%s) failed: %s", dns, err.Error())
	}
	dbConnected = true
	createAllModels()
	startConfigListener(port)
	return nil
}

func createAllModels() {
	for _, model := range allModels {
		modelName := reflect.TypeOf(model).Elem().Name()
		err := db.AutoMigrate(model)
		if err != nil {
			caller := registerInfo[modelName]
			// added by zhuchen
			if !strings.Contains(err.Error(), "already exists") {
				//
				panic(fmt.Sprintf("[recorder] AutoMigrate(%s) failed, registered at %q, err: %s\n", modelName, caller, err.Error()))
			}
		} else {
			fmt.Printf("[recorder] AutoMigrate(%s) succeded\n", modelName)
		}
	}
}

/*
Record: insert value into db asynchronously;

	 If the caller wants to confirm whether the insertion is completed, it can pass in resultC to receive the result;
	 If resultC is nil, nothing will be reported
	 eg:
		c := make(chan error, 1)
		recorder.Record(someVal, c) // run asynchronously
		select {
		case err := <-c:
			if err != nil {
				// something wrong
			} else {
				// everything goes well
			}
		case <-timer.C:
			// timeout
		}
*/
func Record(value interface{}, resultC chan error) {
	if dbConnected {
		modelName := reflect.TypeOf(value).Elem().Name()
		accessLock.RLock()
		allAccess := accessConfig["All"]
		modelAccess := accessConfig[modelName]
		accessLock.RUnlock()
		if allAccess && modelAccess {
			go safeGoroutine(func() error {
				result := db.Create(value)
				if result.Error != nil {
					data, _ := json.Marshal(value)
					return fmt.Errorf("[recorder] Record(%s), data=%s, got err: %s", modelName, data, result.Error.Error())
				}
				return nil
			}, resultC)
		} else if resultC != nil {
			resultC <- fmt.Errorf("[recorder] %q not allowed to record", modelName)
		}
	} else if resultC != nil {
		resultC <- fmt.Errorf("[recorder] database not connected")
	}
}

// GetConfigValue: get the configured value, which can be updated through endpoint [PUT] /config/configvalue
func GetConfigValue(key string) (interface{}, bool) {
	configLock.RLock()
	defer configLock.RUnlock()
	val, ok := configValue[key]
	return val, ok
}
