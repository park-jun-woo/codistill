//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what lazyUse 마운트 체인 테스트(Ghost): 식별자 prefix(import 상수) + 인라인 require + 내부 lazyUse → /ghost/api/content/posts
package express

import "testing"

func TestLazyUseMountChain(t *testing.T) {
	dir := t.TempDir()

	// 크로스파일 상수 정의 + 재노출
	writeFile(t, dir, "shared/url-utils.ts", `
const BASE_API_PATH = '/ghost/api';
module.exports.BASE_API_PATH = BASE_API_PATH;
`)

	// content 라우터
	writeFile(t, dir, "api/content/routes.ts", `
const express = require("express");
const router = express.Router();
router.get('/posts', listPosts);
module.exports = router;
`)

	// api 앱: 식별자 string 리터럴 prefix + 인라인 require
	writeFile(t, dir, "api/app.ts", `
const express = require("express");
const apiApp = express();
apiApp.lazyUse('/content/', require('./content/routes'));
module.exports = apiApp;
`)

	// backend: 식별자 prefix(import 상수) + 인라인 require
	writeFile(t, dir, "backend.ts", `
const express = require("express");
const {BASE_API_PATH} = require('./shared/url-utils');
const backendApp = express();
backendApp.lazyUse(BASE_API_PATH, require('./api/app'));
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	if !found["GET /ghost/api/content/posts"] {
		t.Fatalf("expected GET /ghost/api/content/posts, got %v", found)
	}
}
