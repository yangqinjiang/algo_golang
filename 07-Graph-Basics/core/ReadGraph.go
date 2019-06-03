package core

import (
	"algo_golang/04-Heap/core"
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// 读取图算法
type ReadGraph struct {
	
}

func (this *ReadGraph) Read(g Graph,filename string) (error)  {
	content,err := ReadLines(filename)
	if err != nil{
		fmt.Println(err)
		return err
	}

	//fmt.Print(content)
	if(len(content) < 1){
		return errors.New("数据的行数小于1")
	}
	//var line string
	//var V int
	//var E int

	//第一行读取图中的节点个数和边的个数
	V,err := strconv.Atoi(content[0][0])
	if err != nil{
		return err
	}
	E,err := strconv.Atoi(content[0][1])

	if err != nil{
		return err
	}

	if V != g.V(){
		return errors.New("顶点个数不一致")
	}
	//从第二行起,读取每一条边的信息
	for i := 1;i<=E ;i++  {
		a,err := strconv.Atoi(content[i][0])
		if err != nil{
			return err
		}
		b,err := strconv.Atoi(content[i][1])

		if err != nil{
			return err
		}
		core.Assert(a>=0 && a<V)
		core.Assert(b>=0 && b<V)
		g.AddEdge(a,b)

	}
	return nil
}

//readLines reads a whole file into memory
// and returns a slice of its lines
func ReadLines(path string)([][]string,error)  {
	file,err :=os.Open(path)
	if err != nil{
		return nil,err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		//lines[] =append(lines,scanner.Text())

		line  := strings.Split(scanner.Text()," ")
		//fmt.Println("line=",line,scanner.Text())
		lines = append(lines, line)
	}
	return  lines,scanner.Err()
}
func Ioutil(name string) []string{
	if contents,err := ioutil.ReadFile(name);err == nil{
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		return  strings.Split(string(contents),"\n")
		//return strings.Replace(string(contents),"\n","",1)
	}else{
		return nil
	}
}
