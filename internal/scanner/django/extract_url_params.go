//ff:func feature=scan type=extract control=sequence topic=django
//ff:what Django URL 패턴(path 및 re_path)에서 path parameter를 추출한다(무명 캡처그룹은 {paramN})
package django

// extractURLParams extracts URL variable definitions from a Django URL pattern.
// e.g. "users/<int:pk>/posts/<slug:slug>" -> [{name:"pk", converter:"int"}, {name:"slug", converter:"slug"}];
// re_path named groups "(?P<year>[0-9]+)" -> [{name:"year"}];
// unnamed groups "^(\d+)/(.*)" -> [{name:"param1"}, {name:"param2"}].
// Param names and order match djangoURLToOpenAPI exactly via the shared
// normalizeDjangoPath pipeline.
func extractURLParams(path string) []urlParam {
	_, params := normalizeDjangoPath(path)
	return params
}
