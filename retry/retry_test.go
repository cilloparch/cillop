package retry_test

import (
	"errors"
	"testing"

	"github.com/cilloparch/cillop/v2/retry"
)

type retryFuncTestCase struct {
	name    string
	fn      retry.RetryFunc
	wantErr error
}

func TestRun(t *testing.T) {
	secondTestCounter := 0
	tests := []retryFuncTestCase{
		{
			name: "Success on first try",
			fn: func() error {
				return nil
			},
			wantErr: nil,
		},
		{
			name: "Success on second try",
			fn: func() error {
				if secondTestCounter == 0 {
					secondTestCounter++
					return errors.New("failed operation")
				}
				return nil
			},
			wantErr: nil,
		},
		{
			name: "Retries on failure (specified max)",
			fn: func() error {
				return errors.New("failed operation")
			},
			wantErr: errors.New("failed operation"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := retry.Run(tt.fn, retry.DefaultConfig)
			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
