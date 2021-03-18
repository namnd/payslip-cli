package main

import (
	"fmt"
	"log"
	"os"

	"github.com/namnd/payslip-cli/pkg/util"
)

func main() {
	fmt.Println("MYOB Dev Test")

	instruction, err := util.NewInstruction(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instruction)
}
