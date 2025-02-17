package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	NumStrFormatSpec
//
//	The Number String Format Specification contains
//	parameters used to format signed numbers and currency
//	numeric values as number strings.
//
//	The member variables contained in this structure
//	provide the detail specifications required to
//	support number string multinational and multicultural
//	formatting requirements.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Configuring formatting for Number Stings while
//	supporting multinational and multicultural standards
//	necessarily requires a complex series of parameters
//	and specifications.
//
//	Typically, instances of NumStrFormatSpec are created,
//	or constructed, using the 'New' methods documented
//	below. Many of these methods provide input parameters
//	capable of detailing all the Number Symbol features
//	required to support any multinational and
//	multicultural Number Symbol formatting requirement.
//
//	For those only interested in a quick and simple means
//	of generating Number String formatting configurations,
//	the following methods provide defaults which greatly
//	simplify the Number String Format creation process:
//
//		NumStrFormatSpec.NewCurrencyBasic()
//		NumStrFormatSpec.SetCurrencyBasic()
//		NumStrFormatSpec.NewSignedNumBasic()
//		NumStrFormatSpec.SetSignedNumBasic()
//		NumStrFormatSpec.SetCurrencySimple()
//		NumStrFormatSpec.NewSignedNumSimple()
//		NumStrFormatSpec.SetSignedNumSimple()
//		NumStrFormatSpec.NewCurrencyDefaultsFrance()
//		NumStrFormatSpec.NewCurrencyDefaultsFrance()
//		NumStrFormatSpec.NewSignedNumDefaultsFrance()
//		NumStrFormatSpec.NewCurrencyDefaultsGermany()
//		NumStrFormatSpec.NewSignedNumDefaultsGermany()
//		NumStrFormatSpec.NewCurrencyDefaultsUKMinusOutside()
//		NumStrFormatSpec.NewSignedNumDefaultsUKMinus()
//		NumStrFormatSpec.NewCurrencyDefaultsUSParen()
//		NumStrFormatSpec.SetSignedNumDefaultsUSParen()
//		NumStrFormatSpec.NewSignedNumDefaultsUSMinus()
//		NumStrFormatSpec.SetSignedNumDefaultsUSMinus()
//
//		The 'Basic' and 'Simple' should suffice for the
//		majority of national and cultural number symbol
//		formatting requirements.
//
//	If more granular control is required to meet
//	specialized requirements for multinational or
//	multicultural number string formatting, consider
//	using one of the following methods:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.SetCurrencyParams()
//		NumStrFormatSpec.NewSignedNumParams()
//		NumStrFormatSpec.SetSignedNumParams()
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//
//	These granular methods involve more complexity
//	and a greater number of input parameters, but
//	they do provide greater flexibility and
//	customization in formatting number symbols for
//	number string.
type NumStrFormatSpec struct {
	decSeparator DecimalSeparatorSpec
	//	Contains the decimal separator character
	//	or characters which will separate integer
	//	and fractional digits in a floating point
	//	numbers within a formatted Number String.
	//
	//	The decimal separator is also known as the
	//	radix point.
	//
	//	In the US, UK, Australia and most of
	//	Canada, the decimal separator is the period
	//	character ('.') known as the decimal point.

	intSeparatorSpec IntegerSeparatorSpec
	//	Integer Separator Specification. This
	//	parameter specifies the type of integer
	//	grouping and integer separator characters
	//	which will be applied to the number string
	//	formatting operations.

	numberFieldSpec NumStrNumberFieldSpec
	//	This Number String Number Field Specification
	//	contains the field length and text justification
	//	parameter necessary to display a numeric value
	//	within a number field for display as a number
	//	string.
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
	//				Field Length Examples
	//
	//					Example-1
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 18
	//						fieldJustification = TxtJustify.Center()
	//						Text Field String =
	//							"   5672.1234567   "
	//
	//					Example-2
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 18
	//						fieldJustification = TxtJustify.Center()
	//						Text Field String =
	//							"      5672.1234567"
	//
	//					Example-3
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = -1
	//						fieldJustification = TxtJustify.Center()
	//							// Text Justification Ignored. Field
	//							// Length Equals Title Line String Length
	//						Text Field String =
	//							"5672.1234567"
	//
	//					Example-4
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 2
	//						fieldJustification = TxtJustify.Center()
	//							// Justification Ignored because Field
	//							// Length Less Than Title Line String Length.
	//						Text Field String =
	//							"5672.1234567"
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
	//
	//				Text Justification Examples
	//
	//					Example-1
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 18
	//						fieldJustification = TxtJustify.Center()
	//						Text Field String =
	//							"   5672.1234567   "
	//
	//					Example-2
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 18
	//						fieldJustification = TxtJustify.Center()
	//						Text Field String =
	//							"      5672.1234567"
	//
	//					Example-3
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = -1
	//						fieldJustification = TxtJustify.Center()
	//							// Text Justification Ignored. Field
	//							// Length Equals Title Line String Length
	//						Text Field String =
	//							"5672.1234567"
	//
	//					Example-4
	//  			        Number String = "5672.1234567"
	//						Number String Length = 12
	//						fieldLength = 2
	//						fieldJustification = TxtJustify.Center()
	//							// Justification Ignored because Field
	//							// Length Less Than Title Line String Length.
	//						Text Field String =
	//							"5672.1234567"
	//		}

	numberSymbolsGroup NumStrNumberSymbolGroup
	//	This member variable is used to configure Number
	//	Symbols required in converting numeric values to
	//	Number Strings.
	//
	//	NumStrNumberSymbolGroup contains three instances of
	//	NumStrNumberSymbolSpec defining the Number
	//	Symbols to be used with positive numeric values,
	//	negative numeric values and zero numeric values.
	//
	//	type NumStrNumberSymbolGroup struct {
	//
	//		negativeNumberSign NumStrNumberSymbolSpec
	//
	//			The Number String Negative Number Sign
	//			Specification is used to configure negative
	//			number sign symbols for negative numeric
	//			values formatted and displayed in number
	//			stings.
	//
	//			For currency presentations, the currency
	//			symbol is combined with the negative number
	//			sign.
	//
	//			Example-1: Leading Number Symbols
	//				Leading Number Symbols for Negative Values
	//
	//				Leading Symbols: "- "
	//				Number String:   "- 123.456"
	//
	//			Example-2: Leading Number Symbols With Currency
	//				Leading Number Symbols for Negative Values
	//
	//				Leading Symbols: "$-"
	//				Number String:   "$-123.456"
	//
	//
	//			Example-3: Trailing Number Symbols
	//				Trailing Number Symbols for Negative Values
	//
	//				Trailing Symbols: " -"
	//				Number String:   "123.456 -"
	//
	//			Example-4: Trailing Number Symbols
	//				Trailing Number Symbols for Negative Values
	//
	//				Trailing Symbols: "-€"
	//				Number String:   "123.456-€"
	//
	//		positiveNumberSign NumStrNumberSymbolSpec
	//
	//			Positive number signs are commonly implied
	//			and not specified. However, the user has
	//			the option to specify a positive number sign
	//			character or characters for positive numeric
	//			values using a Number String Positive Number
	//			Sign Specification.
	//
	//			For currency presentations, the currency
	//			symbol is combined with the positive number
	//			sign.
	//
	//			Example-1: Leading Number Symbols
	//				Leading Number Symbols for Positive Values
	//
	//				Leading Symbols: "+ "
	//				Number String:   "+ 123.456"
	//
	//			Example-2: Leading Number Symbols
	//				Leading Number Symbols for Positive Values
	//
	//				Leading Symbols: "$+"
	//				Number String:   "$+123.456"
	//
	//			Example-3: Leading Number Symbols
	//				Leading Number Symbols for Positive Values
	//
	//				Leading Symbols: "$"
	//				Number String:   "$123.456"
	//
	//			Example-4: Trailing Number Symbols
	//				Trailing Number Symbols for Positive Values
	//
	//				Trailing Symbols: " +"
	//				Number String:   "123.456 +"
	//
	//			Example-5: Trailing Number Symbols
	//				Trailing Number Symbols for Positive Values
	//
	//				Trailing Symbols: "+€"
	//				Number String:   "123.456+€"
	//
	//			Example-6: Trailing Number Symbols
	//				Trailing Number Symbols for Positive Values
	//
	//				Trailing Symbols: " €"
	//				Number String:   "123.456 €"
	//
	//		zeroNumberSign NumStrNumberSymbolSpec
	//
	//			The Number String Zero Number Symbol
	//			Specification is used to configure number
	//			symbols for zero numeric values formatted
	//			and displayed in number stings. Zero number
	//			signs are commonly omitted because zero
	//			does not technically qualify as either a
	//			positive or negative value. However,
	//			currency symbols may be required for zero
	//			values.
	//
	//			For currency presentations, the currency
	//			symbol is often used as either a leading
	//			or trailing symbol for zero numeric
	//			values.
	//
	//			Example-1: Leading Number Symbols
	//				Leading Number Symbols for Zero Values
	//
	//				Leading Symbols: "$"
	//				Trailing Symbols: ""
	//				Number String:   "$0.00"
	//
	//			Example-2: Trailing Number Symbols
	//				Trailing Number Symbols for Zero Values
	//
	//				Leading Symbols: ""
	//				Trailing Symbols: " €"
	//				Number String:   "0.00 €"

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrFormatSpec ('incomingSignedNumFmt')
//	to the data fields of the current NumStrFormatSpec
//	instance ('numStrFmtSpec').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values in the current
//	NumStrFormatSpec instance ('numStrFmtSpec') will
//	be deleted and replaced.
//
//	No data validation is performed on input parameter,
//	'incomingSignedNumFmt'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingSignedNumFmt		*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		This method will NOT change the values of internal
//		member variables contained in this instance.
//
//		All data values in this NumStrFormatSpec instance
//		will be copied to current NumStrFormatSpec
//		instance ('signedNumFmtSpec').
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
//		7.  IBasicErrorPrefix
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
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) CopyIn(
	incomingSignedNumFmt *NumStrFormatSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		copySignedNumberFormatSpec(
			numStrFmtSpec,
			incomingSignedNumFmt,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"incomingSignedNumFmt"))
}

//	CopyOut
//
//	Returns a deep copy of the current NumStrFormatSpec instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//		7.  IBasicErrorPrefix
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
// Return Values
//
//	deepCopySignedNumFmtSpec	NumStrFormatSpec
//
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of
//		the current NumStrFormatSpec instance.
//
//
//	err							error
//
//		If the method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error occurs, the text value of input parameter
//		'errorPrefix' will be inserted or prefixed at the
//		beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopySignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopySignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		copySignedNumberFormatSpec(
			&deepCopySignedNumFmtSpec,
			numStrFmtSpec,
			ePrefix.XCpy(
				"deepCopySignedNumFmtSpec<-"+
					"numStrFmtSpec"))

	return deepCopySignedNumFmtSpec, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrFormatSpec to their initial or zero
//	values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of NumStrFormatSpec.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (numStrFmtSpec *NumStrFormatSpec) Empty() {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	new(numStrFmtSpecAtom).empty(
		numStrFmtSpec)

	numStrFmtSpec.lock.Unlock()

	numStrFmtSpec.lock = nil
}

//	Equal
//
//	Receives a pointer to another instance of
//	NumStrFormatSpec and proceeds to compare its internal
//	member variables to those of the current
//	NumStrFormatSpec instance in order to determine if
//	they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingSignedNumFmt		*NumStrFormatSpec
//
//		A pointer to an external instance of
//		NumStrFormatSpec. The internal member variable
//		data values in this instance will be compared to those
//		in the current instance of NumStrFormatSpec. The
//		results of this comparison will be returned to the
//		calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the internal member variable data values contained in
//		input parameter 'incomingSignedNumFmt' are equivalent
//		in all respects to those contained in the current
//		instance of 'NumStrFormatSpec', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (numStrFmtSpec *NumStrFormatSpec) Equal(
	incomingSignedNumFmt *NumStrFormatSpec) bool {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return new(numStrFmtSpecAtom).equal(
		numStrFmtSpec,
		incomingSignedNumFmt)
}

//	GetCurrencySymbolSpec
//
//	Returns the Currency Number Symbol Specification
//	configured for the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a deep
//		copy of the currency number symbol specification
//		configured for the current NumStrFormatSpec
//		instance will be returned.
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
func (numStrFmtSpec *NumStrFormatSpec) GetCurrencySymbolSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetCurrencySymbolSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.numberSymbolsGroup.currencySymbol.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.numberSymbolsGroup" +
				".currencySymbol"))
}

