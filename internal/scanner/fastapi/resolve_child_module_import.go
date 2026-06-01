//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what fi.imports에서 childModule과 일치하는 import를 찾아 서브모듈 경로를 해석한다
package fastapi

// resolveChildModuleImport searches fi.imports for an entry whose name matches
// childModule, appends childModule to the import's module path, and resolves
// the resulting sub-module path to a file system path. Returns "" if no
// matching import is found or the path cannot be resolved.
func resolveChildModuleImport(absRoot, referrerDir, childModule string, imports []importInfo) string {
	for _, imp := range imports {
		if imp.name != childModule {
			continue
		}
		subModule := joinModulePath(imp.module, childModule)
		resolved := resolveImportPath(referrerDir, subModule)
		if resolved == "" {
			resolved = resolveAbsoluteImportPath(absRoot, subModule)
		}
		return resolved
	}
	return ""
}
