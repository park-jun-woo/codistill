//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 두 엔드포인트 집합을 (method,path) 기준으로 중복 제거하며 합친다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

// mergeDedupEndpoints returns base followed by the entries of extra whose
// (Method, Path) pair is not already present in base. base entries are always
// kept as-is so the raw express extraction path is never altered.
func mergeDedupEndpoints(base, extra []scanner.Endpoint) []scanner.Endpoint {
	seen := make(map[string]bool, len(base))
	for _, ep := range base {
		seen[ep.Method+" "+ep.Path] = true
	}
	result := base
	for _, ep := range extra {
		key := ep.Method + " " + ep.Path
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, ep)
	}
	return result
}
