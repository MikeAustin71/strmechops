package strmech

import "sync"

type CharSearchResultsDto struct {
	FoundSearchTarget bool // Signals a successful search
	//                        outcome. If 'true' the Test
	//                        String character or characters
	//                        were found in the Target Search
	//                        String.

	TargetStringStartingIndex int // The index at which the search
	//                                    operation commenced.

	FirstFoundTargetSearchStrIndex int // The first character position in
	//                                    the Target Search String occupied
	//                                    by the first character in the Test
	//                                    String.

	LastFoundTargetSearchStrIndex int // The last character position in
	//                                   the Target Search String occupied
	//                                   by the last character in the Test
	//                                   String.

	LastFoundTestStrIndex int // The index number of the last Test String
	//                           Character found in Target String.

	TestStrDescription1 string // Optional description string describing
	//                            Test Characters used in the search

	TestStrDescription2 string // Optional description string describing
	//                            Test Characters used in the search

	CollectionTestObjIndex int // The collection index of the object
	//                            containing the Test String which was
	//                            found in Target Search String

	SearchType CharacterSearchType // An enumeration value signaling
	//                                the type of search algorithm which
	//                                was used to conduct this search.

	lock *sync.Mutex
}

func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResults.FoundSearchTarget = false

	charSearchResults.FirstFoundTargetSearchStrIndex = -1

	charSearchResults.LastFoundTargetSearchStrIndex = -1

	charSearchResults.LastFoundTestStrIndex = -1

	charSearchResults.TestStrDescription1 = ""

	charSearchResults.TestStrDescription2 = ""

	charSearchResults.CollectionTestObjIndex = -1

	// Set to default Search Type
	charSearchResults.SearchType =
		CharSearchType.LinearTargetStartingIndex()

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil
}
