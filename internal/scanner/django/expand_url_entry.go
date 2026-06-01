//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 단일 urlEntry를 prefix와 결합하고 include면 하위 모듈로 재귀 전개한다
package django

// expandURLEntry composes a urlEntry with the current prefix; if it is an include,
// it recurses into the referenced module, otherwise it yields a single combined entry.
func expandURLEntry(entry urlEntry, prefix string, byModule map[string][]urlEntry, visited map[string]bool) []urlEntry {
	combined := combinePath(prefix, entry.pattern)
	if !entry.isInclude {
		return []urlEntry{{pattern: combined, viewName: entry.viewName, methodActions: entry.methodActions}}
	}
	if entry.includeRouterVar != "" {
		return []urlEntry{{pattern: combined, includeRouterVar: entry.includeRouterVar}}
	}
	if len(entry.includeInline) > 0 {
		var out []urlEntry
		for _, child := range entry.includeInline {
			out = append(out, expandURLEntry(child, combined, byModule, visited)...)
		}
		return out
	}
	mod, ok := resolveIncludeModule(entry.includeModule, byModule)
	if !ok {
		return nil
	}
	return expandURLModule(mod, combined, byModule, visited)
}
