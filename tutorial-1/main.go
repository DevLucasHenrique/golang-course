package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// STRUCTS
type Product struct {
	id uint
	name string
	price float32
	shippingInfo
}

type shippingInfo struct {
	cepDest string
	cepClient string
	address string
	num int
}

type gasEngine struct {
	mpg uint
	gallons uint
}

type electricEngine struct {
	kfwpm uint
	klw uint
}

func (e gasEngine) milesLeft() uint {
	return e.gallons*e.mpg
} 

func (e electricEngine) milesLeft() uint {
	return e.kfwpm*e.klw
}

type engine interface {
	milesLeft() uint
}

func canMakeIt(e engine, miles uint) {
	if miles<=e.milesLeft() {
		fmt.Println("Can take it")
	} else {
		fmt.Println("Can't take it")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Println(myEngine.milesLeft())
	canMakeIt(myEngine, 50)
}



func structs() {
	// var product Product = Product{id: 1, name: "Camisa", price: 119.90}
	var product Product = Product{1,"Camisa", 119.90, shippingInfo{"123-4", "123-4", "av tal", 1}}
	product.price = 129.90
	fmt.Println(product)
	fmt.Println(product.id, product.name, product.price)
}

func stringsF() {
	var str string = "résumé"
	var indexed = str[0]
	fmt.Println(indexed)
	for i, v := range str {
		fmt.Println(i, v)
	}

	var stringsSlice = []string{"a", "b", "c"}

	var strBuilder strings.Builder

	for i := range stringsSlice {
		strBuilder.WriteString(stringsSlice[i])
	}

	var catString = strBuilder.String()
	fmt.Println(catString)
}

func lesson1() {
	intArray := [...]int32{1, 2, 3}
	fmt.Println(intArray)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 2}
	fmt.Printf("\n the length is %v and the capacity is %v", len(intSlice2), cap(intSlice2))
	intSlice2 = append(intSlice2, intSlice...)
	fmt.Printf("\n the length is %v and the capacity is %v", len(intSlice2), cap(intSlice2))

	var slice3 []int32 = make([]int32, 0)
	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice3)

	var num1, num2 = 2, 2
	var result, remainder, err = division(num1, num2)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("\n o resultado de %v dividido por %v é %v", num1, num2, result)
	default:
		fmt.Printf("\n o resultado de %v dividido por %v é %v com resto de %v ", num1, num2, result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was closed")
	default:
		fmt.Println("The division was not closed")
	}

}

func division(n int, d int) (int, int, error) {
	var err error
	if d == 0 {
		err = errors.New("Cannot divided by 0")
		return 0, 0, err
	}

	var r = n / d
	var re = n % d
	return r, re, err
}

func ampersand() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(&intArr[1])
}

func usingMap() {
	var person = map[string]uint8{"Lucas": 14, "Arthur": 13, "John": 34}
	// var age, ok = person["Lucas"]

	/* if ok {
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	} */

	delete(person, "John")
	fmt.Println(person["Lucas"])
	for name, age := range person {
		fmt.Println(name, age)
	}
}

func loops() {
	intArr := []int32{1, 2, 3}
	for index, value := range intArr {
		fmt.Printf("\n Index %v: %v", index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("LOOP!")
	}
}

func testSpeed() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without: %v\n ", timeLoop(testSlice, n))
	fmt.Printf("Total time with: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
