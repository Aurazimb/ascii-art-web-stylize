package asciiart

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetT(inputText, style string) (string, error) {
	if inputText == "" {
		return "", nil
	}

	if style != "standard" && style != "shadow" && style != "thinkertoy" {
		return "", errors.New("Bad request. Wrong select a style")
	}

	if !CheckHesh(style) {
		return "", errors.New("Code: 500, bad hash")
	}

	textArr := strings.Split(inputText, "\r\n")
	var resArr []string
	for i := 0; i < len(textArr); i++ {
		asciiArt := AsciiReturner(textArr[i], style)
		resArr = append(resArr, asciiArt)
	}
	var resStr string
	for i := 0; i < len(resArr); i++ {
		resStr = resStr + resArr[i] + "\n"
	}
	resStr = resStr[:len(resStr)-1]
	return resStr, nil
}

func CheckHesh(s string) bool {
	path := ("styles/" + s + ".txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error with open file")
		os.Exit(1)
	}
	defer file.Close()
	hesher := sha256.New()
	if _, err := io.Copy(hesher, file); err != nil {
		fmt.Println("Error on creating hesh")
		return false
	}
	hashInBites := hesher.Sum(nil)
	hashString := hex.EncodeToString(hashInBites)
	var rightHash []string
	rightHash = append(rightHash, "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf", "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73", "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3")
	if s == "standard" && hashString == rightHash[0] {
		return true
	} else if s == "shadow" && hashString == rightHash[1] {
		return true
	} else if s == "thinkertoy" && hashString == rightHash[2] {
		return true
	} else {
		fmt.Println("Bad Hash")
		return false
	}
}

func AsciiReturner(text, graphical string) string {
	if text == "" {
		return ""
	}
	lines := BannersReader(graphical)
	var ress []string
	for i := 0; i < len(text); i++ {
		textNum := int(text[i] - 32)
		if len(ress) == 0 {
			for j := 0; j < 8; j++ {
				ress = append(ress, lines[(textNum+1)+(textNum*8)+j])
			}
		} else {
			for j := 0; j < 8; j++ {
				ress[j] = ress[j] + lines[(textNum+1)+(textNum*8)+j]
			}
		}
	}
	var res string
	for i := 0; i < len(ress); i++ {
		res = res + ress[i] + "\n"
	}
	res = res[:len(res)-1]
	return res
}

func BannersReader(s string) []string {
	var filePath string = "styles/" + s + ".txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
