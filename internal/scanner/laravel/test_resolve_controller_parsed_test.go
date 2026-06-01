//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveController_Parsed 테스트
package laravel

import "testing"

func TestResolveController_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if resolveController("/root", "UserController", nil, parsed) == nil {
		t.Fatal("expected via parsed")
	}
}
