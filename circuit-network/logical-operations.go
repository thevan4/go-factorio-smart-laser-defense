package circuitnetwork

type LogicalOperation = func(inFirst int64, inSecond int64) (compareResult bool)

type CompareType int

const (
	InvalidCompareType CompareType = iota
	BoolCompareType                //0 or 1
	//SumOfInputsToOutputCompareType            //TODO sum input values
)

//GT ...
func GT(f int64, s int64, compareType CompareType) (compareResult bool) {
	return f > s
}

//LT ...
func LT(f int64, s int64) (compareResult bool) {
	return f < s
}

//GTE ...
func GTE(f int64, s int64) (compareResult bool) {
	return f >= s
}

//LTE ...
func LTE(f int64, s int64) (compareResult bool) {
	return f <= s
}

//EQ ...
func EQ(f int64, s int64) (compareResult bool) {
	return f == s
}

//NEQ ...
func NEQ(f int64, s int64) (compareResult bool) {
	return f != s
}
