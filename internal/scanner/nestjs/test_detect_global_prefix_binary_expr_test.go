//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_BinaryExpr 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_BinaryExpr(t *testing.T) {
	dir := t.TempDir()
	// urlPrefix is unresolved (runtime value); only the literal 'api' survives.
	mainTs := `
import { NestFactory } from '@nestjs/core';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix(urlPrefix + 'api');
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	if prefix := detectGlobalPrefix(dir); prefix != "api" {
		t.Fatalf("expected api, got %q", prefix)
	}
}
