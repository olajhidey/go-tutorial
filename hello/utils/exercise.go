package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func RunHelloExercise() {
	fmt.Println("Hello World :)")
}

func RunVariablesExercise(name string, age int) {
	fmt.Printf("My name is %s and I'm %d years old", name, age)
}

func RunOperatorExercise() {
	var numOfPackages = 100
	var numOfCustomers = 4

	fmt.Printf("I have %d packages to deliver\n", numOfPackages)

	var deliveredPackages = 20
	var remPackages = numOfPackages - deliveredPackages
	fmt.Printf("I have delivered %d packages\n", deliveredPackages)
	fmt.Printf("I have %d remaining packages to deliver\n", remPackages)

	var perCustomerItem = numOfPackages / numOfCustomers
	fmt.Printf("Packages are going to be distributed equally between %v customers. This means %d packages per customer\n", numOfCustomers, perCustomerItem)

}

func RunDataTypesExercise() {
	var food = "Pizza"
	var slices = 4
	var pineappleOnPizza = true
	fmt.Printf("Food is a %T type, slices is a %T type, pineappleOnPizza is a %T type \n", food, slices, pineappleOnPizza)
}

func RunUserInputExercise() {
	var name string
	var age int

	fmt.Println("Enter your name")
	fmt.Scan(&name)

	fmt.Println("Enter your age")
	fmt.Scan(&age)

	RunVariablesExercise(name, age)
}

func RunPackagesExercise() {
	var currentTime = time.Now()

	fmt.Printf(" The time now is %s\n", currentTime)

	var randNum = rand.Intn(100)
	fmt.Printf("The Random number is %d\n", randNum)

	var numSqr = math.Sqrt(9)
	fmt.Printf("The Square root of 9 is %g\n", numSqr)

}

func RunLogicalExercise() {
	x := 2017
	result1 := x > 50 || x < 2020
	result2 := x > 50 || x%2 == 0
	result3 := x%2 == 1 || x+3 == 2025

	fmt.Println(result1) //true
	fmt.Println(result2) // true
	fmt.Println(result3) // true

	x += 5
	result1 = x > 50 || x < 3000
	result2 = x > 50 || x < 3000
	result3 = x > 50 || x < 3000
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

}

func RunConditionalExercise() {
	var randNum = rand.Intn(100)

	if randNum > 50 {
		fmt.Printf("The Random number is %d It's closer to 100", randNum)
	} else if randNum < 50 {
		fmt.Printf("The Random number is %d It's closer to 0", randNum)
	} else if randNum == 50 {
		fmt.Println("It's 50!")
	} else if randNum > 50 && randNum/2 == 0 {
		fmt.Println("It's closer to 100 and it's even!")
	}
}

func RunSwitchExercise() {
	var randNum = rand.Intn(10)

	switch randNum {
	case 1:
		fmt.Println("It's the loneliest number since the number one")
	case 2:
		fmt.Println("Two can be as bad as one. It's the loneliest number since the number one")
	default:
		fmt.Println("Sorry, no mention of the number in the song")
	}
}

func RunSplitStringExercise() {
	name := "Alexander Supertramp"
	fmt.Println("Firstname", strings.Fields(name)[0])
	fmt.Println("Surname", strings.Fields(name)[1])
}
