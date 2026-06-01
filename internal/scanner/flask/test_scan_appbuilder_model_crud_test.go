//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 ModelRestApi 서브클래스의 표준 CRUD 5엔드포인트를 합성하는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_Appbuilder_ModelRestApiCRUD(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "api.py"), []byte(`from flask_appbuilder.api import ModelRestApi

class BarApi(ModelRestApi):
    datamodel = SQLAInterface(Bar)
    base_url = "/api/v1/bar"
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = true
	}
	want := []string{
		"GET /api/v1/bar",
		"GET /api/v1/bar/{pk}",
		"POST /api/v1/bar",
		"PUT /api/v1/bar/{pk}",
		"DELETE /api/v1/bar/{pk}",
	}
	for _, w := range want {
		if !got[w] {
			t.Fatalf("missing %q; endpoints: %+v", w, got)
		}
	}
}
