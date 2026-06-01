//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveConstStringIdentifier_CrossFile 테스트 (import된 const 경로 해석)
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveConstStringIdentifier_CrossFile(t *testing.T) {
	dir := t.TempDir()

	constSrc := `export const HEALTH_CHECK_ROUTE = 'health';`
	writeFile(t, dir, "src/constants.ts", constSrc)

	ctrlSrc := `
import { HEALTH_CHECK_ROUTE } from '../constants';

@Controller(HEALTH_CHECK_ROUTE)
export class HealthController {
  @Get()
  check() {}
}
`
	ctrlFile := filepath.Join(dir, "src/health/health.controller.ts")
	writeFile(t, dir, "src/health/health.controller.ts", ctrlSrc)

	root, err := parseTypeScript([]byte(ctrlSrc))
	if err != nil {
		t.Fatal(err)
	}
	controllers := extractControllers(root, []byte(ctrlSrc), "src/health/health.controller.ts", ctrlFile, dir)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	if controllers[0].prefix != "health" {
		t.Fatalf("cross-file prefix: want %q got %q", "health", controllers[0].prefix)
	}
}
