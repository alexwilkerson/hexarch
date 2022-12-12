package arithmetic

type Arith struct {
}

func New() *Arith {
	return &Arith{}
}

func (arith Arith) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arith Arith) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arith Arith) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arith Arith) Division(a, b int32) (int32, error) {
	return a / b, nil
}
