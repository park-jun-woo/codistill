//ff:func feature=scan type=parse control=sequence topic=laravel
//ff:what relPath 후보를 파싱 캐시에서 조회하거나, miss 시 파싱해 캐시에 채워 반환한다(메모이즈)
package laravel

import (
	"path/filepath"
	"strings"
)

// loadCachedFile resolves a candidate relative path against the parse cache.
// On a cache hit it returns the already-parsed fileInfo; on a miss it parses
// the file on demand and inserts it into parsedFiles so a later reference to the
// same class is a cache hit. It returns nil when the file does not exist or
// fails to parse (e.g. oversized). parsedFiles is the lazy-parse cache; this is
// the memoization point that keeps stage-2 work bounded by referenced classes.
func loadCachedFile(absRoot, relPath string, parsedFiles map[string]*fileInfo) *fileInfo {
	key := strings.ReplaceAll(relPath, "\\", "/")
	if fi, ok := parsedFiles[key]; ok {
		return fi
	}
	fi := parseControllerCandidate(absRoot, filepath.Join(absRoot, filepath.FromSlash(key)))
	if fi == nil {
		return nil
	}
	parsedFiles[fi.relPath] = fi
	return fi
}
