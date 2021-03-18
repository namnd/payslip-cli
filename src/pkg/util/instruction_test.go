package util

import (
	"bytes"
	"testing"
)

func TestNewInstruction(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Instruction
	}{
		{
			"NakedCommand\n",
			&Instruction{
				Command: "NakedCommand",
				Params:  "",
			},
		},
		{
			"CommandWithSingleParam Param\n",
			&Instruction{
				Command: "CommandWithSingleParam",
				Params:  "Param",
			},
		},
		{
			"CommandWithMultipleParams FirstParam SecondParam\n",
			&Instruction{
				Command: "CommandWithMultipleParams",
				Params:  "FirstParam SecondParam",
			},
		},
	}

	var stdin bytes.Buffer
	for _, testCase := range testCases {
		stdin.Write([]byte(testCase.input))

		instruction, _ := NewInstruction(&stdin)
		if instruction.Command != testCase.expected.Command {
			t.Errorf("Expected command %v, got command %v", testCase.expected.Command, instruction.Command)
		}
		if instruction.Params != testCase.expected.Params {
			t.Errorf("Expected params %s, got params %s", testCase.expected.Params, instruction.Params)
		}
	}

}
