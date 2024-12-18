package captcha

import (
	"context"
	"github.com/mojocn/base64Captcha"
)

var Store = base64Captcha.DefaultMemStore

// Info represents a captcha info: captcha ID, picture(base64), and answer.
type Info struct {
	CaptchaID string
	Picture   string
	Answer    string
}

// GetCaptcha returns a captcha info.
func GetCaptcha(ctx context.Context) (*Info, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, ans, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	return &Info{
		CaptchaID: id,
		Picture:   b64s,
		Answer:    ans,
	}, nil
}
