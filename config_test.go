package vertexai

import (
	"reflect"
	"testing"
)

func Test_newClientConfig(t *testing.T) {
	type args struct {
		opts []clientConfigOption
	}
	tests := []struct {
		name string
		args args
		want clientConfig
	}{
		{
			name: "test new client config",
			args: args{
				opts: []clientConfigOption{
					setAuthToken("test_token"),
				},
			},
			want: clientConfig{
				authToken: "test_token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newClientConfig(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClientConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
