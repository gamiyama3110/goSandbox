package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	talk("Hello, world")
	fmt.Println("---------------")
	talk("fuck you !!!")
	talk("fuck! sit!! go to hell!!!")
	talk("HELLO!, JAVA!!!")
	talk("ぶっ殺す")
}

/**
指定メッセージをコンソールにプリント
NGワードはフィルターする。
*/
func talk(message string) {
	newMessage := replaceWord(message)
	fmt.Println(newMessage)
}

/**
ワード中のNGワードを置換
*/
func replaceWord(word string) string {
	for _, keyword := range ngWords() {
		replaceKeyword := makeReplaceKeyword(keyword)
		word = strings.Replace(word, keyword, replaceKeyword, -1)
	}

	return word
}

/**
ワードの置換文字列を作成
*/
func makeReplaceKeyword(word string) string {
	replaceKeyword := ""
	// count := utf8.RuneCountInString(word)
	for _, val := range word {
		if unicode.IsSpace(val) {
			replaceKeyword += " "
		} else {
			replaceKeyword += "*"
		}
	}
	return replaceKeyword
}

/**
NGワード
*/
func ngWords() []string {
	words := []string{"fuck", "sit", "go to hell", "JAVA", "殺す"}
	return words
}
