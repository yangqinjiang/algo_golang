package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func ReadFile(filename string) ([]string,error) {
	fmt.Println("read file",filename)
	// 文件读取
	//line := "";
	contents := ""
	fin,err0 := os.OpenFile(filename,os.O_RDONLY,0666)
	defer fin.Close()
	if err0 != nil{
		return nil,err0
	}

	buf := bufio.NewReader(fin)
	for{
		line ,err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		if err != nil{
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil,err
			}
		}
		contents += line

	}

	// 简单分词
	// 这个分词方式相对简陋, 没有考虑很多文本处理中的特殊问题
	// 在这里只做demo展示用
	// 按空隔分离
	return strings.Split(contents," "), nil
}