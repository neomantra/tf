package gotf

import (
	"reflect"
	"testing"
	"time"
)

func TestEpochToTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{{
		name:    "Test convert seconds",
		args:    args{"1637356129"},
		want:    time.Unix(1637356129, 0),
		wantErr: false,
	}, {
		name:    "Test convert milliseconds",
		args:    args{"1637356129123"},
		want:    time.Unix(1637356129, 123*1000000),
		wantErr: false,
	}, {
		name:    "Test convert microseconds",
		args:    args{"1637356129123456"},
		want:    time.Unix(1637356129, 123456*1000),
		wantErr: false,
	}, {
		name:    "Test convert nanoseconds",
		args:    args{"1637356129123456789"},
		want:    time.Unix(1637356129, 123456789),
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EpochToTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("EpochToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EpochToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTimes(t *testing.T) {
	type args struct {
		str         string
		outFormat   string
		globalMatch bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "Simple passthrough",
		args: args{"askldjaklsdj aslkdj askljda s . 12321312. ", time.RFC3339, true},
		want: "askldjaklsdj aslkdj askljda s . 12321312. ",
	}, {
		name: "Simple timveval",
		args: args{"1637356129 ", time.RFC3339, true},
		want: "2021-11-19T16:08:49-05:00 ",
	}, {
		name: "Global match off",
		args: args{"1637356129 1637356130", time.RFC3339, false},
		want: "2021-11-19T16:08:49-05:00 1637356130",
	}, {
		name: "Global match on",
		args: args{"1637356129 1637356130", time.RFC3339, true},
		want: "2021-11-19T16:08:49-05:00 2021-11-19T16:08:50-05:00",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ConvertTimes(tt.args.str, tt.args.outFormat, tt.args.globalMatch); got != tt.want {
				t.Errorf("ConvertTimes() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
