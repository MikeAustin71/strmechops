package strmech

import "sync"

type CharSearchResultsDto struct {
	FoundSearchTarget bool // Signals a successful search
	//                        outcome. If 'true' the Test
	//                        String character or characters
	//                        were found in the Target Search
	//                        String.

	FirstFoundTargetSearchStrIndex int // The first character position in
	//                                    the Target Search String occupied
	//                                    by the frist character in the Test
	//                                    String.

	LastFoundTargetSearchStrIndex int // The last character position in
	//                                   the Target Search String occupied
	//                                   by the last character in the Test
	//                                   String.

	TestStrDescription string // Optional description string describing
	//                            Test Characters used in the search

	CollectionTestObjIndex int // The collection index of the object
	//                            containing the Test String which was
	//                            found in Target Search String

	lock *sync.Mutex
}

func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResults.FoundSearchTarget = false

	charSearchResults.FirstFoundTargetSearchStrIndex = -1

	charSearchResults.lock.Unlock()

}
