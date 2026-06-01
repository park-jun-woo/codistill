//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestResolvePathArgInterfaceConstant 테스트 — cross-class 인터페이스 상수 + 다중 결합 해석
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolvePathArgInterfaceConstant(t *testing.T) {
	dir := t.TempDir()

	// JaxrsResource: 인터페이스 상수(암묵적 static final) + 재귀 결합 + 정규식 토큰.
	resourceIface := `package x;
interface R {
  String API_VERSION = "/1.0";
  String API_POSTFIX = "/kb";
  String PREFIX = API_VERSION + API_POSTFIX;
  String ACCOUNTS = "accounts";
  String ACCOUNTS_PATH = PREFIX + "/" + ACCOUNTS;
  String UUID_PATTERN = "\\w+-\\w+-\\w+-\\w+-\\w+";
}`
	writeFile(t, dir, "R.java", resourceIface)

	handler := `package x;
class AccountResource {
  @Path(R.ACCOUNTS_PATH) void cls(){}
  @Path(R.ACCOUNTS_PATH + "/{id}") void get(){}
  @Path("/{accountId:" + R.UUID_PATTERN + "}") void byId(){}
}`
	writeFile(t, dir, "AccountResource.java", handler)
	abs := filepath.Join(dir, "AccountResource.java")
	src := []byte(handler)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)

	want := map[string]string{
		"cls":  "/1.0/kb/accounts",
		"get":  "/1.0/kb/accounts/{id}",
		"byId": "/{accountId}",
	}
	methods := findAllByType(root, "method_declaration")
	seen := map[string]bool{}
	for _, m := range methods {
		nameNode := findChildByType(m, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, src)
		w, ok := want[name]
		if !ok {
			continue
		}
		seen[name] = true
		ann := findAnnotation(m, src, AnnPath)
		var got string
		if ann != nil {
			got = resolvePathArg(ann, src, imports, abs, dir)
		}
		if got != w {
			t.Errorf("method %s: got %q want %q", name, got, w)
		}
	}
	for n := range want {
		if !seen[n] {
			t.Errorf("method %s not found in AST", n)
		}
	}
}
