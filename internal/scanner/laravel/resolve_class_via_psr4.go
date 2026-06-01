//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 단축 클래스명을 소스 파일 use 임포트+composer psr-4로 해석해 on-demand 파싱한다
package laravel

// resolveClassViaPSR4 resolves a short class name to its parsed fileInfo by
// (1) looking the short name up in the source file's `use` imports to recover
// the FQCN, then (2) mapping that FQCN to a file path via composer PSR-4, then
// (3) parsing it on demand (memoized in parsedFiles). srcFI is the file whose
// imports name the class — the route file for a controller, the controller file
// for a FormRequest/Resource. It returns nil when the import or PSR-4 mapping
// is absent or the file does not exist, leaving the caller's hardcoded
// fallbacks to run. This is the precision recovery for classes that live in
// non-standard PSR-4 locations the old full-parse linear scan used to cover.
func resolveClassViaPSR4(absRoot, className string, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *fileInfo {
	if srcFI == nil {
		return nil
	}
	fqcn, ok := importMap(srcFI)[className]
	if !ok {
		return nil
	}
	relPath, ok := fqcnToFile(fqcn, composerPSR4(absRoot))
	if !ok {
		return nil
	}
	fi := loadCachedFile(absRoot, relPath, parsedFiles)
	if fi == nil {
		return nil
	}
	if !classMatches(fi, className) {
		return nil
	}
	return fi
}
