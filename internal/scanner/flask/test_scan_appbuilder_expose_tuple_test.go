//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 @expose의 tuple methods=("POST",)를 POST 1건으로 추출하는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_Appbuilder_ExposeTupleMethods(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "api.py"), []byte(`from flask_appbuilder.api import ModelRestApi, expose

class ChartRestApi(ModelRestApi):
    datamodel = SQLAInterface(Chart)
    base_url = "/api/v1/chart"
    @expose("/", methods=("POST",))
    def post(self):
        return {}
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	postRoot := 0
	for _, ep := range result.Endpoints {
		if ep.Method == "POST" && ep.Path == "/api/v1/chart" {
			postRoot++
		}
	}
	// The @expose POST override suppresses the synthesized CRUD POST, so exactly one.
	if postRoot != 1 {
		t.Fatalf("expected exactly 1 POST /api/v1/chart, got %d; endpoints: %+v", postRoot, result.Endpoints)
	}
}
