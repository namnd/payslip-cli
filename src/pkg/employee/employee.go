package employee

import (
	"errors"
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
