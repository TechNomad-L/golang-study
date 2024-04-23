// 第一次作业：
// 功能1：先实现一个交互功能   程序运行 之后提示输入姓名  输入完成回车 返回  你好 xxx
// 功能2：然后 在提示输入英文字母  要将小写字幕准换成大写字母 把大写字母转换成小写字母


package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    input := bufio.NewReader(os.Stdin)

    fmt.Print("Please enter your name:") 
    name, _ := input.ReadString('\n')
    name = strings.TrimSpace(name)
//    fmt.Printf("Hello, %s\n", name)

// 调用isEnglishName函数判断输入的名字是不是英文，否则直接提示输入名称不合规
    if isEnglishName(name) {
        fmt.Printf("你好，%s\n", name)
    } else {
        fmt.Println("请输入英文姓名。")
        os.Exit(0)
    } 

// 如果输入的英文就直接调用change函数修改大小写转换
    fmt.Println("大小写互换后:", change(name))
}

// 判断是不是英文名称以及是否包含空格
func isEnglishName(s string) bool {
    for _, char := range s {
        if char != ' ' && (char < 'A' || (char > 'Z' && char < 'a') || char > 'z') {
            return false
        }
    }
    return true
}

// 大小写转换函数
func change(s string) string  {
    result := ""
    for _, char := range s {
        //小写转大写
        if char >= 'a' && char <= 'z' {
            result += string(char - 32)
        } else if char >= 'A' && char <= 'Z'{
        //大写转小写
            result += string(char + 32)
        } else {
        //如果不是字母就保持原样
            result += string(char)
        }
    }
    return result
}

