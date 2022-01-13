package nfs

import "testing"

func TestCopyNew(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"aaa", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyNew(); (err != nil) != tt.wantErr {
				t.Errorf("CopyNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
