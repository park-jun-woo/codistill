//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_ConfigFallback 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_ConfigFallback(t *testing.T) {
	dir := t.TempDir()
	// Runtime-dependent arg (no const, no literal) must not be code-resolved;
	// Phase044 .env fallback stays in effect.
	mainTs := `
import { NestFactory } from '@nestjs/core';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix(cfg.get('PREFIX'));
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	writeFile(t, dir, ".env.example", "API_PREFIX=api\n")
	if prefix := detectGlobalPrefix(dir); prefix != "api" {
		t.Fatalf("expected api from .env fallback, got %q", prefix)
	}
}
