# dirtywords
基于字典树算法的脏词检测，过滤。时间复杂度与字典大小无关

本库并不承担脏词库文件加载等工作，使用者自行从需要的途径（文件，数据库等）加载后创建字典树即可

## 使用方法
```go
package main

import (
  "fmt"

  "github.com/zxc122333/dirtywords"
)

func main() {
  badWords := [][]rune{
    []rune("脏词1"),
    []rune("脏词2"),
    []rune("bad words"),
  }
  skipLetters := []rune{' ', '.', '-', '*', '#', '@', ',', '/', '=', '_'}
  tree := dirtywords.BuildTree(badWords, skipLetters)
  tests := []string{
    "这句话是正常的",
    "这句话包含了脏词1",
    "这句话包含了脏词1和脏词2",
    "这句话脏 词 1加了空格",
    "这句话脏, *词 2混合加了各种符号",
    "这句话包含了英文bad words",
  }
  for _, test := range tests {
    fmt.Printf(
      "原文：%-20s\t是否包含：%v\t星号过滤：%s\n", 
      test, 
      tree.Check(test), 
      tree.Replace(test, '*'),
    )
  }
}

```
输出：
```
原文：这句话是正常的                   是否包含：false 星号过滤：这句话是正常的
原文：这句话包含了脏词1                是否包含：true  星号过滤：这句话包含了***
原文：这句话包含了脏词1和脏词2          是否包含：true  星号过滤：这句话包含了***和***
原文：这句话脏 词 1加了空格             是否包含：true  星号过滤：这句话*****加了空格
原文：这句话脏, *词 2混合加了各种符号    是否包含：true  星号过滤：这句话*******混合加了各种符号
原文：这句话包含了英文bad words        是否包含：true  星号过滤：这句话包含了英文*********
```
