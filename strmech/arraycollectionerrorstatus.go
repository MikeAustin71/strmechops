package strmech

// ArrayColErrorStatus
//
// This type is used to convey the detailed error status
// after completion of an operation related to an array
// or a collection of objects.
type ArrayColErrorStatus struct {
	IsProcessingError bool
	// When set to 'true', this parameter signals
	// that an error was encountered during an
	// array or object collection processing
	// operation. In this case an appropriate
	// error message describing the error is
	// recorded in data element 'ProcessingError'.

	IsIndexOutOfBounds bool
	// When set to 'true', this parameter signals
	// that the index value used to access the array
	// or object collection was less than zero or
	// greater than the last index in the
	// array/collection.

	IsArrayCollectionEmpty bool
	// When set to 'true', this parameter signals
	// that array or objects collections is empty.

	IsErrorFree bool
	// When set to 'true', this parameter signals that
	// no errors were encountered in the most recent
	// array or collection operation. This also means
	// that data element 'ProcessingError' is set to
	// 'nil'.

	ProcessingError error
	// If no errors were encountered in the most recent
	// array or object collection processing operation,
	// this error parameter will be set to nil.
	//
	// If errors are encountered during an array or
	// object collection processing operation, this
	// error Type will encapsulate an appropriate error
	// message.
	//
	// If an error prefix was specified as an input
	// parameter, this returned error message will
	// incorporate the method chain and text passed by
	// the error prefix input parameter.
	//
	// The Error Prefix text will be prefixed or
	// attached to the beginning of the error message.

}
