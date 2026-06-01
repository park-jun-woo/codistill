//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestEvalPathExpr 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestEvalPathExpr(t *testing.T) {
	dir := t.TempDir()
	// 별도 클래스 파일의 상수가 결합식(PREFIX+"/"+USAGES)인 Kill-Bill 핵심 케이스.
	jaxrs := `package x;
public class JaxrsResource {
  public static final String API_VERSION="/1.0";
  public static final String API_POSTFIX="/kb";
  public static final String PREFIX=API_VERSION+API_POSTFIX;
  public static final String USAGES="usages";
  public static final String USAGES_PATH=PREFIX+"/"+USAGES;
}`
	writeFile(t, dir, "JaxrsResource.java", jaxrs)

	referrer := `package x;
public class UsageResource {
  @Path(JaxrsResource.USAGES_PATH) void g(){}
}`
	writeFile(t, dir, "UsageResource.java", referrer)
	abs := filepath.Join(dir, "UsageResource.java")
	src := []byte(referrer)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)

	ann := findAnnotation(findAllByType(root, "method_declaration")[0], src, AnnPath)
	got := resolvePathArg(ann, src, imports, abs, dir)
	if got != "/1.0/kb/usages" {
		t.Fatalf("scoped recursive constant: got %q want /1.0/kb/usages", got)
	}
}
