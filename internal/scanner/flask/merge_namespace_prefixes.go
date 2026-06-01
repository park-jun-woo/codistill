//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 파일별 ns→prefix 맵을 누적 맵에 병합한다(비어있지 않은 prefix 우선)
package flask

// mergeNamespacePrefixes merges per-file namespace prefixes into the accumulator.
// A non-empty prefix always wins; an existing entry is only filled when absent so
// a later empty registration does not clobber a known prefix.
func mergeNamespacePrefixes(acc, add namespacePrefix) {
	for nsVar, prefix := range add {
		_, exists := acc[nsVar]
		if !exists || prefix != "" {
			acc[nsVar] = prefix
		}
	}
}
