//ff:func feature=scan type=parse control=sequence topic=django
//ff:what 바이트가 식별자 시작 문자(A-Za-z_)인지 판정한다
package django

// isIdentStart reports whether c can start a Python identifier: [A-Za-z_].
func isIdentStart(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}
