package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtSpecAtom - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrFormatSpec and proceeds to reset the
// data values for all member variables to their
// initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'SignedNumFmtSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	SignedNumFmtSpec           *NumStrFormatSpec
//	   - A pointer to an instance of NumStrFormatSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (numStrFmtSpecAtom *numStrFmtSpecAtom) empty(
	signedNumFmtSpec *NumStrFormatSpec) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec == nil {
		return
	}

	signedNumFmtSpec.decSeparator.Empty()

	signedNumFmtSpec.intSeparatorSpec.Empty()

	signedNumFmtSpec.numberSymbolsGroup.Empty()

	signedNumFmtSpec.numberFieldSpec.Empty()
}

//	equal
//
//	Receives a pointer to two instances of
//	NumStrFormatSpec and proceeds to compare their member
//	variables in order to determine if they are
//	equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both instances
//	are equal in all respects, this flag is set to 'true'.
//	Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmtSpec1			*NumStrFormatSpec
//
//		An instance of NumStrFormatSpec. Internal member
//		variables from 'signedNumFmtSpec1' will be
//		compared to those of 'signedNumFmtSpec2' to
//		determine if both instances are equivalent.
//
//
//	signedNumFmtSpec2			*NumStrFormatSpec
//
//		An instance of NumStrFormatSpec. Internal member
//		variables from 'signedNumFmtSpec2' will be
//		compared to those of 'signedNumFmtSpec1' to
//		determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the comparison of 'signedNumFmtSpec1' and
//		'signedNumFmtSpec2' shows that all internal
//		member variables are equivalent, this method
//		will return a boolean value of 'true'.
//
//		If the two instances are NOT equal, this method
//		will return a boolean value of 'false' to the
//		calling function.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) equal(
	signedNumFmtSpec1 *NumStrFormatSpec,
	signedNumFmtSpec2 *NumStrFormatSpec) bool {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec1 == nil ||
		signedNumFmtSpec2 == nil {
		return false
	}

	if !signedNumFmtSpec1.decSeparator.Equal(
		&signedNumFmtSpec2.decSeparator) {

		return false
	}

	areEqual,
		_ := signedNumFmtSpec1.intSeparatorSpec.Equal(
		&signedNumFmtSpec2.intSeparatorSpec,
		nil)

	if !areEqual {

		return false
	}

	if !signedNumFmtSpec1.numberSymbolsGroup.Equal(
		&signedNumFmtSpec2.numberSymbolsGroup) {

		return false
	}

	if !signedNumFmtSpec1.numberFieldSpec.Equal(
		&signedNumFmtSpec2.numberFieldSpec) {

		return false
	}

	return true
}

// setDecimalSeparatorParams - Deletes and resets the member
// variable data value for 'NumStrFormatSpec.decSeparator'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		This instance contains the member variable
//		'signedNumFmt.decSeparator' which will be reset
//		to the value of the decimal separator specified
//		by input parameter, 'decSeparator'.
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
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setDecimalSeparatorParams(
	signedNumFmt *NumStrFormatSpec,
	decSeparator []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setDecimalSeparatorParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.decSeparator.Empty()

	err = signedNumFmt.decSeparator.SetDecimalSeparatorRunes(
		decSeparator,
		ePrefix.XCpy(
			"signedNumFmt.decSeparator<-"+
				"decSeparator"))

	return err
}

// setDecimalSeparatorSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.decSeparator'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'DecimalSeparatorSpec' and
// copies the member variable data values to
// 'signedNumFmt.decSeparator'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	decSeparatorSpec			DecimalSeparatorSpec
//
//		An instance of DecimalSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.decSeparator'.
//
//		In the US, UK, Australia and most of Canada, the
//		Decimal Separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setDecimalSeparatorSpec(
	signedNumFmt *NumStrFormatSpec,
	decSeparatorSpec DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setDecimalSeparatorSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.decSeparator.Empty()

	err = new(decimalSepSpecNanobot).
		copyDecimalSeparator(
			&signedNumFmt.decSeparator,
			&decSeparatorSpec,
			ePrefix.XCpy(
				"signedNumFmt.decSeparatorSpec<-"+
					"decSeparatorSpec"))

	return err
}

