# 1.hello,world的测试学习

1. 我们学习到了编写测试，测试文件规则如下:
   * 程序需要在一个名为xxx_test.go的文件中编写
   * 测试函数的命名必须以单词Test开始
   * 测试函数只接受一个参数t *testing.T
2. 用参数和返回类型声明函数
   * func functionName(params) (returnName returnType){}
3. if,else,switch
4. 声明变量和常量

# 2.测试分组
1. 对一个「事情」进行分组测试，然后再对不同场景进行子测试非常有效。
2. 这种方法的好处是，你可以建立在其他测试中也能够使用的共享代码。
3. 代码如下
```azure
func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
```