package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// mathIntHelperAtom
//
// Provides helper methods for type MathIntHelper

type mathIntHelperAtom struct {
	lock *sync.Mutex
}

//	intNumToSignedPureNumStr
//
//	Receives one of several types of integer values
//	and converts that value to a signed pure number
//	string containing the signed numeric value extracted
//	from the original integer input parameter,
//	'intNumericValue'.
//
//	A signed pure number string consists entirely of
//	numeric digit text characters. In the case of
//	negative numeric values, the negative numeric value
//	is prefixed with a leading minus sign ('-').
//
//	Since the input parameter is an integer value, the
//	returned signed pure number string will NEVER contain
//	fractional digits or a radix point such as a decimal
//	point ('.').
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intNumericValue				interface{}
//
//		Integer numeric values passed by means of this
//		empty interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(currently equivalent to int32)
//			int64
//			uint8
//			uint16
//			uint32
//			uint (currently equivalent to uint32)
//			uint64
//			*big.Int
//
//		If parameter 'intNumericValue' is NOT convertible
//		to one of the types listed above, an error will
//		be returned.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	signedPureNumStr			string
//
//		If this method completes successfully, this
//		parameter will return a signed pure number string
//		representing the numeric value extracted from
//		input parameter 'intNumericValue'.
//
//		A signed pure number string consists entirely of
//		numeric digit text characters. In the case of
//		negative numeric values, the negative numeric value
//		is prefixed with a leading minus sign ('-').
//
//		Since the input parameter is an integer value, the
//		returned signed pure number string will NEVER contain
//		fractional digits or a radix point such as a decimal
//		point ('.').
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
func (mathIntHelpAtom *mathIntHelperAtom) intNumToSignedPureNumStr(
	intNumericValue interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	signedPureNumStr string,
	err error) {

	if mathIntHelpAtom.lock == nil {
		mathIntHelpAtom.lock = new(sync.Mutex)
	}

	mathIntHelpAtom.lock.Lock()

	defer mathIntHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathIntHelperAtom."+
			"intNumToSignedPureNumStr()",
		"")

	if err != nil {

		return signedPureNumStr, err

	}

	var ok bool

	switch intNumericValue.(type) {

	case int8, int16, int, int32, int64,
		uint8, uint16, uint, uint32, uint64:

		signedPureNumStr = fmt.Sprintf("%v",
			intNumericValue)

	case *big.Int:

		var bigIntNum *big.Int

		bigIntNum, ok = intNumericValue.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Int cast to 'bigIntNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr =
			fmt.Sprintf("%v",
				bigIntNum.Text(10))

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'intNumericValue' is an invalid type!\n"+
			"'intNumericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", intNumericValue))

		return signedPureNumStr, err
	}

	return signedPureNumStr, err
}
