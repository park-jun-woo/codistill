//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 모노레포 packages/*/src 및 apps/*/src에서 prefix 설정 파일을 모은다
package nestjs

// collectMonorepoPrefixFiles scans common monorepo layout dirs (packages/* and
// apps/*) under root and returns prefix-setting .ts files found by recursing
// each child package's src/ (collectLayoutPrefixFiles). This covers the case
// where the scan root is a monorepo root and the bootstrap call lives in a
// sibling package's src tree. Absent layout dirs yield nil.
func collectMonorepoPrefixFiles(root string) []string {
	var files []string
	for _, layout := range []string{"packages", "apps"} {
		files = append(files, collectLayoutPrefixFiles(root, layout)...)
	}
	return files
}
