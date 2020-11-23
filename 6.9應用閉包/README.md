# 應用閉包 將函式作為回傳值

### 範例

```go
func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	var f = f()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func f() func(int) int {
	var i int
	return func(delta int) int {
		i += delta
		return i
	}
}
```

範例中有點像是將 `Add()`宣告成`p2` 而`p2` 又是回傳一個閉包 在傳入閉包中所需的參數 才能顯示計算出的結果

在`f()`當中 由於宣告了一個全域變數 只要這次的`main`尚未結束前 `i`就會隨者呼叫的次數增加


### 尋找func 閉包位置

```go
where := func() {
    _, file, line, _ := runtime.Caller(1)
    log.Prinf(%s:%d, file, line)
}
```

### 計算韓式執行時間 尋找程式貧頸

```go
start := time.Now()
//do something
end := time.Now()
cost = end.Sub(start)
fmt.Println(`do something cost time is` : cost)
```