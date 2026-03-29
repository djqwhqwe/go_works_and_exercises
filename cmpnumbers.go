func LessThan[T float32 | float64](op1, op2 T) bool {
	op1Big := new(big.Float).SetFloat64(float64(op1))
	op2Big := new(big.Float).SetFloat64(float64(op2))
	return bop1Big.Cmp(op2Big) < 0
}

func GreaterThan[T float32 | float64](op1, op2 T) bool {
	op1Big := new(big.Float).SetFloat64(float64(op1))
	op2Big := new(big.Float).SetFloat64(float64(op2))
	return op1Big.Cmp(op2Big) > 0
}

func Equal[T float32 | float64](op1, op2 T) bool {
	op1Big := new(big.Float).SetFloat64(float64(op1))
	op2Big := new(big.Float).SetFloat64(float64(op2))
	return op1Big.Cmp(op2Big) == 0
}

func LessThanOrEqual[T float32 | float64](op1, op2 T) bool {
	if LessThan(op1,op2) || Equal(op1,op2) {
		return true
	}
	return false
}

func GreaterThanOrEqual[T float32 | float64](op1, op2 T) bool {
	if GreaterThan(op1,op2) || Equal(op1,op2) {
		return true
	}
	return false
}



type byteptr *byte
type intptr *int
type intptr8 *int8
type intptr16 *int16
type intptr32 *int32
type intptr64 *int64

type uintptr *uint
type uintptr8 *uint8
type uintptr16 *uint16
type uintptr32 *uint32
type uintptr64 *uint64

type float32ptr *float32
type float64ptr *float64

type stringptr *string