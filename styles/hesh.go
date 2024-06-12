package styles

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

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
