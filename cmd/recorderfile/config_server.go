package recorderfile

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	serverPort   uint16                         = 9527 // 端口号
	accessLock   sync.RWMutex                          // 读写锁（accessConfig）
	//初始化accessConfig为全部开关开启
	accessConfig = map[string]bool{
		"All": true,
		"TransactionPoolInputThroughputF": true,
		"NetP2PTransmissionLatencyF": true,
		"PeerMessageThroughputF": true,
		"DbStateWriteRateF": true,
		"DbStateReadRateF": true,
		"TxQueueDelayF": true,
		"BlockCommitDurationStartF": true,
		"BlockCommitDurationEndF": true,
		"BlockValidationEfficiencyStartF": true,
		"BlockValidationEfficiencyEndF": true,
		"TxDelayStartF": true,
		"TxDelayEndF": true,
		"BlockTxConflictRateF": true,
		"ContractExecuteEfficiencyF": true,
		"ConsensusTbftCostF": true,
		"Cpumem": true
	}
	configLock   sync.RWMutex                          // 读写锁（configValue）
	configValue  = make(map[string]interface{})        // accessconfig.yml(使用interface是因为Value可能是各种数据类型)
)

// 启动配置监听器，如果用户命令行传入的端口号>0就使用传入的端口号
func startConfigListener(port uint16) {
	if port > 0 {
		serverPort = port
	}
	go safeGoroutine(runServer, nil)
}

func runServer() error {
	gin.SetMode(gin.ReleaseMode) // 错误信息将输出在log中
	router := gin.Default()
	config := router.Group("/config") // 前缀是config
	config.GET("/registerinfo", getRegisterInfo)
	config.GET("/accessconfig", getAccessConfig)
	config.PUT("/accessconfig", updateAccessConfig)
	config.GET("/configvalue", getConfigValue)
	config.PUT("/configvalue", updateConfigValue)
	return router.Run(fmt.Sprintf(":%d", serverPort))
}

func getRegisterInfo(c *gin.Context) {
	c.YAML(http.StatusOK, registerInfo)
}

func getAccessConfig(c *gin.Context) {
	accessLock.RLock()
	defer accessLock.RUnlock()
	c.YAML(http.StatusOK, accessConfig)
}

func updateAccessConfig(c *gin.Context) {
	accessLock.Lock()
	defer accessLock.Unlock()
	err := c.BindYAML(accessConfig)
	if err != nil {
		c.String(http.StatusBadRequest, "update accessConfig err: %s", err.Error())
		return
	}
	c.YAML(http.StatusOK, accessConfig)
}

func getConfigValue(c *gin.Context) {
	configLock.RLock()
	defer configLock.RUnlock()
	c.YAML(http.StatusOK, configValue)
}

func updateConfigValue(c *gin.Context) {
	configLock.Lock()
	defer configLock.Unlock()
	err := c.BindYAML(configValue)
	if err != nil {
		c.String(http.StatusBadRequest, "update configValue err: %s", err.Error())
		return
	}
	c.YAML(http.StatusOK, configValue)
}
