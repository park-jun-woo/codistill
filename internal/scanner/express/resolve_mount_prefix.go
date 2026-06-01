//ff:func feature=scan type=extract control=selection topic=express
//ff:what 마운트 prefix 인자를 정적 문자열로 해소한다. string 리터럴이면 그대로, 식별자면 상수 해소. 실패 시 ok=false
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveMountPrefix(prefixNode *sitter.Node, fi *fileInfo, absRoot string, aliases map[string]string) (string, bool) {
	switch prefixNode.Type() {
	case "string":
		return unquoteTS(nodeText(prefixNode, fi.Src)), true
	case "identifier":
		// 식별자 prefix는 동일 파일/import 상수 리터럴로만 정적 해소한다.
		// 해소 실패 시 가짜 prefix 합성을 막기 위해 마운트를 스킵한다.
		name := nodeText(prefixNode, fi.Src)
		if lit := resolvePrefixIdentifier(fi, name, absRoot, aliases); lit != "" {
			return lit, true
		}
		return "", false
	}
	return "", false
}
