package main

func main() {
	println(isAnagram("abc", "cba"))
	println(isAnagram("abce", "cbad"))
}

func isAnagram(s1, s2 string) bool {
	counts := make(map[int32]int)
	for _, v := range s1 {
		counts[v]++
	}
	for _, v := range s2 {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}
