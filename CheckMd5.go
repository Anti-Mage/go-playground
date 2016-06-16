package main

import "io"
import "os"
import "fmt"
import "crypto/md5"
import "crypto/sha1"

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: check [md5|sha1] file\n\n")
		fmt.Printf("  %4s: %s\n", "md5", "计算文件md5值")
		fmt.Printf("  %4s: %s\n", "sha1", "计算文件sha1值")
		return
	}
	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Printf("Sorry, file open failed\n")
		return
	}
	buf := make([]byte, 1024*1024*16*16)
	switch os.Args[1] {
	case "md5":
		h := md5.New()
		for {
			n, err := f.Read(buf)
			h.Write(buf[:n])
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%x\n", h.Sum(nil))
	case "sha1":
		h := sha1.New()
		for {
			n, err := f.Read(buf)
			h.Write(buf[:n])
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%x\n", h.Sum(nil))
	default:
		fmt.Printf("参数有误\n")
	}
}