// GetDecSeparatorRunes - Returns an array of runes containing the
// Decimal Separator character or characters configured for the
// current instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]rune
//
//		An array of runes containing the Decimal Separator character
//		or characters configured for the current instance of
//		NumStrFormatSpec.
//
//		If Decimal Separator character(s) have not yet been
//		configured, this method will return 'nil'.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorRunes() []rune {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.decSeparator.GetDecimalSeparatorRunes()
}

// GetDecSeparatorSpec - Returns a deep copy of the Decimal
// Separator Specification configured for the current instance
// of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	DecimalSeparatorSpec
//
//		If this method completes successfully, a deep copy of
//		the Decimal Separator Specification configured for the
//		current instance of NumStrFormatSpec will be
//		returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorSpec(
	errorPrefix interface{}) (
	DecimalSeparatorSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetDecSeparatorSpec()",
		"")

	if err != nil {
		return DecimalSeparatorSpec{}, err
	}

	return numStrFmtSpec.decSeparator.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.decSeparator"))
}

// GetDecSeparatorStr - Returns a string containing the Decimal
// Separator character or characters configured for the current
// instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		A string containing the Decimal Separator character or
//	 	characters configured for the current instance of
//	 	NumStrFormatSpec.
//
//		If Decimal Separator character(s) have not yet been
//		configured, this method will return an empty string.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorStr() string {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.decSeparator.GetDecimalSeparatorStr()
}

// GetIntSeparatorSpec - Returns a deep copy of the Integer
// Grouping Specification configured for the current instance
// of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a deep copy of
//		the Integer Separator Specification configured for the
//		current instance of NumStrFormatSpec will be
//		returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorSpec(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetIntSeparatorSpec()",
		"")

	if err != nil {
		return IntegerSeparatorSpec{}, err
	}

	return numStrFmtSpec.intSeparatorSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.intSeparatorSpec"))
}

// GetIntSeparatorChars - Returns a string containing the Integer
// Separator character or characters configured for the current
// instance of NumStrFormatSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (",").
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period (".") is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method will return a string containing the Integer
//		Separator character or characters configured for the
//	 	current instance of NumStrFormatSpec.
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorChars() string {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	str,
		_ := numStrFmtSpec.intSeparatorSpec.GetIntSeparatorStr(
		nil)

	return str
}

// GetIntegerSeparatorSpec - Returns an instance of
// IntegerSeparatorSpec based on the configuration parameters
// contained within the current instance of
// NumStrFormatSpec.
//
// IntegerSeparatorSpec is used by low level number string
// formatting functions to complete the number string generation
// operation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a copy
//		of the integer grouping specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetIntegerSeparatorSpec(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newIntSeparatorSpec IntegerSeparatorSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetIntegerSeparatorSpec()",
		"")

	if err != nil {
		return newIntSeparatorSpec, err
	}

	newIntSeparatorSpec,
		err =
		numStrFmtSpec.intSeparatorSpec.CopyOut(
			ePrefix.XCpy(
				"<-numStrFmtSpec.intSeparatorSpec"))

	return newIntSeparatorSpec, err
}

// GetIntSeparatorRunes - Returns a rune array containing the
// Integer Separator character or characters configured for the
// current instance of NumStrFormatSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (',').
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period ('.') is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]rune
//
//		This method will return a rune array containing the
//		Integer Separator character or characters configured
//		for the	current instance of NumStrFormatSpec.
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorRunes() []rune {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	intSepRunes,
		_ := numStrFmtSpec.intSeparatorSpec.GetIntSeparatorChars(
		nil)

	return intSepRunes
}

// GetNegativeNumSymSpec - Returns the Negative Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the negative number sign specification
//		configured for the current NumStrFormatSpec
//		instance will be returned.
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
func (numStrFmtSpec *NumStrFormatSpec) GetNegativeNumSymSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetNegativeNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.numberSymbolsGroup.negativeNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.numberSymbols" +
				".negativeNumberSign"))
}

// GetNumberFieldSpec - Returns the Number Field Specification
// currently configured for this instance of
// NumStrFormatSpec.
//
// The Number Field Specification includes parameters for
// field length and text justification ('Right, Center,
// Left').
//
// Numeric digits are formatted using the Number Field
// Specification within a number field specified by field
// length. The numeric digits string is then justified
// within the number field according to a text justification
// specification of 'Right', 'Center' or 'Left'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	NumStrNumberFieldSpec
//
//		If this method completes successfully, a copy
//		of the Number Field Specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
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
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetNumberFieldSpec(
	errorPrefix interface{}) (
	NumStrNumberFieldSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetNumberFieldSpec()",
		"")

	if err != nil {
		return NumStrNumberFieldSpec{}, err
	}

	return numStrFmtSpec.numberFieldSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.numberFieldSpec"))
}

//	GetNumberSymbolsGroup
//
//	Returns a deep copy of the NumStrNumberSymbolGroup
//	configured as a member data element in the current
//	instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
func (numStrFmtSpec *NumStrFormatSpec) GetNumberSymbolsGroup(
	errorPrefix interface{}) (
	NumStrNumberSymbolGroup,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetNumberSymbolsGroup()",
		"")

	if err != nil {
		return NumStrNumberSymbolGroup{}, err
	}

	return numStrFmtSpec.numberSymbolsGroup.CopyOut(
		ePrefix.XCpy("<-numStrFmtSpec." +
			"numberSymbolsGroup"))
}

// GetPositiveNumSymSpec - Returns the Positive Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the positive number sign specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetPositiveNumSymSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetPositiveNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.numberSymbolsGroup.positiveNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.numberSymbolsGroup.positiveNumberSign"))
}

// GetZeroNumSymSpec - Returns the Zero Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the zero number sign specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetZeroNumSymSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetZeroNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.numberSymbolsGroup.zeroNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.zeroNumberSign"))
}

// IsNOP
//
//	Stands for 'Is No Operation'. This method returns
//	a boolean value signaling whether the current
//	instance of NumStrFormatSpec is engaged, valid, fully
//	configured and capable of performing number string
//	formatting operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//		If this method returns 'true', it signals
//		that the current Number String Format
//		Specification ('NumStrFormatSpec') is simply an
//		empty placeholder and is NOT capable of
//		performing number string formatting operations.
//
//		If this method returns 'false', it signals that
//		the current 'NumStrFormatSpec' instance is fully
//		populated, valid, functional and ready to perform
//		number string formatting operations.
func (numStrFmtSpec *NumStrFormatSpec) IsNOP() bool {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return new(numStrFmtSpecNanobot).isNOP(
		numStrFmtSpec)
}

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current NumStrFormatSpec
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return a boolean value of 'false'.
//
//	If all data elements are determined to be valid,
//	this method returns a boolean value of 'true'.
//
//	This method is functionally equivalent to
//	NumStrFormatSpec.IsValidInstanceError() with
//	the sole exceptions being that this method takes
//	no input parameters and returns a boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal data values contained in
//		the current instance of NumStrFormatSpec are
//		found to be invalid, this method will return a
//		boolean value of 'false'.
//
//		If all internal member data values contained in
//		the current instance of NumStrFormatSpec are
//		found to be valid, this method returns a boolean
//		value of 'true'.
func (numStrFmtSpec *NumStrFormatSpec) IsValidInstance() bool {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	isValid,
		_ := new(numStrFmtSpecElectron).
		testValidityNumStrFormatSpec(
			numStrFmtSpec,
			nil)

	return isValid
}

// IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current NumStrFormatSpec
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//		If any of the internal member data variables
//		contained in the current instance of
//		NumStrFormatSpec are found to be invalid, this
//		method will return an error with an appropriate
//		message identifying the invalid member data
//		variable.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (numStrFmtSpec *NumStrFormatSpec) IsValidInstanceError(
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(numStrFmtSpecElectron).
		testValidityNumStrFormatSpec(
			numStrFmtSpec,
			ePrefix.XCpy(
				"numStrFmtSpec"))

	return err
}

//	NewCountryCurrencyNumFormat
//
//	Creates and returns a new, fully populated instance
//	of NumStrFormatSpec based on a Number String Country
//	Culture Specification passed as an input paramter.
//
//	This method will produce a new NumStrFormatSpec
//	configured for Currency Numeric Values according to
//	the designated country or culture specified by input
//	parameter 'countryCultureFormat'.
//
//	For signed number formats see method:
//
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		currency formatting information for the
//		designated country or culture.
//
//		This method will NOT change the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameter 'numberFieldSpec'
//		to construct and return a new instance of
//		NumStrFormatSpec configured for Currency Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		The data values in 'numberFieldSpec' will be
//		combined with input parameter
//		'countryCultureFormat' to construct and return
//		a new instance of NumStrFormatSpec.
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
//	NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new instance of
//		NumStrFormatSpec configured as a Currency
//		Number String Formatting Specification.
//
//		This configuration is based on input paramter
//		'countryCultureFormat' which provides Currency
//		Number String Formatting Specifications for
//		the designated country or culture.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCountryCurrencyNumFormat(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	NumStrFormatSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumStrFmtSpec NumStrFormatSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCountryCurrencyNumFormat()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	if countryCultureFormat.CurrencyNumStrFormat.
		numberSymbolsGroup.IsNOPCurrencySymbols() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'countryCultureFormat.CurrencyNumStrFormat'\n"+
			"is invalid. The Currency Symbols Specification is empty!\n",
			ePrefix.String())

		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		&newNumStrFmtSpec,
		countryCultureFormat.CurrencyNumStrFormat.decSeparator,
		countryCultureFormat.CurrencyNumStrFormat.intSeparatorSpec,
		countryCultureFormat.CurrencyNumStrFormat.numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("newCurrencyNumFmtSpec<-"))

	return newNumStrFmtSpec, err
}

//	NewCountrySignedNumFormat
//
//	Creates and returns a new, fully populated instance
//	of NumStrFormatSpec based on a Number String Country
//	Culture Specification passed as an input paramter.
//
//	This method will produce a new NumStrFormatSpec
//	configured for Signed Numeric Values according to
//	the designated country or culture specified by input
//	parameter 'countryCultureFormat'.
//
//	For currency formats see method:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		currency formatting information for a
//		designated country or culture.
//
//		This method will NOT change the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameter 'numberFieldSpec'
//		to construct and return a new instance of
//		NumStrFormatSpec configured for Signed Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		The data values in 'numberFieldSpec' will be
//		combined with input parameter
//		'countryCultureFormat' to construct and return
//		a new instance of NumStrFormatSpec.
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
//	NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new instance of
//		NumStrFormatSpec configured as a Signed
//		Number String Formatting Specification.
//
//		This configuration is based on input paramter
//		'countryCultureFormat' which provides Signed
//		Number String Formatting Specifications for
//		the designated country or culture.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCountrySignedNumFormat(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	NumStrFormatSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumStrFmtSpec NumStrFormatSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCountrySignedNumFormat()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		&newNumStrFmtSpec,
		countryCultureFormat.SignedNumStrFormat.decSeparator,
		countryCultureFormat.SignedNumStrFormat.intSeparatorSpec,
		countryCultureFormat.SignedNumStrFormat.numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newNumStrFmtSpec, err
}

//	NewCurrencyBasic
//
//	Creates and returns a new instance of NumStrFormatSpec
//	configured using a 'basic' set of Number String Format
//	Specification parameters.
//
//	The returned NumStrFormatSpec instance will include
//	format specifications for currency, positive numeric
//	values, zero numeric values and negative numeric
//	values based on the input parameters described below.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingCurrencySymbol			string
//
//		The character or characters which comprise the
//		leading Currency Symbol. The leading Currency
//		Symbol will be positioned at the beginning or
//		left side of the number string.
//
//			Example: $ 123.45
//
//		If a space between the currency symbol
//		and the first digit of the number string
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Leading Currency Symbol: "$ "
//				Formatted Number String: "$ 123.45"
//
//	trailingCurrencySymbol			string
//
//		The character or characters which comprise the
//		trailing Currency Symbol. The trailing Currency
//		Symbol will be positioned at the end of, or
//		right side of, the number string.
//
//			Example: 123.45 €
//
//		If a space between the last digit of the
//		number string and the currency symbol
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Trailing Currency Symbol: " €"
//				Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This parameter determines whether the currency
//		symbol will be positioned inside or outside
//		the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency and number
//		sign characters relative to a Number Field in
//		which a number string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec configured with
//		currency symbols.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyBasic(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyInsideNumSymbol bool,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyBasic()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setCurrencyBasic(
			&newNumStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingCurrencySymbol),
			[]rune(trailingCurrencySymbol),
			currencyInsideNumSymbol,
			[]rune(leadingNegativeNumSign),
			[]rune(trailingNegativeNumSign),
			numSymbolFieldPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec"))

	return newNumStrFmtSpec, err
}

