//ff:func feature=scan type=extract control=sequence
//ff:what (path,method) 충돌에서 preferEndpoint로 보존 op를 결정하고 폐기 손실 레코드를 만든다
package scanner

// resolvePathConflict — 동일 (oaPath, m)에 candidate(ep)와 incumbent(prev)가 겹칠 때
// preferEndpoint로 남길 쪽을 결정적으로 고르고, 폐기된 핸들러를 pathConflict로 기록한다.
// candidate가 우선되면 keep=true(호출자가 op/incumbent를 교체), 아니면 false(이전 op 유지).
func resolvePathConflict(oaPath, m string, ep, prev Endpoint) (keep bool, conflict pathConflict) {
	if preferEndpoint(ep, prev) {
		return true, pathConflict{
			Path:        oaPath,
			Method:      m,
			KeptFile:    ep.File,
			KeptLine:    ep.Line,
			KeptHandler: ep.Handler,
			DropFile:    prev.File,
			DropLine:    prev.Line,
			DropHandler: prev.Handler,
		}
	}
	return false, pathConflict{
		Path:        oaPath,
		Method:      m,
		KeptFile:    prev.File,
		KeptLine:    prev.Line,
		KeptHandler: prev.Handler,
		DropFile:    ep.File,
		DropLine:    ep.Line,
		DropHandler: ep.Handler,
	}
}
