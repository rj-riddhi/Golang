package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const Name = "Hello" // variable name with capital letter will consider as a public var
func main() {
	fmt.Println("John!!!")

	// variables
	var name string = "John"
	var surname = "Dao"
	age := 20 // := will not used out side of function

	// use of three different Print functions
	fmt.Println(".........................Go print functions.........................")
	fmt.Println(name, surname)
	fmt.Print(name, " ", surname, "\n")
	fmt.Println(age)
	// Printf function with the use of verbs
	fmt.Printf("The name is %v and the type is %T and name with go syntax %#v \n", name, name, name)

	// datatypes: bool(default value is false), numeric (float32, int), string
	fmt.Println(".........................Go Datatypes.........................")
	var user_name string = "Hello"
	var value int = 2
	var float_val float32 = 15.5
	var flag bool = true
	fmt.Println(user_name)
	fmt.Println(value)
	fmt.Println(float_val)
	fmt.Println(flag)

	/*
		Arrays
		1. using var keyword syntax is as below
		var array_name = [length]datatype{values} OR
		var array_name = [...]datatypes{values}
	*/
	fmt.Println(".........................Go print Array.........................")
	var ages = [2]int{1, 2}
	fmt.Println(ages)
	var ages_without_length = [...]int{1, 2}
	fmt.Println("The length of the array is:", len(ages_without_length))

	/*
		2. using := and the syntax is as below
		array := [length]datatype{values} OR
		array := [...]datatype{values}
	*/
	values := [2]int{1} //intialized with one values OR partially initialised array so in the op remaining elements will become 0 (int => 0, string => "", float32 => 0.0)
	fmt.Println(values)
	// if we want to initialize only some specific values at index
	specific_val_init := [3]int{2: 5}
	fmt.Println(specific_val_init)

	/* Slices
	1. create slice with given syntax var slice = []datatype{} OR slice := []datatype{}
	*/
	fmt.Println(".........................Go Slice.........................")
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// 2. create slice from an array
	array1 := [4]int{1, 2, 3, 4}
	slice_from_array := array1[1:3]
	fmt.Println(slice_from_array)

	// 3. create a slice with the make() function
	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// append element to the slice
	fmt.Println(".........................Go slice operations.........................")
	myslice4 = append(myslice4, 1, 2)
	fmt.Println(myslice4)

	// appending one slice with another slice elements
	myslice5 := []int{}
	myslice5 = append(myslice1, myslice3...)
	fmt.Println(myslice5)

	/* operators
	1. Arithmetic operator (+,-,*,/,%,++,--)
	*/
	fmt.Println(".........................Go OPerators.........................")
	a := 1
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// 2. Assignment Operator (=, +=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=)
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)
	a &= b
	fmt.Println(a)
	a |= b
	fmt.Println(a)
	a ^= b
	fmt.Println(a)
	a >>= b
	fmt.Println(a)
	a <<= b
	fmt.Println(a)

	// 3. Comparision Operator
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// 4. Logical Operator
	fmt.Println(a <= 1 && a > 0)
	fmt.Println(a < 2 || a > 1)
	fmt.Println(!(a < 2))

	// 5. Bitwise Operator
	fmt.Println(a & b) // a = 001, b = 010 result = 000 = 0
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

	// Conditional Statements
	/* if and else
	valid syntax:
		if {

		} else {

		}
	invalid syntax:
		if {
		}
		else {
		}
	*/
	fmt.Println(".........................Go Conditional statements.........................")
	if a < b {
		fmt.Printf("%v is greater then %v \n", b, a)
	} else {
		fmt.Printf("%v is greate then %v \n", a, b)
	}

	/* else if condition
	if condition1 {
	} else if condition2 {
	} else {
	}
	*/
	if a > 2 {
		fmt.Printf("%v is greater then 2 \n", a)
	} else if b > 1 {
		fmt.Printf("%v is greater then 1 \n", b)
	} else {
		fmt.Printf("default result")
	}

	/* Switch Statements
	   1. Single switch case
	*/
	fmt.Println(".........................Go switch statements.........................")
	day := 4
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursady")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not in week days")
	}

	// 2. Multi Switch case
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// for loop
	fmt.Println(".........................Go for loop.........................")
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	// use of continue
	fmt.Println(".........................Go continue keyword.........................")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// use of break
	fmt.Println(".........................Go break keyword.........................")
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	// use of range keyword
	fmt.Println(".........................Go range keyword.........................")
	fruits := [2]string{"orange", "apple"}
	for idx, val := range fruits {
		fmt.Printf("%v %v \n", idx, val)
	}

	// use of function call
	fmt.Println(".........................Go function and function call.........................")
	functionCalled()

	fmt.Println(".........................Go function with parameters.........................")
	functionWithParameter("John", "Dao")

	fmt.Println(".........................Go function with return value.........................")
	sum := functionWithReturnVal(1, 2)
	fmt.Println(sum)

	fmt.Println(".........................Go funcParmartion with return multiple value.........................")
	fmt.Println(functionWithMultipleReturns(1, "xyz"))
	result1, result2 := functionWithMultipleReturns(1, "xyz")
	fmt.Println(result1, result2)

	fmt.Println(".........................Go struct/structure.........................")
	var person Person
	person.name = "John"
	person.age = 20
	person.job = "Software Developer"
	person.sal = 50000
	functionStructUse(person)

	fmt.Println(".........................Go map.........................")
	// create map using var or :=
	map1 := map[string]string{"name": "john", "age": "20"}
	fmt.Println(map1)

	// create map using make() function
	map2 := make(map[string]string)
	map2["name"] = "John"
	map2["age"] = "21"
	fmt.Println(map2)
	// access map's element
	fmt.Println(map2["name"])
	// update map value
	map2["age"] = "22"
	fmt.Println(map2["age"])
	// remove element from map
	delete(map2, "age")
	fmt.Println(map2)

	// comma ok, comma err (input, _ := or input, err := ) syntax and user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter our pizza rating...")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating ", input)

	// sum of two numbers
	// method 1
	// fmt.Println("Enter two numbers")
	// input1, _ := reader.ReadString('\n')
	// input2, _ := reader.ReadString('\n')
	// input1 = strings.TrimSpace(input1)
	// input2 = strings.TrimSpace(input2)
	// num2, _ := strconv.ParseInt(input2, 10, 64)
	// num1, _ := strconv.ParseInt(input1, 10, 64)
	// sumation := num1 + num2
	// fmt.Println("sum of two numbers are ... ", sumation)

	// Use of Scanln to read user input
	// method 2 (sum of two numbers)
	var num1, num2 int

	// Prompt the user for input
	fmt.Println("Enter two numbers:")

	// Read two integers directly from user input
	_, err1 := fmt.Scanln(&num1)
	_, err2 := fmt.Scanln(&num2)

	// Check for errors
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please enter valid integers.")
		return
	}

	// Calculate the sum
	newsum := num1 + num2

	// Output the result
	fmt.Printf("The sum is: %d\n", newsum)

	// conversion// fmt.Println("Enter two numbers")
	// input1, _ := reader.ReadString('\n')
	// input2, _ := reader.ReadString('\n')
	// input1 = strings.TrimSpace(input1)
	// input2 = strings.TrimSpace(input2)
	// num2, _ := strconv.ParseInt(input2, 10, 64)
	// num1, _ := strconv.ParseInt(input1, 10, 64)
	// sumation := num1 + num2
	// fmt.Println("sum of two numbers are ... ", sumation)

	// Time package
	presentTime := time.Now()
	fmt.Println("Current time is ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday")) //we need to use thi inbuilt time format to format our time into desire format

}

type Person struct {
	name string
	age  int
	job  string
	sal  int
}

func functionCalled() {
	fmt.Println("I am getting executed")
}

func functionWithParameter(firstname string, lastname string) {
	fmt.Printf("My name is %v %v!! \n", firstname, lastname)
}

func functionWithReturnVal(x int, y int) int {
	return x + y
}

func functionWithMultipleReturns(x int, y string) (numeric int, text string) {
	numeric = x + 10
	text = "Hello " + y
	return
}

func functionStructUse(person Person) {
	fmt.Printf("name is %v \nage is %v \njob is %v \nsalary is %v \n", person.name, person.age, person.job, person.sal)
}
