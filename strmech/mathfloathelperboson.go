package strmech

import "sync"

type mathFloatHelperBoson struct {
	lock *sync.Mutex
}

// bigFloatFromPureNumStr
//
// Creates and returns a big.Float floating point
// numeric value by extracting said numeric value
// from a Pure Number String.
func (*mathFloatHelperBoson) bigFloatFromPureNumStr(
	pureNumberValueStr string,
	numOfExtraDigitsPrecisionBuffer int64) (
	BigFloatDto,
	error) {

}
