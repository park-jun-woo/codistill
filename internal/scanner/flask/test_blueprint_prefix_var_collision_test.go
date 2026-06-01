//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what 동일 변수명 bp가 다른 패키지에서 재정의돼도 라우트가 자기 패키지 prefix를 유지한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlueprintPrefix_VarNameCollision(t *testing.T) {
	dir := t.TempDir()
	// Two packages both bind a blueprint to the local name "bp" with different
	// prefixes; the trigger route must keep "/triggers", not the other "/api".
	triggerInit := `from flask import Blueprint

bp = Blueprint("trigger", __name__, url_prefix="/triggers")

from . import webhook
`
	triggerWebhook := `from controllers.trigger import bp

@bp.route("/webhook/<webhook_id>")
def handle_webhook(webhook_id):
    return {}
`
	webInit := `from flask import Blueprint

bp = Blueprint("web", __name__, url_prefix="/api")

@bp.route("/ping")
def ping():
    return {}
`
	os.MkdirAll(filepath.Join(dir, "controllers", "trigger"), 0o755)
	os.MkdirAll(filepath.Join(dir, "controllers", "web"), 0o755)
	os.WriteFile(filepath.Join(dir, "controllers", "trigger", "__init__.py"), []byte(triggerInit), 0o644)
	os.WriteFile(filepath.Join(dir, "controllers", "trigger", "webhook.py"), []byte(triggerWebhook), 0o644)
	os.WriteFile(filepath.Join(dir, "controllers", "web", "__init__.py"), []byte(webInit), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{
		"/triggers/webhook/{webhook_id}": false,
		"/api/ping":                      false,
	}
	for _, ep := range result.Endpoints {
		if _, ok := want[ep.Path]; ok {
			want[ep.Path] = true
		}
	}
	for p, seen := range want {
		if !seen {
			t.Errorf("expected endpoint path %s, got %v", p, result.Endpoints)
		}
	}
}