// NewCurrencyBasicRunes
//
// Creates and returns a new instance of NumStrFormatSpec
// configured using a 'basic' set of Number String Format
// Specification parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol for the returned
//		instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		This rune array contains one or more text
//		characters used to separate groups of integers.
//		This separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//				'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingCurrencySymbol			[]rune
//
//		This rune array contains the character or
//		characters which comprise the leading Currency
//		Symbol. The leading Currency Symbol will be
//		positioned at the beginning or left side of the
//		number string.
//
//			Example: $ 123.45
//
//		If a space between the currency symbol
//		and the first digit of the number string
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Leading Currency Symbol: "$ "
//				Formatted Number String: "$ 123.45"
//
//	trailingCurrencySymbol			[]rune
//
//		This rune array contains the character or
//		characters which comprise the trailing Currency
//		Symbol. The trailing Currency Symbol will be
//		positioned at the end of, or right side of, the
//		number string.
//
//			Example: 123.45 €
//
//		If a space between the last digit of the
//		number string and the currency symbol
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Trailing Currency Symbol: " €"
//				Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This parameter determines whether the currency
//		symbol will be positioned inside or outside
//		the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency and number
//		sign characters relative to a Number Field in
//		which a number string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec configured with
//		basic currency and number sign symbol default
//		values.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyBasicRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingCurrencySymbol []rune,
	trailingCurrencySymbol []rune,
	currencyInsideNumSymbol bool,
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyBasicRunes()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setCurrencyBasic(
			&newNumStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingCurrencySymbol,
			trailingCurrencySymbol,
			currencyInsideNumSymbol,
			leadingNegativeNumSign,
			trailingNegativeNumSign,
			numSymbolFieldPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec"))

	return newNumStrFmtSpec, err
}

//	NewCurrencyDefaultsFrance
//
//	Returns a new instance of NumStrFormatSpec
//	configured with Currency Number String formatting
//	conventions typically applied in France.
//
//	The default French currency symbol is a trailing
//	Euro sign ('€').
//
//		French Example:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with French Currency
//	Number formatting specifications. New data values
//	will be configured for the positive, zero and
//	negative number sign symbols as well as the currency
//	symbol.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Currency Number
//	String formatting standards implemented by either
//	France or Germany.
//
//	For information on German Currency Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.NewCurrencyDefaultsGermany()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1:
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2:
//		1 000 000 000
//
//	Currency Symbols
//
//	The currency symbol used in the France is the
//	trailing Euro Currency Symbol ('€').
//
//		French Example-3:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied. No positive
//	number is applied, only the trailing Euro Currency
//	Symbol.
//
//		French Example-4:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Zero Numeric Values
//
//	The zero number format has no number sign, but the
//	currency symbol is set to a trailing Euro Currency
//	Symbol.
//
//		French Example-5:
//			Zero Numeric Currency Value
//				0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro Currency Symbol
//	("€").
//
//		French Example-6:
//			Negative Numeric Currency Value
//				-1 000 000,00 €
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		French Example-7:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123,45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " -123,45 €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec					NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
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
//
//	 errorPrefix                	interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newFrenchCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured for
//		Currency Number String formatting parameters
//		typically applied in France.
//
//	err								error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newFrenchCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsFrance()",
		"")

	if err != nil {
		return newFrenchCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyDefaultsFrance(
		&newFrenchCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newFrenchCurrencyNumFmtSpec<-"))

	return newFrenchCurrencyNumFmtSpec, err
}

//	NewCurrencyDefaultsGermany
//
//	Returns a new instance of NumStrFormatSpec
//	configured with Currency Number String
//	formatting conventions typically applied
//	in Germany.
//
//	The default German currency symbol is a trailing
//	Euro sign ('€').
//
//		German Example:
//			Positive Numeric Currency Value
//				1.000.000,00 €
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with German Currency
//	Number formatting specifications. New data values
//	will be configured for the positive, zero and
//	negative number sign symbols as well as the currency
//	symbol.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Currency Number
//	String formatting standards implemented by either
//	Germany or France.
//
//	For information on French Currency Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.NewCurrencyDefaultsFrance()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	Currency Symbols
//
//	The currency symbol used in the Germany is the
//	trailing Euro symbol ('€').
//
//		German Example-3:
//		1.000.000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-4:
//		1.000.000 €
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-5:
//			0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-6:
//		1.000.000- €
//
//	The negative signed number is configured with a
//	trailing minus sign ('-') appended at the end of
//	the number string. The negative number sign and the
//	currency symbol will be positioned inside the number
//	field:
//
//		German Example-5:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: 123,45-
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " 123,45- €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec					NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
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
//
//	 errorPrefix                	interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newGermanCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in Germany.
//
//	err								error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newGermanCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsGermany()",
		"")

	if err != nil {
		return newGermanCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyDefaultsGermany(
		&newGermanCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newGermanCurrencyNumFmtSpec<-"))

	return newGermanCurrencyNumFmtSpec, err
}

//	NewCurrencyDefaultsUKMinusInside
//
//	Returns a new instance of NumStrFormatSpec configured
//	for the United Kingdom (UK) Currency Number String
//	formatting conventions.
//
//	The default UK currency symbol is a leading Pound
//	sign ('£').
//
//		UK Example:
//			Positive Numeric Currency Value
//				£ 123.45  Positive Value
//
//	The term "MinusInside" in the method name means that
//	the minus sign ('-') configured for negative numeric
//	values will be inside, or to the right of, the Pound
//	sign ('£').
//
//		UK Example:
//			Negative Numeric Currency Value
//				£ -123.45  Negative Value
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with UK Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		UK Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		UK Example-2:
//			1,000,000
//
//	Currency Symbol
//
//	The default currency symbol used in the UK is the
//	leading Pound symbol ('£').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-3:
//			Positive Numeric Currency Value
//				£ 123.45
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-4:
//			Zero Numeric Currency Value
//				£ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the Pound sign ('£')
//	such that any minus sign configured for negative
//	numeric values will be inside, or to the right of,
//	the Pound sign ('£').
//
//		UK Example-5:
//			Negative Numeric Currency Value
//				£ -123.45  Negative Value
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		UK Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " £ -123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUKCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United Kingdom.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsUKMinusInside(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUKCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return newUKCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyDefaultsUKMinusInside(
		&newUKCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newUKCurrencyNumFmtSpec<-"))

	return newUKCurrencyNumFmtSpec, err
}

//	NewCurrencyDefaultsUKMinusOutside
//
//	Returns a new instance of NumStrFormatSpec configured
//	for the United Kingdom (UK) Currency Number String
//	formatting conventions.
//
//	The default UK currency symbol is a leading Pound
//	sign ('£').
//
//		UK Example:	Positive Numeric Currency Value
//			£ 123.45
//
//	The term "MinusOutside" in the method name means that
//	the minus sign ('-') configured for negative numeric
//	values will be outside, or to the left of, the Pound
//	sign ('£').
//
//	UK Example:	Negative Numeric Currency Value
//			- £123.45
//
//	Default values will be used to configure the returned
//	instance of NumStrFormatSpec with UK Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		UK Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		UK Example-2:
//			1,000,000
//
//	Currency Symbol
//
//	The default currency symbol used in the UK is the
//	leading Pound symbol ('£').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-3:
//			Positive Numeric Currency Value
//				£ 123.45  Positive Value
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-4:
//			Zero Numeric Currency Value
//				£ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the minus sign ('-')
//	such that any minus sign configured for negative
//	numeric values will be outside, or to the left of,
//	the Pound sign ('£').
//
//		UK Example-5:
//			Negative Numeric Currency Value
//				- £123.45
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		UK Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " - £123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUKCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United Kingdom.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsUKMinusOutside(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUKCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return newUKCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyDefaultsMinusOutside(
		&newUKCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newUKCurrencyNumFmtSpec<-"))

	return newUKCurrencyNumFmtSpec, err
}

//	NewCurrencyDefaultsUSMinus
//
//	Returns a new instance of NumStrFormatSpec configured
//	for the United States (US) Currency Number String
//	formatting conventions.
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example
//			Negative Numeric Currency Value
//				$ -123
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with US Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		US Example-2:
//			1,000,000
//
//	Currency Symbols
//
//	The default currency symbol used in the US is the
//	leading Dollar symbol ('$').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-3:
//			Positive Numeric Currency Value
//				$ 123.45
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-4:
//			Zero Numeric Currency Value
//				$ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the Dollar sign ('$')
//	such that any minus sign configured for negative
//	numeric values will be inside, or to the right of,
//	the Dollar sign ('$').
//
//		US Example-5:
//			Negative Numeric Currency Value
//				$ -123.45
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		US Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " $ -123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUSCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United States.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsUSMinus(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUSCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsUSMinus()",
		"")

	if err != nil {
		return newUSCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setCurrencyDefaultsUSMinus(
			&newUSCurrencyNumFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy(
				"newUSCurrencyNumFmtSpec<-"))

	return newUSCurrencyNumFmtSpec, err
}

