//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 경로 슬라이스에서 중복을 제거하고 첫 등장 순서를 보존한다
package nestjs

// dedupePaths returns paths with duplicates removed, preserving first-seen
// order. Used to merge recursive, monorepo and workspace prefix-file candidates
// without re-parsing the same file twice.
func dedupePaths(paths []string) []string {
	seen := make(map[string]bool, len(paths))
	var out []string
	for _, p := range paths {
		if seen[p] {
			continue
		}
		seen[p] = true
		out = append(out, p)
	}
	return out
}
