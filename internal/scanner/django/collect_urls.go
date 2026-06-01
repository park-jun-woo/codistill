//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urls.py에서 모듈별 urlpatterns 맵을 수집한다
package django

// collectURLs extracts urlpatterns from all parsed files, keyed by dotted module path.
func collectURLs(files []fileInfo) map[string][]urlEntry {
	byModule := make(map[string][]urlEntry)
	for _, fi := range files {
		entries := collectURLsFromFile(fi)
		if len(entries) == 0 {
			continue
		}
		qualifyRouterVars(entries, fi.module)
		byModule[fi.module] = append(byModule[fi.module], entries...)
	}
	return byModule
}
