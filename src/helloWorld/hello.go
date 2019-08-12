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
func replaceWord(word string) (filterdWord string) {
	filterdWord = word
	for _, keyword := range ngWords() {
		replaceKeyword := makeReplaceKeyword(keyword)
		filterdWord = strings.Replace(filterdWord, keyword, replaceKeyword, -1)
	}
	return
}

/**
ワードの置換文字列を作成
*/
func makeReplaceKeyword(word string) (replaceKeyword string) {
	for _, val := range word {
		if unicode.IsSpace(val) {
			replaceKeyword += " "
		} else {
			replaceKeyword += "*"
		}
	}
	return
}

/**
NGワード
*/
func ngWords() []string {
	return []string{"fuck", "sit", "go to hell", "JAVA", "殺す"}
}
