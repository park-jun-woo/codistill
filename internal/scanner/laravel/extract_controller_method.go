//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 컨트롤러 파일에서 지정 메서드를 찾아 controllerMethod로 추출한다
package laravel

// extractControllerMethod finds and extracts a specific method from a controller file.
func extractControllerMethod(fi *fileInfo, methodName string) *controllerMethod {
	method := findAnyClassMethod(fi, methodName)
	if method == nil {
		return nil
	}
	cm := &controllerMethod{
		name:  methodName,
		src:   fi.src,
		srcFI: fi,
	}
	if formalParams := findChildByType(method, "formal_parameters"); formalParams != nil {
		cm.params = extractMethodParams(formalParams, fi.src)
		cm.formRequestRef = findFormRequestParam(cm.params)
	}
	cm.returnNodes = findAllByType(method, "return_statement")
	cm.methodNode = method
	return cm
}
