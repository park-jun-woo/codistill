//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 add_resource 등록과 Resource 메서드를 결합해 엔드포인트를 생성하는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_AddResource(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "resources.py"), []byte(`from flask_restful import Resource

class EventsResource(Resource):
    def get(self):
        return {}
    def post(self):
        return {}, 201
`), 0o644)
	os.WriteFile(filepath.Join(dir, "routes.py"), []byte(`from flask import Flask
from flask_restful import Api
app = Flask(__name__)
api = Api(app)
api.add_resource(EventsResource, "/data/events/last")
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	var got = map[string]bool{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = true
	}
	if !got["GET /data/events/last"] {
		t.Errorf("missing GET /data/events/last; endpoints: %+v", got)
	}
	if !got["POST /data/events/last"] {
		t.Errorf("missing POST /data/events/last; endpoints: %+v", got)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(result.Endpoints), got)
	}
}
