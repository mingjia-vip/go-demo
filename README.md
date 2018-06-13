参照java来学习golang

最近在学习使用golang，由于是名java程序员，所以在学习其他语言上都喜欢和自己熟悉的语言进行比较，而且我发现通过对方的方式学习也有助于更好的学习，所以就以”温故知新“的方式来记录一下学习过程。

一，语言界定
  编译型or解释型
    java，其实是混合型的语言，java代码先编译型字节码，然后再在java虚拟机中解释执行，严格来说还是解释型的；
    go，编译型语言

  动态or静态
    java，在运行时不能动态改变，属于静态语言，同时也是静态类型语言；（很多人认为解释型的都是动态语言（或动态类型），这是错误的，如java）
    go，也是静态语言，也是静态类型语言，但没有java那么明显（@@@需要举例）

  强类型or弱类型
    java，强类型
    go，强类型

  面向对象


  所以，go和java一样，是一种编译型语言，它结合了解释型语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。


二，安装，
  之前简单记录过golang的安装，和java的安装一样简单，同样需要配置环境变量：
  GOROOT 对应JAVA_HOME
  GOPATH,相当于CLASSPATH


三，helloword
  helloword使我们在学习一种语言的第一步，简单对比一下java和go的helloword

  java：
  //包
  package com.java.demo;
  //引用包（举例，程序未使用）
  import java.util.List;
  //类
  public class HelloWorld {
      /**
        第一个Java程序
       */
      public static void main(String []args) {
          System.out.println("Hello World"); // 打印 Hello World
      }
  }

  go：
  //包
  package main
  //引用包
  import "fmt"
  //main函数
  func main() {
     /**
        第一个go程序
       */
     fmt.Println("Hello, World!")
  }

  对于java程序员来说，感觉上还是挺好的，结构大致相同：都有package、import、main函数三个部分，而且关键字完全相同，挺好！
  不同：
  1，行末尾没有“；”，go语言不用显示写“；”，直接换行就好，编译器会自动解析，但如果要在一行写多条语句，那么就要显示的用“；”分隔了。
  2，首先是包路径，java中包都是同包（src下的目录）结构一致的，但是go的包不是这样的，需要根据是否是可执行程序，如果是那么package只能是main，如果不是那么可以根据个人喜欢明明，与目录结构无关；
  3，导入包的写法不同，go通过字符串引入；
  4，go的helloword的main方法不用放到类里，go中也和java中的class相似的结构，叫struct，这个以后再说；
  5，main函数的定义有出入，这个也是java和go在函数定义语法上的不同。
  6，还有一个与java不同的地方就是，没有用到的包是不能引用的，如在java的hellowrod中引用了List，只引用没有用，没问什么问题，但是这种情况在go中是会报错的；

