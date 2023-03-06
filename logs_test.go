package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logBuilder(t *testing.T) {
	type args struct {
		LogType string
	}
	type want struct {
		Message string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "logBuilder when logInfo is called",
			args: args{
				LogType: logInfo,
			},
			want: want{
				Message: "[LOGS]-[INFO] message ",
			},
		},
		{
			name: "logBuilder when logTrace is called",
			args: args{
				LogType: logTrace,
			},
			want: want{
				Message: "[LOGS]-[TRACE] testing.go:1446:tRunner(): message ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			messageLogBuilder := logService.logBuilder(tt.args.LogType, caller, "message")

			assert.Equal(t, tt.want.Message, messageLogBuilder)
		})
	}
}

func Test_logDecorator(t *testing.T) {
	type args struct {
		MessageLevel string
		CallerLevel  string
	}
	type want struct {
		Message string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "logDecorator when callerLevel is 3",
			args: args{
				MessageLevel: logInfo,
				CallerLevel:  caller,
			},
			want: want{
				Message: "[LOGS]-[INFO] asm_amd64.s:1594:goexit(): message ",
			},
		},
		{
			name: "logDecorator when callerLevel is 4",
			args: args{
				MessageLevel: logInfo,
				CallerLevel:  callerDefault,
			},
			want: want{
				Message: "[LOGS]-[INFO] :0:(): message ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			messageLogBuilder := logService.logDecorator(tt.args.MessageLevel, tt.args.CallerLevel, "message")

			assert.Equal(t, tt.want.Message, messageLogBuilder)
		})
	}

}

func Test_getLevelCaller(t *testing.T) {
	type args struct {
		extraMessage []string
	}
	type want struct {
		Level string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "getLevelCaller when level is 3",
			args: args{
				extraMessage: []string{"3"},
			},
			want: want{
				Level: caller,
			},
		},
		{
			name: "getLevelCaller when level is 4",
			args: args{
				extraMessage: []string{"4"},
			},
			want: want{
				Level: callerDefault,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			levelCaller := getLevelCaller(tt.args.extraMessage...)

			assert.Equal(t, tt.want.Level, levelCaller)
		})
	}
}

func TestService_General(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "General",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "Test",
				URL:      "http://localhost:8080",
				FileLog:  true,
				fileName: "Logs",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Trace(tt.args.message)
		})
	}
}

func TestService_Trace(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Trace",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Trace(tt.args.message)
		})
	}
}

func Test_Trace(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Trace",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trace(tt.args.message)
		})
	}
}

func TestService_Info(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Info",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Info(tt.args.message)
		})
	}
}

func Test_Info(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Info",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.message)
		})
	}
}

func TestService_Warning(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Warning",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Warning(tt.args.message)
		})
	}
}

func Test_Warning(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Warning",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warning(tt.args.message)
		})
	}
}

func TestService_Error(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Error",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Error(tt.args.message)
		})
	}
}

func Test_Error(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Error",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.message)
		})
	}
}

func TestService_Fatal(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Fatal",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Fatal(tt.args.message)
		})
	}
}

func Test_Fatal(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Fatal",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.message)
		})
	}
}

func TestService_Debug(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Debug",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Debug(tt.args.message)
		})
	}
}

func Test_Debug(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Debug",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.message)
		})
	}
}

func TestService_Notice(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Notice",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logService := NewService(Service{
				NameApp:  "",
				URL:      "",
				FileLog:  false,
				fileName: "",
				ShowDate: false,
				ShowTime: false,
			})
			logService.Notice(tt.args.message)
		})
	}
}

func Test_Notice(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Notice",
			args: args{
				message: "message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Notice(tt.args.message)
		})
	}
}
