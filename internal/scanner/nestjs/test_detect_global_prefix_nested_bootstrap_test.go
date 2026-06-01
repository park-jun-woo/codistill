//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_NestedBootstrap 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_NestedBootstrap(t *testing.T) {
	dir := t.TempDir()
	// Gauzy-style: call site is nested below src/ (not src/ direct), with a
	// same-file const. Recursive search must reach it.
	writeFile(t, dir, "src/main.ts", `import { bootstrap } from './lib/bootstrap';
bootstrap();
`)
	writeFile(t, dir, "src/lib/bootstrap/index.ts", `export async function bootstrap() {
  const globalPrefix = 'api';
  app.setGlobalPrefix(globalPrefix);
}
`)
	if prefix := detectGlobalPrefix(dir); prefix != "api" {
		t.Fatalf("expected api from nested bootstrap, got %q", prefix)
	}
}
