package strand

import (
	"reflect"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	type args struct {
		len int
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomBytes(tt.args.len)

			if (err != nil) != tt.wantErr {
				t.Errorf("RandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	type args struct {
		len int
	}

	tests := []struct {
		name    string
		args    args
		wantStr string
		wantErr bool
	}{
		// TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := RandomString(tt.args.len)

			if (err != nil) != tt.wantErr {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotStr != tt.wantStr {
				t.Errorf("RandomString() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}
