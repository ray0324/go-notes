package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"time"
)


// var t = time.Now()

func main() {
	t := time.Now()
	// y := t.Year()
	// m := t.Month()
	// d := t.Day()
	// d1 := t.Date(y)
	layout := "2006年01月02日03时04分05秒-07时区08"
	fmt.Println(t.Format(layout))
	hash := sha256.New()
	hash2 := md5.New()
	hash.Write([]byte(layout))
	fmt.Printf("%x\n", hash.Sum(nil))
	fmt.Printf("%x\n", hash2.Sum(nil))
	fmt.Println(time.Hour * 5)
	// fmt.Printf("%d-%02d-%02d\n", t.Year(), t.Month(), t.Day())
	t1 := time.Now()
	t2 := t1.Add(time.Hour * 3)
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
	fmt.Println(&t2)
}
