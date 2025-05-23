package calculator

type Subtraction struct{}

func (s Subtraction) Apply(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func (s Subtraction) Symbol() string {
	return "-"
}
