//ff:type feature=scan type=model topic=django
//ff:what Django URL 정규식 정규화 진행 상태(입력 룬/출력 버퍼/누적 param/무명 그룹 카운터)
package django

import "strings"

// pathNormalizer carries the mutable state of normalizeDjangoPath's single pass.
type pathNormalizer struct {
	runes   []rune
	out     strings.Builder
	params  []urlParam
	unnamed int // count of unnamed capture groups seen so far
}
