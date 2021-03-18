package employee

import (
	"testing"
)

func TestNewEmployee(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Employee
		errMsg   string
	}{
		{
			"\"John Doe\" 1000",
			&Employee{
				Name:         "\"John Doe\"",
				AnnualSalary: 1000,
			},
			"",
		},
		{
			"John 2000",
			&Employee{
				Name:         "John",
				AnnualSalary: 2000,
			},
			"",
		},
		{
			"John",
			&Employee{
				Name: "John",
			},
			"Invalid format",
		},
		{
			"\"John Doe\" not a number",
			nil,
			"Invalid format",
		},
	}

	for _, testCase := range testCases {
		employee, err := NewEmployee(testCase.input)
		if err != nil {
			if err.Error() != testCase.errMsg {
				t.Errorf("Expected error %s, got %s", testCase.errMsg, err.Error())
			}
		} else if *employee != *testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, employee)
		}
	}

}
