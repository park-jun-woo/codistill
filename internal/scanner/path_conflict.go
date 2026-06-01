//ff:type feature=scan type=model
//ff:what (path,method) 충돌로 폐기된 operation 1건의 손실 레코드
package scanner

// pathConflict — 동일 (path, method)에 두 핸들러가 등록돼 한쪽이 폐기된 충돌 1건.
// kept는 최종 보존된 핸들러, dropped는 폐기된 핸들러를 가리킨다.
// preferEndpoint 정책으로 결정적으로 선택되므로 입력 순서와 무관하다.
type pathConflict struct {
	Path        string
	Method      string
	KeptFile    string
	KeptLine    int
	KeptHandler string
	DropFile    string
	DropLine    int
	DropHandler string
}
