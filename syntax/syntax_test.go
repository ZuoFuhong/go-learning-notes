package syntax

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
 	变量
	  1.变量的声明
      2.变量的初始化（默认值）
	  3.短变量
      4.多重赋值
      5.匿名变量
*/
func TestVariableSyntax(t *testing.T) {
	// 变量声明（自动对内存区域进行初始化操作，初始化默认值）
	var name string
	// 批量声明
	var (
		age  int
		flag bool
	)
	// 声明变量时，赋予默认值
	var contury string = "china"

	// 编译器推导类型的格式
	var city = "wuhan"

	// 短变量（只能方法中使用）
	gender := "wucang"

	fmt.Println(name, age, flag, contury, city, gender)

	// 多重赋值（变量的左值和右值按从左到右的顺序赋值）
	city, gender = gender, city
	fmt.Println("city: "+city, "gender: "+gender)

	// 匿名变量
	_ = time.Now().Day()
}

/*
	数据类型
      1.整型
      2.浮点型
      3.布尔型
      4.字符型
*/
func TestDataTypeSyntax(t *testing.T) {
	// 整形
	var number int8 = 23
	fmt.Println("number: ", number)
	fmt.Println("int8 range：", math.MinInt8, "~", math.MaxInt8)
	var count uint8 = 24
	fmt.Println("count: ", count)
	fmt.Println("uint8 maxValue: ", math.MaxUint8)

	fmt.Println("*********************************")

	// 浮点型
	var price = math.Pi
	fmt.Println("pi：", price)
	// 2位精度输出
	fmt.Printf("pi：%.2f\n", price)

	fmt.Println("*********************************")

	// 布尔类型
	aBool := true
	fmt.Println("aBool: ", aBool)

	fmt.Println("*********************************")

	// 字符串
	str := "hello world"
	fmt.Println("str: " + str)
	fmt.Printf("str: %s\n", str)

	// 定义多行字符串，使用反引号
	str = `第一行
	第二行
第三行
`
	fmt.Println("mutiple line：", str)

	/*
			Go语言的字符有以下两种：
		      1）一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
		      2）另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。
		         rune 类型等价于 int32 类型。

			  byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，
		      字符使用单引号括起来。
	*/

	var a byte = 'a'
	fmt.Printf("a: %d, 变量的类型：%T\n", a, a)
	var b rune = '你'
	fmt.Printf("b：%d, 变量的类型：%T\n", b, b)
}

/*
	数据类型转换
*/
func TestTypeConv(t *testing.T) {
	// 1.浮点数在转换为整型时，会将小数部分去掉，只保留整数部分。
	var c = 3.22
	fmt.Println("c：", int8(c))

	// 2.整型转字符串类型，使用strconv.Itoa(int) 函数
	var d = 23
	sd := strconv.Itoa(d)
	fmt.Printf("sd变量类型  %T\n", sd)

	// 3.字符转整型
	num, err := strconv.Atoi(sd)
	if err == nil {
		fmt.Printf("type:%T value:%#v\n", num, num)
	} else {
		fmt.Printf("%v 转换失败！", num)
	}

	// 4.Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("转换成功 b = %v\n", b)
	} else {
		fmt.Printf("转换失败")
	}

	// 5.Format 系列函数用于将给定类型数据格式化为字符串类型的功能，
	// 其中包括 FormatBool()、FormatInt()、FormatUint()、FormatFloat()。
	sBool := strconv.FormatBool(true)
	fmt.Println("sBool: ", sBool)

	// 6.Append 系列函数用于将指定类型转换成字符串后追加到一个切片中，
	// 其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint()。
}

/*
	指针语法
      1.指针地址
      2.指针类型
*/
func TestPointerSyntax(t *testing.T) {
	var hint string = "hello"

	// 对字符串取地址, ptr类型为*string
	var ptr *string = &hint

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址，指针值带有0x的十六进制前缀。
	fmt.Println("ptr value：", ptr)

	// 对指针进行取值
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value：%s\n", value)

	// 使用指针修改值
	*ptr = "world"
	fmt.Println("修改后的hint：", hint)

	// 使用new()创建指针
	str := new(string)
	fmt.Println("指向默认值（空字符串）：", *str)

	*str = "hello"
	fmt.Println("更新后的值：", *str)

	var a = "hello"
	var b = a
	// 输出结果：内存地址不一样
	fmt.Println("a 地址：", &a)
	fmt.Println("b 地址：", &b)
}

