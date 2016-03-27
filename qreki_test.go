package qreki

import (
	"reflect"
	"testing"
	"time"
)

func TestNewQreki(t *testing.T) {
	testcases := []struct {
		in  string
		out Qreki
	}{
		// end of 2015, during 2016, first of 2017
		{"2015-12-31", Qreki{2015, 11, 21, false}},
		{"2016-01-01", Qreki{2015, 11, 22, false}},
		{"2016-02-07", Qreki{2015, 12, 29, false}},
		{"2016-02-08", Qreki{2016, 1, 1, false}},
		{"2016-03-08", Qreki{2016, 1, 30, false}},
		{"2016-03-09", Qreki{2016, 2, 1, false}},
		{"2016-04-06", Qreki{2016, 2, 29, false}},
		{"2016-04-07", Qreki{2016, 3, 1, false}},
		{"2016-05-06", Qreki{2016, 3, 30, false}},
		{"2016-05-07", Qreki{2016, 4, 1, false}},
		{"2016-06-04", Qreki{2016, 4, 29, false}},
		{"2016-06-05", Qreki{2016, 5, 1, false}},
		{"2016-07-03", Qreki{2016, 5, 29, false}},
		{"2016-07-04", Qreki{2016, 6, 1, false}},
		{"2016-08-02", Qreki{2016, 6, 30, false}},
		{"2016-08-03", Qreki{2016, 7, 1, false}},
		{"2016-08-31", Qreki{2016, 7, 29, false}},
		{"2016-09-01", Qreki{2016, 8, 1, false}},
		{"2016-09-30", Qreki{2016, 8, 30, false}},
		{"2016-10-01", Qreki{2016, 9, 1, false}},
		{"2016-10-30", Qreki{2016, 9, 30, false}},
		{"2016-10-31", Qreki{2016, 10, 1, false}},
		{"2016-11-28", Qreki{2016, 10, 29, false}},
		{"2016-11-29", Qreki{2016, 11, 1, false}},
		{"2016-12-28", Qreki{2016, 11, 30, false}},
		{"2016-12-29", Qreki{2016, 12, 1, false}},
		{"2017-01-01", Qreki{2016, 12, 4, false}},
		{"2017-01-27", Qreki{2016, 12, 30, false}},
		{"2017-01-28", Qreki{2017, 1, 1, false}},

		// leap month
		{"2017-06-23", Qreki{2017, 5, 29, false}},
		{"2017-06-24", Qreki{2017, 5, 1, true}},
		{"2017-07-22", Qreki{2017, 5, 29, true}},
		{"2017-07-23", Qreki{2017, 6, 1, false}},

		{"2033-12-21", Qreki{2033, 11, 30, false}},
		{"2033-12-22", Qreki{2033, 11, 1, true}},
		{"2034-01-19", Qreki{2033, 11, 29, true}},
		{"2034-01-20", Qreki{2033, 12, 1, false}},

		{"2020-01-01", Qreki{2019, 12, 7, false}},
		{"2030-01-01", Qreki{2029, 11, 28, false}},
	}

	for _, tc := range testcases {
		in, _ := time.Parse("2006-01-02", tc.in)
		got := NewQreki(in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}
