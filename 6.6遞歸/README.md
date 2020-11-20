# 遞歸
> 如果數字太大 要小心不要超過範圍了

### 基本型

```
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
```

### 計算n階 
> 用int 到 12階就會溢出

```go
func homeWork(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = n * homeWork(n-1)
	}
	return
}
```