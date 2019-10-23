package strand

import (
	"reflect"
	"testing"
)

func TestRandomBytess(t *testing.T) {
	type args struct {
		len int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Verify RandomBytes length result",
			args: args{
				len: 32,
			},
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := RandomBytes(tt.args.len)
			got := len(res)

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
		want    int
		wantErr bool
	}{
		{
			name: "Verify RandomString length result",
			args: args{
				len: 32,
			},
			want: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := RandomString(tt.args.len)
			gotStr := len(res)

			if (err != nil) != tt.wantErr {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotStr != tt.want {
				t.Errorf("RandomString() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}
