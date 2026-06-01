//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestResolveInheritedConstant 테스트 — extends/implements 체인의 bare 상수 해석
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolveInheritedConstant(t *testing.T) {
	dir := t.TempDir()

	// 상수 정의 인터페이스.
	writeFile(t, dir, "R.java", `package x;
interface R {
  String PAYMENT_METHODS = "paymentMethods";
}`)
	// R을 implements 하는 추상 베이스.
	writeFile(t, dir, "Base.java", `package x;
public abstract class Base implements R {
}`)
	// Base를 extends 하는 핸들러: bare 상수 PAYMENT_METHODS 사용.
	handler := `package x;
class AccountResource extends Base {
  @Path("/{accountId}/" + PAYMENT_METHODS + "/refresh") void refresh(){}
}`
	writeFile(t, dir, "AccountResource.java", handler)

	abs := filepath.Join(dir, "AccountResource.java")
	src := []byte(handler)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)

	methods := findAllByType(root, "method_declaration")
	var got string
	for _, m := range methods {
		ann := findAnnotation(m, src, AnnPath)
		if ann != nil {
			got = resolvePathArg(ann, src, imports, abs, dir)
		}
	}
	want := "/{accountId}/paymentMethods/refresh"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
