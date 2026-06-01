//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindResourceFile_Parsed 테스트
package laravel

import "testing"

func TestFindResourceFile_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserResource {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if findResourceFile("/root", "UserResource", nil, parsed) == nil {
		t.Fatal("expected found")
	}
}
