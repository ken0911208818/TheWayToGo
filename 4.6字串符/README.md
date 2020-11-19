# string 應用 
> 為了方面說明 以下 s 代表某串字串
## 目錄
- [判斷前綴](###判斷前綴)
- [判斷後綴](###判斷後綴)
- [判斷整個](###字串當中是否包含)
### 判斷前綴

```go
strings.HasPrefix(s, prefix) bool 
// 判斷s 字串中 preifx 字串是否有包含在前綴 
```

### 判斷後綴

```go
strings.HasSuffix(s, Suffix) bool
// 判斷s 字串中 preifx 字串是否有包含在後綴
```

### 字串當中是否包含
```go
strings.Contains(s, Contains) bool
//判斷 s 字串中 是否有包含 Contains字串 return bool
```

### 字串當中第一個substr位置
```go
strings.Index(s, substr)
// 以單字第一個字為主  s := "this is a apple" substr := apple 則回傳 index = 10 
```

### 字串當中最後一個substr位子(index)
```go
strings.LastIndex(s, substr)
// 以單字第一個字為主  s := "this is a apple" substr := s 則回傳 index = 6
```
### 字串替換
```go
strings.RePlace(s, old, new, int) 
```
> int : -1 會將字串所有的有關係的 old 替換成 new  白話文 -1 全部替換
> int : 代表替換的次數 

### 字串計算
```go
strings.Count(s, substr)
```
> 計算 s字串中 substr 出現過幾次

### 字串修剪
```go
strings.Trim("This is a apple This", "This") // " is a apple "
strings.TrimLeft("This is a apple This", "This") // " is a apple This"
strings.TrimRight("This is a apple This", "This") // "This is a apple "
```