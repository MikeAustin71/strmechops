package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock this mutex before accessing any of these maps.
var preProcessPathCodeMapsLock sync.Mutex

var mPreProcessPathStringToCode = map[string]PreProcessPathCode{
	"None":          PreProcessPathCode(0),
	"PathSeparator": PreProcessPathCode(1),
	"AbsolutePath":  PreProcessPathCode(2),
}

var mPreProcessPathLwrCaseStringToCode = map[string]PreProcessPathCode{
	"none":          PreProcessPathCode(0),
	"pathseparator": PreProcessPathCode(1),
	"absolutepath":  PreProcessPathCode(2),
}

var mPreProcessPathCodeToString = map[PreProcessPathCode]string{
	PreProcessPathCode(0): "None",
	PreProcessPathCode(1): "PathSeparator",
	PreProcessPathCode(2): "AbsolutePath",
}

var mValidPreProcessPathCodeToString = map[PreProcessPathCode]string{
	PreProcessPathCode(1): "PathSeparator",
	PreProcessPathCode(2): "AbsolutePath",
}

type PreProcessPathCode int

var preProcessPathCodeLock sync.Mutex

// None - Take No Action
func (preProcPathCde PreProcessPathCode) None() PreProcessPathCode {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	return PreProcessPathCode(0)
}

// PathSeparator - Convert path separators to the default value for the
// host operating system.
func (preProcPathCde PreProcessPathCode) PathSeparator() PreProcessPathCode {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	return PreProcessPathCode(1)
}

// AbsolutePath - Convert path string to an absolute path.
func (preProcPathCde PreProcessPathCode) AbsolutePath() PreProcessPathCode {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	return PreProcessPathCode(2)
}

// IsValid
//
// Returns an error value signaling whether the current
// PreProcessPathCode value is valid.
//
// If the current PreProcessPathCode is invalid, this
// method returns an error containing an appropriate
// error message.
//
// If the current PreProcessPathCode is valid, this
// method returns an error value of 'nil'.
//
// Be advised, the enumeration value "None" is considered
// an INVALID selection for 'PreProcessPathCode'.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (preProcPathCde PreProcessPathCode) IsValid() error {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	preProcessPathCodeMapsLock.Lock()

	defer preProcessPathCodeMapsLock.Unlock()

	ePrefix := "PreProcessPathCode.IsValid() "

	_, ok := mValidPreProcessPathCodeToString[preProcPathCde]

	var err error

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: Pre Process Code is invalid!\n"+
			"PreProcessPathCode Integer = %v\n"+
			"PreProcessPathCode String = %v\n",
			ePrefix,
			int(preProcPathCde),
			preProcPathCde.String())

		return err
	}

	return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'PreProcessPathCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//	string - The string label or description for the current enumeration
//	         value. If, the PreProcessPathCode value is invalid, this
//	         method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= PreProcessPathCode(0).AbsolutePath()
//	str := t.String()
//	    str is now equal to "AbsolutePath"
func (preProcPathCde PreProcessPathCode) String() string {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	preProcessPathCodeMapsLock.Lock()

	defer preProcessPathCodeMapsLock.Unlock()

	label, ok := mPreProcessPathCodeToString[preProcPathCde]

	if !ok {
		return ""
	}

	return label
}

// Value - Returns the value of the PathFileTypeCode
// instance as type PathFileTypeCode.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (preProcPathCde PreProcessPathCode) Value() PreProcessPathCode {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	return preProcPathCde
}

// ParseString
//
// Receives a string and returns an instance of
// PreProcessPathCode associated with that string.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (preProcPathCde PreProcessPathCode) ParseString(
	valueString string,
	caseSensitive bool) (PreProcessPathCode, error) {

	preProcessPathCodeLock.Lock()

	defer preProcessPathCodeLock.Unlock()

	preProcessPathCodeMapsLock.Lock()

	defer preProcessPathCodeMapsLock.Unlock()

	ePrefix := "PreProcessPathCode.ParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return PreProcessPathCode(0),
			fmt.Errorf("%v\n"+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
				"valueString='%v'\n",
				ePrefix,
				valueString)
	}

	var ok bool

	var preProcessPathCode PreProcessPathCode

	if caseSensitive {

		preProcessPathCode, ok = mPreProcessPathStringToCode[valueString]

		if !ok {

			return PreProcessPathCode(0),
				fmt.Errorf("%v\n"+
					"Invalid Permission Code!\n",
					ePrefix)
		}

	} else {

		valueString = strings.ToLower(valueString)

		preProcessPathCode, ok = mPreProcessPathLwrCaseStringToCode[valueString]

		if !ok {
			return PreProcessPathCode(0),
				fmt.Errorf("%v\n"+
					"Invalid Permission Code!\n",
					ePrefix)
		}

	}

	return preProcessPathCode, nil
}

// PreProcPathCode
//
// This public global variable allows easy access to
// the enumerations of the PreProcessPathCode type using
// the dot operator.
//
//	Example:
//
//		PreProcPathCode.None()
//		PreProcPathCode.PathSeparator()
//		PreProcPathCode.AbsolutePath()
const PreProcPathCode = PreProcessPathCode(0)
