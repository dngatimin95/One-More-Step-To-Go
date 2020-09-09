package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int { //or add(x, y int) if both variables are the same
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} //empty return statement returns named variables

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //if statement can start with a short statement to execute before the condition
		return v
	}
	return lim
}

var c, python, java bool
var i, j int = 1, 2

type Vertex struct {
	X int
	Y int
}

// concurrency
func count(item string) {
	for i := 0; true; i++ {
		fmt.Println(i, item)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(rand.Intn(20), rand.Intn(30)))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var x int
	fmt.Println(x, c, python, java)
	fmt.Println(split(17))

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	l, m := 1, 2
	k := 3
	fmt.Println(l, m, k, c, python, java)

	const World = "世界" // const cannot be declared with :=
	fmt.Println("Hello", World)

	sum := 0 // only for loop avaialble in go
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	newSum := 1
	for newSum < 100 { // or for ; sum< 100;
		newSum += newSum
	}
	fmt.Println(newSum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//defer fmt.Println("world") // defer statement defers the execution of a function until the surrounding function returns. Deferred function calls are pushed onto a stack
	//fmt.Println("hello")

	/*fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")*/

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	d := &v
	d.X = 1e9
	fmt.Println(v)

	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4]
	fmt.Println(s)

	ra := []int{2, 3, 5, 7, 11, 13}
	ra = ra[:0]
	fmt.Println(s)
	ra = ra[1:4]
	fmt.Println(s)
	ra = ra[:2]
	fmt.Println(s)
	ra = ra[1:]
	fmt.Println(s)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	//range form of the for loop iterates over a slice or map. returns two value: first is the index, and the second is a copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	temp := make([]int, 10)
	for i := range temp {
		temp[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range temp {
		fmt.Printf("%d\n", value)
	}

	go count("sheep")
	count("fish")

}
