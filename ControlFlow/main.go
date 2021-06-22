package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("main start")
	//deferExample1()
	//deferExample2()
	//panicExample1()
	recoverExample1()
	fmt.Println("main end")
}

// defer runs at the end of the method
func deferExample1() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}

// defer refers to the value of variables passed to it.
func deferExample2() {
	a := "Hello"
	fmt.Println("print a : ", a)
	defer fmt.Println("defer print a : ", a)
	a = "World"
}

// order of execution is defer -> panic
func panicExample1() {
	fmt.Println("panicExample1 method start")
	defer fmt.Println("defer before panic")
	panic("Panicking!!!")
	fmt.Println("panicExample1 method end")
}

//
func recoverExample1() {
	fmt.Println("Recover method start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panicExample1()
	fmt.Println("Recover method end")
}
