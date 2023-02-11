package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
	"time"
)

func TestTextLineSpecAverageTime_AddDurationEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecAverageTime_AddDurationEvent_000100()",
		"")

	mockDurations := make([]int64, 5)

	mockDurations[0] = 15415462
	mockDurations[1] = 13605311
	mockDurations[2] = 17321123
	mockDurations[3] = 18681792
	mockDurations[4] = 16415443

	var err error

	var avgTimer TextLineSpecAverageTime

	for i := 0; i < len(mockDurations); i++ {

		err = avgTimer.AddDurationEvent(
			time.Duration(mockDurations[i]),
			ePrefix.XCpy(
				fmt.Sprintf("mockDurations[%v]= %v\n",
					i, mockDurations[i])))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

	}

	var avgTimeTxt string

	avgTimeTxt,
		err = avgTimer.GetFormattedText(
		ePrefix.XCpy("avgTimeTxt<-avgTimer"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := "Test #1 Verify Return Of Avg Time String\n"

	if len(avgTimeTxt) == 0 {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: Average Time Text has a length of zero!\n",
			ePrefix.String(),
			testName)

		return

	}

	var actualAvgDuration,
		actualMaximumTimeDuration,
		actualMinimumTimeDuration,
		actualNumberOfTimingEvents,
		expectedAvgDuration,
		expectedMaximumTimeDuration,
		expectedMinimumTimeDuration,
		expectedNumberOfTimingEvents int64

	actualAvgDuration,
		actualMaximumTimeDuration,
		actualMinimumTimeDuration,
		actualNumberOfTimingEvents,
		err = avgTimer.CalcAvgTimeDuration(
		ePrefix.XCpy("avgTimer"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedAvgDuration = 16287826

	testName = "Test #2 Verify Average Time Duration\n"

	if actualAvgDuration != expectedAvgDuration {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualAvgDuration NOT EQUAL TO expectedAvgDuration\n"+
			"    actualAvgDuration = '%v'\n"+
			"expectedAvgDuration   = '%v'\n",
			ePrefix.String(),
			testName,
			actualAvgDuration,
			expectedAvgDuration)

		return

	}

	expectedMaximumTimeDuration = 18681792

	testName = "Test #3 Verify Maximum Time Duration\n"

	if actualMaximumTimeDuration != expectedMaximumTimeDuration {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualMaximumTimeDuration NOT EQUAL TO expectedMaximumTimeDuration\n"+
			"    actualMaximumTimeDuration = '%v'\n"+
			"expectedMaximumTimeDuration   = '%v'\n",
			ePrefix.String(),
			testName,
			actualMaximumTimeDuration,
			expectedMaximumTimeDuration)

		return

	}

	expectedMinimumTimeDuration = 13605311

	testName = "Test #4 Verify Minimum Time Duration\n"

	if actualMinimumTimeDuration != expectedMinimumTimeDuration {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualMinimumTimeDuration NOT EQUAL TO expectedMinimumTimeDuration\n"+
			"    actualMinimumTimeDuration = '%v'\n"+
			"expectedMinimumTimeDuration   = '%v'\n",
			ePrefix.String(),
			testName,
			actualMinimumTimeDuration,
			expectedMinimumTimeDuration)

		return

	}

	expectedNumberOfTimingEvents = 5

	testName = "Test #5 Verify Number Of TimingEvents\n"

	if actualNumberOfTimingEvents != expectedNumberOfTimingEvents {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberOfTimingEvents NOT EQUAL TO expectedNumberOfTimingEvents\n"+
			"    actualNumberOfTimingEvents = '%v'\n"+
			"expectedNumberOfTimingEvents   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberOfTimingEvents,
			expectedNumberOfTimingEvents)

		return

	}

	var actualAllocatedAvgDuration,
		actualAllocatedMaxDuration,
		actualAllocatedMinDuration TimeDurationDto

	expectedAllocatedAvgDuration := TimeDurationDto{
		TotalNanoseconds:     16287826,
		NumberOfDays:         0,
		NumberOfHours:        0,
		NumberOfMinutes:      0,
		NumberOfSeconds:      0,
		NumberOfMilliseconds: 16,
		NumberOfMicroseconds: 287,
		NumberOfNanoseconds:  826,
	}

	actualAllocatedAvgDuration,
		actualAllocatedMaxDuration,
		actualAllocatedMinDuration,
		actualNumberOfTimingEvents,
		err = avgTimer.CalcAvgTimeDurationDetail(
		ePrefix.XCpy("avgTimer"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test #6 Verify AllocatedAvgDuration\n"

	if !actualAllocatedAvgDuration.Equal(&expectedAllocatedAvgDuration) {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualAllocatedAvgDuration NOT EQUAL TO expectedAllocatedAvgDuration\n",
			ePrefix.String(),
			testName)

		return

	}

	expectedAllocatedMaxDuration := TimeDurationDto{
		TotalNanoseconds:     18681792,
		NumberOfDays:         0,
		NumberOfHours:        0,
		NumberOfMinutes:      0,
		NumberOfSeconds:      0,
		NumberOfMilliseconds: 18,
		NumberOfMicroseconds: 681,
		NumberOfNanoseconds:  792,
	}

	testName = "Test #7 Verify AllocatedMaxDuration\n"

	if !actualAllocatedMaxDuration.Equal(
		&expectedAllocatedMaxDuration) {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualAllocatedMaxDuration NOT EQUAL TO expectedAllocatedMaxDuration\n",
			ePrefix.String(),
			testName)

		return

	}

	expectedAllocatedMinDuration := TimeDurationDto{
		TotalNanoseconds:     13605311,
		NumberOfDays:         0,
		NumberOfHours:        0,
		NumberOfMinutes:      0,
		NumberOfSeconds:      0,
		NumberOfMilliseconds: 13,
		NumberOfMicroseconds: 605,
		NumberOfNanoseconds:  311,
	}

	testName = "Test #8 Verify AllocatedMinDuration\n"

	if !actualAllocatedMinDuration.Equal(
		&expectedAllocatedMinDuration) {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualAllocatedMinDuration NOT EQUAL TO expectedAllocatedMinDuration\n",
			ePrefix.String(),
			testName)

		return

	}

	return
}
