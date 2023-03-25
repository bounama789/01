package main

import (
	"fmt"
	"io"
	"os"
	"push-swap/checker/checker"
	"push-swap/utils"
)

func main() {
	if len(os.Args) == 2 {
		a := utils.CheckEntry(os.Args[1])

		opEtries, _ := io.ReadAll(os.Stdin)

		operations := utils.CheckOperations(string(opEtries))

		var b utils.Stack

		a, b = checker.ApplyOperations(operations, a)

		if utils.IsSortedAsc(utils.Compare, a.Data) && len(b.Data) == 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("KO")
		}
	}
}