四，go的包
  go其实没有项目的概念，只有包的概念，所有的包都放在$GOPATH/src 下，GO命令和编译器会在 $GOPATH/src 目录下搜索相应的包；
  包分为两种，可执行包和不可执行包，上边的helloword的package名字是main，说明就是可执行包，固定写法；如果不是main就说明是不可执行包，那么不可执行包的命名有没有规则呢？没有，哈哈可以随便写，和java不同，包的名字和和工程的目录是没有关系的。那么好，随便写，怎么解决包冲突的问题呢？这个时候就用到目录结构了，上边说了，所有的包都会放到$GOPATH/src下，在src下是根据包的目录存储的，如：
  $GOPATH/src
    --you
      --sort
    --my
      --sort
  同样是sort包，分别是you和my两个工程下的，这样我们在引用的时候就可以这样写：
  import (
    yousort "you/sort" 
    mysort "my/sort"  
  )
  在impot的字符串中指定目录信息，通过为相同名字的包指定别名，在使用的时候就可以使用别名调用，如：mysort.Sort(...),yousort.Sort(...)。
  到这里就又会发现一个问题，相同包下面的代码文件中不能出现重复的函数名称，这个需要注意一下。

  包的权限问题：
    包中的方法的访问权限是根据名称首字母的大小写来区分的(这点跟java很不相同）,如果首字母是小写，则表示是包内私有的（相当于java的private），如果首字母是大写，其他包就可以访问（public）在go的权限控制中也会说到（@@详见go语法）


  main函数和init函数
  go中也有main函数，作用和java中的main函数相同，是可执行包（package main）的入口函数（如果有init函数，会先执行init），可执行包下有且只有一个main函数。

  init函数是初始化函数，不论是不是可执行包下都可以有init函数，并且报下的所有代码文件都可以有init函数，而且在一个代码文件中都可以有多个init函数，主要是完成初始化操作使用，作用类似java类中的static{}代码段，不过不建议创建多个init函数，多个代码文件中的init函数的执行顺序不可控，再有在说go和java的helloword的时候提到，如果引用的包没有使用时会报错的，解决办法除了移除外，也可以给未使用的包设个别名“_”,这样就不会报错了，这个就与init方法有关，当设置成”_“别名后，编译器就会去执行该报的init方法，来完成初始化操作，不会人为该包无用。
  这里的别名“_”并不是随便取的，“_”在go中是有意义的，后边会提到（@@‘_’）

 
=================================================================


五，go语法
  go语法上会java是大小写敏感的，java的所有组成部门都要有标识符（名字），如类名，变量名，方法名，参数名等；go也是大小写敏感的

  1，定义变量
    go中变量的定义有多种方式,以string类型为例：
    //定义一个string类型的变量
    var vname string
    //定义三个string类型的变量
    var vname1, vname2, vname3 string
    //定义并初始化
    var vname string = “value”
    //同时初始化多个变量
    var vname1, vname2, vname3 string= “v1”, ”v2“, “v3”
    //有点繁琐？没关系，可以直接忽略类型声明（Go会根据其相应值的类型来帮你初始化它们），那么上面的代码变成这样了
    var vname1, vname2, vname3 = “v1”, ”v2“, “v3”
    //还繁琐？好吧，继续简化，可以省略var,换成“:=”：
    vname1, vname2, vname3 := “v1”, ”v2“, “v3”

    :=这个符号直接取代了var和type,这种形式叫做简短声明。不过它只能用在函数内部；函数外部使用会报错，所以一般用var方式来定义全局变量；

    变量定义在不同语言上都有自己的语法，对于java开发者来说可能感觉最不习惯的是吧类型放到了变量名的后面（这个真的好别扭），其他都可以接受，var的声明，类型自动匹配等在js上也都比较习惯了。

    “_”变量，这个是go的一个特殊变量名，这个变量名表示“被丢弃的”或者“无用的”，这个其实不用太在意，知道有这么个变量就好了。需要提到的一点就是，和包类似，go的编译器对于声明了但是没有使用的变量会报错的，当然"_"除外，比如：
    func main() {
       _, b := 1, 2
    }
    这个时候编译器就汇报b的错误，对于_则不会，之前在将helloword的时候，把未使用的包取个“_”的别名就不会报错是同一个原理。


  2，常量
    go中声明常量的关键字是const ，java中没有，相当于final修饰符。
    go的声明如下：
    const Pi = 3.1415926

    iota
    在go中有一个特殊的常量，可以认为是一个可以被编译器修改的常量，通常用在枚举中。
    在每一个使用const关键定义时，iota的值被重置为0，在本次const定义中，每使用一次iota，其值就会自动增加1，如：
    const (
        a = iota
        b = iota
        c = iota
    )
    这样a，b，c的值分别是0，1，2，还可以简写如下：
    const (
        a = iota  //0
        b         //1
        c         //2
    )

  3，基本类型
    3.1 布尔型
      布尔类型当然是都有的了，和java的没啥区别，没啥好说的，直接看吧：
      java：
        boolean isOk=false;

      go：
        var isOk bool   //or isOk:=false

    3.2 数值型
    整型
      java中的数值类型是有符号的，比如：
        int a=-1,b=1;
      在go中，将属性类型分开了，比如int分为了int（有符号）和uint（无符号），上边的例子用go写如下：
        var a int=-1  //有符号
        var b uint=1  //无符号

      
      说到数值类型，就一定要说范围，我们知道，java中的int是有范围的：-2147483648 ~ 2147483647
      java：
        int a = -2147483648;
        a = 2147483647;
        a = 2147483648; //报错，超出范围了

      go：
        var a int = -2147483648
        a = 2147483647
        a = 2147483648  //正确

      通过上面的测试，你会发现go的int没有超出范围报错，是因为go的数值类型没有范围吗？
      不是的，Go里面也有直接定义好位数的类型：int8, int16, rune/int32, int64 和 byte/uint8, uint16, uint32, uint64 (rune是int32的别称，byte是uint8的别称)
      如下，go：
      //int 32
      var a1 int32 = 2147483647   //正确
      //a1 = 2147483648         //错误，超出范围

      //uint32
      var b1 uint32 = 2147483647*2+1  //最大范围
      //b1 = 2147483647*2+2     //报错,超出范围

      通过上面的测试，发现int32和java中的int范围相同，而且有符号和无符号的长度相同。
      那go中的int是多少位呢？go的数值类型具体长度会根据系统平台来决定，比我测试环境是64机器，那么int的范围与 int64相同（范围就和java中的long范围相同了,但不用像java中要以l结尾），如果是32位系统，则int的范围与int32相同。

      说到与java类型的对应关系，java整型基本类型中还有short，在go中就相当于int16，当然go还有无符号的uint16；在java中还有byte，对应go中的int8，在go中，也有byte，和java的byte不同的是，go的byte是uint8，是无符号。

      注意：虽然在64位系统中int长度和int64相同，但是仍然是两种类型，类型间不能相互转化：
        var a int         //32位系统
        var a1 int32
        var a3 rune = 10
        a1=a3             //正确
        a=a3              //错误
      不同类型不能转化，rune和int32可以，因为rune和int32是相同类型，只是有两个名称而已。


    浮点型
      java中的浮点型有float和double，分别是32位和64位，在java中，float的数值必须以f或F结尾，与double区分，而double的数值可以加d也可以不加
      go中浮点型只有float32和float64（没有float），默认是float64，分别对应java中的float和double

    复数
      复数，这个在java中不支持，在golang中是支持的，有两中长度，分别是complex128（64位实数+64位虚数）和complex64（32位实数+32位虚数）
      go：
        var a complex64 = 5+5i
        fmt.Printf("Value is: %v", a)   //output: (5+5i)
        //内置函数complex从指定的实部和虚部构建复数，内置函数real和imag用来获取复数的实部和虚部：
        var b complex128 = complex(-5, 5) // -5+5i
        fmt.Println(real(b))           // "-5"
        fmt.Println(imag(b))           // "5"

      对于复数，可能平时用到的机会比较少，用java虽然没有支持该类型，但相关的复数操作可以自己实现。


      这里可以列个表，对比一下go和java中的数值类型：
类型              java                golang
整型            byte               int8
整型            int               int32/rune
整型            long               int64
整型                               int、int16、uint、uint8/byte、uint16, uint32, uint64
浮点型            float               float32
浮点型            double               float64
复数                                complex64、  complex128                         

    
    3.1 字符串
      字符串很常用，但是在java中字符串（String）不是基本类型，java中针对字符的基本类型是char，一个单一的16位Unicode字符；
      在go中字符串类型（string）是基本类型，固定长度不可以改变，采用UTF-8字符集编码，用一对双引号（""）或反引号（``）括起来作为定义：
      //字符串初始化
      var s = "hello"
      //多行字符使用``
      s1:= `"this 
      is 
      a
      test"
      `

      学java的人都知道，java中的Stirng是不可变类，在go中，string也是不可变的，但可以采用切片的方式来修改字符串：
      var s = "hello"
      s = "ha" + s[2:]  // 字符串虽不能更改，但可进行切片操作,切片后面会说到
      fmt.Printf(s)     // 变成hallo

      @@@切片是go中内置的一种特殊可变数组，这个以后会说到。
      @@@再有，说到变量，常量等就会涉及到一个作用于的问题，这个我们也在右面的时候说。

      






六，Go内置类型（派生类型）
集合（array、slice、map）

1，数组
go和java中肯定都有数组了，只不过在定定义上和定义变量类似，把类型放到了后边（还是觉得很难受有没有？）：
对比一下定义：
//java：
int[] arr;            //没有初始化
arr = new int[1]{};   //初始化长度为1的数组
int[] arr1 = new int[]{1,2,3};   //初始化三个值
//go:
var arr [1]int        //同样
arr = [1]int{}       //初始化长度为1的数组
arr1 := [3]int{1,2,3}          //同样初始化三个值
arr2 := [...]int{1,2,3}          //同样初始化三个值

我们看到不光数组类型放到了变量后边，类型名称都放到了"[]"后边，和java完全是颠倒的（真的好别扭）；
注意，在初始化arr1的时候，go的[]中指定了3（也可以不写，但是不写就是切片了，之后会说到），但是在java中的[]中是不能指定的，否则会报错。
再有，在go初始化arr2的时候，[]中写额是“...”,三个点表示自动计算长度，用户用户保持长度的正确性，但是这种只能是在静态初始化的时候（指定数组内容），如下也会报错的：
var arr2 [...]string //报错：use of [...] array outside of array literal

数组在使用上也都一样，都是使用“[]”来对元素进行操作:
var arr [3]string = [3]string{"1","2","3"}  //固定长度
arr1 := [3]string{"1","2","3"}              //还是这样写简单
arr[0] = "3"
arr1[len(arr1)-1] = "1"
fmt.Println("arr=",arr)
fmt.Println("arr1=",arr1)

数组长度也属于类型的一部分，所以我们看第一行，定义并初始化的时候，长度必须一致才可以，否则会报错，因为不同长度的数组是两个类型。
再有数组的长度是不可变的。数组之间的赋值是值的赋值，当数组作为参数的时候，其实是数组的拷贝，不是指针，如果要使用指针，那就要使用切片（slice）了，这个稍后会说到。

说到数组，当然不能只是一维，多维数组如何定义呢？跟java没两样，只是要记得go的定义风格就好了，请看：
//go
arr_1 := [1][2]int{{1,2}}
多维数组操作上就不多说了。

2，切片(slice)
切片也是go原生的类型，也是一种数组，叫动态数组，其实就相当于java中的ArrayList，可变数组。
再说数组初始化的时候说到，数组的定义不分是包含长度的（中括号[]中的值），在定义的时候如果不写这个长度值，就相当于定义个一个切片（定义了长度就是数组），如：
//go
var slice []int 
slice = []int{1,2,3} 
slice1 := [3]int{1,2,3}          
可以看到除了[]中没有指定值外，和数组没有什么区别，但是在使用上就灵活很多，比如：
//go
slice := []int{1,2,3} 
slice2 := slice[0:2]  //从一个slice初始化，0：2表示从开始的位置截止到第三个。
slice2 = slice[:2]    //如果是从0开始，可以简写为[:2]
slice2 = slice[:]    //拷贝slice
我们看到可以从一个slice截取一部分作为初始化和赋值，当然也可以使用这种方式来截取数组。
另外切片和数组的不同是，切片其实是一种引用类型，底层都是一个数组的引用，所以在改变引用值得时候，切片的值也会改变的：
arr := [3]int{1,2,3} 
slice = arr[:]    //copy arr
arr[0] = 10 
fmt.Println("arr=",arr)       //[10,2,3]
fmt.Println("slice=",slice) //[10,2,3]

在使用切片的时候，还经常会用到几个内置函数：
cap 得到slice的最大容量 ,这个我们在使用List的时候也经常会给集合一个初始化容量，只不过我们没有直接的方法获得这个初始化容量的值
len 获取slice的长度
append 添加一个或者多个元素，返回相同类型的切片
copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数。(go中，函数是可以返回多个值的，哈哈，这个功能是不是感觉挺不错，如果java中也可以就好了)

说到cap函数，提到了切片的容量，就要说一下切片的操作方式的三个参数的形式：
arr := [10]int{1,2,3,4,5,6,7,8,9,10}
slice3 := arr[:4:6]
如上slice3的初始化的[]中多了一个参数，这个参数就相当于我们在java中使用集合给定一个初始容量一样，给这个切片定义一个初始容量。

接下来就对比一下go的slice和java中的ArrayList:
1，本质上都是数组，ArrayList使用数组来存储，slice操作的也是数组，并且都可以自动扩容；
2，因为使用的是数组，所以他们也都是有序的，可重复的。
3，slice是引用类型，所以作为参数传递的时候，效果和java中的List传参是一样的。
@@@在go的container容器包中还提供了list、heap、ring等，在以后我会专门的进行对比。


3，Map
Map，也叫字典或集合，和java中的map是一样的，纪实一个键值对的集合，但在go的定义方式上和java就有一些区别了，先看一下java的初始化：
//java
Map<Object,Object> map = new HashMap<>(10);
在java中所有对象的父类都是Object，所以这个集合的key和value可以是任意对象，但通常很少这么用，我们都会根据key和value的类型进行分别指定，这样在处理的过程中也方便使用，不用类型转换；
再有java中的Map是接口，初始化的时候需要指定实现类，在go中，没有类，但是有结构体（struct），也可定义接口，也可实现集成，这个后面再说，先来看下go中map的定义：
//go
var map1 map[int]string //声明，没有初始化的时候是nil，此时是不能进行key操作的
map1 = map[int]string{} //初始化之后才能进行操作，初始化一个空的集合
map1 = map[int]string{0:"zero",1:"1"} //也可以在此添加初始数据
//map1 = make(map[int]string) //还可以使用make函数来初始化
map1[1] = "one"
map1[2] = "two"
fmt.Println("map1",map1)  //map1 map[0:zero 1:one 2:two]
delete(map1, 2)           //删除key
fmt.Println("map1",map1)  //map1 map[0:zero 1:one]
map的定义需要使用map关键字，[]中的是key的类型，后面跟着的是value的类型。
map在声明之后，初始化之前是不能进行键值对操作的，但是可以使用len、range等函数（后面会说到）进行操作，并不会报错，而如果是java，这个时候肯定会报错的，因为还是null。
我们看到map添加数据就像使用数组一样，只不过[]中不是index，而是key，不用显式的get、add，很方便，不过在移除的时候需要用到delete方法。

同样，我们对go原生的map和java中的HashMap（真的很像）进行对比：
1，都不能通过index进行操作，只能通过key进行操作
2，go中的map是无序的，类似java中的HashMap，也是无需的
3，go的map不是基本类型，是线程不安全的，同样类似于HashMap。（java中有对应的线程安全实现HashTable）
4，map是引用类型，所以在作为参数使用的时的效果依然是引用传参的效果，和java中的map传参效果相同。



接下来说一下内置的range函数，说了上边的数组、切片还有map后，当然就会想到怎么遍历，rang就是干这个用的，直接上例子：
//遍历数组
arr:= []int{1,2,3,4,5}
fmt.Println(arr)
for v:=range arr{
fmt.Println(v)
}

//遍历slice
slice := arr[:3]
fmt.Println(slice)
for v:= range slice{
fmt.Println(v)
}

//遍历map
var map1 map[int]string = make(map[int]string) //还可以使用make函数来初始化
map1[1] = "one"
map1[2] = "two"
fmt.Println("map1",map1)
for k,v := range map1{
fmt.Println(k,"=",v)
}
又看到了多个返回值，用起来真的好爽啊，哈哈







今天说的这些，引出了引用类型（slice和map），引用类型




六.2 make和new
前面说到go的内置引用类型：slice（切片）和map（集合），另外还有一个引用类型就是channel（通道）；前面也提到这些引用类型可以通过make函数进行初始化，但不是只能通过make函数，在go语言中也是有new函数的，这个new函数更接近java中的new函数，作用也是一样的，接下来我们就说一下make和new函数的区别：

make函数：
返回的是值，而不是指针；
只能用来创建slice、map和channel

new函数：
返回的是指针，而不是值；
用来创建struct（结构体，相当于java中的Class）

接下来再看一个函数，就是上面map初始化时提到的make函数，这是个内建函数，作用就是为go的引用类型进行初始化，所以不光可以初始化map，slice也是可以用make初始化（go的引用类型除了今天说的slice和map，还有channel，这个以后再说），make函数有三个参数，第一个是类型，第二个是存储空间，最后一个是预留存储空间，如：
var slice = make([]string,5,10)//三个参数
fmt.Println("slice",slice)
var map1 map[int]string = make(map[int]string,10) //两个参数
fmt.Println("map1",map1)

也可以使用new来创建slice和map：




今天我们就来对比一下java和go中的引用类型和基本类型，以及引用传参、值传参、和指针的使用；

首先我们来说一下java中的引用类型，也就是对象类型，与之相对的就是基本类型，比如，int、double等就是基本类型，对应的引用类型就是Integer和Double，关键就在于存储的是值还是


make函数


六.3 指针




七，error类型，异常处理，枚举


八，lamda表达式


九，go的优缺点：
  多返回值，在一定程度上为开发者带来了一定的便利性。试想，为了返回两到三个值，不得不封装一个对象，或者抹去业务名称使用Map、List等集合类，高级一点用Apache的Pair和Triple，虽然可行，但始终不如Go的实现来得优雅；
  Go也统一了异常的返回方式，不用再去纠结是通过抛异常还是错误码来判断是否成功，多返回值的最后一个是Error就行了。

  Go在语言的原生类型中支持了常用的一些结构，比如map和slice，而其他语言中它们更多是存在于库中，这也体现了这门语言是从实践角度出发的特点，既然人人都需要，为什么不在语言层面支持它呢。函数作为一等公民出现在了Go语言里，不过Java在最近的Java 8中也有了Lambda表达式，也算是有进步了。


  其他的一些特性，则属于锦上添花型的，比如不定参数，早在2004年的Java 1.5中就对varargs有支持了；多重赋值在Ruby中也有出现，但除了多返回值赋值，以及让你在变量交换值时少写一个中间变量，让代码更美观一些之外，其他的作用着实不是怎么明显。

  说了这么多Go的优点，当然它也有一些问题，比如GC，说到它，Java不得不露出洁白的牙齿，虽然在大堆GC上G1还有些不尽如人意，但Java的GC已经发展了很多年，各种策略也比较成熟，CMS或G1足以应付大多数场景，实在有要求还能用Azul Zing JVM。不过从最新的Go 1.5的消息来看，Go的GC实现有了很大地提升，顺便一提的是GOMAXPROCS默认也从1变成了CPU核数，看来官方对Go在多核的利用方面更有信心了。