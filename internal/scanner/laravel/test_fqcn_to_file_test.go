//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFqcnToFile 테스트
package laravel

import "testing"

func TestFqcnToFile(t *testing.T) {
	psr4 := map[string]string{
		"App\\":        "app/",
		"App\\Domain\\": "src/Domain/",
	}
	got, ok := fqcnToFile("App\\Http\\Controllers\\Api\\UserController", psr4)
	if !ok || got != "app/Http/Controllers/Api/UserController.php" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
	// longest prefix wins
	got, ok = fqcnToFile("App\\Domain\\Order", psr4)
	if !ok || got != "src/Domain/Order.php" {
		t.Fatalf("longest prefix: got %q ok=%v", got, ok)
	}
	// no matching prefix
	if _, ok := fqcnToFile("Vendor\\Pkg\\Thing", psr4); ok {
		t.Fatal("expected no match for unknown namespace")
	}
}
