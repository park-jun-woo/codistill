//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 한 파일에서 Resource 서브클래스를 찾아 카탈로그에 등록한다
package flask

// collectFileResourceClasses scans one parsed file for Resource subclasses and
// records each into the shared catalog keyed by class name. Import aliases are
// resolved once per file so aliased Resource imports are recognized.
func collectFileResourceClasses(fi fileInfo, catalog resourceClassCatalog) {
	aliases := collectImportAliases(fi.root, fi.src)
	for _, cls := range findAllByType(fi.root, "class_definition") {
		if !isResourceSubclass(classSuperclasses(cls, fi.src), aliases) {
			continue
		}
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, fi.src)
		catalog[name] = resourceClassInfo{
			name:    name,
			file:    fi.relPath,
			methods: classHTTPMethods(cls, fi.src),
			node:    cls,
		}
	}
}
