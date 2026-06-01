//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Resource 클래스를 담은 파일을 캐시/PSR-4/하드코딩 경로에서 찾는다
package laravel

// findResourceFile locates the file containing the Resource class. Resolution
// order: parse cache, then PSR-4 (the controller file's `use` import + composer
// psr-4), then the conventional app/Http/Resources hardcoded candidate. srcFI
// is the controller file naming the resource type; it may be nil (PSR-4 step
// skipped).
func findResourceFile(absRoot, className string, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *fileInfo {
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}
	if fi := resolveClassViaPSR4(absRoot, className, srcFI, parsedFiles); fi != nil {
		return fi
	}
	candidates := []string{
		absRoot + "/app/Http/Resources/" + className + ".php",
	}
	for _, candidate := range candidates {
		fi, err := parseFile(absRoot, candidate)
		if err == nil {
			return fi
		}
	}
	return nil
}
