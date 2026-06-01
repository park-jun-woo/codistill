//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestLastNamespaceSegment 테스트
package laravel

import "testing"

func TestLastNamespaceSegment(t *testing.T) {
	if got := lastNamespaceSegment("App\\Http\\Controllers\\UserController"); got != "UserController" {
		t.Errorf("got %q", got)
	}
	if got := lastNamespaceSegment("Plain"); got != "Plain" {
		t.Errorf("got %q", got)
	}
	if got := lastNamespaceSegment("App\\Foo\\"); got != "Foo" {
		t.Errorf("trailing slash: got %q", got)
	}
}
