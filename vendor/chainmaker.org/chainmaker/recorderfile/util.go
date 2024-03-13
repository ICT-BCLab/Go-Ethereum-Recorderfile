package recorderfile

import (
	"fmt"
	"runtime"
)

// 在一个新的goroutine中运行传入的函数，并在函数发生panic时捕获panic，防止程序崩溃
func safeGoroutine(f func() error, resultC chan<- error) {
	var err error
	defer func() {
		// 尝试恢复panic从而收集具体的错误日志同时保证原来的进程没有被挂掉
		if pErr := recover(); pErr != nil {
			fmt.Printf("[recorder] got panic: %+v", pErr)
			if err == nil {
				err = fmt.Errorf("[recorder] got panic: %+v", pErr)
			}
		}
		if resultC != nil {
			resultC <- err
		}
	}()
	err = f()
}

// 返回调用者的文件名、行号和函数名
//getCaller reports file:line number:function name information about function invocation of the parent function
func getCaller() string {
	funcName, fileName, lineNo, _ := caller(3)
	return fmt.Sprintf("%s:%d:%s", fileName, lineNo, funcName)
}

func caller(skip int) (funcName, fileName string, lineNo int, ok bool) {
	pc, file, line, ok := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name(), file, line, ok
}
