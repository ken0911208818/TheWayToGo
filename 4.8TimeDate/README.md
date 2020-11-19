# Time& Date

> 由於筆者曾寫過Laravel 有個套件Carbon 滿常使用的 而語法與 golang  time 大同小異

> >簡單介紹幾個常用的 


### 現在時間

```go
time.Now()
```

### 轉UNIX(now)

```go
time.Now().Unix()
```

### time.Sleep(d Duration)
### time.Ticker

> 如果需要在一定週期執行某個任務或是暫停的話 由於暫停時間都是使用int64 所以大部分計算時間的方式都是使用 time 的方式去拼湊時間並加以套入要執行的goroutine 或是只能進行多久 或是暫停 