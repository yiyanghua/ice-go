package main

import (
	"errors"
	"ice-go"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type TestFun func(url string) string

type DefaultFactory struct {
	test map[string]TestFun
}

func (factory *DefaultFactory) regist(key string, fun TestFun) {
	factory.test[key] = fun
}

func main() {
	ice.Test()
	readFromyaml("main/serverdemo.yaml")

	m := make(map[string]TestFun, 2)
	factory := DefaultFactory{test: m}

	factory.regist("1111", func(url string) string {
		return url;
	})

	factory.regist("2222", func(url string) string {
		return url;
	})

	test := "test12";
	if testFun, ok := factory.test["1111"]; ok {
		value := testFun(test);
		fmt.Printf("########%s\n", value)
	}

}

func readFromyaml(path string) error {
	filename, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("can not find the file:%s\n", filename)
		return errors.New("can not find the file :" + filename)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read config file fail. file: %s, error: %s\n", filename, err.Error())
		return errors.New("read config file fail. " + err.Error())
	}

	m := make(map[string]interface{}, 32)

	yaml.Unmarshal([]byte(data), &m)

	for k, v := range m {
		fmt.Printf("%s=%s \n", k, v)
	}

	return nil
}
