package sensitivewall

import (
	"bufio"
	"fmt"
	"gorobbs/util"
	"os"
)

var trie *util.Trie

func Trieinit() {
	trie = util.NewTrie()

	file, err := os.Open("./static/sensitive_words/words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, word := range words {
		fmt.Println("非法内容包含：", word)
		trie.Add(word, nil)
	}
}

func Check(text, replace string) (result string, hit bool) {
	return trie.Check(text, replace)
}
