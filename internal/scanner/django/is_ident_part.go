//ff:func feature=scan type=parse control=sequence topic=django
//ff:what 바이트가 식별자 구성 문자(A-Za-z0-9_)인지 판정한다
package django

// isIdentPart reports whether c can appear after the first char of a Python
// identifier: [A-Za-z0-9_].
func isIdentPart(c byte) bool {
	return isIdentStart(c) || (c >= '0' && c <= '9')
}
