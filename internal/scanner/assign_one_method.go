//ff:func feature=scan type=extract control=sequence
//ff:what 단일 (path,method)에 op를 할당하되 충돌이면 preferEndpoint로 보존을 결정하고 손실 레코드를 반환한다
package scanner

// assignOneMethod — paths[oaPath][m]에 op를 할당한다.
// 이미 op가 있으면(충돌) resolvePathConflict로 보존 op를 결정적으로 고르고,
// 폐기된 핸들러를 가리키는 *pathConflict를 반환한다(충돌 없으면 nil).
func assignOneMethod(
	paths map[string]map[string]any,
	incumbent map[string]map[string]Endpoint,
	oaPath, m string,
	ep Endpoint,
	op map[string]any,
) *pathConflict {
	if _, dup := paths[oaPath][m]; dup {
		keep, conflict := resolvePathConflict(oaPath, m, ep, incumbent[oaPath][m])
		if keep {
			paths[oaPath][m] = op
			incumbent[oaPath][m] = ep
		}
		return &conflict
	}
	paths[oaPath][m] = op
	incumbent[oaPath][m] = ep
	return nil
}
