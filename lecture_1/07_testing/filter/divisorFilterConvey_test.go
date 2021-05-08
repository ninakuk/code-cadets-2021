package filter_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"//tocka znaci sve stavi u globalni name space, npr so je iz convey, ne moram pisat convey.So, al nije preporuka koristit

	"code-cadets-2021/lecture_1/07_testing/filter"
	//filter2 "stdlib/filter" - da se izbjegne name clash, oba bi referencirli s filter.xyz
)

// NOTE - Convey infix in the function name is here just to prevent a name
// clash with the method in `divisorFilter_test.go`
func TestConveyGetDivisibleFromRange(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {//procitaj o conveyu u dokumentaciji malo

			actualOutput, actualErr := filter.GetDivisibleFromRange(tc.inputStart, tc.inputEnd, tc.inputDivisor)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
