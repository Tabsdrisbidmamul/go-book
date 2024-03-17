package main

import (
	"fmt"
	"os"
)

func main() {
	var value = "a string";
	anotherValue := "another value";

	fmt.Printf("value: %s\nand another value is %s\n", value, anotherValue);

	for i:= 0; i < 3; i++ {
		fmt.Printf("i is %d\n", i)
	}

	var numbers = [5]int32{1, 2, 3, 4, 5};
	
	fmt.Printf("numbers: %v\n", numbers);

	// var slice = make([]int32, 5);
	var slice = numbers[0:5];

	fmt.Printf("slice: %v\n", slice);

	var add = func(x, y int) int {
		return x + y;
	}

	fmt.Println(add(1, 2));

	var f, err = os.Open("test.txt");
	defer f.Close();

	if(err != nil) {
		fmt.Printf("ERROR %s\n", err);
	}

	// defer func() {
	// 	var err_msg = recover();
	// 	fmt.Printf("err msg %s\n", err_msg);
	// }()

	// panic("PANIC");

	var x = 5;
	zero(&x);
	fmt.Printf("x is %d\n", x);

}

func zero(ptr *int) {
	*ptr = 0;
}