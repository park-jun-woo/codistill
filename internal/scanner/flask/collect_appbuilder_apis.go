//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 전 파일에서 Flask-AppBuilder API 클래스(base_url 포함)를 수집한다
package flask

// collectAppbuilderAPIs walks every parsed file and collects class_definitions
// that derive from a Flask-AppBuilder API base (BaseApi / ModelRestApi /
// *RestApi / *Api). Per-file collection is delegated to fileAppbuilderAPIs so
// import aliases are resolved once per file. Used by the Flask-AppBuilder route
// scanner (extractAppbuilderRoutes).
func collectAppbuilderAPIs(files []fileInfo) []appbuilderAPIInfo {
	var apis []appbuilderAPIInfo
	for _, fi := range files {
		apis = append(apis, fileAppbuilderAPIs(fi)...)
	}
	return apis
}
