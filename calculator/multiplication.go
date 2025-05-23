package calculator

type Multiplication struct{}

func (m Multiplication) Apply(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

func (m Multiplication) Symbol() string {
	return "*"
}