/*
	常量和const关键字
      1.定义常量
      2.使用const常量配合iota模拟枚举
*/
func TestConstSyntax(t *testing.T) {
	const size = 4
	const (
		pi = 3.141592
		e  = 2.718281
	)
	fmt.Println("size:", size)
	fmt.Println("pi: ", pi, " e: ", e)

	fmt.Println("*******************************")

	// Go语言中现阶段没有枚举，可以使用 const 常量配合 iota 模拟枚举

	// 将 int 定义为 Weapon 类型
	type Weapon int
	// 标识后，const 下方的常量可以是默认类型的，默认时，默认使用前面指定的类型作为常量类型。该行使用 iota 进行常量值自动生成。
	// iota 起始值为 0，一般情况下也是建议枚举从 0 开始，让每个枚举类型都有一个空值，方便业务和逻辑的灵活使用。

	// 一个 const 声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加。这种模式有点类似于电子表格中的单元格自动填充。
	// 只需要建立好单元格之间的变化关系，拖动右下方的小点就可以自动生成单元格的值。
	const (
		Arrow Weapon = iota // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println("Arrow: ", Arrow)
	fmt.Println("Blower: ", Blower)
}

/*
	使用type关键字 定义 类型别名
*/
func TestTypeAliasSyntax(t *testing.T) {
	// 将NewInt定义为int类型
	type NewInt int

	// 将int取一个别名叫IntAlias
	type IntAlias = int

	// 将a声明为NewInt类型
	var a NewInt

	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)

	// 将a2声明为IntAlias类型
	var a2 IntAlias

	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	// 结果显示a的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型。a2 类型是 int。IntAlias 类型只会在代码中存在，
	// 编译完成时，不会有 IntAlias 类型。
}

/*
	数组
*/
func TestArraySyntax(t *testing.T) {
	// var 数组变量名 [元素数量]T
	var arr [4]int

	// 数组长度
	fmt.Println("arr length: ", len(arr))
	// 默认值
	fmt.Println("arr value：", arr)

	// 数组赋值，数组越界，运行时会报出宕机
	arr[0] = 1
	fmt.Println("arr value: ", arr)

	// 需要确保大括号后面的元素数量与数组的大小一致
	var nameList = [3]string{"dazuo", "wang", "li"}
	fmt.Println(nameList)

	// 让编译器在编译时，根据元素个数确定数组大小
	var ageList = [...]int{1, 3, 4}
	fmt.Println(ageList)
	fmt.Println("ageList length: ", len(ageList))

	// 等效，编译时确定数组大小
	var colorList = []string{"red", "yellow", "bule"}
	fmt.Println(colorList)

	// 遍历数组
	for k, v := range colorList {
		fmt.Println(k, v)
	}
	// 多维数组
	var multiArr = [][]string{{"dazuo", "age"}, {"jin", "age"}}
	fmt.Println(multiArr)

	// 字符类型上下等效
	barr := []byte{'h', 'e', 'l', 'l', 'o'}
	barr2 := []byte("hello")
	fmt.Println("arrb: ", barr)
	fmt.Println("barr2: ", barr2)
}

