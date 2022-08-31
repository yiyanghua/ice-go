package main

import (
	"fmt"
	"reflect"
)

type Name string
type Bird struct {
	Name           Name
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying....")
}

/**
//接口数据  =====》 反射对象
1.Reflection goes from interface value to reflection object.

//反射对象 ===> 接口数据
2.Reflection goes from reflection object to interface value.

// 倘若数据可更改，可通过反射对象来修改它
3.To modify a reflection object,the value must be settable.

https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487027&amp;idx=1&amp;sn=9acf079c7991288277debe0f635ebed2&source=41#wechat_redirect
*/
func main() {
	b := Bird{"Sparrow", 3}
	v := reflect.ValueOf(b)
	t := reflect.TypeOf(b)

	// 通过值反射拿到
	fmt.Println(t) // type
	fmt.Println(v) // value
	fmt.Println(v.Type(), v.Kind() == reflect.Struct)

	// 通过反射后的value，转化为数据对象
	o := v.Interface().(Bird)
	fmt.Println(o)
	fmt.Println(reflect.DeepEqual(b, o))

	// 倘若数据可更改，可通过反射对象来修改它
	// 在Go中，任何函数的参数都是值的拷贝，而非原数据
	// 反射函数 reflect.ValueOf()也不例外。我们目前得到的反射对象，都是原对象的copy的反射对象，而非原对象本身，所以不可以修改到原对象
	e := v.Interface().(Bird)
	fmt.Println(e)
	e.Name = "xx"
	fmt.Println(b, e)

	f := reflect.ValueOf(&b).Elem()
	f.FieldByName("Name").SetString("yy")
	fmt.Println(b, f.Interface().(Bird))

}
