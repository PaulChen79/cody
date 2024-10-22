package handler

import (
	"testing"
)

func TestTestFunc1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc1(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestFunc(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestFunc2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc2(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
