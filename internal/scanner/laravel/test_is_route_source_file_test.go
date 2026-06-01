//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestIsRouteSourceFile 테스트
package laravel

import "testing"

func TestIsRouteSourceFile(t *testing.T) {
	cases := []struct {
		rel  string
		want bool
	}{
		{"routes/api.php", true},
		{"routes/web/hardware.php", true},
		{"app/Providers/RouteServiceProvider.php", true},
		{"app/Providers/Route.php", true},
		{"packages/Webkul/Shop/src/Providers/ShopServiceProvider.php", true},
		{"packages/Webkul/Shop/src/Routes/api.php", true},
		{"app/Http/Controllers/UserController.php", false},
		{"app/Models/User.php", false},
		{"database/migrations/2024_create_users.php", false},
		{"resources/views/index.php", false},
	}
	for _, c := range cases {
		if got := isRouteSourceFile(c.rel); got != c.want {
			t.Errorf("isRouteSourceFile(%q) = %v, want %v", c.rel, got, c.want)
		}
	}
}
