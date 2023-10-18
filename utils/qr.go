package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/png"
)

//CreateQrImage 创建二维码图片

func CreateQrImage(data any) string {
	qr, err := qrcode.New(data.(string), qrcode.Medium)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	qr.ForegroundColor = color.RGBA{0, 0, 0, 255}       // 前景色（黑色）
	qr.BackgroundColor = color.RGBA{255, 255, 255, 255} // 背景色（白色）

	//获取图片
	img := qr.Image(256)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fileToBase64(img)
}

func fileToBase64(file image.Image) string {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, file); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
