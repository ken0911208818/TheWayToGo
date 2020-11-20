# func


### 注意事項
- 順序性
    由於go是`編譯型語言` 所以函式撰寫順序是無關緊要,最好把`main`的func寫在文件前面
    
- DRY原則
    `Don't Repeat Yourself` 不要重複你自己 特定代碼只在func裡面執行一次

- `return`
    - 函式中: 退出
    - 可以帶有多個或`0`的參數 當成回傳值
    - 結束`for`迴圈或是`goroutine`
    
### 傳遞變長參數

可以用於不確定參數的長度使用
```go
func myFunc(a, b, arg ...int) {}
```

範例
```go
func main() {
	Greeting("Name: ", "joe", "Alice", "ken")
}

func Greeting(prefix string, arg ...string) {
	if len(arg) == 0 {
		fmt.Println("is error for no arg")
		return
	}
	for _, val := range arg {
		fmt.Println(prefix, val)
	}
	return
}
```

可以使用於不確定型別的可變參數
```go
func myFunc(a, b, arg ...interface{}) {}
//建議使用 switch 去選擇型別
```

### defer與追蹤
> 先說心得, 筆者對於defer的定義就是當前func結束時候 最後執行的動作 
>> 若一開始就defer 那當前defer的func會等到所屬的func結束時執行

範例
```go
package main
import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}
// out
//In Function1 at the top
//In Function1 at the bottom!
//Function2: Deferred until the end of the calling function!
```

前方高能 有坑請注意

```go
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
//output 
// 0
```
> 由於當下defer 輸入的i 是 0 所以延遲到func結束時 才做執行

#### 多個defer 順序

> 結論: 後進先出 堆疊的概念

```go
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
//output
// 4 3 2 1 0
```
#### 打印 func 的參數與返回值
```go
func test1(s string) (n int, err error) {
	defer func() {
		log.Printf("test1(%q) = %d, %v", s, n ,err)
	}()
	return 7, io.EOF
}
// 可以檢測input output
```
#### homework 

```go
package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
```


答案
> entering: b
> in b
> entering: a
> in a
> leaving: a
> leaving: b