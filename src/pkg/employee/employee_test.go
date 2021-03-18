package employee_test

import (
	"testing"

	"github.com/namnd/payslip-cli/pkg/employee"
)

func TestNewEmployee(t *testing.T) {
	testCases := []struct {
		input    string
		expected *employee.Employee
		errMsg   string
	}{
		{
			"\"John Doe\" 1000",
			&employee.Employee{
				Name:         "\"John Doe\"",
				AnnualSalary: 1000,
			},
			"",
		},
		{
			"John 2000",
			&employee.Employee{
				Name:         "John",
				AnnualSalary: 2000,
			},
			"",
		},
		{
			"John",
			&employee.Employee{
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
		employee, err := employee.NewEmployee(testCase.input)
		if err != nil {
			if err.Error() != testCase.errMsg {
				t.Errorf("Expected error %s, got %s", testCase.errMsg, err.Error())
			}
		} else if *employee != *testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, employee)
		}
	}

}
