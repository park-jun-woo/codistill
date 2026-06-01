//ff:func feature=scan type=test control=sequence topic=flask
//ff:what fastapi와 flask가 모두 있으면 fastapi가 우선 감지된다
package flask

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectFlask_WithFastAPI(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi==0.100.0\nflask==3.0.0\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	// FastAPI takes priority; should detect as exactly fastapi, not flask
	if fw != "fastapi" {
		t.Errorf("expected fastapi when both present, got %q", fw)
	}
}
