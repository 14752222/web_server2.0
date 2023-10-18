package utils

import (
	"fmt"
	"os"
)

type HtmlOptions struct {
	Title    string
	Body     string
	FileName string
	data     string
}

func (h *HtmlOptions) CreateHtml() bool {
	ok := h.createFile()
	if !ok {
		fmt.Println(`创建文件失败`)
		return false
	}
	h.data = h.createContent()

	writeErr := h.writeFile()

	if writeErr != nil {
		fmt.Println(`写入文件失败`)
		return false
	}
	return true
}

func (h *HtmlOptions) createContent() string {
	return fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>%s</title>
			</head>
			<body>
				%s  
			</body>
			</html>`, h.Title, h.Body)
}

func (h *HtmlOptions) createFile() bool {
	file, err := os.Create(h.FileName)
	if err != nil {
		fmt.Println(`无法创建文件 =>`, err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(`关闭文件失败 =>`, err)
		}
	}(file)

	return true
}

func (h *HtmlOptions) writeFile() error {
	// 打开文件
	file, err := os.OpenFile(h.FileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(`打开文件失败 =>`, err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(`文件打开失败`)
		}
	}(file)

	_, writeErr := file.WriteString(h.data)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func (h *HtmlOptions) DeleteFile() {
	err := os.Remove(h.FileName)
	if err != nil {
		fmt.Println(`删除文件失败 =>`, err)
	}
}
