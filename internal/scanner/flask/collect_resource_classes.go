//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 전 파일에서 Resource 서브클래스를 수집해 클래스명→카탈로그 맵을 만든다
package flask

// collectResourceClasses walks every parsed file and collects class_definition
// nodes whose superclass list contains Flask-RESTful's Resource (directly,
// via dotted attribute like flask_restful.Resource, or via an import alias).
// Each qualifying class is recorded in the catalog keyed by class name with its
// HTTP method defs (classHTTPMethods) and source file. Used by Phase170's
// add_resource resolution.
func collectResourceClasses(files []fileInfo) resourceClassCatalog {
	catalog := make(resourceClassCatalog)
	for _, fi := range files {
		collectFileResourceClasses(fi, catalog)
	}
	return catalog
}
