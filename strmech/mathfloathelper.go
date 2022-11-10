package strmech

import "sync"

type MathFloatHelper struct {
	lock *sync.Mutex
}

//	FloatToIntFracRunes
//
//	Receives one of several types of floating point
//	values and converts that value to an integer digit
//	rune array and a fractional digit rune array.
//
//	The integer and fractional digit rune arrays
//	represent and absolute value of the original floating
//	point number.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	floatingPointNumber interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
func (mathFloatHelper *MathFloatHelper) FloatNumToIntFracRunes(
	floatingPointNumber interface{},
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	errorPrefix interface{}) (
	numberSign NumericSignValueType,
	err error) {

}