// setIntegerGroupingParams - Deletes and resets the member
// variable data value for 'NumStrFormatSpec.intSeparatorSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.intSeparatorSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	intGroupingChars			[]rune
//		One or more characters used to separate groups of
//		integer digits. This separator is also known as the
//		'thousands' separator. It is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the
//		standard integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length array
//		and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setIntegerGroupingParams(
	signedNumFmt *NumStrFormatSpec,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setIntegerGroupingParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.intSeparatorSpec.Empty()

	err = signedNumFmt.intSeparatorSpec.SetIntGroupEnumRunes(
		intGroupingType,
		intGroupingChars,
		ePrefix.XCpy(
			"signedNumFmt.intSeparatorSpec<-"))

	return err
}

// setIntegerGroupingSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.intSeparatorSpec'
// contained in the instance of NumStrFormatSpec
// passed as an input parameter.
//
// This method receives an instance of IntegerSeparatorSpec
// and copies the member variable data values to
// 'signedNumFmt.intSeparatorSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		All the member variable data values in this instance
//		will be deleted and reset to the values contained
//		in the Integer Separator Specification. supplied by
//		input parameter, 'intSeparatorSpec'.
//
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		An instance of IntegerSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.intSeparatorSpec'.
//
//		In the United States, the Integer Group Specification
//		character is a comma (',') and integer grouped in
//		thousands.
//			United States Example: 1,000,000,000
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setIntegerGroupingSpec(
	signedNumFmt *NumStrFormatSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setIntegerGroupingSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	return signedNumFmt.intSeparatorSpec.CopyIn(
		&intSeparatorSpec,
		ePrefix.XCpy(
			"signedNumFmt.intSeparatorSpec<-"+
				"intSeparatorSpec"))

}

// setNegativeNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.negativeNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.negativeNumberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	negativeNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.negativeNumberSign'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setNegativeNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNegativeNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeNumberSign.IsNOP() {

		signedNumFmt.numberSymbolsGroup.negativeNumberSign.SetNOP()

		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.numberSymbolsGroup.negativeNumberSign,
			&negativeNumberSign,
			ePrefix.XCpy(
				"signedNumFmt.numberSymbolsGroup."+
					"negativeNumberSign<-"+
					"negativeNumberSign"))

	return err
}

// setNStrFmtComponents
//
// Deletes and resets the member variable data values
// stored in the instance of NumStrFormatSpec passed
// as input parameter 'signedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec instance passed by input parameter
//	'numStrFmtSpec'  will be deleted and replaced by
//	values generated from the other input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec					*NumStrFormatSpec
//
//	A pointer to a NumStrFormatSpec instance. All  member
//	variable data fields in this object will be replaced
//	by data values configured from the input parameter
//	described below.
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
//	numberSymbolsGroup					NumStrNumberSymbolGroup
//
//		This instance of NumStrNumberSymbolGroup contains the
//		Number Symbol Specifications for negative numeric
//		values, positive numeric values and zero numeric
//		values.
//
//	numberFieldSpec					NumStrNumberFieldSpec
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
//	errPrefDto						*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setNStrFmtComponents(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	numberSymbolsSpec NumStrNumberSymbolGroup,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNStrFmtComponents()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = numStrFmtSpec.decSeparator.CopyIn(
		&decSeparatorSpec,
		ePrefix.XCpy(
			"decSeparatorSpec->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.intSeparatorSpec.CopyIn(
		&intSeparatorSpec,
		ePrefix.XCpy(
			"intSeparatorSpec->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.numberFieldSpec.CopyIn(
		&numberFieldSpec,
		ePrefix.XCpy(
			"numberFieldSpec->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.numberSymbolsGroup.CopyIn(
		&numberSymbolsSpec,
		ePrefix.XCpy(
			"numStrFmtSpec.numberSymbols<-"+
				"numberSymbols"))

	return err
}

// setNumberFieldParams - Deletes and resets the member variable data
// value for 'NumStrFormatSpec.numberFieldSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.numberFieldSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	fieldLength					int
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//		An enumeration which specifies the justification of the
//		numeric value string within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//		          TextJustify(0).Left()
//		          TextJustify(0).Right()
//		          TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//		          TxtJustify.Left()
//		          TxtJustify.Right()
//		          TxtJustify.Center()
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setNumberFieldParams(
	signedNumFmt *NumStrFormatSpec,
	fieldLength int,
	fieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNumberFieldParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	numFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        fieldLength,
		fieldJustification: fieldJustification,
	}

	err = numFieldSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"numFieldSpec"))

	if err != nil {
		return err
	}

	signedNumFmt.numberFieldSpec.Empty()

	err = signedNumFmt.numberFieldSpec.SetFieldSpec(
		fieldLength,
		fieldJustification,
		ePrefix.XCpy(
			"signedNumFmt.numberFieldSpec<-"))

	if err != nil {
		return err
	}

	return err
}

// setNumberFieldSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.numberFieldSpec'
// contained in the instance of NumStrFormatSpec
// passed as an input parameter.
//
// This method receives an instance of
// 'NumStrNumberFieldSpec' and copies the member variable
// data values to 'signedNumFmt.numberFieldSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		All the member variable data values in this instance
//		will be deleted and reset to the values contained
//		in the Number Field Specification supplied by
//		input parameter, 'numberFieldSpec'.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		An instance of NumStrNumberFieldSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.numberFieldSpec'.
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
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setNumberFieldSpec(
	signedNumFmt *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.numberFieldSpec.Empty()

	return new(numStrNumberFieldSpecNanobot).
		copyNStrNumberFieldSpec(
			&signedNumFmt.numberFieldSpec,
			&numberFieldSpec,
			ePrefix.XCpy(
				"signedNumFmt<-"+
					"numberFieldSpec"))
}

//	setNumberSymbolGroup
//
//	Receives a pointer to an instance of NumStrFormatSpec
//	and proceeds to reconfigure all the internal member
//	variable data values based on an instance of
//	NumStrNumberSymbolGroup passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the Number Symbol data field
//	contained in input parameter 'numStrFmtSpec'
//	will be deleted and replaced by Number Symbol
//	specifications passed through input parameter,
//	'numberSymbolSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
//
//	numberSymbolSpec			NumStrNumberSymbolGroup
//
//		An instance of NumStrNumberSymbolGroup which is used
//		to reconfigure the NumStrNumberSymbolGroup member
//		variable contained in input parameter,
//		'numStrFmtSpec'.
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setNumberSymbolGroup(
	numStrFmtSpec *NumStrFormatSpec,
	numberSymbolSpec NumStrNumberSymbolGroup,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNegativeNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = numStrFmtSpec.numberSymbolsGroup.CopyIn(
		&numberSymbolSpec,
		ePrefix.XCpy(
			"numStrFmtSpec.numberSymbols<-"))

	return err
}

// setPositiveNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.positiveNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.positiveNumberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	positiveNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.positiveNumberSign'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setPositiveNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setPositiveNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if positiveNumberSign.IsNOP() {

		signedNumFmt.numberSymbolsGroup.positiveNumberSign.SetNOP()

		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.numberSymbolsGroup.positiveNumberSign,
			&positiveNumberSign,
			ePrefix.XCpy(
				"signedNumFmt.numberSymbols"+
					".positiveNumberSign<-"+
					"positiveNumberSign"))

	return err
}

// setZeroNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.zeroNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.zeroNumberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	zeroNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.zeroNumberSign'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrFmtSpecAtom *numStrFmtSpecAtom) setZeroNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrFmtSpecAtom.lock == nil {
		numStrFmtSpecAtom.lock = new(sync.Mutex)
	}

	numStrFmtSpecAtom.lock.Lock()

	defer numStrFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setZeroNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroNumberSign.IsNOP() {

		signedNumFmt.numberSymbolsGroup.zeroNumberSign.SetNOP()

		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.numberSymbolsGroup.zeroNumberSign,
			&zeroNumberSign,
			ePrefix.XCpy(
				"signedNumFmt.zeroNumberSign<-"+
					"zeroNumberSign"))

	return err
}
