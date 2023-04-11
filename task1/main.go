package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("win-1251.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Возникла ошибка при открытии файла", err)
		return
	}
	defer file.Close()

	data := make([]byte, 64)
	utf8Bytes := make([]byte, 64)

	for {
		symbolByte, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:symbolByte]))
	}

	win1251Decoder := charmap.Windows1251.NewDecoder()
	utf8Encoder := transform.UTF8.NewEncoder()

	for i := 0; i < len(data); {
		unicodeBytes, _, errDec := transform.Bytes(win1251Decoder, data[i:])
		if errDec != nil {
			fmt.Println("Ошибка декодирования:", errDec)
			return
		}
		utf8Bytes, _, errEnc := transform.Bytes(utf8Encoder, unicodeBytes)
		if errEnc != nil {
			fmt.Println("Ошибка кодирования:", errEnc)
			return
		}
		copy(data[i:], utf8Bytes)
		i += len(utf8Bytes)
	}

	fmt.Println(utf8Bytes)

	var name string
	fmt.Fscan(os.Stdin, &name)
}
