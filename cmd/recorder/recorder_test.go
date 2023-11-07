package recorder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	c "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

// var dns = "chainmaker:chainmaker@tcp(127.0.0.1:3306)/chainmaker_recorder?charset=utf8mb4&parseTime=True&loc=Local"

var (
	resultC = make(chan error, 1)
)

type SampleTransaction struct {
	ID        uint      `gorm:"primarykey"`
	TxID      string    `gorm:"type:char(64)"`
	CreatedAt time.Time `gorm:"index"`
	Source    int
}
type SampleStruct struct {
	ID        uint      `gorm:"primarykey"`
	TxID      string    `gorm:"type:char(64)"`
	CreatedAt time.Time `gorm:"index"`
	Source    int
}

func TestMain(m *testing.M) {
	RegisterModel(&SampleTransaction{})
	RegisterModel(&SampleStruct{})
	m.Run()
}

func TestRegisterModel(t *testing.T) {
	c.Convey("TestRegisterModel", t, func() {
		c.Convey("RegisterObject", func() {
			defer func() {
				err := recover()
				fmt.Printf("\nRegisterObject err: %+v\n", err)
				c.So(err, c.ShouldNotBeNil)
			}()
			RegisterModel(SampleTransaction{})
		})

		c.Convey("RegisterDuplicated", func() {
			defer func() {
				err := recover()
				fmt.Printf("\nRegisterDuplicated err: %+v\n", err)
				c.So(err, c.ShouldNotBeNil)
			}()
			RegisterModel(&SampleTransaction{})
		})
	})
}

func connectDB() {
	if !dbConnected {
		Start("sqlite_test.db", 9527)
	}
}

func TestRecord(t *testing.T) {
	connectDB()
	c.Convey("TestRecord", t, func() {
		hasTable := db.Migrator().HasTable(&SampleTransaction{})
		c.So(hasTable, c.ShouldBeTrue)

		// without set allow-config to be true, struct is not allowed to record
		transaction := &SampleTransaction{
			TxID:      strconv.Itoa(int(time.Now().UnixNano())),
			CreatedAt: time.Now().Add(-time.Hour * 13),
			Source:    1,
		}
		Record(transaction, resultC)
		err := <-resultC
		c.So(err, c.ShouldNotBeNil)
		fmt.Printf("\nnot allowed to record: %v\n", err)

		// set allow-config to be true
		accessConfig["All"] = true
		accessConfig["SampleTransaction"] = true
		Record(transaction, resultC)
		err = <-resultC
		c.So(err, c.ShouldBeNil)

		// SampleTransaction.ID should be unique
		Record(transaction, resultC)
		err = <-resultC
		c.So(err, c.ShouldNotBeNil)
		fmt.Printf("\nUNIQUE constraint failed: %v\n", err)

		accessConfig = make(map[string]bool)
	})
}

func TestConfigServer(t *testing.T) {
	connectDB()
	/*
		TestMain 中通过 InitWithDNS() 调用了startConfigListener, 启动了 ConfigServer，在单独的 goroutine 中运行；
		主程序阻塞在这里，ConfigServer 不会退出
		通过 `go test -run ^TestConfigServer$ -timeout 30s` 命令可以指定当前test的超时时间
	*/
	c.Convey("TestConfigServer", t, func() {

		c.Convey("config/registerinfo", func() {
			resp, err := http.Get("http://localhost:9527/config/registerinfo")
			c.So(err, c.ShouldBeNil)
			data, err := ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)

			registerConfig := make(map[string]string)
			yaml.Unmarshal(data, registerConfig)
			fmt.Println("\nregisterinfo:")
			fmt.Println(string(data))
			c.So(len(registerConfig["SampleTransaction"]), c.ShouldBeGreaterThan, 0)
		})

		c.Convey("config/accessconfig", func() {
			configs := make(map[string]bool)
			resp, err := http.Get("http://localhost:9527/config/accessconfig")
			c.So(err, c.ShouldBeNil)
			data, err := ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)
			yaml.Unmarshal(data, configs)
			fmt.Println("\nallowconfig:")
			fmt.Println(string(data))
			c.So(len(configs), c.ShouldEqual, 0)

			data, err = ioutil.ReadFile("accessconfig.yml")
			c.So(err, c.ShouldBeNil)
			buffer := bytes.NewBuffer(data)
			request, err := http.NewRequest("PUT", "http://localhost:9527/config/accessconfig", buffer)
			c.So(err, c.ShouldBeNil)
			resp, err = http.DefaultClient.Do(request)
			c.So(err, c.ShouldBeNil)
			data, err = ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)
			yaml.Unmarshal(data, configs)
			fmt.Println("\nafter put, accessconfig:")
			fmt.Println(string(data))
			c.So(configs["All"], c.ShouldBeTrue)
			accessConfig = make(map[string]bool)
		})

		c.Convey("config/configvalue", func() {
			configs := make(map[string]interface{})
			resp, err := http.Get("http://localhost:9527/config/configvalue")
			c.So(err, c.ShouldBeNil)
			data, err := ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)
			yaml.Unmarshal(data, configs)
			fmt.Println("\nallowconfig:")
			fmt.Println(string(data))
			c.So(len(configs), c.ShouldEqual, 0)

			data, err = ioutil.ReadFile("configvalue.yml")
			c.So(err, c.ShouldBeNil)
			buffer := bytes.NewBuffer(data)
			request, err := http.NewRequest("PUT", "http://localhost:9527/config/configvalue", buffer)
			c.So(err, c.ShouldBeNil)
			resp, err = http.DefaultClient.Do(request)
			c.So(err, c.ShouldBeNil)
			data, err = ioutil.ReadAll(resp.Body)
			c.So(err, c.ShouldBeNil)
			yaml.Unmarshal(data, configs)
			fmt.Println("\nafter put, configvalue:")
			fmt.Println(string(data))
			c.So(configs["Ratio"], c.ShouldEqual, 12.2)
			configValue = make(map[string]interface{})
		})
	})
	// select {}
	/* release the upper comment, and run `go test -timeout 300s -run ^TestConfigServer$`
	then the config server will start and keep running for 300s, where we can use Makefile to test the config server
	*/
}
