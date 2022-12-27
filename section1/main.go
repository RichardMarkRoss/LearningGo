package main

import "fmt"

func main() {
var m1 int32;
var m2 int64;
m3 := 4;

	m1 = 2;
	m2 = 3;
	fmt.Println(m1); 
	fmt.Println(m2);
	fmt.Println(m3);
	fmt.Println(int64(m1)+m2); // casting the values
	fmt.Println(int32(m2)+m1);
	// fmt.Println(m3+m1); // has to be casted
	// fmt.Println(m3+m2); // has to be casted
	fmt.Println(int32(m3)+m1);
	fmt.Println(int64(m3)+m2);
}