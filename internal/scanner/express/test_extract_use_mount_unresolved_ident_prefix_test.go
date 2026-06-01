//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractUseMount_UnresolvedIdentPrefix 테스트: 해소 불가 식별자 prefix는 마운트 스킵(가짜 합성 안 함)
package express

import "testing"

func TestExtractUseMount_UnresolvedIdentPrefix(t *testing.T) {
	fi := mustParse(t, []byte(`app.lazyUse(UNKNOWN, userRouter);`))
	m := extractUseMount(firstCallExpr(t, fi), fi, map[string]bool{"app": true}, map[string]string{"userRouter": "./u.ts"}, "", nil)
	if m != nil {
		t.Fatalf("expected skip on unresolved ident prefix, got %+v", m)
	}
}