/*
	切片
*/
func TestSliceSyntax(t *testing.T) {
	var nameList = [...]string{"dazuo", "wang", "li"}
	// 取出元素不包含结束位置对应的索引
	var subList = nameList[1:2]
	fmt.Println(subList)
	fmt.Println("中间至尾部：", nameList[1:])
	fmt.Println("开头至中间：", nameList[:2])
	fmt.Println("全部：", nameList[:])

	fmt.Println("********************************")

	// 声明字符串切片
	var strList []string
	fmt.Println(strList)
	// 声明但未使用的切片的默认值是 nil
	// 本来会在{}中填充切片的初始化元素，这里没有填充，所以切片是空的。但此时 numListEmpty 已经被分配了内存，但没有元素。
	// 因此和 nil 比较时是 false。
	fmt.Println(strList == nil)

	fmt.Println("********************************")

	a := make([]int, 2)
	fmt.Println("a len: ", len(a))
	fmt.Println("a value: ", a)

	b := make([]int, 2, 10)
	fmt.Println("b len: ", len(b))

	// 每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不能容纳足够多的元素时，切片就会进行“扩容”
	c := append(b, 1)
	fmt.Println("c value: ", c)

	// 修改切片值，操作的是同一片内存空间
	c[0] = 3
	fmt.Println(b)

	fmt.Println("********************************")

	// 复制切片
	var d = make([]int, 1, 10)
	i := copy(d, c[:])
	fmt.Println("copy count：", i)
	fmt.Println("d value: ", d)
}

/*
	map映射
*/
func TestMapSyntax(t *testing.T) {
	// 使用make创建map
	s := make(map[string]string)
	s["name"] = "dazuo"
	fmt.Println(s)
	fmt.Println(s["name"])

	// 声明map时填充元素
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(m)

	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 删除map中的键值对
	// 清空map的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	delete(m, "W")
	fmt.Println(m)

	// 并发环境中的map
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

/*
  	list的初始化有两种方法
	  1.通过 container/list 包的 New() 函数初始化 list
        变量名 := list.New()
	  2.通过 var 关键字声明初始化 list
        var 变量名 list.List{}
*/
func TestListSyntax(t *testing.T) {
	// 1.通过 container/list 包的 New() 函数初始化 list
	myList := list.New()
	// 将 fist 字符串插入到列表的尾部，此时列表是空的，插入后只有一个元素。
	myList.PushBack("first")

	// 将数值 68 放入列表。此时，列表中已经存在 fist 元素，67 这个元素将被放在 fist 的前面。
	myList.PushFront(68)

	// 保存元素句柄
	element := myList.PushFront(67)

	fmt.Println("element: ", element)
	fmt.Println(myList)

	// 获取链表首部元素 67，然后找到链表下一个元素 68
	fmt.Println(myList.Front().Next().Value)

	// 通过句柄移除元素
	myList.Remove(element)

	fmt.Println("********************************")

	// 遍历列表
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// 2.通过 var 关键字声明初始化 list
	ageList := list.List{}
	ageList.PushBack(true)
	ageList.PushBack("dazuo")
	fmt.Print(ageList.Front().Next().Value)
}

/*
	流程控制
      1.if else（分支结构）
      2.for（循环结构）
      3.switch语句
      4.goto语句
*/
func TestProcessControlSyntax(t *testing.T) {
	num := 10
	if num > 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// if 的特殊写法
	// 可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
	if num++; num > 0 {
		fmt.Println("ye")
	}

	fmt.Println("*********************************")

	// for（循环结构）
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	var i int
	for ; ; i++ {
		if i > 10 {
			fmt.Println("for i = ", i)
			break
		}
	}

	// 无限循环
	var w = 0
	for {
		if w > 10 {
			break
		}
		w++
	}

	// 自由一个条件的循环
	var n = 0
	for n < 10 {
		if n > 11 {
			break
		}
		n++
	}

	fmt.Println("*********************************")

	// switch结构
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	// 一个分支多个值
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	// 分支表达式
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	fmt.Println("*********************************")

	// 使用goto退出多层循环
	for x := 0; x < 10; x++ {
		if x == 2 {
			goto breakHere
		}
	}

	fmt.Println("in coming")
	// 手动返回，避免执行进入标签；此处如果不手动返回，在不满足条件时，也会执行标签处代码
	return

	// 标签
breakHere:
	fmt.Println("done")

	fmt.Println("*******************************")

	fmt.Println("跳出指定循环")

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}

	fmt.Println("********************************")
	// continue 将结束当前循环，开启下一次的外层循环，而不是内层的循环。
OuterFlag:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterFlag
			}
		}
	}
}