//	NewCurrencyDefaultsUSParen
//
//	Returns a new instance of NumStrFormatSpec configured
//	for United States (US) Currency Number String
//	formatting conventions.
//
//	The default US currency symbol is a leading Dollar
//	sign ('$').
//
//		US Example
//			Positive Numeric Currency Value
//				$ 123.45
//
//	The term 'Paren' in the method name signals that a
//	surrounding parentheses ('()') will be used to
//	designate negative numeric values.
//
//		US Example
//			Negative Numeric Currency Value
//				$ (123)
//
//	Default values will be used to configure the returned
//	instance of NumStrFormatSpec with US Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separators
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Separators
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		US Example-2
//			1,000,000
//
//	Currency Symbols
//
//	The default currency symbol used in the US is the
//	leading Dollar symbol ('$').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-3:
//			Positive Numeric Currency Value
//				$ 123.45
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-4:
//			Zero Numeric Currency Value
//				$ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to surrounding
//	parentheses ('()').
//
//	This method will configure the Dollar sign ('$')
//	such that the leading parenthesis ('(') configured
//	for negative numeric values will be inside, or to the
//	right of, the Dollar sign ('$').
//
//		US Example-5:
//			Negative Numeric Currency Value
//				$ (123.45)
//
//	The negative signed number is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be prefixed with a
//	leading parenthesis symbol ('(') and suffixed with a
//	trailing, or closing, parenthesis symbol (')'). The
//	negative number sign symbols and the currency symbol
//	will be positioned inside the number field:
//
//		US Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 11
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " $ (123.45)"
//				Number Field Index:------>01234567890
//				Total Number String Length: 11
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUSCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United States.
//
//		The default US currency symbol is a leading
//		Dollar sign ('$').
//
//		This method will configure the Dollar sign ('$')
//		such that the leading parenthesis sign ('(')
//		configured for negative numeric values will be
//		inside, or to the right of, the Dollar sign ('$').
//
//			US Example:
//				Negative Numeric Currency Value
//					$ (123.45)
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyDefaultsUSParen(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUSCurrencyNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyDefaultsUSParen()",
		"")

	if err != nil {
		return newUSCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyDefaultsUSParen(
		&newUSCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newUSCurrencyNumFmtSpec<-"))

	return newUSCurrencyNumFmtSpec, err
}

//	NewCurrencyParams
//
//	Creates and returns a new instance of
//	NumStrFormatSpec based on the Positive, Negative,
//	Zero and Currency symbol specifications passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	trailingPosNumSign				string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
//
//	trailingZeroNumSign				string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to an empty string.
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to an empty string.
//
//	currencyInsideNumSymbol			bool
//
//		This boolean parameter determines whether the
//		currency symbol will be positioned inside or
//		outside the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right Justified
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right Justified
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix                interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumFmtSpec				NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number Symbols.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyParams(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	leadingCurrencySymbols string,
	trailingCurrencySymbols string,
	currencyInsideNumSymbol bool,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyParams()",
		"")

	if err != nil {
		return newNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setCurrencyParams(
			&newNumFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			[]rune(leadingCurrencySymbols),
			[]rune(trailingCurrencySymbols),
			currencyInsideNumSymbol,
			currencyNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumFmtSpec<-"))

	return newNumFmtSpec, err
}

//	NewCurrencyParamsRunes
//
//	Creates and returns a new instance of
//	NumStrFormatSpec based on the Positive, Negative,
//	Zero and Currency symbol specifications passed as
//	rune array input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPosNumSign				[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the Dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//				Example: $125.34
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to a 'nil' or empty
//		rune array.
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to a 'nil' or empty
//		rune array.
//
//	currencyInsideNumSymbol			bool
//
//		This boolean parameter determines whether the
//		currency symbol will be positioned inside or
//		outside the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') and Trailing
//		Currency Symbol ('trailingCurrencySymbol')
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right Justified
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right Justified
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'numFieldJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number Symbols.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyParamsRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	leadingCurrencySymbols []rune,
	trailingCurrencySymbols []rune,
	currencyInsideNumSymbol bool,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencyParamsRunes()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setCurrencyParams(
			&newSignedNumFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			leadingCurrencySymbols,
			trailingCurrencySymbols,
			currencyInsideNumSymbol,
			currencyNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewCurrencySimple
//
//	Creates and returns and instance of NumStrFormatSpec
//	configured for Currency Number String formatting.
//
//	If currency number symbol formatting IS NOT
//	required, see method:
//
//		NumStrFormatSpec.NewSignedNumSimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Simple Currency Defaults
//
//	Integer Grouping
//		Integers are grouped by thousands or groups
//		of three integers.
//
//		Example: 1,000,000,000
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values are assumed.
//
//		Positive Numeric Value Currency Examples:
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//		Zero Numeric Value Currency Examples:
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency.
//		This currency formatting will be configured in
//		the new instance of NumStrFormatSpec returned by
//		this method.
//
//	leadingCurrencySymbols		bool
//
//		Controls the positioning of Currency Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Currency
//		Symbols at the beginning or left side of the
//		number string. Such Currency Symbols are therefore
//		configured as leading Currency Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number Strings:
//			"$ 123.456"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the first digit or
//				leading minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure Currency Symbols on the right side of
//		the number string. Currency Number Symbols are
//		therefore configured as trailing Number Symbols.
//		This is the positioning format used in France,
//		Germany and many other countries in the European
//		Union.
//
//			Example Number Strings:
//				"123.456 €"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the last digit or
//				trailing minus sign.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencySimple(
	decSeparatorChars string,
	intSeparatorChars string,
	currencySymbols string,
	leadingCurrencySymbols bool,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencySimple()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setCurrencySimple(
			&newNumStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			[]rune(currencySymbols),
			leadingCurrencySymbols,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

	return newNumStrFmtSpec, err
}

//	NewCurrencySimpleRunes
//
//	Creates and returns and instance of NumStrFormatSpec
//	configured for Currency Number String formatting.
//
//	If currency number symbol formatting IS NOT
//	required, see method:
//
//		NumStrFormatSpec.
//			NewSignedNumSimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Simple Currency Defaults
//
//	Integer Grouping
//		Integers are grouped by thousands or groups
//		of three integers.
//
//		Example: 1,000,000,000
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values are assumed.
//
//		Positive Numeric Value Currency Examples:
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//		Zero Numeric Value Currency Examples:
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			[]rune
//
//		This rune array contains one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the United
//		States, the standard integer digits separator is
//		the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	currencySymbols				[]rune
//
//		This rune array contains the symbol or symbols
//		used to format currency. This currency formatting
//		will be configured in the new instance of
//		NumStrFormatSpec returned by this method.
//
//	leadingCurrencySymbols		bool
//
//		Controls the positioning of Currency Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Currency
//		Symbols at the beginning or left side of the
//		number string. Such Currency Symbols are therefore
//		configured as leading Currency Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number Strings:
//			"$ 123.456"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the first digit or
//				leading minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure Currency Symbols on the right side of
//		the number string. Currency Number Symbols are
//		therefore configured as trailing Number Symbols.
//		This is the positioning format used in France,
//		Germany and many other countries in the European
//		Union.
//
//			Example Number Strings:
//				"123.456 €"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the last digit or
//				trailing minus sign.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencySimpleRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	currencySymbols []rune,
	leadingCurrencySymbols bool,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewCurrencySimpleRunes()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setCurrencySimple(
			&newNumStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			currencySymbols,
			leadingCurrencySymbols,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

	return newNumStrFmtSpec, err
}

//	NewNumFmtComponents
//
//	Creates and returns a new instance of NumStrFormatSpec
//	generated from Number String formatting input
//	components passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	numberSymbolsGroup  				NumStrNumberSymbolGroup
//
//		This instance of NumStrNumberSymbolGroup contains the
//		Number Symbol Specifications for negative numeric
//		values, positive numeric values and zero numeric
//		values.
//
//		type NumStrNumberSymbolGroup struct {
//
//			negativeNumberSign NumStrNumberSymbolSpec
//
//				The Number String Negative Number Sign
//				Specification is used to configure negative
//				number sign symbols for negative numeric
//				values formatted and displayed in number
//				stings.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Negative
//					Values
//
//					Leading Symbols: "- "
//					Number String:   "- 123.456"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Negative
//					Values
//
//					Leading Symbols: "-"
//					Number String:   "-123.456"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Negative
//					Values
//
//					Trailing Symbols: " -"
//					Number String:   "123.456 -"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Negative
//					Values
//
//					Trailing Symbols: "-"
//					Number String:   "123.456-"
//
//			positiveNumberSign NumStrNumberSymbolSpec
//
//				Positive number signs are commonly implied
//				and not specified. However, the user has
//				the option to specify a positive number sign
//				character or characters for positive numeric
//				values using a Number String Positive Number
//				Sign Specification.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Positive
//					Values
//
//					Leading Symbols: "+ "
//					Number String:   "+ 123.456"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Positive
//					Values
//
//					Leading Symbols: "+"
//					Number String:   "+123.456"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Positive
//					Values
//
//					Trailing Symbols: " +"
//					Number String:   "123.456 +"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Positive
//					Values
//
//					Trailing Symbols: "+"
//					Number String:   "123.456+"
//
//			zeroNumberSign NumStrNumberSymbolSpec
//
//				The Number String Zero Number Sign
//				Specification is used to configure number
//				sign symbols for zero numeric values formatted
//				and displayed in number stings. Zero number
//				signs are commonly omitted because zero
//				does not technically qualify as either a
//				positive or negative value. However,
//				the user has the option to configure number
//				sign symbols for zero values if necessary.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Zero Values
//
//					Leading Symbols: "+"
//					Trailing Symbols: ""
//					Number String:   "+0.00"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Zero Values
//
//					Leading Symbols: "+ "
//					Trailing Symbols: ""
//					Number String:   "+ 0.00"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Zero Values
//
//					Leading Symbols: ""
//					Trailing Symbols: " +"
//					Number String:   "0.00 +"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Zero Values
//
//					Leading Symbols: ""
//					Trailing Symbols: "+"
//					Number String:   "0.00+"
//
//			currencySymbol NumStrNumberSymbolSpec
//
//				A Currency Symbol positioned next to a numeric
//				value designates the number as a monetary amount.
//
//				Examples of Currency Symbols include the Dollar
//				sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//				This instance of NumStrNumberSymbolSpec is used
//				to configure leading Currency Symbols, trailing
//				Currency Symbols or both leading and trailing
//				Currency Symbols.
//
//				Example-1: Leading Currency Symbols
//
//					Leading Currency Symbols: "$ "
//					Number String:   "$ 123.456"
//
//				Example-2: Leading Currency Symbols
//
//					Leading Currency Symbols: "$"
//					Number String:   "$123.456"
//
//				Example-3: Trailing Currency Symbols
//					Trailing Currency Symbols for Positive Values
//
//					Trailing Currency Symbols: "€"
//					Number String:   "123.456€"
//
//				Example-4: Trailing Currency Symbols
//					Trailing Currency Symbols for Positive Values
//
//					Trailing Currency Symbols: " €"
//					Number String:   "123.456 €"
//		}
//
//	numberFieldSpec			NumStrNumberFieldSpec
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
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewNumFmtComponents(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	numberSymbolsGroup NumStrNumberSymbolGroup,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewNumFmtComponents()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		&newSignedNumFmtSpec,
		decSeparatorSpec,
		intSeparatorSpec,
		numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewNumFmtElements
//
//	Creates and returns a new instance of NumStrFormatSpec
//	generated from Number String formatting data
//	specifications passed as input parameters.
//
//	Options include customizing for currency symbols,
//	integer separation, number sign	management, radix
//	point symbols, and floating point number rounding.
//
//	If required, users have the options of
//	implementing the India or Chinese Numbering
//	Systems for integer separation.
//
//	This method offers the maximum degree of granular
//	control over all aspects of the Number String
//	formatting operation.
//
//	In particular, it offers maximum flexibility in
//	configuring integer separator characters and
//	integer grouping sequences.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If Currency Symbol formatting is NOT required,
//	set the Currency Symbol parameter to a 'NOP'.
//	'NOP' is a computer science term for
//	'No Operation'.
//
//		currencySymbol :=
//			new(NumStrNumberSymbolSpec).NewNOP()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	negativeNumberSign				NumStrNumberSymbolSpec
//
//		The Number String Negative Number Sign
//		Specification is used to configure negative
//		number sign symbols for negative numeric
//		values formatted and displayed in number
//		stings.
//
//		If this parameter is submitted as an empty or
//		invalid Negative Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	positiveNumberSign 				NumStrNumberSymbolSpec
//
//		Positive number signs are commonly implied
//		and not specified. However, the user has
//		the option to specify a positive number sign
//		character or characters for positive numeric
//		values using this input parameter.
//
//		If this parameter is submitted as an empty or
//		invalid Positive Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//	zeroNumberSign					NumStrNumberSymbolSpec
//
//		The Number String Zero Number Sign
//		Specification is used to configure number
//		sign symbols for zero numeric values formatted
//		and displayed in number stings. Zero number signs
//		are commonly omitted because zero does not
//		technically qualify as either a positive or
//		negative value. However, the user has the option
//		to configure number sign symbols for zero values
//		if necessary.
//
//		If this parameter is submitted as an empty or
//		invalid Zero Number Sign Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+ "
//			Trailing Symbols: ""
//			Number String:   "+ 0.00"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: "+"
//			Number String:   "0.00+"
//
//	currencySymbol					NumStrNumberSymbolSpec
//
//		A Currency Symbol positioned next to a numeric
//		value designates the number is a monetary amount.
//
//		The Number String Currency Symbol Specification
//		is used to configure currency symbols for
//		positive, negative and zero numeric values
//		formatted and displayed in number stings.
//
//		If Currency Symbol formatting is NOT required,
//		set this parameter to a 'NOP'. 'NOP' is a
//		computer science term for 'No Operation'.
//
//			currencySymbol :=
//				new(NumStrNumberSymbolSpec).NewNOP()
//
//		If this parameter is submitted as an empty or
//		invalid Currency Symbol Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Examples of Currency Symbols include the Dollar
//		sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//		This instance of NumStrNumberSymbolSpec is used
//		to configure leading Currency Symbols, trailing
//		Currency Symbols or both leading and trailing
//		Currency Symbols.
//
//		Example-1: Leading Currency Symbols
//
//			Leading Currency Symbols: "$ "
//			Number String:   "$ 123.456"
//
//		Example-2: Leading Currency Symbols
//
//			Leading Currency Symbols: "$"
//			Number String:   "$123.456"
//
//		Example-3: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: "€"
//			Number String:   "123.456€"
//
//		Example-4: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: " €"
//			Number String:   "123.456 €"
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
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewNumFmtElements(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbol NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewNumFmtElements()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtElements(
		&newSignedNumFmtSpec,
		decSeparatorSpec,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		currencySymbol,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumBasic
//
//	Returns a new instance of NumStrFormatSpec configured
//	with basic Signed Number formatting specification
//	default values.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	This method is one step above NewSignedNumSimple()
//	in terms of complexity and customization options.
//
//	This method provides additional options for multiple
//	integer grouping types including India and Chinese
//	Numbering.
//
//	Negative number sign symbols may be configured as
//	leading negative number signs, trailing negative
//	number signs or both. The combination of leading
//	and trailing negative number signs allows for the
//	configuration symbols like parentheses for the
//	formatting of negative numbers.
//
//		Example: (125.34)
//
// ----------------------------------------------------------------
//
// # Signed Number Basic Defaults
//
//	No Currency Symbols Included:
//
//	No currency symbols will be configured since by
//	definition, signed numbers do not contain currency
//	symbols.
//
//	Positive Numeric Values:
//
//	The positive number sign is implied. No positive
//	number sign is applied.
//
//		Example-1:
//			Positive Numeric Signed Number Value
//				123.456
//
//	Zero Numeric Values:
//
//	The zero number value is neither positive nor
//	negative. Therefore, no number sign is applied to
//	zero numeric values.
//
//		Example-2:
//			Zero Numeric Signed Number Value
//				0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumBasic(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumBasic()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setSignedNumBasic(
			&newSignedNumFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingNegativeNumSign),
			[]rune(trailingNegativeNumSign),
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumBasicRunes
//
//	Returns a new instance of NumStrFormatSpec configured
//	with a Signed Number String Format using a basic set
//	Signed Number String Format Specification parameters.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	This method is one step above NewSignedNumSimple()
//	in terms of complexity and customization options.
//
//	This method provides additional options for multiple
//	integer grouping types including India and Chinese
//	Numbering.
//
//	Negative number sign symbols may be configured as
//	leading negative number signs, trailing negative
//	number signs or both. The combination of leading
//	and trailing negative number signs allows for the
//	configuration symbols like parentheses for the
//	formatting of negative numbers.
//
//		Example: (125.34)
//
// ----------------------------------------------------------------
//
// # Signed Number Basic Defaults
//
//	No Currency Symbols Included:
//
//	No currency symbols will be configured since by
//	definition, signed numbers do not contain currency
//	symbols.
//
//	Positive Numeric Values:
//
//	The positive number sign is implied. No positive
//	number sign is applied.
//
//		Example-1:
//			Positive Numeric Signed Number Value
//				123.456
//
//	Zero Numeric Values:
//
//	The zero number value is neither positive nor
//	negative. Therefore, no number sign is applied to
//	zero numeric values.
//
//		Example-2:
//			Zero Numeric Signed Number Value
//				0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol for the returned
//		instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		This rune array contains one or more text
//		characters used to separate groups of integers.
//		This separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		basic currency and number sign symbol default
//		values.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumBasicRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumBasicRunes()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setSignedNumBasic(
			&newSignedNumFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingNegativeNumSign,
			trailingNegativeNumSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumDefaultsFrance
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using French
//	Number String formatting conventions.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrFormatSpec will be
//	empty and configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on German Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.NewSignedNumDefaultsGermany()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		French Example-4
//		1 000 000 000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		French Example-5
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec configured with French number
//		formatting specifications.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumDefaultsFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNumDefaultsFrance(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumDefaultsGermany
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using German
//	Number String formatting conventions.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrFormatSpec will be
//	empty and configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on French Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.NewSignedNumDefaultsFrance()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-3
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		German Example-4
//		1.000.000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		German Example-5
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec configured with German number
//		formatting specifications.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumDefaultsGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNumDefaultsGermany(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumDefaultsUKMinus
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using United Kingdom
//	(UK) Number String formatting conventions.
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		UK Example: Negative Numeric Value
//				-123
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols. Therefore, Currency Symbols WILL NOT BE
//	INCLUDED in the returned number symbol
//	specifications. The Currency member variable in the
//	returned NumStrFormatSpec will be empty and
//	configured as a 'NOP' or empty placeholder. 'NOP' is
//	a computer science term which stands for
//	'No Operation'.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://freeformatter.com/united-kingdom-standards-code-snippets.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United Kingdom Example-2
//			1,000,000,000
//
//	Negative Number Formatting
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United Kingdom Example-3
//			-1,000,000,000
//
//	Positive Number Formatting
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-4
//			1,000,000,000
//
//	Zero Number Formatting
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumDefaultsUKMinus(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumDefaultsUKMinus()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNumDefaultsUKMinus(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumDefaultsUSMinus
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using US
//	(United States) Number String formatting
//	conventions.
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrFormatSpec will be
//	empty and configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyBasic()
//		NumStrFormatSpec.NewCurrencySimple()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator:
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1:
//			123.45
//
//	Integer Group Separator:
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		US Example-2:
//			1,000,000,000
//
//	Negative Number Formatting:
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		US Example-3:
//		NumFieldSymPos.InsideNumField()
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//	Positive Number Formatting:
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-4:
//			Positive Numeric Value
//				1,000,000
//
//	Zero Number Formatting:
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-5:
//			Zero Numeric Value
//					0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumDefaultsUSMinus(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNumDefaultsUSMinus(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumDefaultsUSParen
//
//	Returns a new instance of NumStrFormatSpec configured
//	for a Signed Number using United States (US) Number
//	String formatting conventions.
//
//	The term 'Paren' in the method name signals that
//	negative numeric values will be configured with a
//	surrounding parentheses ('()').
//
//		US Example: Negative Numeric Value
//					(123.456)
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications since by definition,
//	signed numbers do not contain currency symbols.
//	Therefore, the Currency member variable in the
//	returned NumStrFormatSpec instance will be empty and
//	configured as a 'NOP' or empty placeholder. 'NOP'
//	stands for 'No Operation'.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewSignedNumParams()
//		NumStrFormatSpec.NewSignedNumBasic()
//		NumStrFormatSpec.NewSignedNumSimple()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator:
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Group Separator:
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		US Example-2
//			1,000,000,000
//
//	Negative Number Formatting:
//
//	The negative signed number symbol is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be surrounded with a
//	leading parenthesis sign ('(') and trailing closing
//	parenthesis sing (')'). The negative number signs
//	will be positioned inside the number field:
//
//		US Example-3
//		NumFieldSymPos.InsideNumField()
//				Number Field Length: 9
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " (123.45)"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//	Positive Number Formatting:
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-4
//			Positive Numeric Value
//				1,000,000
//
//	Zero Number Formatting:
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-5
//			Zero Numeric Value
//					0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		signed number formatting commonly applied in
//		the United States (US). This version of the
//		signed number format will configure negative
//		values with surrounding parentheses:
//				Example: (123.456)
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumDefaultsUSParen(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumDefaultsUSParen()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNumDefaultsUSParen(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumParams
//
//	Creates and returns a new instance of
//	NumStrFormatSpec configured with signed number
//	formatting specifications.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	The new returned instance of NumStrFormatSpec will
//	be reconfigured using the Positive, Negative, and
//	Zero number symbol specifications passed as input
//	parameters.
//
//	The Currency symbol member variables in the new
//	returned instance of NumStrFormatSpec are
//	assigned empty 'NOP' placeholder values since, by
//	definition, signed numbers do not contain currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	trailingPosNumSign				string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
//
//	trailingZeroNumSign				string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix						interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumFmtSpec				NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		signed number formatting specifications.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumParams(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumParams()",
		"")

	if err != nil {
		return newNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setSignedNumParams(
			&newNumFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumFmtSpec<-"))

	return newNumFmtSpec, err
}

//	NewSignedNumParamsRunes
//
//	Creates and returns a new instance of
//	NumStrFormatSpec configured with signed number
//	formatting specifications.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	The new returned instance of NumStrFormatSpec will
//	be reconfigured using the Positive, Negative, and
//	Zero number symbol specifications passed as rune
//	array input parameters.
//
//	The Currency symbol member variables in the new
//	returned instance of NumStrFormatSpec are
//	assigned empty 'NOP' placeholder values since, by
//	definition, signed numbers do not contain currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPosNumSign				[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix                interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumFmtSpec				NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		signed number formatting specifications.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumParamsRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumParamsRunes()",
		"")

	if err != nil {
		return newNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setSignedNumParams(
			&newNumFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumFmtSpec<-"))

	return newNumFmtSpec, err
}

//	NewSignedNumSimple
//
//	Creates and returns and instance of NumStrFormatSpec
//	configured for Signed Number String formatting.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	If currency number symbol formatting IS REQUIRED,
//	see method:
//
//		NumStrFormatSpec.NewCurrencySimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Simple Signed Number Defaults
//
//	Integer Separator Type:
//
//	The integer separator type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumSimple(
	decSeparatorChars string,
	intSeparatorChars string,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumSimple()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setSignedNumSimple(
			&newNumStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

	return newNumStrFmtSpec, err
}

//	NewSignedNumSimpleRunes
//
//	Creates and returns and instance of NumStrFormatSpec
//	configured for Signed Number String formatting.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	If currency number symbol formatting IS REQUIRED,
//	see method:
//
//		NumStrFormatSpec.NewCurrencySimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.NewCountryCurrencyNumFormat()
//		NumStrFormatSpec.NewCountrySignedNumFormat()
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewCurrencyParams()
//		NumStrFormatSpec.NewCurrencyParamsRunes()
//
// ----------------------------------------------------------------
//
// # Simple Signed Number Defaults
//
//	Integer Separator Type:
//
//	The integer separator type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			[]rune
//
//		This rune array contains one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the United
//		States, the standard integer digits separator is
//		the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumSimpleRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumSimpleRunes()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecMechanics).
		setSignedNumSimple(
			&newNumStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

	return newNumStrFmtSpec, err
}

// NewSignedPureNumberStr
//
// Creates and returns and instance of NumStrFormatSpec
// configured with specifications used in formatting
// floating point pure number strings.
//
// A floating point pure number string is defined as
// follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters.
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point. This
//     could be, but is not limited to, a decimal point
//     ('.').
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NOT include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted floating
//		point Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		also known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingNumSymbols			bool
//
//		In Pure Number Strings, positive numeric values
//		are NOT configured with leading or trailing plus
//		signs ('+'). Negative values on the other hand,
//		are always designated by leading or trailing
//		minus sign ('-').
//
//		This parameter, 'leadingNumSymbols', controls
//		the positioning of minus signs for negative
//		numeric values within a	Number String.
//
//		When set to 'true', the returned NumStrFormatSpec
//		instance will configure minus signs for negative
//		numbers at the beginning of, or on the left side
//		of, the numeric value. In these cases, the minus
//		sign is said to be configured as a leading minus
//		sign. This is the positioning format used in the
//		US, UK, Australia and most of Canada. In
//		addition, library functions in 'Go' and other
//		programming languages generally expect leading
//		minus signs for negative numbers.
//
//			Example Leading Minus Sign:
//				"-123.456"
//
//		When parameter 'leadingNumSymbols' is set to
//		'false', the returned instance of NumStrFormatSpec
//		will configure minus signs for negative numbers
//		at the end of, or on the right side of, the
//		numeric value. With this positioning format, the
//		minus sign is said to be configured as a trailing
//		minus sign. This is the positioning format used
//		in France, Germany and many countries in the
//		European Union.
//
//			Example Trailing Minus Sign:
//				"123.456-"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//	newNumStrFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedPureNumberStr(
	decSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newNumStrFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedPureNumberStr()",
		"")

	if err != nil {
		return newNumStrFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setSignedPureNStrSpec(
			&newNumStrFmtSpec,
			decSeparatorChars,
			leadingNumSymbols,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec"))

	return newNumStrFmtSpec, err
}

//	SetCountryCurrencyNumFmt
//
//	Reconfigures the current instance of
//	NumStrFormatSpec based on a Number String Country
//	Culture Specification passed as an input paramter.
//
//	This method will configure the current instance of
//	NumStrFormatSpec for Currency Numeric Values
//	according to the designated country or culture
//	specified by input parameter 'countryCultureFormat'.
//
//
//	For signed number formats see method:
//
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that all data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	with Currency Number String formatting parameters specified
//	by input parameter 'countryCultureFormat'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		currency formatting information for a
//		designated country or culture.
//
//		This method will NOT change the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameter 'numberFieldSpec'
//		to construct and return a new instance of
//		NumStrFormatSpec configured for Currency Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		The data values in 'numberFieldSpec' will be
//		combined with input parameter
//		'countryCultureFormat' to construct and return
//		a new instance of NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) SetCountryCurrencyNumFmt(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCountryCurrencyNumFmt()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		countryCultureFormat.CurrencyNumStrFormat.decSeparator,
		countryCultureFormat.CurrencyNumStrFormat.intSeparatorSpec,
		countryCultureFormat.CurrencyNumStrFormat.numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("newCurrencyNumFmtSpec<-"))
}

//	SetCountrySignedNumFmt
//
//	Reconfigures the current instance of
//	NumStrFormatSpec based on a Number String Country
//	Culture Specification passed as an input paramter.
//
//	This method will configure the current instance of
//	NumStrFormatSpec for Signed Numeric Values
//	according to the designated country or culture
//	specified by input parameter 'countryCultureFormat'.
//
//
//	For currency number formats see method:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that all data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	with Signed Number String formatting parameters specified
//	by input parameter 'countryCultureFormat'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		signed number formatting information for a
//		designated country or culture.
//
//		This method will NOT change the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameter 'numberFieldSpec'
//		to construct and return a new instance of
//		NumStrFormatSpec configured for Signed Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		The data values in 'numberFieldSpec' will be
//		combined with input parameter
//		'countryCultureFormat' to construct and return
//		a new instance of NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) SetCountrySignedNumFmt(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCountrySignedNumFmt()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		countryCultureFormat.SignedNumStrFormat.decSeparator,
		countryCultureFormat.SignedNumStrFormat.intSeparatorSpec,
		countryCultureFormat.SignedNumStrFormat.numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))
}

//	SetCurrencyBasic
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using basic currency and number
//	string formatting default values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingCurrencySymbol			string
//
//		The character or characters which comprise the
//		leading Currency Symbol. The leading Currency
//		Symbol will be positioned at the beginning or
//		left side of the number string.
//
//			Example: $ 123.45
//
//		If a space between the currency symbol
//		and the first digit of the number string
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Leading Currency Symbol: "$ "
//				Formatted Number String: "$ 123.45"
//
//	trailingCurrencySymbol			string
//
//		The character or characters which comprise the
//		trailing Currency Symbol. The trailing Currency
//		Symbol will be positioned at the end of, or
//		right side of, the number string.
//
//			Example: 123.45 €
//
//		If a space between the last digit of the
//		number string and the currency symbol
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Trailing Currency Symbol: " €"
//				Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This parameter determines whether the currency
//		symbol will be positioned inside or outside
//		the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency and number
//		sign characters relative to a Number Field in
//		which a number string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyBasic(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyInsideNumSymbol bool,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyBasic()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setCurrencyBasic(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingCurrencySymbol),
			[]rune(trailingCurrencySymbol),
			currencyInsideNumSymbol,
			[]rune(leadingNegativeNumSign),
			[]rune(trailingNegativeNumSign),
			numSymbolFieldPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec"))
}

//	SetCurrencyBasicRunes
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using basic currency and number
//	string formatting default values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol for the current
//		instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		This rune array contains one or more text
//		characters used to separate groups of integers.
//		This separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//				'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingCurrencySymbol			[]rune
//
//		This rune array contains the character or
//		characters which comprise the leading Currency
//		Symbol. The leading Currency Symbol will be
//		positioned at the beginning or left side of the
//		number string.
//
//			Example: $ 123.45
//
//		If a space between the currency symbol
//		and the first digit of the number string
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Leading Currency Symbol: "$ "
//				Formatted Number String: "$ 123.45"
//
//	trailingCurrencySymbol			[]rune
//
//		This rune array contains the character or
//		characters which comprise the trailing Currency
//		Symbol. The trailing Currency Symbol will be
//		positioned at the end of, or right side of, the
//		number string.
//
//			Example: 123.45 €
//
//		If a space between the last digit of the
//		number string and the currency symbol
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Trailing Currency Symbol: " €"
//				Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This parameter determines whether the currency
//		symbol will be positioned inside or outside
//		the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency and number
//		sign characters relative to a Number Field in
//		which a number string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyBasicRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingCurrencySymbol []rune,
	trailingCurrencySymbol []rune,
	currencyInsideNumSymbol bool,
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyBasicRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setCurrencyBasic(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingCurrencySymbol,
			trailingCurrencySymbol,
			currencyInsideNumSymbol,
			leadingNegativeNumSign,
			trailingNegativeNumSign,
			numSymbolFieldPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec"))
}

//	SetCurrencyDefaultsFrance
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Currency Number String
//	formatting conventions typically applied in France.
//
//	The default French currency symbol is a trailing Euro
//	sign ('€').
//
//		French Example:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with French Currency
//	Number formatting specifications. New data values
//	will be configured for the positive, zero and
//	negative number sign symbols as well as the currency
//	symbol.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either France or
//	Germany.
//
//	For information on German Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.SetCurrencyDefaultsGermany()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1:
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2:
//		1 000 000 000
//
//	Currency Symbols
//
//	The currency symbol used in the France is the
//	trailing Euro Currency Symbol ('€').
//
//		French Example-3:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied. No positive
//	number is applied, only the trailing Euro Currency
//	Symbol.
//
//		French Example-4:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Zero Numeric Values
//
//	The zero number format has no number sign, but the
//	currency symbol is set to a trailing Euro Currency
//	Symbol.
//
//		French Example-5:
//			Zero Numeric Currency Value
//				0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-') and a trailing Euro Currency Symbol
//	("€").
//
//		French Example-6:
//			Negative Numeric Currency Value
//				-1 000 000,00 €
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		French Example-7:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123,45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " -123,45 €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsFrance(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyDefaultsGermany
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Currency Number String
//	formatting conventions typically applied in Germany.
//
//	The default German currency symbol is a trailing Euro
//	sign ('€').
//
//		German Example:
//			Positive Numeric Currency Value
//				1.000.000,00 €
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with German Currency
//	Number formatting specifications. New data values
//	will be configured for the positive, zero and
//	negative number sign symbols as well as the currency
//	symbol.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Currency Number
//	String formatting standards implemented by either
//	Germany or France.
//
//	For information on French Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.SetCurrencyDefaultsFrance()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	Currency Symbols
//
//	The currency symbol used in the Germany is the
//	trailing Euro symbol ('€').
//
//		German Example-3:
//		1.000.000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-4:
//		1.000.000 €
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-5:
//			0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-6:
//		1.000.000- €
//
//	The negative signed number is configured with a
//	trailing minus sign ('-') appended at the end of
//	the number string. The negative number sign and the
//	currency symbol will be positioned inside the number
//	field:
//
//		German Example-5:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: 123,45-
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " 123,45- €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsGermany(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyDefaultsUKMinusInside
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the United Kingdom (UK).
//
//	The default UK currency symbol is a leading Pound
//	sign ('£').
//
//		UK Example:
//			Positive Numeric Currency Value
//				£ 123.45  Positive Value
//
//	The term "MinusInside" in the method name means that
//	the minus sign ('-') configured for negative numeric
//	values will be inside, or to the right of, the Pound
//	sign ('£').
//
//		UK Example:
//			Negative Numeric Currency Value
//				£ -123.45  Negative Value
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with UK Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the United Kingdom (UK).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		UK Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		UK Example-2:
//			1,000,000
//
//	Currency Symbol
//
//	The default currency symbol used in the UK is the
//	leading Pound symbol ('£').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-3:
//			Positive Numeric Currency Value
//				£ 123.45
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-4:
//			Zero Numeric Currency Value
//				£ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the minus sign ('-')
//	such that any minus sign configured for negative
//	numeric values will be inside, or to the right of,
//	the Pound sign ('£').
//
//		UK Example-5:
//			Negative Numeric Currency Value
//				£ -123.45  Negative Value
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		UK Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " £ -123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsUKMinusInside(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsUKMinusInside(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyDefaultsUKMinusOutside
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the United Kingdom (UK).
//
//	The default UK currency symbol is a leading Pound
//	sign ('£').
//
//		UK Example:	Positive Numeric Currency Value
//			£ 123.45
//
//	The term "MinusOutside" in the method name means that
//	the minus sign ('-') configured for negative numeric
//	values will be outside, or to the left of, the Pound
//	sign ('£').
//
//	UK Example:	Negative Numeric Currency Value
//			- £123.45
//
//	Default values will be used to configure the returned
//	instance of NumStrFormatSpec with UK Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the United Kingdom (UK).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		UK Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		UK Example-2:
//			1,000,000
//
//	Currency Symbol
//
//	The default currency symbol used in the UK is the
//	leading Pound symbol ('£').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-3:
//			Positive Numeric Currency Value
//				£ 123.45  Positive Value
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-4:
//			Zero Numeric Currency Value
//				£ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the Pound sign ('£')
//	such that any minus sign configured for negative
//	numeric values will be outside, or to the left of,
//	the Pound sign ('£').
//
//		UK Example-5:
//			Negative Numeric Currency Value
//				- £123.45
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		UK Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " - £123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsUKMinusOutside(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsMinusOutside(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyDefaultsUSMinus
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the United States (US).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example
//			Negative Numeric Currency Value
//				$ -123
//
//	Default values will be used to configure the current
//	instance of NumStrFormatSpec with US Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced by Currency Number String formatting
//	parameters typically applied in the United States
//	(US).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator:
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator:
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		US Example-2:
//			1,000,000
//
//	Currency Symbols:
//
//	The default currency symbol used in the US is the
//	leading Dollar symbol ('$').
//
//	Positive Numeric Values:
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-3:
//			Positive Numeric Currency Value
//				$ 123.45
//
//	Zero Numeric Values:
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-4:
//			Zero Numeric Currency Value
//				$ 0.00
//
//	Negative Numeric Values:
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the Dollar sign ('$')
//	such that any minus sign configured for negative
//	numeric values will be inside, or to the right of,
//	the Dollar sign ('$').
//
//		US Example-5:
//			Negative Numeric Currency Value
//				$ -123.45
//
//	The negative signed number is configured with a
//	leading minus sign ('-') prefixed at the beginning
//	of the number string. The negative number sign and
//	the currency symbol will be positioned inside the
//	number field:
//
//		US Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " $ -123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsUSMinus(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsUSMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsUSMinus(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyDefaultsUSParen
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the United States (US).
//
//	The default US currency symbol is a leading Dollar
//	sign ('$').
//
//		US Example
//			Positive Numeric Currency Value
//				$ 123.45
//
//	The term 'Paren' in the method name signals that a
//	surrounding parentheses ('()') will be used to
//	designate negative numeric values.
//
//		US Example
//			Negative Numeric Currency Value
//				$ (123)
//
//	Default values will be used to configure the returned
//	instance of NumStrFormatSpec with US Currency Number
//	formatting specifications. New data values will be
//	configured for the positive, zero and negative number
//	sign symbols as well as the currency symbol.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced by Currency Number String formatting
//	parameters typically applied in the United States
//	(US).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separators
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Separators
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		US Example-2
//			1,000,000
//
//	Currency Symbols
//
//	The default currency symbol used in the US is the
//	leading Dollar symbol ('$').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-3:
//			Positive Numeric Currency Value
//				$ 123.45
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-4:
//			Zero Numeric Currency Value
//				$ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to surrounding
//	parentheses ('()').
//
//	This method will configure the Dollar sign ('$')
//	such that the leading parenthesis ('(') configured
//	for negative numeric values will be inside, or to the
//	right of, the Dollar sign ('$').
//
//		US Example-5:
//			Negative Numeric Currency Value
//				$ (123.45)
//
//	The negative signed number is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be prefixed with a
//	leading parenthesis symbol ('(') and suffixed with a
//	trailing, or closing, parenthesis symbol (')'). The
//	negative number sign symbols and the currency symbol
//	will be positioned inside the number field:
//
//		US Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 11
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " $ (123.45)"
//				Number Field Index:------>01234567890
//				Total Number String Length: 11
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyDefaultsUSParen(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyDefaultsUSParen()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyDefaultsUSParen(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyParams
//
//	Deletes all data values in the current instance of
//	NumStrFormatSpec and reconfigures that instance
//	based on the Positive, Negative, Zero and Currency
//	symbol specifications passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	trailingPosNumSign				string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
//
//	trailingZeroNumSign				string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to an empty string.
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to an empty string.
//
//	currencyInsideNumSymbol			bool
//
//		This boolean parameter determines whether the
//		currency symbol will be positioned inside or
//		outside the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right Justified
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right Justified
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix                interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyParams(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	leadingCurrencySymbols string,
	trailingCurrencySymbols string,
	currencyInsideNumSymbol bool,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyParams()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyParams(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			[]rune(leadingCurrencySymbols),
			[]rune(trailingCurrencySymbols),
			currencyInsideNumSymbol,
			currencyNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec<-"))
}

//	SetCurrencyParamsRunes
//
//	Deletes all data values in the current instance of
//	NumStrFormatSpec and reconfigures that instance
//	based on the Positive, Negative, Zero and Currency
//	symbol specifications passed as rune array input
//	parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPosNumSign				[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the Dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//				Example: $125.34
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to a 'nil' or empty
//		rune array.
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to a 'nil' or empty
//		rune array.
//
//	currencyInsideNumSymbol			bool
//
//		This boolean parameter determines whether the
//		currency symbol will be positioned inside or
//		outside the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') and Trailing
//		Currency Symbol ('trailingCurrencySymbol')
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right Justified
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right Justified
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'numFieldJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyParamsRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	leadingCurrencySymbols []rune,
	trailingCurrencySymbols []rune,
	currencyInsideNumSymbol bool,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencyParamsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyParams(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			leadingCurrencySymbols,
			trailingCurrencySymbols,
			currencyInsideNumSymbol,
			currencyNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))
}

//	SetCurrencySimple
//
//	Reconfigures the current instance of NumStrFormatSpec
//	for Currency Number String formatting.
//
//	If currency number symbol formatting IS NOT required,
//	see method:
//
//		NumStrFormatSpec.SetSignedNumSimple()
//
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced with Currency Number String Formatting
//	parameters using default values.
//
// ----------------------------------------------------------------
//
// # Simple Currency Defaults
//
//	Integer Grouping:
//
//	The integer grouping type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String:	"123.456- €"
//			US Number String:		"$ -123.456"
//			UK Number String:		"£ -123.45"
//
//	Currency Symbol - Padding Space:
//
//		As a default, one space may be added as padding
//		for the currency symbol.
//
//		If a space is NOT present, a space will be
//		automatically inserted between the currency
//		symbol and the first digit or minus sign.
//
//		Example Number Strings:
//			"$ 123.456"
//			"123.456 €"
//			"$ -123.456"
//			"123.456- €"
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the current instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency.
//		This currency formatting will be configured in
//		the current instance of NumStrFormatSpec.
//
//	leadingCurrencySymbols		bool
//
//		Controls the positioning of Currency Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Currency
//		Symbols at the beginning or left side of the
//		number string. Such Currency Symbols are therefore
//		configured as leading Currency Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number Strings:
//			"$ 123.456"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the first digit or
//				leading minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure Currency Symbols on the right side of
//		the number string. Currency Number Symbols are
//		therefore configured as trailing Number Symbols.
//		This is the positioning format used in France,
//		Germany and many other countries in the European
//		Union.
//
//			Example Number Strings:
//				"123.456 €"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the last digit or
//				trailing minus sign.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencySimple(
	decSeparatorChars string,
	intSeparatorChars string,
	currencySymbols string,
	leadingCurrencySymbols bool,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencySimple()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setCurrencySimple(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			[]rune(currencySymbols),
			leadingCurrencySymbols,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

}

//	SetCurrencySimpleRunes
//
//	Reconfigures the current instance of NumStrFormatSpec
//	for Currency Number String formatting.
//
//	If currency number symbol formatting IS NOT required,
//	see method:
//
//		NumStrFormatSpec.
//			SetSignedNumSimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced with Currency Number String Formatting
//	parameters using default values.
//
// ----------------------------------------------------------------
//
// # Simple Currency Defaults
//
//	Integer Grouping:
//
//	The integer grouping type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String:	"123.456- €"
//			US Number String:		"$ -123.456"
//			UK Number String:		"£ -123.45"
//
//	Currency Symbol - Padding Space:
//
//		As a default, one space may be added as padding
//		for the currency symbol.
//
//		If a space is NOT present, a space will be
//		automatically inserted between the currency
//		symbol and the first digit or minus sign.
//
//		Example Number Strings:
//			"$ 123.456"
//			"123.456 €"
//			"$ -123.456"
//			"123.456- €"
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		current instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			[]rune
//
//		This rune array contains one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the United
//		States, the standard integer digits separator is
//		the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	currencySymbols				[]rune
//
//		This rune array contains the symbol or symbols
//		used to format currency. This currency formatting
//		will be configured in the current instance of
//		NumStrFormatSpec.
//
//	leadingCurrencySymbols		bool
//
//		Controls the positioning of Currency Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Currency
//		Symbols at the beginning or left side of the
//		number string. Such Currency Symbols are therefore
//		configured as leading Currency Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number Strings:
//			"$ 123.456"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the first digit or
//				leading minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure Currency Symbols on the right side of
//		the number string. Currency Number Symbols are
//		therefore configured as trailing Number Symbols.
//		This is the positioning format used in France,
//		Germany and many other countries in the European
//		Union.
//
//			Example Number Strings:
//				"123.456 €"
//
//		NOTE:	If a space is NOT present, a space will
//				be automatically inserted between the
//				currency symbol and the last digit or
//				trailing minus sign.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencySimpleRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	currencySymbols []rune,
	leadingCurrencySymbols bool,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetCurrencySimpleRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setCurrencySimple(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			currencySymbols,
			leadingCurrencySymbols,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

}

//	SetDecimalSeparator
//
//	SpecDeletes and replaces the Decimal Separator
//	Specification for the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec			DecimalSeparatorSpec
//
//		An instance of DecimalSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//
//			'NumStrFormatSpec.decSeparator'.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted, floating
//		point Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		Decimal Separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetDecimalSeparatorSpec(
	decSeparatorSpec DecimalSeparatorSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetDecimalSeparatorSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setDecimalSeparatorSpec(
		numStrFmtSpec,
		decSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-decSeparatorSpec"))
}

// SetIntegerGroupingSpec - Deletes and replaces the Integer
// Grouping Specification for the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		An instance of Integer Separator Specification.
//		The member variable data values contained in
//		this structure will be copied to the current
//		instance of	NumStrFormatSpec :
//
//			'NumStrFormatSpec.intSeparatorSpec'.
//
//		In the United States, the Integer Separator
//		Specification character is a comma (',') with
//		integers grouped in	thousands.
//
//			United States Example: 1,000,000
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetIntegerGroupingSpec(
	intSeparatorSpec IntegerSeparatorSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetIntegerGroupingSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setIntegerGroupingSpec(
		numStrFmtSpec,
		intSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-intSeparatorSpec"))
}

//	SetNegativeNumberFmtSpec
//
//	Deletes and replaces the Negative Number Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.negativeNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a negative numeric value, this formatting
//	specification will be applied.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.negativeNumberSign'.
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetNegativeNumberFmtSpec(
	negativeNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNegativeNumberFmtSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setNegativeNumberSignSpec(
			numStrFmtSpec,
			negativeNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"negativeNumberSign"))
}

// SetNumberFieldSpec - Deletes and replaces the Number Field
// Specification for the current instance of
// NumStrFormatSpec.
//
// The Number Field Specification includes parameters for
// field length and text justification ('Right, Center,
// Left').
//
// Numeric digits are formatted using the Number Field
// Specification within a number field specified by field
// length. The numeric digits string is then justified
// within the number field according to a text justification
// specification of 'Right', 'Center' or 'Left'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		An instance of NumStrNumberFieldSpec. The member
//		variable data values contained in this instance
//		will be copied to the current instance of
//		NumStrFormatSpec :
//			'NumStrFormatSpec.numberFieldSpec'.
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetNumberFieldSpec(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNumberFieldSpec(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-numberFieldSpec"))
}

//	SetPositiveNumberFmtSpec
//
//	Deletes and replaces the Positive Number Sign Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.positiveNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a positive numeric value, this formatting
//	specification will be applied.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.positiveNumberSign'.
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetPositiveNumberFmtSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetPositiveNumberFmtSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setPositiveNumberSignSpec(
			numStrFmtSpec,
			positiveNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"positiveNumberSign"))
}

//	SetNumFmtComponents
//
//	Reconfigures the current instance of NumStrFormatSpec
//	based on Number String formatting components passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced by values generated from the listed
//	input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	numberSymbolsGroup  				NumStrNumberSymbolGroup
//
//		This instance of NumStrNumberSymbolGroup contains the
//		Number Symbol Specifications for negative numeric
//		values, positive numeric values and zero numeric
//		values.
//
//		type NumStrNumberSymbolGroup struct {
//
//			negativeNumberSign NumStrNumberSymbolSpec
//
//				The Number String Negative Number Sign
//				Specification is used to configure negative
//				number sign symbols for negative numeric
//				values formatted and displayed in number
//				stings.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Negative
//					Values
//
//					Leading Symbols: "- "
//					Number String:   "- 123.456"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Negative
//					Values
//
//					Leading Symbols: "-"
//					Number String:   "-123.456"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Negative
//					Values
//
//					Trailing Symbols: " -"
//					Number String:   "123.456 -"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Negative
//					Values
//
//					Trailing Symbols: "-"
//					Number String:   "123.456-"
//
//			positiveNumberSign NumStrNumberSymbolSpec
//
//				Positive number signs are commonly implied
//				and not specified. However, the user has
//				the option to specify a positive number sign
//				character or characters for positive numeric
//				values using a Number String Positive Number
//				Sign Specification.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Positive
//					Values
//
//					Leading Symbols: "+ "
//					Number String:   "+ 123.456"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Positive
//					Values
//
//					Leading Symbols: "+"
//					Number String:   "+123.456"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Positive
//					Values
//
//					Trailing Symbols: " +"
//					Number String:   "123.456 +"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Positive
//					Values
//
//					Trailing Symbols: "+"
//					Number String:   "123.456+"
//
//			zeroNumberSign NumStrNumberSymbolSpec
//
//				The Number String Zero Number Sign
//				Specification is used to configure number
//				sign symbols for zero numeric values formatted
//				and displayed in number stings. Zero number
//				signs are commonly omitted because zero
//				does not technically qualify as either a
//				positive or negative value. However,
//				the user has the option to configure number
//				sign symbols for zero values if necessary.
//
//				Example-1: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Zero Values
//
//					Leading Symbols: "+"
//					Trailing Symbols: ""
//					Number String:   "+0.00"
//
//				Example-2: Leading Number Sign Symbols
//					Leading Number Sign Symbols for Zero Values
//
//					Leading Symbols: "+ "
//					Trailing Symbols: ""
//					Number String:   "+ 0.00"
//
//				Example-3: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Zero Values
//
//					Leading Symbols: ""
//					Trailing Symbols: " +"
//					Number String:   "0.00 +"
//
//				Example-4: Trailing Number Sign Symbols
//					Trailing Number Sign Symbols for Zero Values
//
//					Leading Symbols: ""
//					Trailing Symbols: "+"
//					Number String:   "0.00+"
//
//			currencySymbol NumStrNumberSymbolSpec
//
//				A Currency Symbol positioned next to a numeric
//				value designates the number is a monetary
//				amount.
//
//				Examples of Currency Symbols include the Dollar
//				sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//				This instance of NumStrNumberSymbolSpec is used
//				to configure leading Currency Symbols, trailing
//				Currency Symbols or both leading and trailing
//				Currency Symbols.
//
//				Example-1: Leading Currency Symbols
//
//					Leading Currency Symbols: "$ "
//					Number String:   "$ 123.456"
//
//				Example-2: Leading Currency Symbols
//
//					Leading Currency Symbols: "$"
//					Number String:   "$123.456"
//
//				Example-3: Trailing Currency Symbols
//					Trailing Currency Symbols for Positive Values
//
//					Trailing Currency Symbols: "€"
//					Number String:   "123.456€"
//
//				Example-4: Trailing Currency Symbols
//					Trailing Currency Symbols for Positive Values
//
//					Trailing Currency Symbols: " €"
//					Number String:   "123.456 €"
//		}
//
//	numberFieldSpec			NumStrNumberFieldSpec
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetNumFmtComponents(
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	numberSymbolsGroup NumStrNumberSymbolGroup,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNumFmtComponents()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		numberSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))

	return err
}

//	SetNumFmtElements
//
//	Reconfigures the current instance of NumStrFormatSpec
//	based on Number String formatting specifications
//	passed as input parameters.
//
//	Options include customizing for currency symbols,
//	integer separation, number sign	management, radix
//	point symbols, and floating point number rounding.
//
//	If required, users have the options of
//	implementing the India or Chinese Numbering
//	Systems for integer separation.
//
//	This method offers the maximum degree of granular
//	control over all aspects of the Number String
//	formatting operation.
//
//	In particular, it offers maximum flexibility in
//	configuring integer separator characters and
//	integer grouping sequences.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If Currency Symbol formatting is NOT required,
//	set the Currency Symbol parameter to a 'NOP'.
//	'NOP' is a computer science term for
//	'No Operation'.
//
//		currencySymbol :=
//			new(NumStrNumberSymbolSpec).NewNOP()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced by values generated from the listed
//	input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	negativeNumberSign				NumStrNumberSymbolSpec
//
//		The Number String Negative Number Sign
//		Specification is used to configure negative
//		number sign symbols for negative numeric
//		values formatted and displayed in number
//		stings.
//
//		If this parameter is submitted as an empty or
//		invalid Negative Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	positiveNumberSign 				NumStrNumberSymbolSpec
//
//		Positive number signs are commonly implied
//		and not specified. However, the user has
//		the option to specify a positive number sign
//		character or characters for positive numeric
//		values using this input parameter.
//
//		If this parameter is submitted as an empty or
//		invalid Positive Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//	zeroNumberSign					NumStrNumberSymbolSpec
//
//		The Number String Zero Number Sign
//		Specification is used to configure number
//		sign symbols for zero numeric values formatted
//		and displayed in number stings. Zero number signs
//		are commonly omitted because zero does not
//		technically qualify as either a positive or
//		negative value. However, the user has the option
//		to configure number sign symbols for zero values
//		if necessary.
//
//		If this parameter is submitted as an empty or
//		invalid Zero Number Sign Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+ "
//			Trailing Symbols: ""
//			Number String:   "+ 0.00"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: "+"
//			Number String:   "0.00+"
//
//	currencySymbol					NumStrNumberSymbolSpec
//
//		A Currency Symbol positioned next to a numeric
//		value designates the number is a monetary amount.
//
//		The Number String Currency Symbol Specification
//		is used to configure currency symbols for
//		positive, negative and zero numeric values
//		formatted and displayed in number stings.
//
//		If Currency Symbol formatting is NOT required,
//		set this parameter to a 'NOP'. 'NOP' is a
//		computer science term for 'No Operation'.
//
//			currencySymbol :=
//				new(NumStrNumberSymbolSpec).NewNOP()
//
//		If this parameter is submitted as an empty or
//		invalid Currency Symbol Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Examples of Currency Symbols include the Dollar
//		sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//		This instance of NumStrNumberSymbolSpec is used
//		to configure leading Currency Symbols, trailing
//		Currency Symbols or both leading and trailing
//		Currency Symbols.
//
//		Example-1: Leading Currency Symbols
//
//			Leading Currency Symbols: "$ "
//			Number String:   "$ 123.456"
//
//		Example-2: Leading Currency Symbols
//
//			Leading Currency Symbols: "$"
//			Number String:   "$123.456"
//
//		Example-3: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: "€"
//			Number String:   "123.456€"
//
//		Example-4: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: " €"
//			Number String:   "123.456 €"
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
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec.
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
func (numStrFmtSpec *NumStrFormatSpec) SetNumFmtElements(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbol NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNumFmtElements()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtElements(
		numStrFmtSpec,
		decSeparatorSpec,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		currencySymbol,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))
}

