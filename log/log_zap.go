package main

import (
	"fmt"
	"go.uber.org/zap"
)

var zlogger *zap.Logger
var sugar *zap.SugaredLogger

func InitLogger() {
	zlogger, _ = zap.NewProduction()
	sugar = zlogger.Sugar()
}

func main() {
	InitLogger()
	zlogger.Error("test .", zap.String("test", "testvalue"))
	fmt.Println("xxxx--------------------")
	sugar.Errorf("test %s.", "testvalue")

}
