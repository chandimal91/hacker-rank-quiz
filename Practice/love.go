package main

import (
	"fmt"
	"strings"
)

/* Interview code */

func main() {

	var w, l string
	var count int
	fmt.Scanln(&w)
	fmt.Scanln(&l)
	var f bool
	var s bool
	var q bool

	arr := strings.Fields(l)
	r := strings.Split(w,"")

	for _, word := range arr {
		c := strings.Split(word,"")

		for _, letter := range r{
			f = strings.ContainsAny(word, letter)
			if f == false {
				break;
			}
			s = true
		}

		for _, letter := range c{
			f = strings.ContainsAny(w, letter)
			if f == false {
				break;
			}
			q = true
		}

		if s == true && q == true {
		count++
	    }
	}
	
	fmt.Print(count)
}