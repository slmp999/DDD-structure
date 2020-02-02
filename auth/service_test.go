package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/paybox/cloud/auth"
	"gitlab.com/paybox/cloud/mock"
)

func TestGetToken(t *testing.T) {
	mockRepo := mock.AuthRepository{
		GetTokenFunc: func(tokenID string) (*auth.Token, error) {
			if tokenID != "1" {
				return nil, auth.ErrTokenNotFound
			}

			return &auth.Token{
				ID:        "1",
				ClientID:  1,
				AccountID: 2,
			}, nil
		},
	}

	t.Run("success", func(t *testing.T) {
		repo := mockRepo
		s, err := auth.NewService(&repo)
		assert.NoError(t, err)

		tk, err := s.GetToken("1")
	 
		assert.NoError(t, err)
		assert.NotNil(t, tk)

		assert.True(t, repo.GetTokenInvoked)
		assert.Equal(t, int64(1), tk.ClientID)
		assert.Equal(t, int64(2), tk.AccountID)
	})
}
