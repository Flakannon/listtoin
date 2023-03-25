package main

import (
	"fmt"

	"github.com-Flakannon/Flakannon/file"
	"github.com-Flakannon/Flakannon/sqlClauses"
)

func main() {
	inValues, err := file.GetFileLines("topics.txt")
	if err != nil {
		fmt.Println("Error is = ", err)
	}

	sql := sqlClauses.CreateInClause(inValues)
	fmt.Println(sql)
}
