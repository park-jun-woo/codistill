//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectURIVersioning_Found 테스트
package nestjs

import "testing"

func TestDetectURIVersioning_Found(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
import { VersioningType } from '@nestjs/common';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableVersioning({ type: VersioningType.URI });
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	if enabled, _ := detectURIVersioning(dir); !enabled {
		t.Fatal("expected URI versioning to be detected")
	}
}
