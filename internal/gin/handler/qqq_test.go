package handler

import (
	"math"
	"testing"
)

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
		{
			name: "positive numbers",
			args: args{
				a: 5,
				b: 3,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "negative numbers",
			args: args{
				a: -2,
				b: -3,
			},
			want:    -5,
			wantErr: false,
		},
		{
			name: "zero values",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "one positive one negative",
			args: args{
				a: 5,
				b: -3,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "max int test",
			args: args{
				a: math.MaxInt / 2,
				b: math.MaxInt / 2,
			},
			want:    math.MaxInt - 1,
			wantErr: false,
		},
		{
			name: "min int test",
			args: args{
				a: math.MinInt / 2,
				b: math.MinInt / 2,
			},
			want:    math.MinInt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name: "positive numbers",
			args: args{
				a: 5,
				b: 3,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "negative numbers",
			args: args{
				a: -2,
				b: -3,
			},
			want:    -5,
			wantErr: false,
		},
		{
			name: "zero values",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "one positive one negative",
			args: args{
				a: 5,
				b: -3,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "max int test",
			args: args{
				a: math.MaxInt / 2,
				b: math.MaxInt / 2,
			},
			want:    math.MaxInt - 1,
			wantErr: false,
		},
		{
			name: "min int test",
			args: args{
				a: math.MinInt / 2,
				b: math.MinInt / 2,
			},
			want:    math.MinInt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc1(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc3() = %v, want %v", got, tt.want)
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
		{
			name: "positive numbers",
			args: args{
				a: 5,
				b: 3,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "negative numbers",
			args: args{
				a: -2,
				b: -3,
			},
			want:    -5,
			wantErr: false,
		},
		{
			name: "zero values",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "one positive one negative",
			args: args{
				a: 5,
				b: -3,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "max int test",
			args: args{
				a: math.MaxInt / 2,
				b: math.MaxInt / 2,
			},
			want:    math.MaxInt - 1,
			wantErr: false,
		},
		{
			name: "min int test",
			args: args{
				a: math.MinInt / 2,
				b: math.MinInt / 2,
			},
			want:    math.MinInt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc2(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestFunc3(t *testing.T) {
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
		{
			name: "positive numbers",
			args: args{
				a: 5,
				b: 3,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "negative numbers",
			args: args{
				a: -2,
				b: -3,
			},
			want:    -5,
			wantErr: false,
		},
		{
			name: "zero values",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "one positive one negative",
			args: args{
				a: 5,
				b: -3,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "max int test",
			args: args{
				a: math.MaxInt / 2,
				b: math.MaxInt / 2,
			},
			want:    math.MaxInt - 1,
			wantErr: false,
		},
		{
			name: "min int test",
			args: args{
				a: math.MinInt / 2,
				b: math.MinInt / 2,
			},
			want:    math.MinInt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TestFunc3(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestFunc3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TestFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
