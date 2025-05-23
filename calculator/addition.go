package calculator

type Addition struct{}

func (a Addition) Apply(num1, num2 float64) (float64, error) {
	return num1 + num2, nil

}

func (a Addition) Symbol() string {
	return "+"
}
