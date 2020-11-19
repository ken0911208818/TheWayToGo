# Pointer
> 先說結論 不要太依賴 Pointer 否則效能只會不進則退
> 除非變數的量很大 不然一般變數直接傳值比較快

> tips: fmt.Printf(%p) %p 是用來顯示記憶體位置使用的


### 常識
當一個指針被定義後沒有分配到任何變數時 值為nil 不為0 null (與php不同)

### 記憶體位置
```go
func main() {
	s := "hello word"
	var p *string
	p = &s
	*p = "你好"
	fmt.Printf("原本的s :%s 記憶體位置 : %p 被更改過的Ｐ: %s, 記憶體位置 ： %p", s, &s, *p , p)
}
// 原本的s :你好 記憶體位置 : 0xc000010200 被更改過的Ｐ: 你好, 記憶體位置 ： 0xc000010200

```
由於s p 都是指向某一個記憶體位置 當s 原本宣告 hello word 後來被 共用記憶體位置的p 給更改後 s 與 p 的值就跟者改變了

### 錯誤示範
```
const i = 5
func main() {
ptr := &i
ptr := &10
}
```
不能用pointer的方式去存取const的值 或是 常數 本身 10 就不是的變數 所以就沒有所謂的記憶體位置

```go
var p *int = nil
func main() {
	*p = 0
	fmt.Print(*p)
}
```
在這個實例中 由於已經宣告了 p 的記憶體位置為nil 所以當 *p 去存取值的就會發生問題 也沒辦法做到更改變數