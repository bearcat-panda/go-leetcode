package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
	"testing"
)

/*func TestSlice(t *testing.T)  {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}*/

/*type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestStruct(t *testing.T) {
	//var peo People = &Stduent{}
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}*/

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func TestStruct(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}

func TestAppend(t *testing.T) {
	list := new([]int)
	//list = append(list, 1)
	fmt.Println(list)
}

func TestEqual(t *testing.T) {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1 == sm2 {
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}

}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func TestInterface(t *testing.T) {
	var x *int = nil
	Foo(x)
}

func TestLoop(t *testing.T) {

loop:
	for i := 0; i < 10; i++ {
		println(i)
	}
	goto loop
}

func TestType(t *testing.T) {
	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	//var i1 MyInt1 = i
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func TestImple(t *testing.T) {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func TestScope(t *testing.T) {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

func TestDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func TestString(t *testing.T) {
	fmt.Println(Utf8Index("北京天安门最美丽", "天安门"))
	fmt.Println(strings.Index("北京天安门最美丽", "男"))
	fmt.Println(strings.Index("", "男"))
	fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}

func TestReval(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))

}

func reverse(x int) int {
	var num int
	for x != 0 {
		num = num*10 + x%10
		x /= 10
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

	return num
}

type Test struct {
	Name string
}

var list map[string]Test

func TestMap(t *testing.T) {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	//list["name"].Name = "Hello"
	fmt.Println(list["name"])

	//map中存储结构体可以实现.
}

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

/*func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D

}*/

func TestBool(t *testing.T) {
	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}

	fmt.Println("3")

}

func TestSub(t *testing.T) {
	var array = []int{1, 9, 3, 8, 7, 6, 5, -9}
	var target = 5
	result := sub(array, target)

	/*for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array)-1; j++ {
			if array[j] - array[i] == target {
				fmt.Print(array[i], array[j])
			}
		}
	}*/
	fmt.Printf("%v", result)
}

func sub(array []int, target int) [][]int {
	var result = [][]int{}
	sort.Ints(array)
	var left, right = 0, 1

	for left < right {
		if array[right]-array[left] > target {
			left++
			if right < len(array)-1 {
				right++
			}
		} else if array[right]-array[left] < target {
			if right < len(array)-1 {
				right++
			}
		} else {
			left++
			if left+1 < len(array)-1 {
				right = left + 1
			}

			result = append(result, []int{array[left], array[right]})
		}

	}
	return result
}

func compre() {
	var str = []string{"a", "a", "b", "b", "e", "c", "e"}
	var result []string
	var left, right int = 0, 1
	for right <= len(str)-1 {
		if str[left] == str[right] {
			right++
		} else {
			result = append(result, str[left], fmt.Sprintf("%d", right-left))
			left = right
			if left == len(str)-1 {
				result = append(result, str[left], fmt.Sprintf("%d", 1))
			}
		}
	}
	fmt.Printf("%v", result)
}

func TestCompre(t *testing.T) {
	compre()
}

type DefineError struct {
	msg string
}

func (d *DefineError) Error() string {
	return d.msg
}
func TestError(t *testing.T) {
	//wrap error
	err1 := &DefineError{"this is a define error type"}
	err2 := fmt.Errorf("wrap err2: %w\n", err1)
	err3 := fmt.Errorf("wrap err3: %w\n", err2)
	var err4 *DefineError

	if errors.As(err3, &err4) {
		//errors.As() 顺着错误链,从err3一直找到被包装最底层的错误值err1,并且将
		//err3与其自定义类型var err4 *DefineError匹配成功
		fmt.Println("err1 is a variable of the DefineError type")
		fmt.Println(err4 == err1)
		return
	}
	fmt.Println("err1 is not a variable of the DefineError type")
}

//哨兵错误处理
var (
	ErrinvalidUser  = errors.New("invalid user")
	ErrNotFoundUser = errors.New("not found user")
)

func TestError1(t *testing.T) {
	err1 := fmt.Errorf("warp err1: %w\n", ErrinvalidUser)
	err2 := fmt.Errorf("warp err2: %w\n", err1)

	//golang 1.13新增Is()函数
	if errors.Is(err2, ErrinvalidUser) {
		fmt.Println(ErrinvalidUser)
		return
	}
	fmt.Println("success")

}
