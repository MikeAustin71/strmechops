package strmech

// DataFieldProfileDto - This type is used to encapsulate information
// related to an extracted data field string.
//
type DataFieldProfileDto struct {
	TargetStr                      string                         //  The string from which the data field  is extracted.
	TargetStrLength                int                            //  Length of 'TargetStr'.
	TargetStrStartIndex            int                            //  The index with in 'TargetStr' from which the search for a data field was initiated.
	TargetStrLastGoodIndex         int                            //  Last valid index in target string which is less than the target string length and is NOT an 'End Of Field' or 'End Of Line' Delimiter.
	LeadingKeyWordDelimiter        string                         //  The Leading Key Word Delimiter which is used to identify the beginning of the field search.
	LeadingKeyWordDelimiterIndex   int                            //  Index of the found Leading Key Word Delimiter
	DataFieldStr                   string                         //  The extracted data field string
	DataFieldIndex                 int                            //  The index in 'TargetStr' where the data field begins.
	DataFieldLength                int                            //  The length of the extracted data field string.
	DataFieldTrailingDelimiter     string                         //  The trailing character which marked the end of the data field. A zero value indicates end of string encountered.
	DataFieldTrailingDelimiterType DataFieldTrailingDelimiterType //  A constant or enumeration type used to describe the type of delimiter used to mark the end of a data field.
	NextTargetStrIndex             int                            //  The index in 'TargetStr' immediately following the extracted data field.
	CommentDelimiter               string                         //  If a Comment Delimiter is detected it is stored here.
	CommentDelimiterIndex          int                            //  If a Comment Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
	EndOfLineDelimiter             string                         //  If an End-Of-Line Delimiter is detected it is captured and stored here.
	EndOfLineDelimiterIndex        int                            //  If an End-Of-Line Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
}

// ConvertToErrorState - Prepares the current DataFieldProfileDto instance
// for return as part of an error or null state condition. All references
// to the data field are zeroed.
//
func (dfProfile *DataFieldProfileDto) ConvertToErrorState() {
	dfProfile.LeadingKeyWordDelimiter = ""
	dfProfile.LeadingKeyWordDelimiterIndex = -1
	dfProfile.DataFieldStr = ""
	dfProfile.DataFieldIndex = -1
	dfProfile.DataFieldLength = 0
	dfProfile.NextTargetStrIndex = -1

}

// New - Creates and returns a new instance of DataFieldProfileDto
// containing properly initialized internal data fields.
func (dfProfile DataFieldProfileDto) New() DataFieldProfileDto {
	newDataDto := DataFieldProfileDto{}
	newDataDto.TargetStr = ""
	newDataDto.TargetStrStartIndex = -1
	newDataDto.LeadingKeyWordDelimiter = ""
	newDataDto.LeadingKeyWordDelimiterIndex = -1
	newDataDto.DataFieldStr = ""
	newDataDto.DataFieldIndex = -1
	newDataDto.DataFieldLength = 0
	newDataDto.DataFieldTrailingDelimiter = ""
	newDataDto.DataFieldTrailingDelimiterType = DfTrailDelimiter.Unknown()
	newDataDto.NextTargetStrIndex = -1
	newDataDto.CommentDelimiter = ""
	newDataDto.CommentDelimiterIndex = -1
	newDataDto.EndOfLineDelimiter = ""
	newDataDto.EndOfLineDelimiterIndex = -1
	return newDataDto
}
