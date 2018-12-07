package golang

import "fmt"

func MakeURLSafe(phrase string) string {
	phraseAsRunes := []rune(phrase)

	for i, c := range phraseAsRunes {
		if c == '\'' || c == ' ' || c == '?' {
			phraseAsRunes[i] = '-'
		}
	}

	return string(phraseAsRunes)
}

func PhraseAsRunes() {
	phrase := "How're you guys doing?"
	fmt.Println(MakeURLSafe(phrase))
}
