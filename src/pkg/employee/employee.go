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
	if e.AnnualSalary <= 20000 {
		return 0
	}
	if e.AnnualSalary <= 40000 {
		return ((e.AnnualSalary - 20000) * 0.1) / 12
	}
	if e.AnnualSalary <= 80000 {
		return ((e.AnnualSalary-40000)*0.2 + 2000) / 12
	}
	if e.AnnualSalary <= 180000 {
		return ((e.AnnualSalary-80000)*0.3 + 10000) / 12
	}
	return ((e.AnnualSalary-180000)*0.4 + 40000) / 12
}
