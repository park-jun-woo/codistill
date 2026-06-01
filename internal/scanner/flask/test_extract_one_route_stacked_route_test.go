//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what 한 핸들러에 쌓인 다중 @bp.route 데코레이터의 모든 경로를 추출한다
package flask

import (
	"testing"
)

func TestExtractOneRoute_StackedRoutes(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.route('/health')
@app.route('/healthcheck')
@app.route('/ping')
def health():
    return 'ok'
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	bpPrefixes := make(blueprintPrefix)
	routes := extractRoutes(root, src, bpPrefixes, "health.py")

	if len(routes) != 3 {
		t.Fatalf("expected 3 routes, got %d", len(routes))
	}

	paths := make(map[string]bool)
	for _, r := range routes {
		paths[r.path] = true
		if r.method != "GET" {
			t.Errorf("expected GET for %s, got %s", r.path, r.method)
		}
		if r.handler != "health" {
			t.Errorf("expected handler health for %s, got %s", r.path, r.handler)
		}
	}
	for _, want := range []string{"/health", "/healthcheck", "/ping"} {
		if !paths[want] {
			t.Errorf("expected path %s to be extracted", want)
		}
	}
}
