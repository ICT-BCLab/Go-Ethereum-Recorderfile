package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	ip       string
	port     int
	user     string
	password string
	database string
	dns      string
)

var (
	recorder_dns  string
	recorder_port int
)

// geth --mysqlHost=127.0.0.1 --mysqlPort=9527 --mysqlUser=test --mysqlPassword=123456 --mysqlDatabase=geth_test
// recorder_dns := "test:123456@tcp(127.0.0.1:3306)/chainmaker_test?charset=utf8mb4&parseTime=True&loc=Local"
// recorder_port := 9527
func getDns_Port(ctx *cli.Context) (string, int) {
	ip = ctx.String("mysqlHost")
	port = ctx.Int("mysqlPort")
	user = ctx.String("mysqlUser")
	password = ctx.String("mysqlPassword")
	database = ctx.String("mysqlDatabase")
	dns = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, database)
	return dns, port
}

func get_Port(ctx *cli.Context) int {
	port = ctx.Int("serverPort")
	return port
}
