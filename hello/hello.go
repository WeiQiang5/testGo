package hello

import "fmt"

// 提出常量
// 常量应该可以提高应用程序的性能，它避免了每次调用 Hello 时创建 "Hello, " 字符串实例。
const helloPrefix = "Hello, "

// 语言
const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

// 根据语言返回hello,world
func returnHello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

// 根据语言返回前缀
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello() {
	fmt.Println(returnHello("world", ""))
}