//	SetSignedNumBasic
//
//	Reconfigures the current instance of
//	NumStrFormatSpec with basic Signed Number formatting
//	specification default values.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	Negative number sign symbols may be configured as
//	leading negative number signs, trailing negative
//	number signs or both. The combination of leading
//	and trailing negative number signs allows for the
//	configuration symbols like parentheses for the
//	formatting of negative numbers.
//
//		Example: (125.34)
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Signed Number Basic Defaults
//
//	No Currency Symbols Included:
//
//	No currency symbols will be configured since by
//	definition, signed numbers do not contain currency
//	symbols.
//
//	Positive Numeric Values:
//
//	The positive number sign is implied. No positive
//	number sign is applied.
//
//		Example-1:
//			Positive Numeric Signed Number Value
//				123.456
//
//	Zero Numeric Values:
//
//	The zero number value is neither positive nor
//	negative. Therefore, no number sign is applied to
//	zero numeric values.
//
//		Example-2:
//			Zero Numeric Signed Number Value
//				0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumBasic(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumBasic()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setSignedNumBasic(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingNegativeNumSign),
			[]rune(trailingNegativeNumSign),
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec"))
}

//	SetSignedNumBasicRunes
//
//	Reconfigures the current instance of
//	NumStrFormatSpec with basic Signed Number formatting
//	specification default values.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	Negative number sign symbols may be configured as
//	leading negative number signs, trailing negative
//	number signs or both. The combination of leading
//	and trailing negative number signs allows for the
//	configuration symbols like parentheses for the
//	formatting of negative numbers.
//
//		Example: (125.34)
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Signed Number Basic Defaults
//
//	No Currency Symbols Included:
//
//	No currency symbols will be configured since by
//	definition, signed numbers do not contain currency
//	symbols.
//
//	Positive Numeric Values:
//
//	The positive number sign is implied. No positive
//	number sign is applied.
//
//		Example-1:
//			Positive Numeric Signed Number Value
//				123.456
//
//	Zero Numeric Values:
//
//	The zero number value is neither positive nor
//	negative. Therefore, no number sign is applied to
//	zero numeric values.
//
//		Example-2:
//			Zero Numeric Signed Number Value
//				0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol for the returned
//		instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		This rune array contains one or more text
//		characters used to separate groups of integers.
//		This separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumBasicRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumBasicRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setSignedNumBasic(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingNegativeNumSign,
			trailingNegativeNumSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec"))
}

