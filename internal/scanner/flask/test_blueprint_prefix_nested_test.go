//ff:func feature=scan type=test control=sequence topic=flask
//ff:what 중첩 register_blueprint가 부모+자식 prefix를 누적한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlueprintPrefix_Nested(t *testing.T) {
	dir := t.TempDir()
	src := `from flask import Flask, Blueprint

app = Flask(__name__)
parent = Blueprint("parent", __name__, url_prefix="/api")
child = Blueprint("child", __name__, url_prefix="/v1")

@child.route("/x")
def handler_x():
    return {}

parent.register_blueprint(child)
app.register_blueprint(parent)
`
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d: %v", len(result.Endpoints), result.Endpoints)
	}
	if result.Endpoints[0].Path != "/api/v1/x" {
		t.Errorf("expected /api/v1/x, got %s", result.Endpoints[0].Path)
	}
}
