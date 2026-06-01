//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectURIVersioning_NotFound 테스트
package nestjs

import "testing"

func TestDetectURIVersioning_NotFound(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	if enabled, _ := detectURIVersioning(dir); enabled {
		t.Fatal("expected no URI versioning")
	}
}
