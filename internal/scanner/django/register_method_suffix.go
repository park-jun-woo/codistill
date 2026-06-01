//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what attribute 호출 텍스트가 알려진 register 메서드인지 판정하고 router 변수를 반환한다
package django

import "strings"

// registerMethodSuffixes are the attribute method names treated as router
// registrations: DRF's `.register` and custom router aliases such as Wagtail's
// `WagtailAPIRouter.register_endpoint`.
var registerMethodSuffixes = []string{".register", ".register_endpoint"}

// registerRouterVar returns the router variable for an attribute call text if it
// ends with a known register method suffix, plus whether it matched. e.g.
// "router.register_endpoint" -> ("router", true).
func registerRouterVar(text string) (string, bool) {
	for _, suffix := range registerMethodSuffixes {
		if strings.HasSuffix(text, suffix) {
			return strings.TrimSuffix(text, suffix), true
		}
	}
	return "", false
}
