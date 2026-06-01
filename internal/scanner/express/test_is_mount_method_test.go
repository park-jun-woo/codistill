//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsMountMethod 테스트
package express

import "testing"

func TestIsMountMethod(t *testing.T) {
	if !isMountMethod("use") || !isMountMethod("lazyUse") {
		t.Fatal("use/lazyUse should be mount methods")
	}
	if isMountMethod("get") || isMountMethod("someMethod") {
		t.Fatal("non-mount methods must not be accepted")
	}
}
