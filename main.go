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

var days = [][2]func(string) string{
	{SolveFirst},
	{SolveSecond},
	{SolveThird, SolveThirdP2},
	{SolveFourth, SolveFourthP2},
	{SolveDay5, SolveDay5P2},
	{SolveDay6, SolveDay6P2},
	{SolveDay7, SolveDay7P2},
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("expected day and part")
	}

	day, _ := strconv.Atoi(os.Args[1])
	if day > len(days) {
		log.Fatal("not yet")
	}
	part, _ := strconv.Atoi(os.Args[2])

	fmt.Println(days[day-1][part-1](getInput(day)))
}
