package 数组

var morseCodes = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
	"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
	"..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	hasMorseSequenceExisted := make(map[string]bool)
	for _, word := range words {
		hasMorseSequenceExisted[getMorseSequenceOfWord(word)] = true
	}
	return len(hasMorseSequenceExisted)
}

func getMorseSequenceOfWord(word string) string {
	sequence := ""
	for i := 0; i < len(word); i++ {
		sequence += getMorseCodeOfLetter(word[i])
	}
	return sequence
}

func getMorseCodeOfLetter(letter byte) string {
	return morseCodes[letter-'a']
}
