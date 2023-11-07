package recorder

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	_ "strconv"
	"strings"
	"testing"

	"time"
)

func TestCpumem(t *testing.T) {

	for i := 0; i <= 300; i++ { //设置计时，默认循环统计300次
		time.Sleep(time.Second)                               //每1s统计一次cpu内存占比
		network := GetCommandLinuxCon("ps -aux | grep 23688") //根据实际，修改进程号，默认23688
		linesNetwork := strings.Split(string(network), " ")
		//fmt.Println(linesNetwork)
		fmt.Println("res", linesNetwork[8], linesNetwork[10])
		cpu, _ := strconv.ParseFloat(linesNetwork[8], 2)
		mem, _ := strconv.ParseFloat(linesNetwork[10], 2)

		transaction := &Cpumem{
			OccurTime: time.Now(),
			Cpupro:    cpu,
			Mempro:    mem,
		}
		resultC := make(chan error, 1)
		Record(transaction, resultC)

	}

}

func GetCommandLinuxCon(commandLinux string) []byte {
	//需要执行命令： commandLinux
	cmd := exec.Command("/bin/bash", "-c", commandLinux)
	// 获取输入
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return nil
	}
	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return nil
	}
	// 读取输出
	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return nil
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return nil
	}
	return bytes
}
