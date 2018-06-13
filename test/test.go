package main

import (
	"fmt"
	"strings"
	"reflect"
	//"math/bits"
)

//常量打印分割线
const Constant_splitStr = "\n------------------------\n"

func StringTest() {
	fmt.Println("desc:", "字符串测试")
	index := 1

	//字符串初始化
	var s = "hello"
	s = "ha" + s[2:] // 字符串虽不a能更改，但可进行切片操作
	fmt.Printf("%d，%s%s", index, s, Constant_splitStr);index++

	//多行字符使用``，:=可以用于初始化及赋值
	s1:= `"this 
is 
a
test"
`
	fmt.Printf("%d，%s%s", index, s1, Constant_splitStr);index++

	//split
	var arr []string = strings.Split(s1," ")
	//type:=reflect.TypeOf(arr)
	fmt.Println(arr)

	_, b := 34, 35
	fmt.Println(b)

	//contains

	//replace

	//subStr

	//indexOf

	//equals

	//toLower\toUper

	//rego

}

func InnorTypeTest() {
	fmt.Println("desc:", "内置类型测试")
	var a int=-1	//有符号
	var b uint=1	//无符号

	//b=-1 //报错
	//a=b  //报错
	//b=a  //报错

	//边界
	//int 取值范围
	a = -2147483648
	a = 2147483647
	a = 2147483648					//正确
	//int 32
	var a1 int32 = 2147483647		//正确
	//a1 = 2147483648				//错误，超出范围

	//b = -2147483648
	b = 2147483647
	b = 2147483647*2+2				//正确
	//uint32
	var b1 uint32 = 2147483647*2+1	//最大范围
	//b1 = 2147483647*2+2			//报错,超出范围

	//int 64 ，-9223372036854775808 9223372036854775807
	a = -9223372036854775808		//正确
	a = 9223372036854775807			//正确
	//a = 9223372036854775808			//错误，超出范围
	//int 64
	var a2 int64 = -9223372036854775808
	a2 = 9223372036854775807		//正确
	//a2 = 9223372036854775808		//错误，超出范围

	var a3 rune
	a3=a1							//正确

	atype:=reflect.TypeOf(a)
	btype:=reflect.TypeOf(b)
	fmt.Println("a=",a," ",atype)
	fmt.Println("b=",b," ",btype)

	a1type:=reflect.TypeOf(a1)
	b1type:=reflect.TypeOf(b1)
	fmt.Println("a1=",a1," ",a1type)
	fmt.Println("b1=",b1," ",b1type)

	a2type:=reflect.TypeOf(a2)
	fmt.Println("a2=",a2," ",a2type)

	a3type:=reflect.TypeOf(a3)
	fmt.Println("a3=",a3," ",a3type)

	var c float32
	var c1 bool
	ctype:=reflect.TypeOf(c)
	fmt.Println("c=",c," ",ctype)
	c1type:=reflect.TypeOf(c1)
	fmt.Println("c1=",c1," ",c1type)

	var byt byte
	byt=123
	s := "hello"
	bytes := []byte(s) // 将字符串 s 转换为 []byte 类型
	fmt.Println("bytes=",bytes)
	byt = bytes[0]
	bytType:=reflect.TypeOf(byt)
	fmt.Println("byte=",byt," ",bytType)



}

func ArrayTest() {
	fmt.Println("desc:", "数组测试")

	var arr [3]string = [3]string{"1","2","3"}
	var arr111 [3]string
	arr1 := [3]string{}
	arr11 := [...]int{1,2,3}
	fmt.Println("arr=",arr)
	fmt.Println("arr1=",arr1)
	arr[0] = "3"
	arr1[len(arr1)-1] = "1"
	fmt.Println("arr=",arr)
	fmt.Println("arr1=",arr1)
	fmt.Println("arr11=",arr11)
	fmt.Println("arr111=",arr111)

	arr_1 := [1][2]int{{1,2}}
	fmt.Println("arr_1=",arr_1)
	fmt.Println("arr_1[0]=",arr_1[0])
	fmt.Println("arr_1[0][1]=",arr_1[0][1])
}

func SliceTest(){

	var slice []int
	slice = []int{1,2,3}
	slice1 := []int{1,2,3}
	fmt.Println("slice=",slice)
	fmt.Println("slice1=",slice1)

	//new
	slice2:=slice[0:2]
	slice2=slice[:2]
	fmt.Println("slice2=",slice2)

	slice[0] = 10
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)


	fmt.Println("cap=",cap(slice2))
	slice2 = append(slice2, slice1[0])
	fmt.Println("cap=",cap(slice2))

	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	slice3 := arr[:4:6]
	fmt.Println("slice3=",slice3[3])



}


func MapTest(){

	//map的初始化方式有两种
	var map1 map[int]string	//声明，没有初始化的时候是nil，此时是不能进行key操作的
	map1 = map[int]string{} //初始化之后才能进行操作，初始化一个空的集合
	map1 = map[int]string{0:"zero",1:"1"} //也可以在此添加
	//map1 = make(map[int]string) //还可以使用make函数来初始化
	map1[1] = "one"
	map1[2] = "two"
	fmt.Println("map1",map1)
	delete(map1, 2)
	fmt.Println("map1",map1)
}

func RangeTest(){

	arr:= []int{1,2,3,4,5}
	fmt.Println(arr)
	for v:=range arr{
		fmt.Println(v)
	}

	slice := arr[:3]
	fmt.Println(slice)
	for v:= range slice{
		fmt.Println(v)
	}

	//map的初始化方式有两种
	var map1 map[int]string	= make(map[int]string) //还可以使用make函数来初始化
	map1[1] = "one"
	map1[2] = "two"
	fmt.Println("map1",map1)
	for k,v := range map1{
		fmt.Println(k,"=",v)
	}

}



func MakeAndNewTest(){
	var slice1 = make([]string,5,10)//make函数有三个参数，第一个是类型，第二个是存储空间，最后一个是预留存储空间
	fmt.Println("slice1",slice1)
	var map1 map[int]string	= make(map[int]string,10) //还可以使用make函数来初始化
	fmt.Println("map1",map1)


	var slice2 = new([]string)//参数就是一个，就是Type
	fmt.Println("slice2",slice2)
	var map2 *map[int]string = new(map[int]string)
	fmt.Println("map2",map2)

}


func main(){
	//StringTest()
	//InnorTypeTest()
	//ArrayTest()
	//SliceTest()
	//MapTest()
	//RangeTest()
	MakeAndNewTest()
}