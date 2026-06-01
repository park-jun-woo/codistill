//ff:func feature=scan type=render control=iteration dimension=1
//ff:what (path,method) 충돌 손실 레코드들을 사람이 읽는 요약 리포트 문자열로 집계한다
package scanner

import (
	"fmt"
	"sort"
	"strings"
)

// formatPathConflictReport — 누적된 (path,method) 충돌을 path별로 그룹지어
// "총 N건 + 어디서 무엇이 소실됐는지"를 한 덩어리 텍스트로 만든다.
// 충돌이 없으면 빈 문자열을 반환한다(회귀 없음 보장).
// path 정렬로 출력이 결정적이다.
func formatPathConflictReport(conflicts []pathConflict) string {
	if len(conflicts) == 0 {
		return ""
	}

	byPath := map[string][]pathConflict{}
	for _, c := range conflicts {
		byPath[c.Path] = append(byPath[c.Path], c)
	}
	paths := make([]string, 0, len(byPath))
	for p := range byPath {
		paths = append(paths, p)
	}
	sort.Strings(paths)

	var b strings.Builder
	fmt.Fprintf(&b, "warning: %d operation(s) lost to (path,method) collisions — check route prefix composition\n",
		len(conflicts))
	for _, p := range paths {
		group := byPath[p]
		sort.Slice(group, func(i, j int) bool {
			if group[i].Method != group[j].Method {
				return group[i].Method < group[j].Method
			}
			return group[i].DropFile < group[j].DropFile
		})
		for _, c := range group {
			fmt.Fprintf(&b,
				"  %s %s: dropped %q at %s:%d (kept %q at %s:%d)\n",
				strings.ToUpper(c.Method), c.Path,
				c.DropHandler, c.DropFile, c.DropLine,
				c.KeptHandler, c.KeptFile, c.KeptLine)
		}
	}
	return b.String()
}
