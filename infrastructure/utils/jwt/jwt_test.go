package jwt

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateAndParseToken(t *testing.T) {

	mockConf := &config.Default{
		Auth: config.Auth{
			JwtSecret:     "test-secret",
			JwtExpireTime: "1h",
			Issuer:        "xxx",
		},
	}
	patch := gomonkey.ApplyFunc(config.Get, func() *config.Default {
		return mockConf
	})
	defer patch.Reset()
	fmt.Println("mock config.Get() =>", config.Get()) // 如果是 nil，说明 patch 没生效

	id := uint(999)

	token, err := GenerateToken(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedID, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, id, parsedID)
}

func TestParseToken_Invalid(t *testing.T) {

	invalidToken := "xxx.yyy.zzz"
	_, err := ParseToken(invalidToken)
	assert.Error(t, err)
}
