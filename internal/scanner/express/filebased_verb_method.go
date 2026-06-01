//ff:func feature=scan type=extract control=sequence topic=express
//ff:what export 식별자가 httpMethods 화이트리스트의 대문자 VERB인지 판정하고 해당 메서드를 반환한다
package express

import "strings"

// filebasedVerbMethod reports whether ident is an uppercase HTTP verb that
// exists in the httpMethods whitelist (GET/POST/PUT/PATCH/DELETE), returning the
// uppercase method. The identifier must be exactly the uppercase verb so that
// non-method exports like "config" or "AUTHENTICATE" are rejected.
func filebasedVerbMethod(ident string) (string, bool) {
	lower := strings.ToLower(ident)
	method, ok := httpMethods[lower]
	if !ok || method == "all" {
		return "", false
	}
	if ident != method {
		return "", false
	}
	return method, true
}
