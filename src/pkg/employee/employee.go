package employee

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Employee struct {
	Name         string
	AnnualSalary float64
}

func NewEmployee(input string) (*Employee, error) {
	matched, err := regexp.MatchString(`^(\"(.)*\"|[\w]*)\s[\d]*$`, input)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, errors.New("Invalid format")
	}

	salaryStr := input[strings.LastIndex(input, " ")+1:]
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		return nil, err
	}
	name := strings.Trim(input, salaryStr)
	return &Employee{
		Name:         strings.TrimSpace(name),
		AnnualSalary: salary,
	}, nil
}

func (e *Employee) GenerateMonthlyPayslip() string {
	return fmt.Sprintf(`
Monthly Payslip for: %s
Gross Monthly Income: $%.0f
Monthly Income Tax: $%.0f
Net Monthly Income: $%.0f
`,
		e.Name,
		e.GetGrossMonthlyIncome(),
		e.GetMonthlyIncomeTax(),
		e.GetGrossMonthlyIncome()-e.GetMonthlyIncomeTax())
}

func (e *Employee) GetGrossMonthlyIncome() float64 {
	return e.AnnualSalary / 12
}

func (e *Employee) GetMonthlyIncomeTax() float64 {
	return 500.0
}
