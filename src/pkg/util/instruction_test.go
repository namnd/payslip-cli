package util_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/namnd/payslip-cli/pkg/util"
)

func TestNewInstruction(t *testing.T) {
	testCases := []struct {
		input    string
		expected *util.Instruction
	}{
		{
			"NakedCommand\n",
			&util.Instruction{
				Command: "NakedCommand",
				Params:  "",
			},
		},
		{
			"CommandWithSingleParam Param\n",
			&util.Instruction{
				Command: "CommandWithSingleParam",
				Params:  "Param",
			},
		},
		{
			"CommandWithMultipleParams FirstParam SecondParam\n",
			&util.Instruction{
				Command: "CommandWithMultipleParams",
				Params:  "FirstParam SecondParam",
			},
		},
	}

	var stdin bytes.Buffer
	for _, testCase := range testCases {
		stdin.Write([]byte(testCase.input))

		instruction, _ := util.NewInstruction(&stdin)
		if *instruction != *testCase.expected {
			t.Errorf("Expected command %v, got command %v", testCase.expected.Command, instruction.Command)
		}
	}
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func TestExecuteCommand(t *testing.T) {
	john := &Person{Name: "John"}

	greet := util.ExecuteCommand(john, "Say")
	if greet != "Hello, I'm John" {
		t.Error("Expect a Person to Say something")
	}

}
