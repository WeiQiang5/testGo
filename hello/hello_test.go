package hello

import "testing"

// 子测试
func TestHello(t *testing.T) {
	//抽取的函数
	assertCorrectMessage := func(t *testing.T, got, want string) {
		/**
		t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号
		将在函数调用中而不是在辅助函数内部。这将帮助其他开发人员更容易地跟踪问题。
		*/
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := returnHello("golang", "")
		want := "Hello, golang"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := returnHello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := returnHello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
}
