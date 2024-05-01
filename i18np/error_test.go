package i18np

import "testing"

func TestError(t *testing.T) {
	type args struct {
		key    string
		params P
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Error",
			args: args{
				key:    "error.key",
				params: P{"param1": "value1"},
			},
			want: "error.key param1: value1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewError(tt.args.key, tt.args.params)
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Test Error implements error interface", func(t *testing.T) {
		var _ error = NewError("error.key")
		if false {
			t.Errorf("Error does not implement error interface")
		}
	})
}
