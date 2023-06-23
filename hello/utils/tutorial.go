package utils

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
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

func extend(val string) func() string {
	i := 0
	return func() string {
		i++
		return val[:i]
	}
}

func RunFunctionAsAVariableTutorial() {

	ca := "cloudacademy"

	word := extend(ca)

	for i := 0; i < len(ca); i++ {
		fmt.Println(word())
	}

}

type person struct {
	firstname string
	surname   string
}

type lecture struct {
	name       string
	instructor person
	duration   int //seconds
}

func RunStructTutorial() {
	lectures := []lecture{
		{"Structs", person{"Jeremy", "Cook"}, 3000},
		{"Pointers", person{"Jeremy", "Cook"}, 3000},
		{"Functions", person{"Jeremy", "Cook"}, 3000},
	}

	for _, value := range lectures {
		name := value.name
		instructor := fmt.Sprintf("%s %s", value.instructor.firstname, value.instructor.surname)
		duration := value.duration

		fmt.Printf("Lecture: '%s', Author: %s, Duration: %d secs\n", name, instructor, duration)
	}
}

func RunPointerHowTutorial() {
	var num1 int = 100
	var num2 int = 200

	var str1 string = "blah"

	var ptr1 *int = &num1

	fmt.Printf("mem address of num1 is %p\n", &num1)
	fmt.Printf("mem address of num2 is %p\n", &num2)
	fmt.Printf("mem address of str1 is %p\n", &str1)

	fmt.Printf("ptr1 points to mem address %p\n", ptr1)

	*ptr1 = 101
	fmt.Println(num1)

	ptr1 = &num2
	fmt.Printf("ptr1 points to mem address %p\n", ptr1)
	fmt.Println(*ptr1)

	ptr2 := new(int)
	ptr2 = ptr1
	fmt.Println(*ptr2)
}

func notString(msg string) {
	msg = fmt.Sprintf("not%s", msg)
}

func notStringPtr(msg *string) {
	*msg = fmt.Sprintf("not%s", *msg)
}

func RunPointerWhyTutorial() {
	message := "cloudacademy"

	notString(message)
	fmt.Println(message)

	notStringPtr(&message)
	fmt.Println(message)
}

type user struct {
	firstname string
	surname   string
	age       int
}

func (p *user) fullname() string {
	return p.firstname + " " + p.surname
}

func (p *user) canDrive() bool {
	if p.age >= 20 {
		return true
	} else {
		return false
	}
}

func (p *user) updateAge(newAge int) {
	p.age = newAge
}

func RunStructFunction() {
	person1 := user{"Ola", "Sam", 80}
	person2 := user{"Jude", "Stark", 16}

	fmt.Printf("%s can drive: %t\n", person1.fullname(), person1.canDrive())
	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())

	person2.updateAge(person2.age + 4)
	fmt.Println(person2.age)

	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())
}

type upstring string

func (msg upstring) up() string {
	s := string(msg)
	return strings.ToUpper(s)
}

func RunNonStructFunctionsTutorial() {
	message := upstring("Mint")
	fmt.Println(message.up())
}

type device interface {
	turnOn() string
	turnOff() string
}

type iphone struct {
	name  string
	model string
}

type imac struct {
	name  string
	model string
}

func (phone iphone) turnOn() string {
	return "iOS starting up..."
}

func (phone iphone) turnOff() string {
	return "iOS shutting down..."
}

func (mac imac) turnOn() string {
	return "macOS starting up..."
}

func (mac imac) turnOff() string {
	return "macOS shutting down..."
}

func RunInterfaceTutorial() {
	dev1 := iphone{"iPhone", "11"}
	dev2 := imac{"iMac", "27 5K Retina"}

	devices := []device{dev1, dev2}

	for _, device := range devices {
		fmt.Println(device.turnOff())
	}
}

func circleArea(radius float32) (float32, error) {
	if radius < 0 {
		return 0, errors.New("Radius should be a positive value")
	} else {
		return math.Pi * radius * radius, nil
	}
}

