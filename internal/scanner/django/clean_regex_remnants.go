//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what 구조 정규화 후 남은 정규식 잔재(백슬래시 이스케이프/문자클래스/수량자)를 세그먼트별로 보수 정리한다
package django

import "strings"

// cleanRegexRemnants performs per-segment cleanup of leftover regex syntax after
// groups/anchors have been replaced. Segments that already contain a "{param}"
// placeholder, or are pure literals, are preserved verbatim; only segments still
// carrying regex metacharacters are cleaned. A lone "." (e.g. "jwks.json") is a
// valid literal and is never treated as a metacharacter trigger.
func cleanRegexRemnants(path string) string {
	if !needsRegexCleanup(path) {
		return path
	}
	segments := strings.Split(path, "/")
	for i, seg := range segments {
		if strings.Contains(seg, "{") || !needsRegexCleanup(seg) {
			continue
		}
		segments[i] = cleanRegexSegment(seg)
	}
	return strings.Join(segments, "/")
}
