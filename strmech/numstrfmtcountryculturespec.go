package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type NumStrFmtCountryCultureSpec struct {
	IdNo uint64
	//	Optional
	//	An identification number used to differentiate
	//	and track multiple Country Culture Specification
	//	objects.

	IdString string
	//	Optional
	//	An identification string of text characters
	//	used to differentiate and track multiple
	//	Country Culture Specification objects.

	Description string
	//	Optional
	//	This string contains descriptive text describing
	//	the Country Culture Specification instance.

	Tag string
	//	Optional
	//	This string contains descriptive text describing
	//	the Country Culture Specification instance.

	CountryIdNo uint64
	//	Optional
	//	A number identifying the specific country
	//	or culture specified by the current Country
	//	Culture Specification instance.

	CountryIdString string
	//	Optional
	//	A string of text characters identifying the
	//	specific country or culture specified by the
	//	current Country	Culture Specification instance.

	CountryDescription string
	//	Optional
	//	A string of characters providing a	narrative
	//	text description of the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryTag string
	//	Optional
	//	A string containing a brief text description
	//	 of the country or culture identified by the
	//	 current Country Culture Specification instance.

	CountryCultureName string
	//	Required
	//	The ISO 3166 name of the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryCultureOfficialStateName string
	//	Optional
	//	The ISO 3166 official state name of the country
	//	or culture identified by the current Country
	//	Culture Specification instance.

	CountryAbbreviatedName string
	//	Optional
	//	An abbreviated name for the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryAlternateNames []string
	//	Optional
	//	An alternate or additional name for the country
	//	or culture identified by the current Country
	//	Culture Specification instance.

	CountryCodeTwoChar string
	//	Optional
	//	The unique ISO 3166-1 alpha-2 Two Character code
	//	identifying the country or culture associated
	//	with the current Country Culture Specification
	//	instance.

	CountryCodeThreeChar string
	//	Optional
	//	The unique ISO 3166-1 alpha-3 Three Character code
	//	identifying the country or culture associated with
	//	the current Country Culture Specification instance.

	CountryCodeNumber string
	//	Optional
	//	The official ISO 3166-1 numeric code identifier
	//	for the country or culture associated with the
	//	current Country Culture Specification instance.

	CurrencyCode string
	//	Optional
	//	The official ISO 4217 currency code associated
	//	with the country or culture identified by the
	//	current Country Culture Specification instance.

	CurrencyCodeNo string
	//	Optional
	//	The official ISO 4217 currency code number
	//	associated with the country or culture identified
	//	by the current Country Culture Specification
	//	instance.

	CurrencyName string
	//	Optional
	//	The official ISO 4217 currency name associated
	//	with the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	CurrencySymbols []rune
	//	Optional. Not required for Currency Number
	//	String Formatting.
	//
	//	The official ISO 4217 currency symbol or symbols
	//	for the country or culture identified by the
	//	current Country Culture Specification instance.

	CurrencyDecimalDigits uint
	//	The number of fractional digits typically used
	//	in formatting currency numeric values

	MinorCurrencyName string
	//	Optional
	//	The name of the minor currency associated
	//	with the country or culture identified by
	//	the current Country Culture Specification
	//	instance. In the United States, the minor
	//	currency name is "Cents"

	MinorCurrencySymbols []rune
	//	Optional
	//	The Minor Currency symbol or symbols. In
	//	the United States, the minor currency
	//	name is "Cents" and the minor currency
	//	symbol is "¢".

	CurrencyNumStrFormat NumStrFormatSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	This NumStrFormatSpec instance
	//	contains all the parameters necessary
	//	to produce a formatted Currency
	//	Number String.

	SignedNumStrFormat NumStrFormatSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	This NumStrFormatSpec instance
	//	contains all the parameters necessary
	//	to produce a formatted Signed
	//	Number String.

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrFmtCountryCultureSpec ('sourceCountryCultureSpec')
//	to the data fields of the current
//	NumStrFmtCountryCultureSpec instance
//	('nStrFmtCountryCultureSpec').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values in the current
//	NumStrFmtCountryCultureSpec instance
//	('nStrFmtCountryCultureSpec') will	be deleted and
//	replaced.
//
//	No data validation is performed on input parameter,
//	'incomingSignedNumFmt'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceCountryCultureSpec		*NumStrFmtCountryCultureSpec
//
//		A pointer to an instance of NumStrFmtCountryCultureSpec.
//		This method will NOT change the values of internal
//		member variables contained in this instance.
//
//		All data values in this NumStrFmtCountryCultureSpec
//		instance will be copied to current
//		NumStrFmtCountryCultureSpec instance
//		('nStrFmtCountryCultureSpec').
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) CopyIn(
	sourceCountryCultureSpec *NumStrFmtCountryCultureSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecMech).
		copyCountryCulture(
			nStrFmtCountryCultureSpec,
			sourceCountryCultureSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	NumStrFmtCountryCultureSpec instance.
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of
//		the current NumStrFmtCountryCultureSpec instance.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) CopyOut(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNStrFmtCountryCulture NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return newNStrFmtCountryCulture, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		copyCountryCulture(
			&newNStrFmtCountryCulture,
			nStrFmtCountryCultureSpec,
			ePrefix.XCpy(
				"newNStrFmtCountryCulture<-"))

	return newNStrFmtCountryCulture, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrFmtCountryCultureSpec to their
