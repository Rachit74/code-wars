package main

func vowel_count(str string) int {
	count := 0
	for i := 0; i <= len(str)-1; i++ {
		if str[i] == ('a') || str[i] == ('e') || str[i] == ('i') || str[i] == ('o') || str[i] == ('u') {
			count += 1
		}
	}
	return count
}
