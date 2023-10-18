package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetDocument(fileName string) (string, error) {
	filePath := "./tmp/" + fileName

	outPutFileName := "./assets/" + fileName + ".html"

	//创建一个文件在output
	file, err := os.Create(outPutFileName)

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	defer file.Close()

	cmd := exec.Command("pandoc", filePath, "-o", outPutFileName)

	combinedOutput, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(combinedOutput))
		return "", err
	}

	fmt.Println("Conversion completed successfully")

	return outPutFileName, nil
}
