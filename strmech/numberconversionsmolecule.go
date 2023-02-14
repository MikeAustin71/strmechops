package strmech

import (
	"math"
	"sync"
)

type numberConversionsMolecule struct {
	lock *sync.Mutex
}

// convertDecimalToOctal
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
//			new(numberConversionsMolecule).
//			convertDecimalToOctal(initialDecimalValue)
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
//	octalNumber					int
//
//		The converted octal (Base 8) number. Input
//		parameter 'decimalNumber' (Base 10) will be
//		converted to an octal number and returned
//		through this parameter.
func (numConvertMolecule *numberConversionsMolecule) convertDecimalToOctal(
	decimalNumber int) (octalNumber int) {

	if numConvertMolecule.lock == nil {
		numConvertMolecule.lock = new(sync.Mutex)
	}

	numConvertMolecule.lock.Lock()

	defer numConvertMolecule.lock.Unlock()

	counter := 1
	remainder := 0

	for decimalNumber != 0 {
		remainder = decimalNumber % 8
		decimalNumber = decimalNumber / 8
		octalNumber += remainder * counter
		counter *= 10
	}

	return octalNumber
}

// convertOctalToDecimal
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
//	decimalNumber				int
//
//		The converted decimal (Base 10) number. Input
//		parameter 'octalNumber' (Base 8) will be
//		converted to a decimal number and returned
//		through this parameter.
func (numConvertMolecule *numberConversionsMolecule) convertOctalToDecimal(
	octalNumber int) (decimalNumber int) {

	if numConvertMolecule.lock == nil {
		numConvertMolecule.lock = new(sync.Mutex)
	}

	numConvertMolecule.lock.Lock()

	defer numConvertMolecule.lock.Unlock()

	counter := 0.0
	remainder := 0

	for octalNumber != 0 {
		remainder = octalNumber % 10
		decimalNumber += remainder * int(math.Pow(8.0, counter))
		octalNumber = octalNumber / 10
		counter++
	}

	return decimalNumber
}
