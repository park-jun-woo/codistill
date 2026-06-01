//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_ConstIdentifier 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_ConstIdentifier(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
import { NestFactory } from '@nestjs/core';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const globalPrefix = 'api';
  app.setGlobalPrefix(globalPrefix);
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	if prefix := detectGlobalPrefix(dir); prefix != "api" {
		t.Fatalf("expected api, got %q", prefix)
	}
}
