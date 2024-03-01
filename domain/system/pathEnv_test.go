package system

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGetStoragePath(t *testing.T) {
	nowTime := time.Now().Unix() + 11

	ticker1 := time.NewTicker(5 * time.Second)
	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 在改协程中阻塞等待，不影响其他协程的执行
			_ = <-t.C

			nowTime1 := time.Now().Unix()
			if nowTime1 > nowTime {
				t.Stop()
			} else {
				fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
			}

		}
	}(ticker1)

	println("Ticker start")
	time.Sleep(30 * time.Second)
	fmt.Println("ok")
}

func TestEventLoopSigtermSignal(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EventLoopSigtermSignal()
		})
	}
}

func TestGetCronPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCronPath(); got != tt.want {
				t.Errorf("GetCronPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDaemonPath(); got != tt.want {
				t.Errorf("GetDaemonPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMainPid(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMainPid(); got != tt.want {
				t.Errorf("GetMainPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRootDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootDir(); got != tt.want {
				t.Errorf("GetRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServerPidFile(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerPidFile(); got != tt.want {
				t.Errorf("GetServerPidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStorageGinLog(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStorageGinLog(); got != tt.want {
				t.Errorf("GetStorageGinLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStoragePath1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStoragePath(); got != tt.want {
				t.Errorf("GetStoragePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDev(); got != tt.want {
				t.Errorf("IsDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLinux(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLinux(); got != tt.want {
				t.Errorf("IsLinux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMacos(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMacos(); got != tt.want {
				t.Errorf("IsMacos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsProd(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProd(); got != tt.want {
				t.Errorf("IsProd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTest(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTest(); got != tt.want {
				t.Errorf("IsTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWindows(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWindows(); got != tt.want {
				t.Errorf("IsWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathExist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveMainPid(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMainPid()
		})
	}
}

func TestSignal(t *testing.T) {
	type args struct {
		sigs os.Signal
		fn   HandleSignal
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Signal(tt.args.sigs, tt.args.fn)
		})
	}
}

func TestEventLoopSigtermSignal1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EventLoopSigtermSignal()
		})
	}
}

func TestGetCronPath1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCronPath(); got != tt.want {
				t.Errorf("GetCronPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDaemonPath1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDaemonPath(); got != tt.want {
				t.Errorf("GetDaemonPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMainPid1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMainPid(); got != tt.want {
				t.Errorf("GetMainPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRootDir1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootDir(); got != tt.want {
				t.Errorf("GetRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServerPidFile1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerPidFile(); got != tt.want {
				t.Errorf("GetServerPidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStorageGinLog1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStorageGinLog(); got != tt.want {
				t.Errorf("GetStorageGinLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStoragePath2(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStoragePath(); got != tt.want {
				t.Errorf("GetStoragePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDev1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDev(); got != tt.want {
				t.Errorf("IsDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLinux1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLinux(); got != tt.want {
				t.Errorf("IsLinux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMacos1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMacos(); got != tt.want {
				t.Errorf("IsMacos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsProd1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProd(); got != tt.want {
				t.Errorf("IsProd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTest1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTest(); got != tt.want {
				t.Errorf("IsTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWindows1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWindows(); got != tt.want {
				t.Errorf("IsWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathExist1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathExist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveMainPid1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveMainPid()
		})
	}
}

func TestSignal1(t *testing.T) {
	type args struct {
		sigs os.Signal
		fn   HandleSignal
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Signal(tt.args.sigs, tt.args.fn)
		})
	}
}
