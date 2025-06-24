package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/repository"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

type AliYunSms struct {
	*dysmsapi.Client
}

func NewAliYunSms() repository.ISMSUtils {
	var client *dysmsapi.Client
	var err error
	conf := config.Get().Phone

	client, err = dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(conf.AccessKeyId),
		AccessKeySecret: tea.String(conf.AccessKeySecret),
		Endpoint:        tea.String(conf.Endpoint),
	})
	if err != nil {
		panic(err)
	}

	return &AliYunSms{
		Client: client,
	}
}

func (ali *AliYunSms) SendCaptcha(phone string, code string) error {

	ma := make(map[string]string, 1)
	ma["code"] = code

	marshal, err := json.Marshal(ma)
	if err != nil {
		return err
	}

	data := string(marshal)

	conf := config.Get().Phone
	sendSmsRequest := &dysmsapi.SendSmsRequest{
		SignName:      tea.String(conf.SignName),
		TemplateCode:  tea.String(conf.TemplateCode),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: &data,
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		_, err = ali.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return err
		}

		return nil
	}()

	// 进行错误断言
	if tryErr != nil {
		var e = &tea.SDKError{}
		var _t *tea.SDKError
		if errors.As(tryErr, &_t) {
			e = _t
		}

		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(e.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err = util.AssertAsString(e.Message)
		if err != nil {
			return err
		}
	}
	return err
}
