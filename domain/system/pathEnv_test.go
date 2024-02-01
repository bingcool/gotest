package system

import "testing"

func TestGetStoragePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试",
			want: "/Users/huangzengbing/Documents/wwwroot/goTest/domain/storage",
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
