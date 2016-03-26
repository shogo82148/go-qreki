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
		{"2016-01-01", Qreki{11, 22, false}},
		{"2020-01-01", Qreki{12, 7, false}},
		{"2030-01-01", Qreki{11, 28, false}},
	}

	for _, tc := range testcases {
		in, _ := time.Parse("2006-01-02", tc.in)
		got := NewQreki(in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}
