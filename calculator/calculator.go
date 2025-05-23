package calculator

type Operation interface {
	Apply(num1, num2 float64) (float64, error)
	Symbol() string
}
