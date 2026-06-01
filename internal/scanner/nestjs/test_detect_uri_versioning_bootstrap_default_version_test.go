//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectURIVersioning_BootstrapDefaultVersion 테스트
package nestjs

import "testing"

// TestDetectURIVersioning_BootstrapDefaultVersion mirrors Novu's layout:
// src/main.ts only calls bootstrap(), while enableVersioning with defaultVersion
// lives in src/bootstrap.ts. The recursive search must find it and extract '1'.
func TestDetectURIVersioning_BootstrapDefaultVersion(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
import { bootstrap } from './bootstrap';
bootstrap();
`
	bootstrapTs := `
import { VersioningType } from '@nestjs/common';
export async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableVersioning({ type: VersioningType.URI, defaultVersion: '1' });
  await app.listen(3000);
}
`
	writeFile(t, dir, "src/main.ts", mainTs)
	writeFile(t, dir, "src/bootstrap.ts", bootstrapTs)

	enabled, defaultVersion := detectURIVersioning(dir)
	if !enabled {
		t.Fatal("expected URI versioning to be detected from src/bootstrap.ts")
	}
	if defaultVersion != "1" {
		t.Fatalf("expected defaultVersion '1', got %q", defaultVersion)
	}
}
