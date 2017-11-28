package cal

import "fmt"
import "os"


func Init() int {
	a,b := 1,2
	return a + b
}

func Initarr(){
	//数组
	//var a [2]string
	var a [2]string
	a[0] = "hello"
	a[1] = "e world"
	//fmt.Println(a[0], a[1]) //按下标取值
	//fmt.Println(a)

	b := []int{1, 2, 3, 4, 5};
	b[4] = b[1] + len(b);
	fmt.Println(b[4])
	fmt.Println(b)

	//for i := 0; i < len(b); i++{
	//	fmt.Println(b[i])
	//}

	//切片
	c := [][]int{{1,2,3},{4,5,6}}
	fmt.Println(c)

	var d []int
	d = make([]int, 2, 8)
	//d = b[1:2]
	d[0] = 3;
	d[1] = 4;
	fmt.Println(d)
}

func Initmap(){
	m1 := map[int]string{1:"a"}
	//var mi map[int]string
	//m1 = make(map[int]string)
	//m1[1] = "aa"

	if v, ok := m1[1]; ok{
		fmt.Println(v);
	} else {
		fmt.Println("Key Not Found")
	}
}

func InitComplex(){
	var v1 complex64
	v1 = 3.2 + 12i
	v2 := 3.2 + 12i
	v3 := complex(3.2, 12)
	fmt.Println(v1, v2, v3)
	fmt.Println(real(v1), imag(v1))
}

func Initpointer(){
	a := 20;
	var b *int
	b = &a
	*b = 30
	fmt.Println(a)
	fmt.Println(*b)
}

func Initrange(){
	var s1 string
	s1 = "hello, word"
	n := len(s1)
	for i := 0; i < n;i++ {
		fmt.Println(i, s1[i])
	}
}

func InitOs(){
	hostname, _ := os.Hostname();
	fmt.Println(hostname)
}

type IFly interface {
	Fly()
}

