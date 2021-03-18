package main

import (
	"fmt"
	"log"
	"os"

	"github.com/namnd/payslip-cli/pkg/employee"
	"github.com/namnd/payslip-cli/pkg/util"
)

func main() {
	fmt.Println("MYOB Dev Test")

	instruction, err := util.NewInstruction(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	employee, err := employee.NewEmployee(instruction.Params)
	if err != nil {
		log.Fatal(err)
	}

	output := util.ExecuteCommand(employee, instruction.Command)
	fmt.Println(output)
}
