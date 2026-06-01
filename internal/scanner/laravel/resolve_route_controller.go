//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 라우트의 controller/action으로 컨트롤러 메서드를 해석한다
package laravel

func resolveRouteController(absRoot string, ri routeInfo, parsedFiles map[string]*fileInfo) *controllerMethod {
	if ri.controller == "" || ri.action == "" {
		return nil
	}
	srcFI := parsedFiles[ri.file]
	controllerFI := resolveController(absRoot, ri.controller, srcFI, parsedFiles)
	if controllerFI == nil {
		return nil
	}
	return extractControllerMethod(controllerFI, ri.action)
}
