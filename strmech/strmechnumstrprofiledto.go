package strmech

// NumStrProfileDto - This type is used to encapsulate information
// on strings of numeric digits which are extracted from larger
// strings.
//
type NumStrProfileDto struct {
	TargetStr          string //  The original target string which is scanned for a number string.
	StartIndex         int    //  The starting index in 'TargetStr' from which the number string search was initiated.
	LeadingSignIndex   int    //  The string index of a leading sign in 'NumStr' below. If a leading sign character is NOT present in 'NumStr' this value is set to -1.
	LeadingSignChar    string //  If a leading sign character (plus '+' or minus '-') exists in data field 'NumStr' (below), it is stored in this string.
	FirstNumCharIndex  int    //  The index in 'TargetStr' (above) where the first character of the extracted number string is located.
	NextTargetStrIndex int    //  The index of the next character in 'TargetStr' immediately following the extracted number string. If no number string is identified or the next index in 'TargetStr' is invalid, the value is set to -1.
	NumStrLen          int    //  The length of the extracted number string. If a number string is not located, this value is set to -1.
	NumStr             string //  The number string extracted from 'TargetStr'.
}

// New - Creates and returns a new instance of NumStrProfileDto
// which is properly initialized.
func (exNumDto NumStrProfileDto) New() NumStrProfileDto {

	newDto := NumStrProfileDto{}
	newDto.TargetStr = ""
	newDto.StartIndex = -1
	newDto.LeadingSignIndex = -1
	newDto.LeadingSignChar = ""
	newDto.FirstNumCharIndex = -1
	newDto.NextTargetStrIndex = -1
	newDto.NumStrLen = 0
	newDto.NumStr = ""
	return newDto
}
