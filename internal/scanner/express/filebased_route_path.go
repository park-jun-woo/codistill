//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Medusa route.ts 파일경로의 src/api 이후 디렉터리 세그먼트를 URL path로 변환한다 ([param]→{param})
package express

import (
	"path/filepath"
	"strings"
)

// filebasedRoutePath derives the OpenAPI URL path from a Medusa file-based
// routing file path. It takes the directory segments after the last "api"
// segment, drops the trailing "route.{ts,js}" file, and converts dynamic
// "[param]" segments to "{param}". Catch-all "[...rest]" segments are
// conservatively converted to "{rest}". The empty path (api/route.ts) yields
// "/". relPath is expected to use forward slashes.
func filebasedRoutePath(relPath string) string {
	segs := strings.Split(filepath.ToSlash(relPath), "/")
	// drop the trailing route.{ts,js} file name.
	if len(segs) > 0 {
		segs = segs[:len(segs)-1]
	}
	// find the last "api" segment and keep what follows it.
	apiIdx := -1
	for i, s := range segs {
		if s == "api" {
			apiIdx = i
		}
	}
	if apiIdx < 0 {
		return "/"
	}
	rest := segs[apiIdx+1:]
	if len(rest) == 0 {
		return "/"
	}
	out := make([]string, 0, len(rest))
	for _, s := range rest {
		out = append(out, filebasedSegment(s))
	}
	return "/" + strings.Join(out, "/")
}
