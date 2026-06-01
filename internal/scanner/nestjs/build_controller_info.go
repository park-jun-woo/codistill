//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 단일 클래스 노드에서 컨트롤러 정보를 생성한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// buildControllerInfo builds a controllerInfo from a class node.
// It returns false if the class does not have a @Controller decorator.
// root is the file AST root and projectRoot the scan root; both are used to
// resolve enum member-expression paths (e.g. @Controller(RouteKey.Asset)).
func buildControllerInfo(cls *sitter.Node, src []byte, file string, absFile string, imports map[string]string, root *sitter.Node, projectRoot string) (controllerInfo, bool) {
	prefix, ok := controllerPrefix(cls, src)
	if !ok {
		return controllerInfo{}, false
	}
	pc := enumPathCtx{root: root, src: src, absFile: absFile, imports: imports, projectRoot: projectRoot}
	prefix = pc.resolve(prefix)
	ci := controllerInfo{
		prefix:  prefix,
		version: controllerVersion(cls, src),
		imports: imports,
	}
	collectClassLevelDecorators(cls, src, &ci)
	direct := extractMethods(cls, src, file)
	inherited := resolveBaseController(cls, src, absFile, imports, file)
	ci.endpoints = mergeEndpoints(inherited, direct)
	resolveEndpointEnumPaths(ci.endpoints, pc)
	applyClassDecorators(&ci)
	return ci, true
}
