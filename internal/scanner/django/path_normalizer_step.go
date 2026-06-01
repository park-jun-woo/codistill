//ff:func feature=scan type=convert control=selection topic=django
//ff:what 정규화 1스텝: 현재 문자 종류에 따라 앵커 제거/그룹·변수 치환/리터럴 출력을 분기한다
package django

// step processes the rune at index i and returns the index to resume from
// (i.e. the last index it consumed). Anchors are dropped; '(' and '<' delegate
// to group/variable handlers; everything else is copied literally.
func (n *pathNormalizer) step(i int) int {
	switch n.runes[i] {
	case '^', '$':
		return i // anchor: emit nothing
	case '(':
		return n.emitGroup(i)
	case '<':
		return n.emitAngle(i)
	default:
		n.out.WriteRune(n.runes[i])
		return i
	}
}
