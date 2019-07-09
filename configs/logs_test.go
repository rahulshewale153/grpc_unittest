package configs

import (
	"context"
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestLogdata_Logger(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx       context.Context
		errortype string
		args      interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		flag   int
	}{
		// TODO: Add test cases.
		{"Error Log", fields{}, args{context.Background(), "error", "error generate"}, 0},
		{"Info Log", fields{}, args{context.Background(), "info", "info generate"}, 0},
		{"Debug Log", fields{}, args{context.Background(), "debug", "debug generate"}, 0},
		{"Warn Log", fields{}, args{context.Background(), "warn", "warn generate"}, 0},
		{"Default Log", fields{}, args{context.Background(), "default", "default generate"}, 0},
		{"File Not present error", fields{}, args{context.Background(), "default", "default generate"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ld := Logdata{
				logger: tt.fields.logger,
			}
			ld.logger = log.New()
			ld.logger.SetLevel(log.TraceLevel)
			ld.logger.Formatter = &log.TextFormatter{}
			ld.Logger(tt.args.ctx, tt.args.errortype, tt.args.args)
		})
	}
}

func TestWithRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		//	want context.Context
	}{
		// TODO: Add test cases.
		{"check with request id", args{context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRequestID(tt.args.ctx); got.Value(REQUESTID) == "" {
				t.Errorf("WithRequestID() = %v", got)
			}
		})
	}
}
