# vowel-go

## purpose
hiragana and katakana convert vowel
## usage
```go
package main 
import(
"fmt"
"github.com/ryomak/vowel-go/vowel"
)

func main() {
	fmt.Println(vowel.GetVowel("日本文字はいいなあ"))//日本文字アイイアア
}
```
