package utility

/**
 * @Author  糊涂的老知青
 * @Date    2022/4/4
 * @Version 1.0.0
 */
import (
	"fmt"
	captcha "github.com/mojocn/base64Captcha"
	"image/color"
)

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverMath {
	driver := &captcha.DriverMath{
		Height:          38,
		Width:           120,
		NoiseCount:      0, // 干扰词数量
		ShowLineOptions: 0,
		Fonts:           []string{"ApothecaryFont.ttf", "DeborahFancyDress.ttf", "Flim-Flam.ttf", "wqy-microhei.ttc"}, //"ApothecaryFont.ttf","DeborahFancyDress.ttf","Flim-Flam.ttf","wqy-microhei.ttc"
		BgColor: &color.RGBA{ // 背景颜色
			R: 128,
			G: 98,
			B: 112,
			A: 0,
		},
	}
	return driver
}

type captchaUtil struct{}

func captchaOperateNum() (string, int, int, int) {
	ops := []string{"+", "-", "*"}
	op := ops[RandomNum(0, 3)]
	firstNum := 0
	secondNum := 0
	resultNum := 0
	switch op {
	case "+":
		firstNum = RandomNum(0, 20)
		secondNum = RandomNum(0, 20)
		resultNum = firstNum + secondNum
		break
	case "-":
		firstNum = RandomNum(1, 20)
		secondNum = RandomNum(0, firstNum)
		resultNum = firstNum - secondNum
		break
	case "*":
		firstNum = RandomNum(0, 20)
		if firstNum < 11 {
			secondNum = RandomNum(0, 10)
			resultNum = firstNum * secondNum
		} else {
			cops := []int{0, 1, 2, 10, 20}
			secondNum = cops[RandomNum(0, 5)]
			resultNum = firstNum * secondNum
		}
		break
	}
	return op, firstNum, secondNum, resultNum
}

// GenerateCaptchaHandler 生成图形验证码
func (*captchaUtil) GenerateCaptchaHandler() (string, int) {
	var driver = NewDriver().ConvertFonts()
	c := captcha.NewCaptcha(driver, store)
	op, firstNum, secondNum, resultNum := captchaOperateNum()
	content := fmt.Sprintf("%d %v %d = ?", firstNum, op, secondNum)
	item, _ := c.Driver.DrawCaptcha(content)
	return item.EncodeB64string(), resultNum
}

var CaptchaUtil = &captchaUtil{}
