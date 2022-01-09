package circuitnetwork

import "time"

const MagicSleepTick int8 = 60

//TickTime - default 60 ticks/ups
var TickTime time.Duration

func init() {
	//Set default tick time
	TickTime = time.Second / 60
}

//SetTickTime ...
func SetTickTime(t time.Duration) {
	TickTime = t
}

type Signal struct {
	Name  string
	Value int
}

//type MathematicalOperation int
//
//const (
//	InvalidMathematical MathematicalOperation = iota
//	Addition                                  //(+)
//	Subtraction                               //(−)
//	Multiplication                            //(*)
//	Division                                  //(/)
//	Modulo                                    //(%)
//	Exponentiation                            //(^)
//	LeftBitShift                              //(<<)
//	RightBitShift                             //(>>)
//	BitwiseAND                                //(&)
//	BitwiseOR                                 //(|)
//	BitwiseXOR                                //(^)
//)
//
//func (mo MathematicalOperation) String() string {
//	switch mo {
//	case Addition:
//		return "addition"
//	case Subtraction:
//		return "subtraction"
//	case Multiplication:
//		return "multiplication"
//	case Division:
//		return "division"
//	case Modulo:
//		return "modulo"
//	case Exponentiation:
//		return "exponentiation"
//	case LeftBitShift:
//		return "leftBitShift"
//	case RightBitShift:
//		return "rightBitShift"
//	case BitwiseAND:
//		return "bitwiseAND"
//	case BitwiseOR:
//		return "bitwiseOR"
//	case BitwiseXOR:
//		return "bitwiseXOR"
//
//	default:
//		return "invalid"
//	}
//}
//
//type LogicalOperation int
//
//const (
//	InvalidLogical LogicalOperation = iota
//	GT                              //(>)
//	LT                              //(<)
//	GTE                             //(>=)
//	LTE                             //(<=)
//	EQ                              //(=)
//	NEQ                             //(!=)
//)
//
//func (lo LogicalOperation) String() string {
//	switch lo {
//	case GT:
//		return "greater than"
//	case LT:
//		return "lower than"
//	case GTE:
//		return "greater or equal"
//	case LTE:
//		return "lower or equal"
//	case EQ:
//		return "equal"
//	case NEQ:
//		return "not equal"
//
//	default:
//		return "invalid"
//	}
//}