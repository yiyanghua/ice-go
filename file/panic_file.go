package main

import (
	"os"
	"syscall"
)

const panicfile = "/Users/yiyanghua/panic.txt"

func initPanicFile() error {
	file, err := os.OpenFile(panicfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		return err
	}
	return nil
}

func init() {
	err := initPanicFile()
	if err != nil {
		println(err)
	}
}
func testPanic() {
	panic("test panic")
}

func main() {
	testPanic()
}
