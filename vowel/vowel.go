package vowel

import (
	"strings"
	"unicode"
)

var hiraToKatakana = unicode.SpecialCase{
	unicode.CaseRange{
		Lo:    0x3041, // ぁ
		Hi:    0x3093, // ん
		Delta: [unicode.MaxCase]rune{0x30a1 - 0x3041, 0, 0},
	},
}

func GetVowel(str string) string {
	word := strings.ToUpperSpecial(hiraToKatakana, str)
	b := make([]byte, 0, 10)
	for _, v := range word {
		b = append(b, toVowel(v)...)
	}
	return string(b)
}

func toVowel(w rune) string {
	if smallKana[w] {
		return string(w)
	}
	word := []rune(kana[w])
	switch word[len(word)-1] {
	case 'a':
		return "ア"
	case 'i':
		return "イ"
	case 'u':
		return "ウ"
	case 'e':
		return "エ"
	case 'o':
		return "オ"
	case 'n':
		return "ン"
	default:
		return string(w)
	}
}

var kana = map[rune]string{
	'ア': "a",
	'イ': "i",
	'ウ': "u",
	'エ': "e",
	'オ': "o",
	'カ': "ka",
	'ガ': "ga",
	'キ': "ki",
	'ギ': "gi",
	'ク': "ku",
	'グ': "gu",
	'ケ': "ke",
	'ゲ': "ge",
	'コ': "ko",
	'ゴ': "go",
	'サ': "sa",
	'ザ': "za",
	'シ': "shi",
	'ジ': "zi",
	'ス': "su",
	'ズ': "zu",
	'セ': "se",
	'ゼ': "ze",
	'ソ': "so",
	'ゾ': "zo",
	'タ': "ta",
	'ダ': "da",
	'チ': "chi",
	'ヂ': "di",
	'ツ': "tsu",
	'ヅ': "du",
	'テ': "te",
	'デ': "de",
	'ト': "to",
	'ド': "do",
	'ナ': "na",
	'ニ': "ni",
	'ヌ': "nu",
	'ネ': "ne",
	'ノ': "no",
	'ハ': "ha",
	'バ': "va",
	'パ': "pa",
	'ヒ': "hi",
	'ビ': "vi",
	'ピ': "pi",
	'フ': "fu",
	'ブ': "bu",
	'プ': "pu",
	'ヘ': "he",
	'ベ': "ve",
	'ペ': "pe",
	'ホ': "ho",
	'ボ': "vo",
	'ポ': "po",
	'マ': "ma",
	'ミ': "mi",
	'ム': "mu",
	'メ': "me",
	'モ': "mo",
	'ヤ': "ya",
	'ユ': "yu",
	'ヨ': "yo",
	'ラ': "ra",
	'リ': "ri",
	'ル': "ru",
	'レ': "re",
	'ロ': "ro",
	'ワ': "wa",
	'ヰ': "wi",
	'ヱ': "we",
	'ヲ': "wo",
	'ン': "n",
	'ヴ': "vu",
}

var smallKana = map[rune]bool{
	'ァ': true,
	'ィ': true,
	'ゥ': true,
	'ェ': true,
	'ォ': true,
	'ッ': true,
	'ャ': true,
	'ュ': true,
	'ョ': true,
	'ヮ': true,
	'ヵ': true,
	'ヶ': true,
}
