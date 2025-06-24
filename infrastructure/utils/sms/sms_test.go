package sms

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSMS(t *testing.T) {
	config.Init()
	a := NewAliYunSms()
	err := a.SendCaptcha("14718121585", "123456")
	assert.NoError(t, err)
}
