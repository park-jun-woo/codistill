//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what class_definition을 Flask-AppBuilder API 정보로 변환한다(비대상은 ok=false)
package flask

import sitter "github.com/smacker/go-tree-sitter"

// appbuilderAPIFromClass converts one class_definition into an appbuilderAPIInfo
// when it derives from a Flask-AppBuilder API base (isAppbuilderAPISubclass).
// ok is false for non-API classes or classes without a name. The base_url is
// resolved via classBaseURL and the ModelRestApi-family flag is carried through
// to trigger standard CRUD synthesis.
func appbuilderAPIFromClass(cls *sitter.Node, fi fileInfo, aliases importAlias) (appbuilderAPIInfo, bool) {
	isAPI, isModel := isAppbuilderAPISubclass(classSuperclasses(cls, fi.src), aliases)
	if !isAPI {
		return appbuilderAPIInfo{}, false
	}
	nameNode := findChildByType(cls, "identifier")
	if nameNode == nil {
		return appbuilderAPIInfo{}, false
	}
	name := nodeText(nameNode, fi.src)
	return appbuilderAPIInfo{
		name:           name,
		file:           fi.relPath,
		baseURL:        classBaseURL(cls, name, fi.src),
		isModelRestApi: isModel,
		node:           cls,
	}, true
}
