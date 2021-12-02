package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func getInput(day int) string {
	all, _ := ioutil.ReadFile(path.Join(".", "inputs", fmt.Sprintf("%d.txt", day)))
	return string(all)
}

var days = []func(string) string{
	func(s string) string { return "" },
	SolveSecond,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no day specified")
	}

	day, _ := strconv.Atoi(os.Args[1])
	if day > len(days) {
		log.Fatal("not yet")
	}
	fmt.Println(days[day-1](getInput(day)))
}
