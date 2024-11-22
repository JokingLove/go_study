package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func main() {
	str := "123456"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))

	fmt.Printf("current time str : %s\n", getTimeStr())

	fmt.Printf("current time unix : %d\n", gettimeUnix())
}

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func gettimeUnix() int64 {
	return time.Now().Unix()
}