//	initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of NumStrFmtCountryCultureSpec.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) Empty() {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	new(numStrFmtCountryCultureSpecAtom).empty(
		nStrFmtCountryCultureSpec)

	nStrFmtCountryCultureSpec.lock.Unlock()

	nStrFmtCountryCultureSpec.lock = nil
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
//	incomingCountryCulture		*NumStrFmtCountryCultureSpec
//
//		A pointer to an external instance of
//		NumStrFormatSpec. The internal member variable
//		data values in this instance will be compared to
//		those in the current instance of
//		NumStrFmtCountryCultureSpec. The results of this
//		comparison will be returned to the calling
//		function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the internal member variable data values
//		contained in input parameter 'incomingCountryCulture'
//		are equivalent in all respects to those contained
//		in the current instance of
//		'NumStrFmtCountryCultureSpec', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) Equal(
	incomingCountryCulture *NumStrFmtCountryCultureSpec) bool {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	return new(numStrFmtCountryCultureSpecAtom).equal(
		nStrFmtCountryCultureSpec,
		incomingCountryCulture)
}

//	New
//
//	Creates and returns a new instance of
//	NumStrFmtCountryCultureSpec.
//
//	The returned instance will be configured with the
//	currency and signed number string formatters. In
//	addition, the Country/Culture name will also be
//	populated.
//
//	While the returned instance is sufficient for
//	generating formatted number strings for currency
//	and signed numbers, all the description member
//	variables will remain empty. These description
//	member variables are public variables and may be
//	populated manually by the user as needed.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureName			string
//
//		It is recommended that this parameter be set
//		to the ISO 3166 name of the country or culture.
//
//	currencyNumStrFormat		NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		currency format associated with the
//		designated Country or Culture.
//
//	signedNumStrFormat			NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		signed number format associated with the
//		designated Country or Culture.
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a partially
//		completed instance of NumStrFmtCountryCultureSpec
//		will be returned.
//
//		While the returned instance is sufficient for
//		generating formatted number strings for currency
//		and signed numbers, all the description member
//		variables will remain empty. These description
//		member variables are public variables and may be
//		populated manually by the user as needed.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) New(
	countryCultureName string,
	currencyNumStrFormat NumStrFormatSpec,
	signedNumStrFormat NumStrFormatSpec,
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	fmtCountryCultureAtom := numStrFmtCountryCultureSpecAtom{}

	err = fmtCountryCultureAtom.copyNumStrFormatSpec(
		&newCountryCultureSpec.CurrencyNumStrFormat,
		&currencyNumStrFormat,
		ePrefix.XCpy(
			"newCountryCultureSpec<-"+
				"currencyNumStrFormat"))

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = fmtCountryCultureAtom.copyNumStrFormatSpec(
		&newCountryCultureSpec.SignedNumStrFormat,
		&signedNumStrFormat,
		ePrefix.XCpy(
			"newCountryCultureSpec<-"+
				"signedNumStrFormat"))

	if err != nil {
		return newCountryCultureSpec, err
	}

	newCountryCultureSpec.CountryCultureName =
		countryCultureName

	return newCountryCultureSpec, err
}

//	NewFrance
//
//	Returns a new instance of NumStrFmtCountryCultureSpec
//	formatted with country and Number String formatting
//	parameters commonly used in France.
//
//	These same formatting specifications are also used by
//	various other member countries in the European Union.
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a new
//		instance of NumStrFmtCountryCultureSpec
//		will be returned configured with country and
//		Number String formatting specifications
//		commonly used in France.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) NewFrance(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"NewFrance()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		setCountryFrance(
			&newCountryCultureSpec,
			ePrefix.XCpy(
				"newCountryCultureSpec<-"))

	return newCountryCultureSpec, err
}

//	NewGermany
//
//	Returns a new instance of NumStrFmtCountryCultureSpec
//	formatted with country and Number String formatting
//	parameters commonly used in Germany.
//
//	These same formatting specifications are also used by
//	many member countries in the European Union.
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a new
//		instance of NumStrFmtCountryCultureSpec
//		will be returned configured with country and
//		Number String formatting specifications
//		commonly used in Germany.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) NewGermany(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"NewGermany()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		setCountryGermany(
			&newCountryCultureSpec,
			ePrefix.XCpy(
				"newCountryCultureSpec<-"))

	return newCountryCultureSpec, err
}

