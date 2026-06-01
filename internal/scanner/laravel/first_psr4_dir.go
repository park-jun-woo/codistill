//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what composer psr-4 값(문자열 또는 문자열 배열)에서 첫 디렉터리를 추출한다
package laravel

import "encoding/json"

// firstPSR4Dir decodes a composer psr-4 mapping value, which may be a single
// string ("app/") or an array of strings (["app/", "src/"]), and returns the
// first directory. It returns "" when the value is neither form or is empty.
func firstPSR4Dir(raw json.RawMessage) string {
	var single string
	if err := json.Unmarshal(raw, &single); err == nil {
		return single
	}
	var many []string
	if err := json.Unmarshal(raw, &many); err == nil && len(many) > 0 {
		return many[0]
	}
	return ""
}
