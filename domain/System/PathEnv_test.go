package System

import "testing"

func TestGetStoragePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试",
			want: "/Users/huangzengbing/Documents/wwwroot/goTest/domain/Storage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStoragePath(); got != tt.want {
				t.Errorf("GetStoragePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyTestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试",
			args: args{a: 4, b: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyTestAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("TestAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
