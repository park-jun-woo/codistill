//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestExtractUseMount_SameFileConstPrefix 테스트: 동일 파일 const 식별자 prefix 해소 + lazyUse 인식
package express

import "testing"

func TestExtractUseMount_SameFileConstPrefix(t *testing.T) {
	fi := mustParse(t, []byte("const P = '/ghost/api';\napp.lazyUse(P, userRouter);"))
	calls := findAllByType(fi.Root, "call_expression")
	var mount *useMount
	for _, c := range calls {
		if m := extractUseMount(c, fi, map[string]bool{"app": true}, map[string]string{"userRouter": "./u.ts"}, "", nil); m != nil {
			mount = m
		}
	}
	if mount == nil || mount.Prefix != "/ghost/api" || mount.VarName != "userRouter" {
		t.Fatalf("got %+v", mount)
	}
}