//	SetSignedNumDefaultsFrance
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Number String
//	formatting conventions typically
//	applied in France.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on German Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.SetSignedNumDefaultsGermany()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in France.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		French Example-4
//		1 000 000 000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		French Example-5
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumDefaultsFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNumDefaultsFrance(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumDefaultsGermany
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Number String
//	formatting conventions typically
//	applied in Germany.
//
//	Within in the European Union, many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on French Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.SetSignedNumDefaultsFrance()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see method:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-3
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		German Example-4
//		1.000.000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		German Example-5
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumDefaultsGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNumDefaultsGermany(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumDefaultsUK
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the UK (United Kingdom).
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://freeformatter.com/united-kingdom-standards-code-snippets.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United Kingdom Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United Kingdom Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumDefaultsUK(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumDefaultsUK()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNumDefaultsUSMinus(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumDefaultsUSMinus
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Signed Number String formatting conventions
//	typically applied in the US (United States).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The reconfigured current instance of NumStrFormatSpec
//	will include signed number symbols for positive, zero
//	and negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	number symbol specifications. The Currency member
//	variable will be configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//		NumStrFormatSpec.SetSignedNumBasic()
//		NumStrFormatSpec.SetSignedNumSimple()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator:
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Group Separator:
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	Negative Number Formatting:
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	Positive Number Formatting:
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	Zero Number Formatting:
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumDefaultsUSMinus(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNumDefaultsUSMinus(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumDefaultsUSParen
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the US (United States).
//
//	The term 'Paren' in the method name signals that
//	negative numeric values will be configured with a
//	surrounding parentheses ('()').
//
//		US Example: Negative Numeric Value
//					(123.456)
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of NumStrFormatSpec will
//	include signed number symbols for positive, zero and
//	negative numeric values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	number symbol specifications. The Currency member
//	variable will be configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewSignedNumParams()
//		NumStrFormatSpec.NewSignedNumBasic()
//		NumStrFormatSpec.NewSignedNumSimple()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator:
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		US Example-1
//			123.45 (The fractional digits are "45")
//
//	Integer Group Separator:
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		US Example-2
//			1,000,000,000
//
//	Negative Number Formatting:
//
//	The negative signed number symbol is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be surrounded with a
//	leading parenthesis sign ('(') and trailing closing
//	parenthesis sing (')'). The negative number signs
//	will be positioned inside the number field:
//
//		US Example-3
//		NumFieldSymPos.InsideNumField()
//				Number Field Length: 9
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " (123.45)"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//	Positive Number Formatting:
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-4
//			Positive Numeric Value
//				1,000,000
//
//	Zero Number Formatting:
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-5
//			Zero Numeric Value
//					0
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumDefaultsUSParen(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumDefaultsUSParen()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNumDefaultsUSParen(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumParams
//
//	Reconfigures the current instance of
//	NumStrFormatSpec with basic Signed Number Symbol
//	formatting specifications passed as input
//	parameters.
//
//	Creates and returns a new instance of
//	NumStrFormatSpec configured with signed number
//	formatting specifications.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	The Currency symbol member variables in the
//	reconfigured current instance of NumStrFormatSpec
//	are assigned empty 'NOP' placeholder values since,
//	by definition, signed numbers do not contain currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	trailingPosNumSign				string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
//
//	trailingZeroNumSign				string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix						interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumParams(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumParams()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setSignedNumParams(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec<-"))
}

//	SetSignedNumParamsRunes
//
//	Creates and returns a new instance of
//	NumStrFormatSpec configured with signed number
//	formatting specifications.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	The new returned instance of NumStrFormatSpec will
//	be reconfigured using the Positive, Negative, and
//	Zero number symbol specifications passed as rune
//	array input parameters.
//
//	The Currency symbol member variables in the new
//	returned instance of NumStrFormatSpec are
//	assigned empty 'NOP' placeholder values since, by
//	definition, signed numbers do not contain currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		returned instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPosNumSign				[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " +123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right Justified
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right Justified
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	errorPrefix                interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumParamsRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumParamsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setSignedNumParams(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec<-"))
}

//	SetSignedNumSimple
//
//	Reconfigures the current instance of NumStrFormatSpec
//	for Signed Number String formatting.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	If currency number symbol formatting IS REQUIRED,
//	see method:
//
//		NumStrFormatSpec.SetCurrencySimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with signed number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number
//	string formatting, use one of the more advanced
//	constructor or 'Set' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced with Signed Number String Formatting
//	parameters using 'simple' default values.
//
// ----------------------------------------------------------------
//
// # Simple Signed Number Defaults
//
//	Integer Grouping Type:
//
//	The integer grouping type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumSimple(
	decSeparatorChars string,
	intSeparatorChars string,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumSimple()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setSignedNumSimple(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

}

//	SetSignedNumSimpleRunes
//
//	Reconfigures the current instance of NumStrFormatSpec
//	for Signed Number String formatting.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	If currency number symbol formatting IS REQUIRED,
//	see method:
//
//		NumStrFormatSpec.SetCurrencySimple()
//
//	Type NumStrFormatSpec is used to convert numeric
//	values to formatted Number Strings.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	generated returned instance of NumStrFormatSpec
//	will be configured with signed number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number
//	string formatting, use one of the more advanced
//	constructor or 'Set' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrFormatSpec.SetCountryCurrencyNumFmt()
//		NumStrFormatSpec.SetCountrySignedNumFmt()
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetSignedNumParams()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced with Signed Number String Formatting
//	parameters using 'simple' default values.
//
// ----------------------------------------------------------------
//
// # Simple Signed Number Defaults
//
//	Integer Grouping Type:
//
//	The integer grouping type defaults to thousands.
//	This means that integer digits will be separated in
//	groups of three using the integer separator character
//	passed as input parameter 'intSeparatorChars'.
//
//		Example Integer Separation-1:
//			intSeparatorChars = ','
//			Integer Value = 1000000
//			Formatted Integer Digits: 1,000,000
//
//		Example Integer Separation-2:
//			intSeparatorChars = '.'
//			Integer Value = 1000000
//			Formatted Integer Digits: 1.000.000
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingMinusSign'.
//
//		Examples:
//
//			Leading Minus Sign: "-123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right Justified
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		current instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars			[]rune
//
//		This rune array contains one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator. It is used to separate groups of
//		integer digits to the left of the decimal
//		separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	leadingMinusSign			bool
//
//		Controls the positioning of the minus sign ('-')
//		in a Number String Format configured with a
//		negative numeric value.
//
//		For NumStrNumberSymbolGroup configured with the
//		Simple Currency Number String formatting
//		specification, the default negative number sign
//		symbol is the minus sign ('-').
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure the minus
//		sign at the beginning or left side of the number
//		string. Such minus signs are therefore configured
//		as leading minus signs.
//
//		Example Number Strings:
//			" -123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure the minus sign ('-') on the right side
//		of the number string. The minus sign is therefore
//		configured as trailing minus sign.
//
//			Example Number Strings:
//				"123.456-"
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumSimpleRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	leadingMinusSign bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumSimpleRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecMechanics).
		setSignedNumSimple(
			numStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			leadingMinusSign,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumStrFmtSpec<-"))

}

// SetSignedPureNumberStr
//
// Reconfigures the current NumStrFormatSpec instance
// with specifications for generating a pure number
// string.
//
// A Signed Floating Point Pure Number String is defined
// as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters.
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point. This
//     could be, but is not limited to, a decimal point
//     ('.').
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NOT include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the current instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted floating
//		point Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		also known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingNumSymbols			bool
//
//		In Pure Number Strings, positive numeric values
//		are NOT configured with leading or trailing plus
//		signs ('+'). Negative values on the other hand
//		are always designated by leading or trailing
//		minus sign ('-').
//
//		This parameter, 'leadingNumSymbols', controls
//		the positioning of minus signs for negative
//		numeric values within a	Number String.
//
//		When set to 'true', the current NumStrFormatSpec
//		instance will configure minus signs for negative
//		numbers at the beginning of, or on the left side
//		of, the numeric value. In these cases, the minus
//		sign is said to be configured as a leading minus
//		sign. This is the positioning format used in the
//		US, UK, Australia and most of Canada. In
//		addition, library functions in 'Go' and other
//		programming languages generally expect leading
//		minus signs for negative numbers.
//
//			Example Leading Minus Sign:
//				"-123.456"
//
//		When parameter 'leadingNumSymbols' is set to
//		'false', the current instance of NumStrFormatSpec
//		will configure minus signs for negative numbers
//		at the end of, or on the right side of, the
//		numeric value. With this positioning format, the
//		minus sign is said to be configured as a trailing
//		minus sign. This is the positioning format used
//		in France, Germany and many countries in the
//		European Union.
//
//			Example Trailing Minus Sign:
//				"123.456-"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedPureNumberStr(
	decSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedPureNumberStr()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setSignedPureNStrSpec(
			numStrFmtSpec,
			decSeparatorChars,
			leadingNumSymbols,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"numStrFmtSpec"))
}

//	SetZeroNumberFmtSpec
//
//	Deletes and replaces the Zero Number Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.zeroNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a zero numeric value, this formatting specification
//	will be applied.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.zeroNumberSign'.
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetZeroNumberFmtSpec(
	zeroNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetZeroNumberFmtSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setZeroNumberSignSpec(
			numStrFmtSpec,
			zeroNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"zeroNumberSign"))

}
