package strmech

import "sync"

type NumStrSignedNum struct {
	lock *sync.Mutex
}

// GetNumStr - Returns a formatted number string for the current
// instance of NumStrSignedNum.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel               NumberStrKernel
//	   - An instance of NumberStrKernel containing the numeric
//	     digits which will be used to create and format the
//	     returned number string.
//
//
//	decSeparator               DecimalSeparatorSpec
//	   - An instance of DecimalSeparatorSpec which supplies the
//	     decimal separator characters, or decimal point
//	     characters, which will be used to separate integer and
//	     fractional digits in the returned number string.
//
//
//	intSeparator               IntegerSeparatorDto
//	   - An instance of IntegerSeparatorDto which contains the
//	     text character or characters used to separate integer
//	     digits. This specification is most commonly used as
//	     the thousand's separator. Example: 1,000,000
//
//
//  roundingSpec               NumStrRoundingSpec
//     - This instance of NumStrRoundingSpec specifies the
//       rounding type and number of digits to be used in the
//       in rounding fractional digits configured in the
//       returned number string.
//
//
//  positiveNumberSign         string
//     - Allows the user to specify a number sign character
//       or characters for positive numeric values. The
//       positive number sign will be positioned in the
//       number string as specified by the
//       'positiveNumFormat' parameter.
//
//
//  leadingNegativeNumberSign  string
//     - Allows the user to specify a leading number sign
//       character for negative numeric values. The
//       leading negative number sign will be positioned
//       in the number string as specified by the
//       'negativeNumFormat' parameter.
//
//
//  trailingNegativeNumberSign string
//     - Allows the user to specify a trailing number sign
//       character for negative numeric values. The
//       trailing negative number sign will be positioned
//       in the number string as specified by the
//       'negativeNumFormat' parameter.
//
//
//  positiveNumFormat          string
//     - If the number to be formatted is a positive number, this
//       parameter MUST BE configured using the following place
//       holders
//       ####### - A numeric place holder consisting of 7 "#"
//                 characters.
//
//       ,       - The integer separator placeholder
//
//       +       - The positive number sign placeholder
//
//       Example Formats:
//        - Number formatted with no integer separators
//           #######
//
//        - Standard thousands separators with no positive
//          number sign.
//           #,###,###
//
//        - Number formatted using the India Numbering System.
//           ##,##,###
//
//
//        - Number formatted using the Chinese Numbering System
//           ###,####
//
//        - Number formatted with Positive Number Sign. Note,
//          for positive number values, the positive number sign
//          is typically implied and not displayed (#######).
//           +#,###,###
//           #,###,###+
//           + #,###,###
//           #,###,### +
//           +##,##,###
//           +###,####
//
//
//  negativeNumFormat          string
//     - If the number to be formatted is a negative number, this
//       parameter MUST BE configured using the following
//       placeholders:
//        - A numeric place holder consisting of 7 "#"
//          characters.
//           #######
//
//        - The integer separator placeholder
//           ,
//
//       - The negative number sign placeholders
//           -
//          ()
//
//       Example Formats:
//        - Number formatted with no integer separators
//           #######
//
//        - Standard thousands separators with no positive
//          number sign.
//           #,###,###
//
//        - Number formatted using the India Numbering System.
//           ##,##,###
//
//        - Number formatted using the Chinese Numbering System
//           ###,####
//
//        - Number formatted with Negative Number Sign.
//           -#,###,###
//           #,###,###-
//           - #,###,###
//           #,###,### -
//           -##,##,###
//           -###,####
//           (#,###,###)
//           (#,###,###)
//           ( #,###,### )
//           ( #,###,### )
//           (##,##,###)
//           (###,####)
//
//
//  numberFieldLength          int
//     - The length of the text field in which the numeric value
//       will be displayed. If 'numberFieldLength' is less than
//       the length of the numeric value, it will be automatically
//       set equal to the numeric value string length.
//
//       To automatically set the value of 'numberFieldLength' to
//       the length of the numeric value, set this parameter to a
//       value of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  numberJustification        TextJustify
//     - An enumeration which specifies the justification of the
//       numeric value within the field specified by
//       'numberFieldLength'.
//
//       Text justification can only be evaluated in the context of
//       a number string, number field length and a
//       'numberJustification' object of type TextJustify. This is
//       because number strings with a number field length equal to
//       or less than the length of the number string never use
//       text justification. In these cases, text justification is
//       completely ignored.
//
//       If the number string length is greater than the length of
//       the number string, text justification must be equal to one
//       of these three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
/*func (nStrSignedNum *NumStrSignedNum) GetNumStr(
	numStrKernel NumberStrKernel,
	decSeparator DecimalSeparatorSpec,
	intSeparator IntegerSeparatorDto,
	roundingSpec NumStrRoundingSpec,
	positiveNumberSign string,
	leadingNegativeNumberSign string,
	trailingNegativeNumberSign string,
	positiveNumFormat string,
	negativeNumFormat string,
	numberFieldLength int,
	numberJustification TextJustify,
	errorPrefix interface{}) (
	string,
	error) {

	var err error

	return "", err
}
*/
