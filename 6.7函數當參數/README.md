# 將函數當作參數

> 可以在參數中定義func()長相 只要符合就可以使用函式當作參數
> 筆者使用gin一定會用的 router 宣告的部分 只要function 符合 gin.Handler 
### gin

```go
func main() {
    r := gin.Default()
    r.Post("/v1/test/", handler.UserCreate)
}
```