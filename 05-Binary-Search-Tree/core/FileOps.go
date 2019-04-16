package core

import (
	"regexp"
	"strings"
)

// 读取文件名称为filename中的内容，并将其中包含的所有词语放进words中
func FirstCharacterIndex(s string,start int) int  {
	for i:=start;i< len(s);i++  {
		if 'a' == s[i]{
			return i;
		}
	}
	return len(s)
}
func IsAlpha(s string) bool  {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(s)
}
// 将字符串s中的所有字母转换成小写之后返回
func LowerS(s string) string  {
	return strings.ToLower(s)
}
// 读取文件名称为filename中的内容，并将其中包含的所有词语放进words中
func ReadFile(filename string) ([]map[string]int,error) {
	// 文件读取
	//line := "";
	//contents := ""

	return nil, nil
}