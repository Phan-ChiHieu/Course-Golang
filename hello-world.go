package main

// const s string = "constant" // constant

func main() {
	// Hello world
	// fmt.Println("hello world")  => In ra terminal hello world

	//===============================================

	// Values
	// fmt.Println("1+1=", 1+1) 1+1=2

	//===============================================

	// Variables
	// var a = "initial"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// var d = true
	// fmt.Println(d)

	// var e int
	// fmt.Println(e)

	// Khai bao Variables nhanh
	// f := "apple"
	// fmt.Println(f)

	//===============================================

	// Constant
	// fmt.Println(s) => line 8

	// const n = 5000000

	// const d = 3e20 / n
	// fmt.Println(d)

	// fmt.Println(int64(d))

	// fmt.Println(math.Sin(n))

	//===============================================

	// For loop

	// cach 1
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// // cach 2
	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	//===============================================

	// If/Else

	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")

	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	//===============================================

	// Switch

	// i := 2
	// fmt.Println("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	//
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("It's a weekday")

	// }

	// t := time.Now() // lay ra gio hien tai
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm a int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	//===============================================

	// Array
	// var a [5]int
	// fmt.Println("emp", a) // [0 0 0 0 0]

	// a[4] = 100
	// fmt.Println("set", a)      // set [0 0 0 0 100]
	// fmt.Println("get", a[4])   // get 100
	// fmt.Println("len", len(a)) // len 5

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl", b) // dcl [1 2 3 4 5]

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d", twoD) // 2d [[0 1 2] [1 2 3]]

	//===============================================

	// Slices

	// var s []string
	// fmt.Println("uninit:", s, s == nil, len(s) == 0) // [] true true

	// s = make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // [  ] len: 3 cap: 3

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)      // [a b c]
	// fmt.Println("get:", s[2])   // get: c
	// fmt.Println("len:", len(s)) // len: 3

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd:", s) // pd: [a b c d e f]

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c) // cpy: [a b c d e f]

	// l := s[2:5]
	// fmt.Println("sl1:", l) // sl1: [c d e]

	// l = s[:5]
	// fmt.Println("sl2:", l) // sl2: [a b c d e] => [0,1,2,3,4] => khong bao gom phan tu thu 5

	// l = s[2:]
	// fmt.Println("sl3:", l) // sl3: [c d e f] => [2,3,4,5] => bao gom luon ca phan tu thu 2

	// t := []string{"g", "h", "i"}
	// fmt.Println("dcl", t) // dcl [g h i]

	// twoD := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]

	// d := []string{"a", "b", "c", "d"}
	// e := d[2:]
	// fmt.Println("e ==", e) // e == [c d]
	// e[1] = "m"
	// fmt.Println("e ==", e) // e == [c m]
	// fmt.Println("d ==", d) // d == [a b c m]

	// addvanced => https://go.dev/blog/slices-intro

	//===============================================

	// Maps
	// m := make(map[string]int)

	// m["k1"] = 7
	// m["k2"] = 13

	// fmt.Println("map:", m) // map: map[k1:7 k2:13]

	// v1 := m["k1"]
	// fmt.Println("v1:", v1) // v1: 7

	// v3 := m["k3"]
	// fmt.Println("v3:", v3) // v3: 0 => If the key doesn’t exist, the zero value of the value type is returned.

	// fmt.Println("len:", len(m)) // len: 2

	// delete(m, "k2")
	// fmt.Println("map:", m) // map: map[k1:7]

	// _, prs := m["k2"]
	// fmt.Println("prs", prs) // prs false

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map", n) // map: map[bar:2 foo:1]
}
