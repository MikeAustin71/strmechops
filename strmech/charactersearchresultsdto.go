package strmech

import "sync"

type CharSearchResultsDto struct {
	FoundSearchTarget bool // Signals a successful search
	//                        outcome. If 'true' the Test
	//                        String character or characters
	//                        were found in the Target Search
	//                        String.

	TargetStringLength int // Actual Full Length of the Target Search
	//                        String

	TargetStringSearchLength int // The actual number of characters within
	//                              Target Search String that are included
	//                              in the actual search operation.

	TargetStringStartingSearchIndex int // The index at which the search
	//                               	   operation commenced.

	TargetStringFirstFoundIndex int // The first character position in the
	//                                 Target Search String occupied by the
	//                                 first character in the Test String.

	TargetStringLastFoundIndex int // The last character position in
	//                                the Target Search String occupied
	//                                by the last character in the Test
	//                                String.

	TargetStringLastSearchIndex int // The index in Target Search String
	//                                 occupied by the last character searched.
	//                                 If the Search Target was found this
	//                                 value is equal to the 'TargetStringLastFoundIndex'.
	//                                 If the Search Target was NOT found this
	//                                 value is equal to the 'TargetStringStartingSearchIndex'

	TargetStringDescription1 string // First optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TargetStringDescription2 string // Second Optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TestStrLength int // The Length of the Test String used in
	//                   this search operation.

	TestStrStartingIndex int // The Test String beginning index used in the
	//                          search operation. Usually zero.

	TestStringFirstFoundIndex int // The index number in Test String of the first
	//                               test character to be located in the Target
	//                               Search String.

	TestStrLastFoundIndex int // The index number in the Test String occupied
	//                           by the last Test Character to be located in the
	//                           Target Search String.

	TestStrDescription1 string // Optional description string describing
	//                            Test Characters used in this search

	TestStrDescription2 string // Optional description string describing
	//                            Test Characters used in the search

	CollectionTestObjIndex int // The collection index of the object
	//                            containing the Test String which was
	//                            found in Target Search String

	NumSignValue NumericSignValueType // An enumeration value classifying the
	//                                   number sign.
	//                                     NumSignVal.None()
	//                                     NumSignVal.Negative()
	//                                     NumSignVal.Zero()
	//                                     NumSignVal.Positive()

	NumSignPosition NumSignSymbolPosition // Used in searches involving
	//                                       positive and negative number
	//                                       signs.
	//                                        NSignSymPos.None()
	//                                        NSignSymPos.Before()
	//                                        NSignSymPos.After()
	//                                        NSignSymPos.BeforeAndAfter()

	SearchType CharacterSearchType // An enumeration value signaling
	//                                the type of search algorithm which
	//                                was used to conduct this search.
	//                                 CharSearchType.None()
	//                                 CharSearchType.LinearTargetStartingIndex() - Default
	//                                 CharSearchType.SingleTargetChar()
	//                                 CharSearchType.LinearEndOfString()

	lock *sync.Mutex
}

func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResults.FoundSearchTarget = false

	charSearchResults.TargetStringLength = -1

	charSearchResults.TargetStringStartingSearchIndex = -1

	charSearchResults.TargetStringFirstFoundIndex = -1

	charSearchResults.TargetStringLastFoundIndex = -1

	charSearchResults.TargetStringLastSearchIndex = -1

	charSearchResults.TargetStringDescription1 = ""

	charSearchResults.TargetStringDescription2 = ""

	charSearchResults.TestStrLength = -1

	charSearchResults.TestStrStartingIndex = -1

	charSearchResults.TestStringFirstFoundIndex = -1

	charSearchResults.TestStrLastFoundIndex = -1

	charSearchResults.TestStrDescription1 = ""

	charSearchResults.TestStrDescription2 = ""

	charSearchResults.CollectionTestObjIndex = -1

	charSearchResults.NumSignValue = NumSignVal.None()

	charSearchResults.NumSignPosition = NSignSymPos.None()

	// Set to default Search Type
	charSearchResults.SearchType =
		CharSearchType.LinearTargetStartingIndex()

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil
}
