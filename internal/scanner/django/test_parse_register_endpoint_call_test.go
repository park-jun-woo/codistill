//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseRegisterCall_RegisterEndpointAlias 테스트
package django

import "testing"

func TestParseRegisterCall_RegisterEndpointAlias(t *testing.T) {
	fi := newTestFileInfo(t, "api_router.register_endpoint('pages', PagesAPIViewSet)\n")
	call := djFirst(t, fi.root, "call")
	reg := parseRegisterCall(call, fi)
	if reg == nil {
		t.Fatal("expected registration")
	}
	if reg.prefix != "pages" {
		t.Errorf("prefix: %q", reg.prefix)
	}
	if reg.viewsetName != "PagesAPIViewSet" {
		t.Errorf("viewsetName: %q", reg.viewsetName)
	}
	if reg.routerVar != "api_router" {
		t.Errorf("routerVar: %q", reg.routerVar)
	}
}
