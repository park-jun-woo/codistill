//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what 쌓인 @app.route alias에서 path param 경로와 정적 경로를 모두 추출한다
package flask

import (
	"testing"
)

func TestExtractOneRoute_StackedAlias(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.route('/avatar/<name>.svg')
@app.route('/avatar/blank.svg')
def avatar(name=None):
    return ''
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	bpPrefixes := make(blueprintPrefix)
	routes := extractRoutes(root, src, bpPrefixes, "blueprint.py")

	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}

	paths := make(map[string]bool)
	for _, r := range routes {
		paths[r.path] = true
		if r.handler != "avatar" {
			t.Errorf("expected handler avatar for %s, got %s", r.path, r.handler)
		}
	}
	if !paths["/avatar/{name}.svg"] {
		t.Errorf("expected path param route /avatar/{name}.svg, got %v", paths)
	}
	if !paths["/avatar/blank.svg"] {
		t.Errorf("expected static route /avatar/blank.svg, got %v", paths)
	}
}
