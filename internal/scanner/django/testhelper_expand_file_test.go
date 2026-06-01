//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what expandFile 테스트 헬퍼 — 파일 urlEntry 수집 후 평탄 pattern→view 맵으로 전개
package django

import "testing"

// expandFile collects a single file's urlEntries and fully expands them into flat
// pattern->view rows, mirroring the runtime include expansion path.
func expandFile(t *testing.T, src string) map[string]string {
	t.Helper()
	fi := newTestFileInfo(t, src)
	byModule := collectURLs([]fileInfo{fi})
	got := map[string]string{}
	for _, e := range byModule[fi.module] {
		for _, x := range expandURLEntry(e, "", byModule, map[string]bool{}) {
			got[x.pattern] = x.viewName
		}
	}
	return got
}
