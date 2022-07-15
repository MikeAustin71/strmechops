package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextFormatterCollection struct {
	fmtCollection []TextFormatterDto
	// Text

	stdTextLineParamCollection []TextFmtParamsLineColumns
	// Standard Format Parameter Collection.
	// Provides standard text formats

	lock *sync.Mutex
}

// AddLine1Col - Adds Field Text and Format Parameters for
// Format Type 'Line1Column' Text. This method will assign
// Format Parameter to this 'Line1Column' Text Field based
// on the Standard Parameters for this Text Type.
//
// The standard parameters must be set before calling this
// method. See SetStandardLine1ColParams() for more
// information.
//
func (txtFmtCollection *TextFormatterCollection) AddLine1Col(
	column1FieldText string) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		_ = ePref.ErrPrefixDto{}.NewIEmpty(
		"",
		"TextFormatterCollection."+
			"AddLine1Col()",
		"")

	// lineColsFormatter TextFmtParamsLineColumns

	foundStdParms,
		stdLineColsFmt :=
		textFormatterCollectionElectron{}.ptr().
			findStdTxtLineParameters(
				txtFmtCollection,
				TxtFieldType.Line1Column())

	if !foundStdParms {

		panic(
			fmt.Sprintf("%v - Error\n"+
				"Could NOT locate Standard Text Line Parameter\n"+
				"for Text Field Type: %v\n",
				ePrefix.String(),
				TxtFieldType.Line1Column().String()))

	}

	newLine1Col := TextFormatterDto{
		FormatType: 0,
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer:     TextFieldSpacerDto{},
		BlankLine:  TextLineBlankDto{},
		SolidLine:  TextLineSolidLineDto{},
		LineColumns: TextLineColumnsDto{
			formatType:    TxtFieldType.Line1Column(),
			col1FieldText: column1FieldText,
			col2FieldText: "",
			col3FieldText: "",
			col4FieldText: "",
			col5FieldText: "",
			col6FieldText: "",
			col7FieldText: "",
			col8FieldText: "",
			fmtParameters: stdLineColsFmt,
			lock:          nil,
		},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newLine1Col)

}

// func (txtFmtCollection *TextFormatterCollection) SetStdLine1ColFormatParams()
