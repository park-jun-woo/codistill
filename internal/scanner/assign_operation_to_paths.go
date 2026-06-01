//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 메서드(any 확장 포함)별로 operation을 paths에 할당하고 (path,method) 충돌은 preferEndpoint로 결정적 보존하며 손실을 수집한다
package scanner

import (
	"strings"
)

// assignOperationToPaths — ep의 op를 paths[oaPath]에 메서드별로 할당한다.
//
// 동일 (path, method)에 이미 operation이 있으면(충돌) 마지막-등록-승(last-wins)으로
// 무조건 덮어쓰지 않고, incumbent에 보존된 Endpoint와 preferEndpoint로 비교해
// **결정적으로** 남길 쪽을 고른다(스캔 파일 순서에 따른 비결정성 제거).
// 폐기된 operation은 pathConflict 레코드로 누적해 스캔 종료 시 요약 리포트로 가시화한다.
//
// incumbent는 paths와 평행한 보존-Endpoint 추적 맵으로, 현재 어떤 Endpoint의 op가
// 각 (path,method)에 실려 있는지 기억한다. 호출자(buildSpecNode)가 소유한다.
// 반환값은 이번 호출에서 발생한 충돌(폐기) 레코드들이다.
func assignOperationToPaths(
	paths map[string]map[string]any,
	incumbent map[string]map[string]Endpoint,
	oaPath string,
	ep Endpoint,
	op map[string]any,
) []pathConflict {
	method := strings.ToLower(ep.Method)
	var conflicts []pathConflict
	if incumbent[oaPath] == nil {
		incumbent[oaPath] = map[string]Endpoint{}
	}
	for _, m := range expandAnyMethod(method) {
		if c := assignOneMethod(paths, incumbent, oaPath, m, ep, op); c != nil {
			conflicts = append(conflicts, *c)
		}
	}
	return conflicts
}
