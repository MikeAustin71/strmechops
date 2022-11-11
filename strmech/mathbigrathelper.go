package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

//	mathBigRatHelper
//
//	This type provides helper methods for Rational
//	Numbers created with the Go Programming Language
//	Big Math package.
//
//	The Big Math package defines type 'Rat' or Rational
//	Number as follows:
//
//	"A Rat represents a quotient a/b of arbitrary
//	precision. The zero value for a Rat represents the
//	value 0."
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/math/big#Rat
type MathBigRatHelper struct {
	lock *sync.Mutex
}

//	RatToBigFloat
//
//	Converts a rational number (type big.Rat) to a
//	floating point number of type big.Float.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	bigRatNum					*big.Rat
//
//		The rational number which will be converted to a
//		big.Float numeric value.
//
//	roundToFractionalDigits 	int
//
//		Controls the number of fractional digits returned
//		as a big float number ('bigFloatNum').
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bigFloatNum					*big.Float
//
//		If this method completes successfully, this
//		parameter will return a big.Float number
//		representing the numeric value of the rational
//		number passed through input parameter,
//		'bigRatNum'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathBigRatHelp *MathBigRatHelper) RatToBigFloat(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	bigFloatNum *big.Float,
	err error) {

	if mathBigRatHelp.lock == nil {
		mathBigRatHelp.lock = new(sync.Mutex)
	}

	mathBigRatHelp.lock.Lock()

	defer mathBigRatHelp.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	bigFloatNum = big.NewFloat(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathBigRatHelper."+
			"RatToBigFloat()",
		"")

	if err != nil {
		return bigFloatNum, err
	}

	bigFloatNum,
		err = new(mathBigRatHelperAtom).
		ratToBigFloat(
			bigRatNum,
			roundToFractionalDigits,
			ePrefix)

	return bigFloatNum, err
}
