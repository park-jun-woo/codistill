//ff:func feature=scan type=convert control=selection topic=django
//ff:what 정규식 그룹 '('를 named({name})·무명({paramN})·비포획(드롭)으로 치환하고 끝 인덱스를 반환한다
package django

import "strconv"

// emitGroup handles a regex group starting at runes[open]. It writes the
// replacement to the output buffer, records any path parameter, and returns the
// index of the group's closing ')' (or open itself when unbalanced).
func (n *pathNormalizer) emitGroup(open int) int {
	end := matchParen(n.runes, open)
	if end < 0 {
		n.out.WriteRune(n.runes[open]) // unbalanced: emit literally
		return open
	}
	inner := string(n.runes[open+1 : end])
	if name, ok := namedGroupName(inner); ok {
		n.out.WriteString("{" + name + "}")
		n.params = append(n.params, urlParam{name: name})
		return end
	}
	if isNonCapturingGroup(inner) {
		return end // non-capturing / lookaround: drop entirely
	}
	n.unnamed++
	name := "param" + strconv.Itoa(n.unnamed)
	n.out.WriteString("{" + name + "}")
	n.params = append(n.params, urlParam{name: name})
	return end
}
