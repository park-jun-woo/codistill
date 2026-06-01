//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what VERB 화이트리스트 매처 테스트: 대문자 GET/POST/PUT/PATCH/DELETE만 통과, all/config 거부
package express

import "testing"

func TestFilebasedVerbMethod(t *testing.T) {
	tests := []struct {
		ident  string
		want   string
		wantOK bool
	}{
		{"GET", "GET", true},
		{"POST", "POST", true},
		{"PUT", "PUT", true},
		{"PATCH", "PATCH", true},
		{"DELETE", "DELETE", true},
		{"get", "", false},
		{"config", "", false},
		{"ALL", "", false},
		{"Get", "", false},
	}
	for _, tt := range tests {
		got, ok := filebasedVerbMethod(tt.ident)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("filebasedVerbMethod(%q) = (%q,%v), want (%q,%v)", tt.ident, got, ok, tt.want, tt.wantOK)
		}
	}
}
