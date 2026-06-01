//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_WorkspacePackage 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_WorkspacePackage(t *testing.T) {
	dir := t.TempDir()
	// Monorepo root: entry imports bootstrap from a workspace package; the
	// prefix call lives in that package's src. Workspace mapping must resolve
	// '@app/core' -> packages/core and reach the call site.
	writeFile(t, dir, "src/main.ts", `import { bootstrap } from '@app/core';
bootstrap();
`)
	writeFile(t, dir, "packages/core/package.json", `{"name":"@app/core"}`)
	writeFile(t, dir, "packages/core/src/lib/bootstrap/index.ts", `export async function bootstrap() {
  const globalPrefix = 'api';
  app.setGlobalPrefix(globalPrefix);
}
`)
	if prefix := detectGlobalPrefix(dir); prefix != "api" {
		t.Fatalf("expected api from workspace package, got %q", prefix)
	}
}
