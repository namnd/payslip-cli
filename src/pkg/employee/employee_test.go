package employee_test

import (
	"fmt"
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
			nil,
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

func TestGenerateMonthlyPayslip(t *testing.T) {
	employee := &employee.Employee{
		Name:         "\"John Doe\"",
		AnnualSalary: 60000,
	}
	expected := fmt.Sprintf(`
Monthly Payslip for: %s
Gross Monthly Income: $%.0f
Monthly Income Tax: $%.0f
Net Monthly Income: $%.0f
`, "\"John Doe\"", 5000.0, 500.0, 4500.0)

	output := employee.GenerateMonthlyPayslip()
	if output != expected {
		t.Errorf("Expected ouput:\n%sActual:\n%s", expected, output)
	}
}

func TestGetGrossMonthlyIncome(t *testing.T) {
	e := &employee.Employee{
		Name:         "John",
		AnnualSalary: 2000,
	}

	expected := 2000.0 / 12
	if result := e.GetGrossMonthlyIncome(); result != expected {
		t.Errorf("Expected result: %.0f, got: %.0f", expected, result)
	}
}

func TestGetMonthlyIncomeTax(t *testing.T) {
	testCases := []struct {
		annualSalary             float64
		expectedMonthlyIncomeTax float64
	}{
		{
			20000,
			0,
		},
		{
			21000,
			100.0 / 12,
		},
		{
			40000,
			2000.0 / 12,
		},
		{
			41000,
			2200.0 / 12,
		},
		{
			80000,
			10000.0 / 12,
		},
		{
			81000,
			10300.0 / 12,
		},
		{
			180000,
			40000.0 / 12,
		},
		{
			181000,
			44000.0 / 12,
		},
	}

	for _, testCase := range testCases {
		e := &employee.Employee{
			Name:         "John",
			AnnualSalary: testCase.annualSalary,
		}

		if result := e.GetMonthlyIncomeTax(); result != testCase.expectedMonthlyIncomeTax {
			t.Errorf("Expected result: %.0f, got: %.0f", testCase.expectedMonthlyIncomeTax, result)
		}
	}
}
