//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 파일기반 라우팅 디렉터리 세그먼트 한 건을 URL 세그먼트로 변환한다 ([id]→{id}, [...rest]→{rest})
package express

import "strings"

// filebasedSegment converts a single directory segment of a Medusa file-based
// route into a URL segment. "[id]" becomes "{id}" and a catch-all
// "[...rest]" becomes "{rest}". Static segments are returned unchanged.
func filebasedSegment(seg string) string {
	if strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]") {
		inner := seg[1 : len(seg)-1]
		inner = strings.TrimPrefix(inner, "...")
		return "{" + inner + "}"
	}
	return seg
}
