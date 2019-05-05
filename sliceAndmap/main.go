package main

import "fmt"

//学习slice与map的区别
func main() {

	fmt.Println(test()) //2
	//test_1()
	//test_2()
	test_3()
}

func test() int {
	var s int
	s++ //s=1
	defer func() {
		s++ //s=3
		fmt.Println("延迟中的s是", s)
	}()
	return s + 1 //s=2
}

func test_1() {
	var num = [5]int{1, 2, 3, 4, 5}
	s := num[2:4]
	s[0] = 0
	fmt.Println("s", cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println("s1", cap(s))
	fmt.Println(s)
	fmt.Println(num)
	ss := s
	ss[0] = 99
	fmt.Println()
	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Println(ss)
	fmt.Println("1", cap(ss))
	ss = append(ss, 1, 1)
	fmt.Println("2", cap(ss))
	ss[0] = 100
	fmt.Println(s)
}

func cat_slice() {
	a := []int{1}
	//	截取slice,可以使用len长度在前面，不会越界，如果大于len的数后都越界
	a = a[len(a):]

}

//直接创slice时也相对于底层引用了数组/所以初始化一个slice把一个指针指向其中一个元素时，当元素的值改变时，切片对应元素也发送改变，当切边扩容后则不会
func test_2() {
	num := make([]int, 5)
	fmt.Println("cap", cap(num))
	fmt.Println(num)
	s := &num[1]
	*s = 1
	fmt.Println("s", s)
	fmt.Println(num)
	num = append(num, 1, 2, 3, 4, 5)
	fmt.Println("cap", cap(num))
	fmt.Println(s)
	*s = 2
	fmt.Println(num)

}

//map是无序的/cap是不能用于map的/map不能被用于取指针，因为扩容后会造成地址不正确/map的key要选择零值不为nil的类型
func test_3() {
	m := make(map[string]int, 5)
	length := len(m)
	fmt.Println(length)
	var s string
	s = "1"
	for i := 0; i < 5; i++ {
		m[s] = i
		s = s + "i"
	}
	length = len(m)
	fmt.Println(length)
	fmt.Println(m)
	m["s"] = 2
	fmt.Println(len(m))
	m1 := make(map[int]string, 5)
	fmt.Println(m1)
}
