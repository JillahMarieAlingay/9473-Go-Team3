package main

import "testing"

func Test_is_scheduled_date(t *testing.T) {
	type args struct {
		list []string
		date string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_scheduled_date(tt.args.list, tt.args.date); got != tt.want {
				t.Errorf("is_scheduled_date() = %v, want %v", got, tt.want)
			}
		})
	}
}
