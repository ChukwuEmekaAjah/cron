package cron

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	now := time.Now()
	type args struct {
		from     time.Time
		schedule string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "every minute",
			args: args{
				from:     time.Now(),
				schedule: "* *  * * *",
			},
			want:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()+1, 0, 0, time.Local),
			wantErr: false,
		}, // * 0/12 * * *
		{
			name: "every 12 hours",
			args: args{
				from:     time.Date(now.Year(), now.Month(), 27, 23, 22, 0, 0, time.Local),
				schedule: "* 0/12  * * *",
			},
			want:    time.Date(time.Now().Year(), time.Now().Month(), 28, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Next(tt.args.from, tt.args.schedule)
			if (err != nil) != tt.wantErr {
				fmt.Println("error is ", err)
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
