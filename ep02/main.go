package main

import "fmt"

func main() {
	// var a [5]int32

	// b := [10]int8{
	// 	1, 2, 3, 4,
	// }

	// var twoD [2][3]int
	// for idx1 := 0; idx1 < 2; idx1++ {
	// 	for idx2 := 0; idx2 < 3; idx2++ {
	// 		twoD[idx1][idx2] = idx1 + idx2
	// 	}
	// }

	// fmt.Println(twoD)

	//slice
	// s := make([]int, 3, 4)

	// s[0] = 0
	// s[1] = 1
	// s[2] = 2

	// fmt.Printf("%p\n", s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s = append(s, 3)
	// fmt.Printf("%p\n", s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s = append(s, 4)
	// fmt.Printf("%p\n", s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// map
	// m := make(map[string]int)

	// m["first"] = 1
	// m["second"] = 2
	// fmt.Println(m)

	// m["first"] = 11
	// fmt.Println(m)

	// delete(m, "second")
	// fmt.Println(m)

	// fmt.Println(m["first"])

	//for range

	// nums := []string{"first", "second"}

	// for index, value := range nums {

	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// }

	// myMap := map[string]string{
	// 	"foo":   "bar",
	// 	"hello": "world",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }

	// result, msg := pluswith2(3)
	// fmt.Println(result, msg)

	// res := closure()
	// fmt.Println(res())
	// fmt.Println(res())

	// fmt.Println(fact(7))

	value := 1
	fmt.Printf("inital: %d\n", value)

	passByValue(value)
	fmt.Printf("after passByValue: %d\n", value)

	passByPointer(&value)
	fmt.Printf("third: %d\n", value)

}

// func pluswith2(num int) (int, string) {
// 	return num + 2, "success"
// }

// func closure() func() int {
// 	index := 0
// 	return func() int {
// 		index++
// 		return index
// 	}
// }

// func fact(num int) int {
// 	if num == 0 {
// 		return 1
// 	}

// 	return num * fact(num-1)
// }

func passByValue(val int) {
	val = 0
}

func passByPointer(val *int) {
	*val = 0
}
