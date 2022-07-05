package strmech

import "sync"

// charSearchTestConfigDtoAtom - Provides helper methods for type
// CharSearchTestConfigDto.
//
type charSearchTestConfigDtoAtom struct {
	lock *sync.Mutex
}

func (searchTestConfigAtom *charSearchTestConfigDtoAtom) empty(
	searchTestConfigDto *CharSearchTestConfigDto) {

	if searchTestConfigAtom.lock == nil {
		searchTestConfigAtom.lock = new(sync.Mutex)
	}

	searchTestConfigAtom.lock.Lock()

	defer searchTestConfigAtom.lock.Unlock()

	if searchTestConfigDto == nil {
		return
	}

	searchTestConfigDto.TestInputParametersName = ""

	searchTestConfigDto.TestStringName = ""

	searchTestConfigDto.TestStringLengthName = ""

	searchTestConfigDto.TestStringStartingIndex = -1

	searchTestConfigDto.TestStringStartingIndexName = ""

	searchTestConfigDto.TestStringDescription1 = ""

	searchTestConfigDto.TestStringDescription2 = ""

	searchTestConfigDto.CollectionTestObjIndex = -1

	searchTestConfigDto.NumValueType = NumValType.None()

	searchTestConfigDto.NumStrFormatType = NumStrFmtType.None()

	searchTestConfigDto.NumSymbolLocation = NumSymLocation.None()

	searchTestConfigDto.NumSymbolClass = NumSymClass.None()

	searchTestConfigDto.NumSignValue = NumSignVal.None()

	searchTestConfigDto.PrimaryNumSignPosition =
		NumSignSymPos.None()

	searchTestConfigDto.SecondaryNumSignPosition =
		NumSignSymPos.None()

	searchTestConfigDto.TextCharSearchType =
		CharSearchType.None()

	return
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoAtom.
//
func (searchTestConfigAtom charSearchTestConfigDtoAtom) ptr() *charSearchTestConfigDtoAtom {

	if searchTestConfigAtom.lock == nil {
		searchTestConfigAtom.lock = new(sync.Mutex)
	}

	searchTestConfigAtom.lock.Lock()

	defer searchTestConfigAtom.lock.Unlock()

	return &charSearchTestConfigDtoAtom{
		lock: new(sync.Mutex),
	}
}
