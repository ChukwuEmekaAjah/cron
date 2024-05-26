package cron

import "testing"

func TestParseSchedule(t *testing.T) {
	type args struct {
		schedule string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Schedule 1",
			args: args{schedule: "* 10 * * *"},
		},
		{
			name: "4:05 am every day",
			args: args{schedule: "5 4 * * *"},
		},
		{
			name: "even numbered hours",
			args: args{schedule: "5 4,6,8,10 * * *"},
		},
		{
			name: "hour ranges",
			args: args{schedule: "5 4-10 * * *"},
		},
		{
			name: "wednesday,thursday and friday",
			args: args{schedule: "5 4-10 * * 3-5"},
		},
		{
			name: "recurring every 5 hours at minute 5",
			args: args{schedule: "5 3/5 * * *"},
		},
		{
			name: "recurring every 3 days starting at day 2 at minute 5",
			args: args{schedule: "5 5 2/3 * *"},
		},
		{
			name: "recurring every 2 hours starting at hour 0 at minute 5",
			args: args{schedule: "5 */2 * * *"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSchedule(tt.args.schedule)
		})
	}
}
