//ff:func feature=scan type=convert control=sequence topic=django
//ff:what Django URL 패턴의 <type:name>·re_path 정규식(앵커/캡처그룹/문자클래스)을 OpenAPI {name}로 완전 정규화한다
package django

// djangoURLToOpenAPI converts Django URL patterns to OpenAPI path format.
// e.g. "users/<int:pk>/" -> "users/{pk}/"; re_path "^articles/(?P<year>[0-9]+)/$" -> "articles/{year}/";
// unnamed groups "^(\d+)/(.*)" -> "{param1}/{param2}"; escapes "^\.well-known/" -> ".well-known/".
func djangoURLToOpenAPI(path string) string {
	normalized, _ := normalizeDjangoPath(path)
	return normalized
}
