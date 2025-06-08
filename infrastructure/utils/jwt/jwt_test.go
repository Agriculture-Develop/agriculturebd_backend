package jwt_test

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/jwt"
	"testing"
	"time"
)

// 初始化配置（根据你的 config 结构做适配）
func init() {
	// 模拟设置 config 值（建议换成 mock 或使用你自己的 config 初始化方法）
	config.Init()
}

func TestGenerateAndParseToken(t *testing.T) {
	id := int64(12345)

	token, exp, err := jwt.GenerateToken(id)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	if token == "" {
		t.Fatal("Expected non-empty token")
	}
	if exp == nil || time.Until(*exp) <= 0 {
		t.Fatal("Expected valid expiration time")
	}

	parsedID, err := jwt.ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken error: %v", err)
	}
	if parsedID != id {
		t.Errorf("Expected ID %d, got %d", id, parsedID)
	}
}

func TestParseToken_Invalid(t *testing.T) {
	invalidToken := "this.is.an.invalid.token"
	_, err := jwt.ParseToken(invalidToken)
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}
