//ff:func feature=scan type=convert control=selection topic=django
//ff:what Django <conv:name>/<name> 변수를 {name}로 치환하고 컨버터를 기록한 뒤 끝 인덱스를 반환한다
package django

// emitAngle handles a Django path variable starting at runes[open] ('<'). On a
// valid "<conv:name>" / "<name>" it writes "{name}", records the parameter, and
// returns the index of '>'. Otherwise it emits '<' literally and returns open.
func (n *pathNormalizer) emitAngle(open int) int {
	end := indexRune(n.runes, open+1, '>')
	if end < 0 {
		n.out.WriteRune(n.runes[open])
		return open
	}
	inner := string(n.runes[open+1 : end])
	name, ok := angleParamName(inner)
	if !ok {
		n.out.WriteRune(n.runes[open])
		return open
	}
	n.out.WriteString("{" + name + "}")
	n.params = append(n.params, urlParam{name: name, converter: angleConverter(inner)})
	return end
}
