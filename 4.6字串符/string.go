package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := "This is an example of a string"
	prefix := strings.HasPrefix(s, "this") //判斷前綴 是否包含prefix 大小寫也會判斷
	fmt.Println(prefix)                    //false

	suffix := strings.HasSuffix(s, "ing") // 判斷後綴 是否包含 suffix
	fmt.Println(suffix)                   // Ture
	//找不到會顯示-1
	contains := strings.Contains(s, "example") //判斷字串中是否包含 substr字串
	fmt.Println(contains)                      //True

	index := strings.Index(s, "s") // 判斷字串Ｓ 最一開始出現的substr字串
	fmt.Println(index)             //3

	last := strings.LastIndex(s, "s") // 判斷字串Ｓ 最後一次出現的substr字串
	fmt.Println(last)                 //24

	//替換字符
	replace := strings.Replace(s, " ", "   ", 2) // -1 全部替換 1: 替換一次 2:替換兩次 從前面開始計數
	fmt.Println(replace)                         // This   is   an example of a string

	//countString
	count := strings.Count(s, " ") // 計算substr 在字符s 出現過幾次
	fmt.Println(count)             // 6

	//Repeat
	repeat := strings.Repeat(s, 3) // 打印機 (要印的字串, 印的次數)
	fmt.Printf("The new repeated string is: %s\n", repeat)

	// 大小寫轉換
	upper := strings.ToUpper(s)
	fmt.Println(upper)
	lower := strings.ToLower(upper)
	fmt.Println(lower)

	// 修剪字串符
	trim := strings.Trim(s, "This") // 修剪 s前後 cutset 的字串   trimRight trimLeft 可以選擇開頭或結尾的修剪字串 常用於修正空白
	fmt.Println(trim)               // is an example of a string

	// 分割字符串

	fields := strings.Fields(s) //以空白做為分隔的陣列 回傳slice
	fmt.Println(fields[1:3])    // is an

	split := strings.Split(s, " ")
	for _, val := range split {
		fmt.Println(val) // This
		// is
		//	an
		//	example
		//	of
		//	a
		//	string
	}

	// 拼接slice 字串
	join := strings.Join(split, ",")
	fmt.Println(join)

	// 讀取string 內容

	reader := strings.NewReader(s) // 返回 Reader 指針  在用下面的方式 每5個字串為一個陣列
	t := make([]byte, 5)
	var writer bytes.Buffer
	for {
		n, err := reader.Read(t)
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		str := string(t[:n])
		fmt.Println(str)
		wr, err := writer.Write([]byte(str))
		fmt.Println(wr) // 寫了多少 byte
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	//簡易的string 轉換 int

	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is : %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is : %d\n", an) // integer %d
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The string is : %s\n", newS) // string %s
}
