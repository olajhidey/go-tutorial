package utils

import (
	"errors"
	"fmt"
)

const Gpopulation int = 10000

func RunBoolTutorial() {
	var isBig bool
	var isFast, hasFourWheels bool = false, true

	fmt.Println(isBig)
	fmt.Println(isFast)
	fmt.Println(hasFourWheels)

	fmt.Println(!hasFourWheels)
	fmt.Println(isFast && hasFourWheels)
	fmt.Println(isFast || hasFourWheels)

	if !isFast {
		fmt.Println("its slow")
	}

	if hasFourWheels {
		fmt.Println("its a car")
	}
}

func RunIntegerTutorial() {
	var count1 int
	var count2 int = 100
	count3 := 200
	count4 := count2 + count3

	fmt.Println(count1) //prints out 0
	count1 = 10
	fmt.Println(count1) // prints out 10

	fmt.Println(count2) // prints out 100
	fmt.Println(count3) // prints out 200
	fmt.Println(count4) // prints out 300

	count3++
	count4--

	fmt.Println(count3) // prints out 201
	fmt.Println(count4) // prints out 299
}

func RunConstantTutorial() {
	const name = "Cloudacademy"

	if true {
		const color = "red"

		fmt.Println(Gpopulation)
		fmt.Println(name)
	}
}

func RunIfElseTutorial() {
	var population int = 5500

	if population < 5500 {
		fmt.Println("small")
	} else if population >= 5000 && population < 7000 {
		fmt.Println("medium")
	} else {
		fmt.Println("large")
	}

	if toggle := true; toggle {
		fmt.Println("its on!")
	}
}

func RunForTutorial() {
	x := 0
	y := 0

	// First for loop usage
	for {
		if x++; x > 2 {
			fmt.Println(x)
			break
		}
	}

	// Second for loop usage
	for y < 3 {
		fmt.Println(y)
		y++
	}

	// Third for loop usage
	for z := 0; z < 10; z++ {
		if z < 8 {
			continue
		}
		fmt.Println(z)
	}
}

func RunSwitchTutorial() {

	// First usage of Switch

	switch signal := -1; signal {
	case 0:
		fmt.Println("red")
	case 1:
		fmt.Println("orange")
	case 2:
		fmt.Println("green")
	default:
		fmt.Println("Can't figure out your selection")
	}

	// Second usage of Switch
	var score int = 63

	switch {
	case score <= 25:
		fmt.Println("beginner")
	case score <= 75:
		fmt.Println("pro")
	default:
		fmt.Println("expert")
	}
}

func RunArrayTutorial() {

	// Regular Array
	var arr1 [4]int
	arr2 := [...]int{3, 1, 5, 10, 100}
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("length arr1", len(arr1))
	fmt.Println("length arr2", len(arr2))

	// Looping through the provided array
	for _, value := range arr2 {
		fmt.Println(value)
	}

	// // using the regular for
	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println(arr2[i])
	// }

	arr1[0] = 10

	fmt.Println(arr1)
	fmt.Println(arr1[0])

	arr3 := [2][2]string{
		{"a1", "a2"},
		{"b1", "b2"},
	}

	fmt.Println(arr3)

}

func RunSpliceTutorial() {
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := []int{1, 3, 5, 6, 11, 13}
	slice2 := slice1[1:2]
	var slice3 = make([]int, 2, 3)

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice3 = slice1[1:4]

	fmt.Println(slice3)
	fmt.Println(len(slice3))

	slice1 = append(slice1, 200, 300, 400)
	fmt.Println(slice1)

	slice2 = append(slice2, []int{7, 8, 9}...)
	fmt.Println(slice2)

	copySlice := make([]int, len(slice1))
	copy(copySlice, slice1)
	fmt.Println(copySlice)
}

type player struct {
	id   int
	rank int
}

func RunMapTutorial() {

	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"

	fmt.Println(map1)

	value1 := map1["key1"]
	fmt.Println(value1)

	map1["key1"] = "value1.1"
	delete(map1, "key2")
	map1["newKey"] = "value4"
	fmt.Println(map1)

	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}
	fmt.Println(team)
}

func RunRangeTutorial() {
	bytes := []int{67, 108, 111, 117, 100, 65, 99, 100, 101, 109, 121}
	for idx, value := range bytes {
		fmt.Print(value, " ")
		if len(bytes)-1 == idx {
			fmt.Println()
		}
	}

	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}
	fmt.Println(team)

	for key, value := range team {
		fmt.Printf("%s -> %d\n", key, value)
	}

	for _, value := range team {
		fmt.Println(value)
	}

	for key := range team {
		fmt.Println("Key: ", key)
	}
}

func sum(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func minus(num1 int, num2 int) (result int) {
	result = num1 - num2
	return
}

func RunFunctionTutorial() {
	fmt.Println(sum(1, 2))
	fmt.Println(minus(10, 2))

	result := func(sum1 int, sum2 int) int {
		sum := sum(sum1, sum2)
		minus := minus(sum1, sum2)
		return sum * minus
	}(5, 3)

	fmt.Println(result)
}

func displayCount(id int, letters ...string) {
	count := 0

	for range letters {
		count++
	}

	fmt.Printf("%d - %d - %T - %s\n", id, count, letters, letters)
}

func RunVariadicFuncTutorial() {
	displayCount(1, "c", "l", "o", "u", "d")
	displayCount(2, "a", "c", "a", "d", "e", "m", "y")

	cloud := []string{"d", "e", "v", "o", "p", "s"}
	displayCount(3, cloud...)
}

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Division by zero not allowed")
	} else {
		return (num1 / num2), nil
	}
}

func RunFunctionWithMultipleValuesTutorial() {
	if result, err := divide(10, 2); err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}

func extend(val string) func() string{
	i:= 0
	return func() string {
		i++
		return val[:i]
	}
}

func RunFunctionAsAVariableTutorial() {

	ca:= "cloudacademy"

	word := extend(ca)

	for i := 0; i < len(ca); i++ {
		fmt.Println(word())
	}

}
