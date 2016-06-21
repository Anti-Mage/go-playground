package main

import (
	"fmt"
	"os"
)

func testDirRename() {
	dir1 := "./test/tmp/"
	dir2 := "./test2/tmp2/"
	err := os.MkdirAll(dir1, 0777)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Rename(dir1, dir2)
	if err != nil {
		fmt.Println("1", err)
	}

	err = os.MkdirAll(dir2, 0777)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Rename(dir1, dir2)
	if err != nil {
		fmt.Println("2", err)
	}
}

func errMsg(line int, err error) {
	fmt.Println("line:", line, " errror:", err)
}

func testFileRename() {
	srcFilename := "test.txt"
	f, err := os.OpenFile(srcFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		errMsg(40, err)
	}

	str := "123456"
	for i := 0; i < 5; i++ {
		_, err = f.WriteString(str)
		if err != nil {
			errMsg(46, err)
		}
	}
	f.Close()

	dstFilename := "dstfile.txt"
	err = os.Rename(srcFilename, dstFilename)
	if err != nil {
		errMsg(52, err)
	}

}

func main() {
	testFileRename()
	testFileRename()
	testFileRename()
}
