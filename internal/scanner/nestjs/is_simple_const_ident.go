//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 경로 인자가 상수 해석 대상 단순 식별자인지 검사한다
package nestjs

// isSimpleConstIdent reports whether arg is a bare JS/TS identifier
// (no dot, paren, bracket, backtick, quote or whitespace) that can name a
// const string declaration — e.g. "HEALTH_CHECK_ROUTE". Enum member
// expressions ("RouteKey.Asset") and string literals are excluded so this
// guard only matches the simple const-reference case. The first rune may not
// be a digit (rejects numbers like "200").
func isSimpleConstIdent(arg string) bool {
	if arg == "" {
		return false
	}
	for i, r := range arg {
		letter := r == '_' || r == '$' ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z')
		digit := r >= '0' && r <= '9'
		ok := letter || (digit && i > 0)
		if !ok {
			return false
		}
	}
	return true
}
