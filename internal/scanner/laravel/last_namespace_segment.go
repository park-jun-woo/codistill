//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what FQCN의 마지막 \-구분 세그먼트(단축 클래스명)를 반환한다
package laravel

import "strings"

// lastNamespaceSegment returns the final \-separated segment of a fully-qualified
// class name (e.g. "App\\Http\\Controllers\\UserController" -> "UserController").
func lastNamespaceSegment(fqcn string) string {
	fqcn = strings.TrimRight(fqcn, "\\")
	if i := strings.LastIndex(fqcn, "\\"); i >= 0 {
		return fqcn[i+1:]
	}
	return fqcn
}
