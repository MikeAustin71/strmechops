package examples

import "fmt"

type SourceCodeComments struct {
	input string
}

//	testPrimaryComments
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	 errorPrefix                interface{}
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
//	error
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
func (srcCodeComments *SourceCodeComments) testPrimaryComments() {
	fmt.Println("HELLO WORLD!")
}

//	testSubsidiaryComments
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
//	error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (srcCodeComments *SourceCodeComments) testSubsidiaryComments() {
	fmt.Println("HELLO WORLD!")

}

//	testRoundingComments
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingSpec				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//
//	roundingType				NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
func (srcCodeComments *SourceCodeComments) testRoundingComments() {
	fmt.Println("HELLO WORLD!")

}

//	testNumStrFormatSpecComments
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec				NumStrFormatSpec
//
//		This structure includes all parameters
//		necessary for formatting a number string.
//		These customization options provide maximum
//		granularity in controlling the formatting
//		of the returned Number String.
//
//		type NumStrFormatSpec struct {
//
//			decSeparator			DecimalSeparatorSpec
//
//				Contains the radix point or decimal
//				separator character(s) which will
//				separate integer and fractional
//				numeric digits in a floating point
//				number.
//
//			intSeparatorSpec 		IntegerSeparatorSpec
//
//				Integer Separator Specification. This
//				parameter specifies the type of integer
//				specifies the type of integer grouping and
//				integer separator characters which will be
//				applied to the number string formatting
//				operations.
//
//			negativeNumberSign		NumStrNumberSymbolSpec
//
//				The Number String Negative Number Sign
//				Specification is used to configure negative
//				number sign symbols for negative numeric values
//				formatted and displayed in number stings.
//
//				This specification can also be used to
//				configured currency symbols.
//
//			numberFieldSpec			NumStrNumberFieldSpec
//
//				This Number String Number Field Specification
//				contains the field length and text
//				justification parameter necessary to display
//				a numeric value within a text number field
//				for display as a number string.
//
//			positiveNumberSign		NumStrNumberSymbolSpec
//
//				Positive number signs are commonly implied
//				and not specified. However, the user as the
//				option to specify a positive number sign
//				character or characters for positive numeric
//				values using a Number String Positive Number
//				Sign Specification.
//
//				This specification can also be used to
//				configure currency symbols.
//
//			zeroNumberSign		NumStrNumberSymbolSpec
//
//				The Number String Zero Number Symbol
//				Specification is used to configure
//				number symbols for zero numeric values
//				formatted and displayed in number stings.
//		}
func (srcCodeComments *SourceCodeComments) testNumStrFormatSpecComments() {
	fmt.Println("HELLO WORLD!")

}

// testNumStrNumberFieldSpecComments0
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
func (srcCodeComments *SourceCodeComments) testNumStrNumberFieldSpecComments() {
	fmt.Println("HELLO WORLD!")

}
