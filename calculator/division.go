package calculator

type Division struct{}

func (d Division) Apply(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, ErrDivideByZero
	}
	return num1 / num2, nil
}

func (d Division) Symbol() string {
	return "/"
}
