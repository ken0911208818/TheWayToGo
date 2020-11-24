# 切片

###宣告
```go
//var indenfier [len]type
var arr1 [5]int
```

### 顯示Array內值
```go
for index, value := range arr1 {
    fmt.Println("Array item", index, "is", value)
}

for index := range arr1 {
    fmt.Println("Array item", index, "is", arr1[index])
}

for i := 0; i < len(arr); i++ {
    fmt.Println("Array item", i, "is", arr1[i])
}
```
> 三種都能顯示Array，建議以需求去決定使用場景