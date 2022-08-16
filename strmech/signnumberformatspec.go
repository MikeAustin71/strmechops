package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// SignedNumberFormatSpec - The Signed Number Format
// Specification is used to format signed numeric values
// as number strings. This type contains all the
// specifications and parameters required to format a
// numeric value for text displays.
type SignedNumberFormatSpec struct {
	decSeparator DecimalSeparatorSpec
	// Contains the decimal separator character or
	// characters which will separate integer and
	// fractional digits in a floating point number

	intGroupingType NumStrIntegerGroupingSpec
	// This parameter specifies the type of integer
	// grouping and integer separator characters
	// which will be applied to the number string
	// formatting operation.

	roundingSpec NumStrRoundingSpec
	// Controls the rounding algorithm applied to
	// floating point numbers.

	positiveNumberSign NumStrPositiveNumberSignSpec
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a positive number sign character or
	// characters for positive numeric values using a
	// Number String Positive Number Sign
	// Specification.

	negativeNumberSign NumStrNegativeNumberSignSpec
	// The Number String Negative Number Sign
	// Specification is used to configure negative
	// number sign symbols for negative numeric values
	// formatted and displayed in number stings.

	numberField NumStrNumberFieldSpec
	// This Number String Number Field Specification
	// contains the field length and text justification
	// parameter necessary to display a numeric value
	// within a number field for display as a number
	// string.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrNumberFieldSpec ('incomingNStrNumFieldSpec')
// to the data fields of the current NumStrNumberFieldSpec
// instance ('nStrNumberFieldSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrNumberFieldSpec
// instance ('nStrNumberFieldSpec') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//			 incomingNStrNumFieldSpec *NumStrNumberFieldSpec
//			   - A pointer to an instance of NumStrNumberFieldSpec.
//			     This method will NOT change the values of internal member
//			     variables contained in this instance.
//
//			     All data values in this NumStrNumberFieldSpec instance
//			     will be copied to current NumStrNumberFieldSpec
//			     instance ('nStrNumberFieldSpec').
//
//			     If parameter 'incomingNStrNumFieldSpec' is determined to
//			     be invalid, an error will be returned.
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (signedNumFmtSpec *SignedNumberFormatSpec) CopyIn(
	incomingSignedNumFmt *SignedNumberFormatSpec,
	errorPrefix interface{}) error {

	if signedNumFmtSpec.lock == nil {
		signedNumFmtSpec.lock = new(sync.Mutex)
	}

	signedNumFmtSpec.lock.Lock()

	defer signedNumFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"SignedNumberFormatSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(signedNumFmtSpecNanobot).
		copySignedNumberFormatSpec(
			signedNumFmtSpec,
			incomingSignedNumFmt,
			ePrefix.XCpy(
				"signedNumFmtSpec<-"+
					"incomingSignedNumFmt"))
}

// signedNumFmtSpecNanobot - This type provides
// helper methods for SignedNumberFormatSpec
type signedNumFmtSpecNanobot struct {
	lock *sync.Mutex
}

// copySignedNumberFormatSpec - Copies all data from input parameter
// 'sourceSignedNumFmtSpec' to input parameter
// 'destinationSignedNumFmtSpec'. Both instances are of type
// SignedNumberFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationSignedNumFmtSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceSignedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationSignedNumFmtSpec  *SignedNumberFormatSpec
//	   - A pointer to a SignedNumberFormatSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceSignedNumFmtSpec'.
//
//	     'destinationSignedNumFmtSpec' is the destination for this
//	     copy operation.
//
//
//	sourceSignedNumFmtSpec       *SignedNumberFormatSpec
//	   - A pointer to another SignedNumberFormatSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationSignedNumFmtSpec'.
//
//	     'sourceSignedNumFmtSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceSignedNumFmtSpec'.
//
//
//	errPrefDto                     *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *signedNumFmtSpecNanobot) copySignedNumberFormatSpec(
	destinationSignedNumFmtSpec *SignedNumberFormatSpec,
	sourceSignedNumFmtSpec *SignedNumberFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumberFieldSpecNanobot.lock == nil {
		nStrNumberFieldSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"signedNumFmtSpecNanobot."+
			"copySignedNumberFormatSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationSignedNumFmtSpec' is invalid!\n"+
			"'destinationSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceSignedNumFmtSpec' is invalid!\n"+
			"'sourceSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(signedNumberFormatSpecAtom).empty(
		destinationSignedNumFmtSpec)

	err = destinationSignedNumFmtSpec.decSeparator.CopyIn(
		&sourceSignedNumFmtSpec.decSeparator,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.decSeparator"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.intGroupingType.CopyIn(
		&sourceSignedNumFmtSpec.intGroupingType,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.intGroupingType"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.roundingSpec.CopyIn(
		&sourceSignedNumFmtSpec.roundingSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.roundingSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.positiveNumberSign.CopyIn(
		&sourceSignedNumFmtSpec.positiveNumberSign,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.positiveNumberSign"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.negativeNumberSign.CopyIn(
		&sourceSignedNumFmtSpec.negativeNumberSign,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.negativeNumberSign"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.numberField.CopyIn(
		&sourceSignedNumFmtSpec.numberField,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.numberField"+
				"<-sourceSignedNumFmtSpec"))

	return err
}

