//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 call_expression에서 app.use/lazyUse("/prefix", router) 마운트 패턴을 감지하고 마운트 정보를 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractUseMount(call *sitter.Node, fi *fileInfo, routers map[string]bool, imports map[string]string, absRoot string, aliases map[string]string) *useMount {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	objName := nodeText(obj, fi.Src)
	if !routers[objName] {
		return nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	if !isMountMethod(nodeText(prop, fi.Src)) {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	m := parseUseMountArgs(args, fi, imports, absRoot, aliases)
	if m != nil {
		m.SourceRouter = objName
	}
	return m
}
