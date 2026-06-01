//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 모노레포 패키지 디렉터리의 package.json#name과 디렉터리를 매핑한다
package nestjs

// collectWorkspacePackages builds a map from workspace package name (the
// package.json#name field, e.g. "@gauzy/core") to that package's directory,
// by scanning common monorepo layout dirs (packages/*, apps/*, libs/*) under
// root (addLayoutPackageNames). This lets a workspace import like '@gauzy/core'
// be resolved to packages/core. Returns an empty map when no layout dirs or
// named packages are present.
func collectWorkspacePackages(root string) map[string]string {
	result := make(map[string]string)
	for _, layout := range []string{"packages", "apps", "libs"} {
		addLayoutPackageNames(root, layout, result)
	}
	return result
}
