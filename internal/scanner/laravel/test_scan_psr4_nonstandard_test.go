//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 비표준 PSR-4 위치의 컨트롤러/FormRequest/Resource가 lazy parse 후에도 동일 해석되는지 회귀 가드
package laravel

import "testing"

// TestScan_PSR4NonstandardLocations is the regression guard the lazy-parse phase
// requires: a controller, FormRequest and Resource that live OUTSIDE the
// hardcoded candidate paths (app/Http/Controllers, app/Http/Requests,
// app/Http/Resources) and are located purely via composer's PSR-4 map. Before
// lazy parse, the full-parse linear scan resolved these; with the parse map
// reduced to route sources, only PSR-4 + use-import resolution recovers them.
func TestScan_PSR4NonstandardLocations(t *testing.T) {
	dir := t.TempDir()
	writePSR4NonstandardProject(t, dir)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	show := findEndpoint(result.Endpoints, "GET", "/api/invoices/{invoice}")
	store := findEndpoint(result.Endpoints, "POST", "/api/invoices")
	if show == nil || store == nil {
		logEndpoints(t, result.Endpoints)
		t.Fatal("expected both invoice endpoints")
	}

	if got := psr4PathParamType(show, "invoice"); got != "integer" {
		t.Errorf("invoice param type = %q, want integer (PSR-4 controller param)", got)
	}
	if got := psr4ResourceFieldCount(show, "InvoiceResource"); got != 2 {
		t.Errorf("resource response fields = %d, want 2 (PSR-4 resource)", got)
	}
	if got := psr4BodyFieldCount(store); got != 2 {
		t.Errorf("form request body fields = %d, want 2 (PSR-4 request)", got)
	}
}
