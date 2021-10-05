package goyyds

import (
	"context"
	"github.com/goyyds/goyyds/v1/plugins/cmd"
	"github.com/goyyds/goyyds/v1/server"
	"reflect"
	"testing"
)

func Test_newOptions(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		{"aaa",
			args{opts: []Option{}},
			Options{
				Name:        "GOYYDS",
				Cmd:         cmd.DefaultCmd,
				Server:      server.DefaultServer,
				BeforeStart: nil,
				BeforeStop:  nil,
				AfterStart:  nil,
				AfterStop:   nil,
				Context:     context.Background(),
				Signal:      true,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
