//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 파일기반 라우팅 경로 변환 테스트: src/api 이후 세그먼트 + [id]→{id}
package express

import "testing"

func TestFilebasedRoutePath(t *testing.T) {
	tests := []struct {
		relPath string
		want    string
	}{
		{"src/api/store/products/route.ts", "/store/products"},
		{"src/api/store/products/[id]/route.ts", "/store/products/{id}"},
		{"packages/medusa/src/api/admin/orders/route.ts", "/admin/orders"},
		{"src/api/route.ts", "/"},
		{"src/api/store/[...rest]/route.ts", "/store/{rest}"},
	}
	for _, tt := range tests {
		got := filebasedRoutePath(tt.relPath)
		if got != tt.want {
			t.Errorf("filebasedRoutePath(%q) = %q, want %q", tt.relPath, got, tt.want)
		}
	}
}
