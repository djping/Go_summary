## 包
  + 每一个Go程序都有是一个包，`main`包是程序的入口
  + `main`包必须有一个`main`函数，它没有返回值和参数
  ******
## 变量的定义
  + 使用`var`关键字定义，每个变量有其特定的类型
  + `var x int = 100`
  + `var n string`; `n = "haha"`
  + 可以一次定义多个: `var x, y = 10, 20`
  + 可以使用分组的方式定义: 
## 内建变量类型
+ bool(true/false)
+ 数值
    + 整数: int/uint
    + 浮点数: float64/float32,complex64/complex128
+ string
******
### 不同类型的值不能做运算
   + 数值的int、float、complex之间不能做运算
   + 不同位的int不能做运算  `int8(10) + int16(20) // wrong`

## 常量与枚举
  ### 常量(const)
  + 常量使用`const`定义且不必大写，`const radius = 2` 常量不必加类型，可以当做任意类型？？？
  + 分组`const (width = 20;height = 40)`
  ### 枚举(enum)
  ```
     const ( 
       orange iota // 关键字
       blue 
       red
      )
  ```
  + 枚举的第一值默认为0,后面的依次增加
  ****
  ****
## 条件语句
  + Go语言中的`if`语句加`{}`
  + `if` 语句可以先赋值再判断
  +     var x int = 10
        if x > 10 {...}
  +     if var y int = 10; y > 20 {..} //此时的y具块级作用域，只能在当前条件语句中使用
***
## 循环
  + ！Go语言中只有 `for` 没有 `while`
  + `for` 的一般形式与js一样,只是没有`()`
    ```
    for varl = value1;var1 <  value2; var1++ {...}
    ```
  + 例子:
      ```
        // eg1:
        var n int
        for n = 0; n < 10; n++ {...}

        //eg2--没有初始语句
        var n int = 10
        for ; n < 10; n++ {...}

        // eg3--没有初始和结尾语句:
        var n int = 20
        for n > 20 {...}
       ```
  + 参考上面`eg3`，当for没有初始与结尾数据的就是`while` 语句
## 函数
  + Go中函数使用关键字`func` 定义, 它可以有多个返回值
  + 函数的参数和返回值需要标明类型
  + 例子:
    ```
      // func_eg1
      func add (n1 int, n2 int) int {
        return n1 + n2
      }
      //当参数类型一样时可简写，上面代码简化为:
      func add (n1, n2 int) int {..}
    ```
    ```
      <!-- func_eg2：多返回值 -->
      func add (n1, n2 int) (int int) {
        return n1 + n2, n1 - n2
      }

      // 为返回值去别名,可直接return
      func add (num int) (numMax int) {
        numMax = num * num
        return
      }
    ```
******
******
## 指针
  + Go按值传递，传基本类型时传其copy，因此在函数中对数值的运算不影响其本身
  ```
    var a int = 10
    func add (num int) {
      num--
    }
    add(a)
    fmt.Printf("%d", a) // 10--不会改变
  ```
  + 改变基本类型的值需要传递其指针而不是copy,通过 `&` 操作符
  ```
    var a int = 10
    func add (num *int) { // 此时num是int型指针
      *num--
    }
    add(&a)
    fmt.Printf("%d", a) // 9 -- 改变
  ```
