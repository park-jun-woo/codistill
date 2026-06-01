//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 한 파일에서 Flask-AppBuilder API 클래스를 찾아 정보 슬라이스로 반환한다
package flask

// fileAppbuilderAPIs scans one parsed file for Flask-AppBuilder API subclasses
// (isAppbuilderAPISubclass), resolving import aliases once. Each qualifying class
// yields an appbuilderAPIInfo with its resolved base_url (classBaseURL) and the
// ModelRestApi-family flag that drives standard CRUD synthesis.
func fileAppbuilderAPIs(fi fileInfo) []appbuilderAPIInfo {
	aliases := collectImportAliases(fi.root, fi.src)
	var apis []appbuilderAPIInfo
	for _, cls := range findAllByType(fi.root, "class_definition") {
		api, ok := appbuilderAPIFromClass(cls, fi, aliases)
		if ok {
			apis = append(apis, api)
		}
	}
	return apis
}
