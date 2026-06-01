//ff:func feature=scan type=parse control=iteration dimension=1 topic=django
//ff:what 룬 슬라이스에서 start 이후 첫 target 룬의 인덱스를 찾는다(없으면 -1)
package django

// indexRune returns the index of the first occurrence of target in runes at or
// after start, or -1 if not found.
func indexRune(runes []rune, start int, target rune) int {
	for i := start; i < len(runes); i++ {
		if runes[i] == target {
			return i
		}
	}
	return -1
}
