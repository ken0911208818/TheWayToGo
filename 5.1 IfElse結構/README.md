# if else 結構控制
> 基本的我就跳過了 我只講與php的差異  
> 註： 筆者本身從php開始學習 


### 沒有括弧
> 在if switch for 當中 是不使用括弧的 好處是看起來非常簡潔 壞處: 需要一段時間才能習慣


### 塞入變數
> 在if switch for 迴圈當中 可以適度的塞入變數 並進行判斷 
> 但只能在該迴圈當中做使用

```go
if v:= true; v == true {   // 也可以塞入function 回傳的值
    //do something
} else{
    // do something
}
```