func RunErrorTutorial() {
	area1, _ := circleArea(3)
	fmt.Println(area1)

	if area2, err := circleArea(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area2)
	}
}

func doSomething(msg string) {
	fmt.Println(msg)
}

func system() int {
	fmt.Println("system started...")

	defer doSomething("cleanup")
	defer doSomething("stop")

	fmt.Println("system finished!")

	return 1

}

func RunDeferTutorial() {

	data := system()
	fmt.Println(data)

}

func system2() int {
	fmt.Println("system started...")

	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}("blah!")

	var data []int
	var x = data[0]
	x++

	fmt.Println("system finished")
	return 1

}

func RunPanicRecoverTutorial() {
	data := system2()
	fmt.Println(data)
	panic("die!")
}

func pause() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomethingGoRoutine(msg string) {
	pause()
	fmt.Println(msg)
}

func RunGoRoutineTutorial() {
	doSomethingGoRoutine("sync1")

	go doSomethingGoRoutine("async1")
	go doSomethingGoRoutine("async2")
	go doSomethingGoRoutine("async3")

	doSomethingGoRoutine("sync2")

	time.Sleep(time.Second * 10)
}

func RunGoChannel() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "Cloud"
		msgChan <- "Academy"
		msgChan <- "2023"
	}()

	msg1 := <-msgChan
	msg2 := <-msgChan
	msg3 := <-msgChan

	fmt.Println(msg1, msg2, msg3)

}

func RunBufferChannel() {
	size := 3
	var buffChan = make(chan int, size)

	// reader
	go func() {
		for {
			_ = <-buffChan
			time.Sleep(time.Second * 3)
		}
	}()

	// Writer
	writer := func() {
		for i := 0; i <= 10; i++ {
			buffChan <- i
			println(i)
		}
	}

	writer()
}

// write only using channel direction notation
func in(channel chan<- string, msg string) {
	channel <- msg
}

// Read only with channel direction notation
func out(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func RunChannelDirections() {
	channel := make(chan string, 1)

	go out(channel)

	for i := 0; i < 10; i++ {
		in(channel, fmt.Sprintf("cloudacademy - %d", i))
	}

	time.Sleep(time.Second * 10)
}

func pauseChannel() {
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Second)
}

func func1(c chan<- string) {
	for {
		pause()
		c <- "cloud"
	}
}

func func2(c chan<- string) {
	for {
		pause()
		c <- "academy"
	}
}

// Run select with channels
func RunChannelSelectTutorial() {
	rand.Seed(time.Now().Unix())

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func1(chan1)
	go func2(chan2)

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}

}

func RunChannelTimeoutTutorial() {
	channel := make(chan string)

	go func(channel chan string) {
		time.Sleep(2 * time.Second)
		channel <- "cloudacademy"
	}(channel)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout!")
	}
}

func ChannelCloseTutorial(work <-chan string, fin chan<- string) {
	var b bytes.Buffer
	for {
		if msg, notClosed := <-work; notClosed {
			fmt.Printf("%s received...\n", msg)
			b.WriteString(msg)
		} else {
			fmt.Println("channel closed")
			fin <- b.String()
			return
		}
	}
}

func RuncChannelCloseTutorial(){
	
	work := make(chan string, 3)
	fin := make(chan string)

	go ChannelCloseTutorial(work, fin)

	word := "cloudacademy"

	for j := 0; j < len(word); j++ {
		letter := string(word[j])
		work <- letter
		fmt.Printf("%s sent ...\n", letter)
	}

	close(work)

	fmt.Printf("result: %s", <-fin)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func RunChannelRangeTutorial() {
	f := squares()
	squares := make(chan int, 20)
	for i := f(); i <= 100; i = f() {
		squares <- i
	}

	close(squares)

	for elem := range squares {
		fmt.Println(elem)
	}
}
