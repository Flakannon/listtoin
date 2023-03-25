package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Property struct {
	PropertyId string
}

func main() {

	file, e := os.Open("list.txt")
	if e != nil {
		fmt.Println("Error is = ", e)
	}

	var inSlice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), ",")
		var id Property
		id.PropertyId = s[0]
		inSlice = append(inSlice, fmt.Sprintf("'%v',", id.PropertyId))

	}

	file.Close()

	fmt.Print(inSlice)
}