// setDecimalSeparator - Deletes and resets the member
// variable data value for 'SignedNumberFormatSpec.decSeparator'
// contained in the instance of SignedNumberFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		signedNumFmt               *SignedNumberFormatSpec
//		    - A pointer to an instance of SignedNumberFormatSpec.
//		      All the member variable data values in this instance
//		      will be deleted and reset according to the data
//		      extracted from the following input parameters.
//
//		 decSeparator              []rune
//		    - This rune array contains the decimal separator
//	       character or characters used to separate integer
//	       digits from fractional digits in floating point
//		      number strings.
//
//	       In the United States, the decimal separator is
//	       referred to as the decimal point.
//
//		 errPrefDto                 *ePref.ErrPrefixDto
//		    - This object encapsulates an error prefix string which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods listed
//		      as a function chain.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      Type ErrPrefixDto is included in the 'errpref' software
//		      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *signedNumFmtSpecNanobot) setDecimalSeparator(
	signedNumFmt *SignedNumberFormatSpec,
	decSeparator []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumberFieldSpecNanobot.lock == nil {
		nStrNumberFieldSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"signedNumFmtSpecNanobot."+
			"setDecimalSeparator()",
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

	if len(decSeparator) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decSeparator' is invalid!\n"+
			"'decSeparator' is empty and has a zero (0) length.\n",
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

// signedNumberFormatSpecAtom - This type provides
// helper methods for SignedNumberFormatSpec
type signedNumberFormatSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// SignedNumberFormatSpec and proceeds to reset the
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
//	SignedNumFmtSpec           *SignedNumberFormatSpec
//	   - A pointer to an instance of SignedNumberFormatSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (signedNumFmtSpecAtom *signedNumberFormatSpecAtom) empty(
	signedNumFmtSpec *SignedNumberFormatSpec) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec == nil {
		return
	}

	signedNumFmtSpec.decSeparator.Empty()

	signedNumFmtSpec.intGroupingType.Empty()

	signedNumFmtSpec.roundingSpec.Empty()

	signedNumFmtSpec.positiveNumberSign.Empty()

	signedNumFmtSpec.negativeNumberSign.Empty()

	signedNumFmtSpec.numberField.Empty()

}

// equal - Receives a pointer to two instances of
// SignedNumberFormatSpec and proceeds to compare their
// member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'. Otherwise,
// this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmtSpec1    *SignedNumberFormatSpec
//	   - An instance of SignedNumberFormatSpec.
//	     Internal member variables from 'signedNumFmtSpec1'
//	     will be compared to those of 'signedNumFmtSpec2' to
//	     determine if both instances are equivalent.
//
//
//	signedNumFmtSpec2    *SignedNumberFormatSpec
//	   - An instance of SignedNumberFormatSpec.
//	     Internal member variables from 'signedNumFmtSpec2'
//	     will be compared to those of 'signedNumFmtSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'signedNumFmtSpec1' and
//	     'signedNumFmtSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (signedNumFmtSpecAtom *signedNumberFormatSpecAtom) equal(
	signedNumFmtSpec1 *SignedNumberFormatSpec,
	signedNumFmtSpec2 *SignedNumberFormatSpec) bool {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec1 == nil ||
		signedNumFmtSpec2 == nil {
		return false
	}

	if !signedNumFmtSpec1.decSeparator.Equal(
		&signedNumFmtSpec2.decSeparator) {

		return false
	}

	if !signedNumFmtSpec1.intGroupingType.Equal(
		&signedNumFmtSpec2.intGroupingType) {

		return false
	}

	if !signedNumFmtSpec1.roundingSpec.Equal(
		&signedNumFmtSpec2.roundingSpec) {

		return false
	}

	if !signedNumFmtSpec1.positiveNumberSign.Equal(
		&signedNumFmtSpec2.positiveNumberSign) {

		return false
	}

	if !signedNumFmtSpec1.negativeNumberSign.Equal(
		&signedNumFmtSpec2.negativeNumberSign) {

		return false
	}

	if !signedNumFmtSpec1.numberField.Equal(
		&signedNumFmtSpec2.numberField) {

		return false
	}

	return true
}
