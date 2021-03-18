package util

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Instruction struct {
	Command string
	Params  string
}

func NewInstruction(reader io.Reader) (*Instruction, error) {
	fmt.Println("Enter your input (e.g GenerateMonthlyPayslip \"Mary Song\" 60000)")

	line, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.Replace(line, "\n", "", -1)

	s := strings.SplitN(line, " ", 2)
	var instruction Instruction
	instruction.Command = strings.TrimSpace(s[0])

	if len(s) > 1 {
		instruction.Params = strings.TrimSpace(s[1])
	}

	return &instruction, nil
}

func ExecuteCommand(i interface{}, command string) interface{} {
	method := reflect.ValueOf(i).MethodByName(command)
	return method.Call([]reflect.Value{})[0].Interface()
}