//	NewUK
//
//	Returns a new instance of NumStrFmtCountryCultureSpec
//	formatted with country and Number String formatting
//	parameters commonly used in the UK (United Kingdom).
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a new
//		instance of NumStrFmtCountryCultureSpec
//		will be returned configured with country and
//		Number String formatting specifications
//		commonly used in the UK (United Kingdom).
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) NewUK(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"NewUK()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		setCountryUK(
			&newCountryCultureSpec,
			ePrefix.XCpy(
				"newCountryCultureSpec<-"))

	return newCountryCultureSpec, err
}

//	NewUS
//
//	Returns a new instance of NumStrFmtCountryCultureSpec
//	formatted with country and Number String formatting
//	parameters commonly used in the US (United States).
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
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a new
//		instance of NumStrFmtCountryCultureSpec
//		will be returned configured with country and
//		Number String formatting specifications
//		commonly used in the US (United States).
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) NewUS(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"NewUS()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		setCountryUS(
			&newCountryCultureSpec,
			ePrefix.XCpy(
				"newCountryCultureSpec<-"))

	return newCountryCultureSpec, err
}

//	SetCurrencyNumberFieldSpec
//
//	Deletes and resets the Currency Number Field
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'currencyNumFieldSpec'.
//
//	For the current instance of NumStrFmtCountryCultureSpec,
//	this method will modify the internal member variable:
//
//	NumStrFmtCountryCultureSpec.CurrencyNumStrFormat.numberFieldSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencyNumFieldSpec		NumStrNumberFieldSpec
//
//		Contains the Number String Number Field Specification
//		which will be copied to the Currency Number String
//		Number Field Specification encapsulated within
//		the current instance of NumStrFmtCountryCultureSpec.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetCurrencyNumberFieldSpec(
	currencyNumFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetCurrencyNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecElectron).
		copyNumberFieldSpec(
			&nStrFmtCountryCultureSpec.CurrencyNumStrFormat.numberFieldSpec,
			&currencyNumFieldSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-currencyNumFieldSpec"))
}

//	SetCurrencyNumStrFormatSpec
//
//	Deletes and resets the Currency Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
//	The Currency Number String Format Specification is
//	used to generate number strings containing currency
//	numeric values using the string format associated
//	with the designated country or culture.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values in the Currency Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec will be deleted and reset
//	to new values.
//
//	No data validation is performed on input parameter,
//	'currencyNumStrFormat'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencyNumStrFormat		NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		currency format associated with the
//		designated Country or Culture.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetCurrencyNumStrFormatSpec(
	currencyNumStrFormat NumStrFormatSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetCurrencyNumStrFormatSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecAtom).
		copyNumStrFormatSpec(
			&nStrFmtCountryCultureSpec.CurrencyNumStrFormat,
			&currencyNumStrFormat,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-"+
					"currencyNumStrFormat"))
}

//	SetSignedNumberFieldSpec
//
//	Deletes and resets the Signed Number Field
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'signedNumFieldSpec'.
//
//	For the current instance of NumStrFmtCountryCultureSpec,
//	this method will modify the internal member variable:
//
//	NumStrFmtCountryCultureSpec.SignedNumStrFormat.numberFieldSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFieldSpec		NumStrNumberFieldSpec
//
//		Contains the Number String Number Field Specification
//		which will be copied to the Signed Number String
//		Number Field Specification encapsulated within
//		the current instance of NumStrFmtCountryCultureSpec.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetSignedNumberFieldSpec(
	signedNumFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetSignedNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecElectron).
		copyNumberFieldSpec(
			&nStrFmtCountryCultureSpec.SignedNumStrFormat.numberFieldSpec,
			&signedNumFieldSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-signedNumFieldSpec"))
}

//	SetSignedNumStrFormatSpec
//
//	Deletes and resets the Signed Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
//	The Signed Number String Format Specification is
//	used to generate number strings containing signed
//	numeric values using the string format associated
//	with the designated country or culture.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values in the Signed Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec will be deleted and reset
//	to new values.
//
//	No data validation is performed on input parameter,
//	'signedNumStrFormat'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumStrFormat		NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		signed format associated with the
//		designated Country or Culture.
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
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetSignedNumStrFormatSpec(
	signedNumStrFormat NumStrFormatSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetSignedNumStrFormatSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecAtom).
		copyNumStrFormatSpec(
			&nStrFmtCountryCultureSpec.SignedNumStrFormat,
			&signedNumStrFormat,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-"+
					"signedNumStrFormat"))
}
