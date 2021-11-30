package sendinblue

import (
	"testing"

	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSendinblueRepository(t *testing.T) {
	type args struct {
		c *config.SMTPConfig
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Initialize an SendinBlueRepo",
			args: args{c: &config.SMTPConfig{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewSendinBlueRepo(tt.args.c))
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		e mail.Email
	}

	tests := []struct {
		name    string
		wantErr bool
		want    mail.Email
		args    args
	}{
		{
			name: "Test",
			args: args{e: mail.Email{
				From:     "Andreas",
				FromName: "Andreas Hansson",
				To:       "andreas@dreamdata.io",
				Subject:  "Test",
				Content:  "This is a test",
			}},
			want: mail.Email{
				From:     "Andreas",
				FromName: "Andreas Hansson",
				To:       "andreas@dreamdata.io",
				Subject:  "Test",
				Content:  "This is a test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sbr := newSendinBlueMockRepo(&config.SMTPConfig{})
			got, err := sbr.Send(tt.args.e)
			if tt.wantErr {
				assert.Error(t, err)
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)

		})
	}
}

type mockSendinBlueRepo struct{}

func newSendinBlueMockRepo(*config.SMTPConfig) *mockSendinBlueRepo {
	return &mockSendinBlueRepo{}
}

func (sbr *mockSendinBlueRepo) Send(e mail.Email) (mail.Email, error) {
	return e, nil
}
