package strmech

import (
	"strings"
	"sync"
	"time"
)

type TextFieldSpecDateTime struct {
	dateTime time.Time // The content of the datetime text.
	fieldLen int       // The length of the text field in which
	//               //  the text label will be positioned.
	dateTimeFormat string // Holds the format parameters used to
	//                        //  format the datetime for text presentation.
	textJustification TextJustify // The specification which controls
	//                            //  how the datetime text will be positioned
	//                            //  within the text field: 'Left', 'Right'
	//                            //  or 'Center'.
	textLineReader *strings.Reader
	lock           *sync.Mutex
}
