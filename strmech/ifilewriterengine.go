package strmech

type IFileWriterEngine interface {
	Close() error

	Write([]byte) (int, error)

	WriteTextOrNumbers(
		charsToWrite interface{},
		writeEndOfLineChars string,
		writeEndOfTextChars string,
		autoCloseOnExit bool,
		errorPrefix interface{}) (
		numOfBytesWritten int64,
		err error)
}
