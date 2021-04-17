package strmech

// SortStrLengthHighestToLowest - Uses to perform two level sort
// on string arrays. The strings are first sorted by string length
// (greatest length to Lowest length) and then by alphabetic sort.
//
// This type is designed to be used in conjunction with 'sort.Sort()'
// Reference the Go Sort Package:
//      https://golang.org/pkg/sort/#Sort
//
// Example Usage:
//   badChars := []string {
//    "aaaaa",
//    "bbbbb",
//    "cccccccccc",
//    "z",
//    "fffffffffff",
//    "xx",
//    "ddddddddd",
//    "eeeeeeeeeee" }
//
//     SortStrLengthLowestToHighest(badChars)
//
//     Output:
//
//        ================================
//        Sort by Length Highest To Lowest
//        Ordered List
//        ================================
//
//        1. fffffffffff
//        2. eeeeeeeeeee
//        3. cccccccccc
//        4. ddddddddd
//        5. bbbbb
//        6. aaaaa
//        7. xx
//        8. z
//
type SortStrLengthHighestToLowest []string

// Len - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Len() int {
	return len(sortStrLenHigh)
}

// Swap - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Swap(i, j int) {
	sortStrLenHigh[i], sortStrLenHigh[j] = sortStrLenHigh[j], sortStrLenHigh[i]
}

// Less - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Less(i, j int) bool {

	lenI := len(sortStrLenHigh[i])
	lenJ := len(sortStrLenHigh[j])
	if lenI == lenJ {
		return sortStrLenHigh[i] > sortStrLenHigh[j]
	}

	return lenI > lenJ
}

// SortStrLengthLowestToHighest - Uses to perform two level sort
// on string arrays. The strings are first sorted by string length
// (smallest length to greatest length) and then by alphabetic sort.
//
// This type is designed to be used in conjunction with 'sort.Sort()'
// Reference the Go Sort Package:
//      https://golang.org/pkg/sort/#Sort
//
// Example Usage:
//      badChars := []string {
//    "aaaaa",
//    "bbbbb",
//    "cccccccccc",
//    "z",
//    "fffffffffff",
//    "xx",
//    "ddddddddd",
//    "eeeeeeeeeee" }
//
//     SortStrLengthLowestToHighest(badChars)
//
//     Output:
//
//       ================================
//       Sort by Length Lowest To Highest
//       Ordered List
//       ================================
//       1. z
//       2. xx
//       3. aaaaa
//       4. bbbbb
//       5. ddddddddd
//       6. cccccccccc
//       7. eeeeeeeeeee
//       8. fffffffffff
//
type SortStrLengthLowestToHighest []string

// Len - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Len() int {
	return len(sortStrLenLow)
}

// Swap - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Swap(i, j int) {
	sortStrLenLow[i], sortStrLenLow[j] = sortStrLenLow[j], sortStrLenLow[i]
}

// Less - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Less(i, j int) bool {

	lenI := len(sortStrLenLow[i])
	lenJ := len(sortStrLenLow[j])
	if lenI == lenJ {
		return sortStrLenLow[i] < sortStrLenLow[j]
	}

	return lenI < lenJ
}
