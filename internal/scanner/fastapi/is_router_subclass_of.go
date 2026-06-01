//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 부모 이름 목록 중 하나라도 APIRouter/알려진 라우터/기수집 서브클래스면 true
package fastapi

// isRouterSubclassOf reports whether any of the given parent class names refers
// to APIRouter, a known router class, or a class already known to be a router
// subclass.
func isRouterSubclassOf(parents []string, subclasses map[string]bool) bool {
	for _, p := range parents {
		if p == "APIRouter" || routerClassNames[p] || subclasses[p] {
			return true
		}
	}
	return false
}
