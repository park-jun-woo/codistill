//ff:func feature=scan type=extract control=sequence topic=express
//ff:what prefix 식별자를 정적 문자열 리터럴로 해소한다 (동일 파일 const → import 원본 파일 const). 실패 시 ""
package express

import "path/filepath"

func resolvePrefixIdentifier(fi *fileInfo, name, absRoot string, aliases map[string]string) string {
	// (a) 동일 파일 내 const X = '...'
	if lit := resolveConstString(fi.Root, fi.Src, name); lit != "" {
		return lit
	}
	// (b) import 바인딩이면 원본 파일에서 export된 상수 리터럴 해소
	dir := filepath.Dir(fi.Path)
	target := resolveDestructuredRequirePath(fi.Root, fi.Src, name, dir, absRoot, aliases)
	if target == "" {
		return ""
	}
	imported, err := parseFile(target)
	if err != nil {
		return ""
	}
	return resolveConstString(imported.Root, imported.Src, name)
}
