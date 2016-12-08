package gotool

import (
	"crypto/md5"
	"io"
	"fmt"
	"crypto/sha1"
	"strconv"
)

//对字符串进行MD5哈希
func StringMd5(data string) string {
	t := md5.New();
	io.WriteString(t, data);
	return fmt.Sprintf("%x", t.Sum(nil));
}

//对字符串进行SHA1哈希
func Sha1(data string) string {
	t := sha1.New();
	io.WriteString(t, data);
	return fmt.Sprintf("%x", t.Sum(nil));
}

//int 类型转 string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

//int64 类型转 string
func Int64ToString(i int64) string {
	return  strconv.FormatInt(i, 10)
}

//string 类型转 int64
func StringToInt64(s string)(int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//float64类型转string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 1, 64)
}


//求百分比
func IntPercentString(a int, b int) string {
	return strconv.FormatFloat(float64(a) / float64(b) * 100, 'f', 1, 64)
}