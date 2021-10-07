package addr

import (
	"testing"
)

func TestIsLocal(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"127.0.0.1",
			args{addr: "127.0.0.1"},
			true,
		},
		{
			"127.0.0.1",
			args{addr: "127.0.0.1:9999"},
			true,
		},
		{
			"localhost",
			args{addr: "localhost"},
			true,
		},
		{
			"localhost",
			args{addr: "localhost:8080"},
			true,
		},
		{
			"other",
			args{addr: "255.255.255.255"},
			false,
		},
		{
			"other",
			args{addr: "47.104.1.1"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLocal(tt.args.addr); got != tt.want {
				t.Errorf("IsLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendPrivateBlocks(t *testing.T) {
	tests := []struct {
		name   string
		addr   string
		expect bool
	}{
		{
			name:   "aaa",
			addr:   "127.0.0.1",
			expect: false,
		},
		{
			name:   "bbb",
			addr:   "47.104.15.255",
			expect: false,
		},
		{
			name:   "bbb",
			addr:   "47.104.16.19",
			expect: true,
		},
		{
			name:   "ccc",
			addr:   "10.1.1.1",
			expect: true,
		},
	}
	AppendPrivateBlocks("47.104.16.0/24")
	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			res := isPrivateIp(tt.addr)
			if res != tt.expect {
				t.Fatalf("expected %t got %t", tt.expect, res)
			}
		})
	}
}
