package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// BaseTypeDto
//
// A data transfer object used to encapsulate elementary
// base types and transmit them between methods as
// required.
//
// Conversion methods using empty interfaces typically
// populate only one of the base data types with
// information extracted from an empty interface
// object.
type BaseTypeDto struct {
	IsAByteArrayDto             bool
	AByteArrayDto               ByteArrayDto
	AByteArrayDtoLength         int
	AByteArrayDtoSourceDataType string
	AByteArrayDtoDesc1          string
	AByteArrayDtoDesc2          string

	IsARuneArrayDto             bool
	ARuneArrayDto               RuneArrayDto
	ARuneArrayDtoLength         int
	ARuneArrayDtoSourceDataType string
	ARuneArrayDtoDesc1          string
	ARuneArrayDtoDesc2          string

	IsAString             bool
	AString               string
	AStringLength         int
	AStringSourceDataType string
	AStringDesc1          string
	AStringDesc2          string

	IsAStringArrayDto             bool
	AStringArrayDto               StringArrayDto
	AStringArrayDtoLength         int
	AStringArrayDtoSourceDataType string
	AStringArrayDtoDesc1          string
	AStringArrayDtoDesc2          string

	lock *sync.Mutex
}

// GetStringFromResult
//
// Analyzes the results and data elements for the current
// instance of BaseTypeDto and then proceeds to convert
// and those data elements as a single string.
//
// If the encapsulated data element resolves as a string
// array, that string array will be consolidated into a
// single string by inserting input parameter
// 'strArrayElementSeparator' between each element in the
// string array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strArrayElementSeparator	string
//
//		If the active or populated internal data element
//		for the current BaseTypeDto instance is a string
//		array, the elements of that string array will
//		be concatenated to generate the returned
//		'convertedString'. If the length of
//		'strArrayElementSeparator' is greater than zero
//		('0'), 'strArrayElementSeparator' will be
//		inserted between each member element of the
//		string array when creating the returned
//		'convertedString'.
//
//
//	isEmptyErrorMessage			string
//
//		This string contains the text of an error
//		message which will be returned if all data
//		elements in the current BaseTypeDto instance
//		are empty.
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
//	convertedString				string
//
//		This parameter returns the string generated
//		by the active data elements contained in the
//		current instance of BaseTypeDto.
//
//		For valid instances of BaseTypeDto, data
//		elements for one of the four basic types,
//		string, string array, rune array or byte
//		array, will be populated. The populated
//		data type will therefore be converted to a
//		string and returned via this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (bTypeDto *BaseTypeDto) GetStringFromResult(
	strArrayElementSeparator string,
	isEmptyErrorMessage string,
	errorPrefix interface{}) (
	convertedString string,
	err error) {

	if bTypeDto.lock == nil {
		bTypeDto.lock = new(sync.Mutex)
	}

	bTypeDto.lock.Lock()

	defer bTypeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"BaseTypeDto."+
			"GetStringFromResult()",
		"")

	if err != nil {
		return convertedString, err
	}

	if bTypeDto.IsAByteArrayDto == true {

		convertedString =
			string(bTypeDto.AByteArrayDto.ByteArray)

	} else if bTypeDto.IsARuneArrayDto == true {

		convertedString =
			string(bTypeDto.ARuneArrayDto.CharsArray)

	} else if bTypeDto.IsAString == true {

		convertedString = bTypeDto.AString

	} else if bTypeDto.IsAStringArrayDto == true {

		convertedString =
			bTypeDto.AStringArrayDto.
				ConcatenateStrings(strArrayElementSeparator)

	} else {

		if len(isEmptyErrorMessage) == 0 {

			isEmptyErrorMessage =
				"Error: The current 'BaseTypeDto' object is invalid.\n" +
					"The data elements contained in this 'BaseTypeDto'\n" +
					"instance are empty and have a zero length.\n"

		}

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"%v",
			ePrefix.String(),
			isEmptyErrorMessage)

	}

	return convertedString, err
}
