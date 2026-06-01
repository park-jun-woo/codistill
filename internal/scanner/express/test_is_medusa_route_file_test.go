//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what Medusa route 파일 가드 테스트: api 하위 route.{ts,js}만 통과, 일반 express 파일 차단
package express

import "testing"

func TestIsMedusaRouteFile(t *testing.T) {
	tests := []struct {
		relPath string
		want    bool
	}{
		{"src/api/store/products/route.ts", true},
		{"src/api/route.js", true},
		{"packages/medusa/src/api/admin/route.ts", true},
		{"src/api/store/products/middlewares.ts", false},
		{"src/routes/products.ts", false},
		{"src/commands/start.ts", false},
		{"src/api-helpers/route.ts", false},
	}
	for _, tt := range tests {
		got := isMedusaRouteFile(tt.relPath)
		if got != tt.want {
			t.Errorf("isMedusaRouteFile(%q) = %v, want %v", tt.relPath, got, tt.want)
		}
	}
}
