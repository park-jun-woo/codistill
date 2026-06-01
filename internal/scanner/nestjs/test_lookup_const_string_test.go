//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestLookupConstString 테스트 (const 문자열 선언 값 추출)
package nestjs

import "testing"

func TestLookupConstString(t *testing.T) {
	src := []byte(`
const HEALTH_CHECK_ROUTE = 'health';
export const V2_BASE_PATH = "api/v2";
const COUNT = 3;
const MADE = makePath();
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := lookupConstString(root, src, "HEALTH_CHECK_ROUTE"); !ok || v != "health" {
		t.Fatalf("HEALTH_CHECK_ROUTE: want health,true got %q,%v", v, ok)
	}
	if v, ok := lookupConstString(root, src, "V2_BASE_PATH"); !ok || v != "api/v2" {
		t.Fatalf("V2_BASE_PATH: want api/v2,true got %q,%v", v, ok)
	}
	// numeric initializer is not a string node → no resolution
	if _, ok := lookupConstString(root, src, "COUNT"); ok {
		t.Fatal("COUNT: want false (non-string)")
	}
	// function-call initializer → no resolution
	if _, ok := lookupConstString(root, src, "MADE"); ok {
		t.Fatal("MADE: want false (non-string)")
	}
	if _, ok := lookupConstString(root, src, "MISSING"); ok {
		t.Fatal("MISSING: want false")
	}
}
