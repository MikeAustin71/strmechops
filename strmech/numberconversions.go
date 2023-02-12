package strmech

import (
	"sync"
)

type NumberConversions struct {
	lock *sync.Mutex
}

// ConvertDecimalToOctal
//
// Utility routine to convert a decimal (base 10) numeric
// value to an octal (base 8) numeric value. Useful in
// evaluating 'os.FileMode' values and associated
// constants.
//
//	Reference:
//
//	https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ----------------------------------------------------------------
//
// Usage:
//
//	 initialDecimalValue := 511
//	 expectedOctalValue := 777
//
//		var actualOctalValue int
//
//	 actualOctalValue =
//			new(NumberConversions).
//				convertDecimalToOctal(initialDecimalValue)
//
//	 'actualOctalValue' is now equal to integer value '777'.
//
// ----------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and converted
// to a decimal value. Therefore, x:= int(0777) will mean
// that 'x' is set equal to 511. If you set x:= int(777), x
// will be set equal to '777'. '777' is the correct result.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decimalNumber				int
//
//		The decimal or base 10 number which will be
//		converted to an octal value.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int
//
//		The converted octal (Base 8) number. Input
//		parameter 'decimalNumber' (Base 10) will be
//		converted to an octal number and returned
//		through this parameter.
func (numConversions *NumberConversions) ConvertDecimalToOctal(
	decimalNumber int) int {

	if numConversions.lock == nil {
		numConversions.lock = new(sync.Mutex)
	}

	numConversions.lock.Lock()

	defer numConversions.lock.Unlock()

	return new(numberConversionsMolecule).
		convertDecimalToOctal(decimalNumber)
}

// ConvertOctalToDecimal
//
// Utility routine to convert an octal (base 8) numeric
// value to a decimal (base 10) numeric value. Useful in
// evaluating 'os.FileMode' values and associated
// constants.
//
//	Reference:
//
//	https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ----------------------------------------------------------------
//
// Usage:
//
//		expectedDecimalValue := 511
//		initialOctalValue := 777
//		actualDecimalValue :=
//			new(numberConversionsMolecule).
//				convertOctalToDecimal(initialOctalValue)
//
//	 actualDecimalValue is now equal to integer value, '511'.
//
// ----------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and
// converted to a decimal value. Therefore, x:= int(0777)
// will mean that 'x' is set equal to 511. If you set
// x:= int(777), x will be set equal to '777'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	octalNumber					int
//
//		The octal or Base 8 number which will be
//		converted to a decimal numeric value (Base 10).
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int
//
//		The converted decimal (Base 10) number. Input
//		parameter 'octalNumber' (Base 8) will be
//		converted to a decimal number and returned
//		through this parameter.
func (numConversions *NumberConversions) ConvertOctalToDecimal(
	octalNumber int) int {

	if numConversions.lock == nil {
		numConversions.lock = new(sync.Mutex)
	}

	numConversions.lock.Lock()

	defer numConversions.lock.Unlock()

	return new(numberConversionsMolecule).
		convertOctalToDecimal(octalNumber)
}
