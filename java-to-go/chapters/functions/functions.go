package main

import "fmt"

func add(x int, y int) int {
	result := x + y
	return result
}

func neg(x *int) {
	neg := (-1) * (*x)
	*x = neg
}

func multReturn(x int, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

func multReturn2(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func area(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero inputs:[%v,%v]", x, y)
	}
	area := x * y
	return &area, nil
}

func printMsg(id string, m string) {
	fmt.Printf("[%v]: %v\n", id, m)
}

func main() {

	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Println("Add", add(1, 2))
	fmt.Println("Sub", sub(1, 2))

	//
	x := 1
	neg(&x)
	fmt.Println("Neg", x)

	//
	sum, sub2 := multReturn(1, 2)
	fmt.Println("Multiple Return SUM", sum)
	fmt.Println("Multiple Return SUB", sub2)

	//
	sum2, _ := multReturn2(1, 2)
	fmt.Println("Multiple Return2 SUM2", sum2)

	//
	area, err := area(2, 1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Area Error Validation", *area)

	//
	msg := "Hello, world!"
	defer printMsg("Defer-0", msg)
	defer printMsg("Defer-1", msg)

	msg = "Goodbye, world!"
	fmt.Println("Main has exited.")
}
