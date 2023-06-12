package vertexai

import "testing"

func Test_getToken(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "unable to retrieve credentials",
			key:     ``,
			wantErr: true,
		},
		{
			name: "unable to retreive token",
			key: `{
  "type": "service_account",
  "project_id": "project_id",
  "private_key_id": "1234567",
  "private_key": "-----BEGIN PRIVATE KEY-----\n456789\n-----END PRIVATE KEY-----\n",
  "client_email": "dev@project.iam.gserviceaccount.com",
  "client_id": "12345",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/dev%40project.iam.gserviceaccount.com"
}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &defaultTokenizer{}
			got, err := tokenizer.getToken(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
