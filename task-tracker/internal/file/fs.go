package file

import (
	"fmt"
	"io/fs"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Read(name string) []byte {
	bs, err := os.ReadFile(name)
	checkError(err)

	return bs
}

func Save(name string, d []byte) {
	err := os.WriteFile(name, d, fs.ModePerm)
	checkError(err)
}