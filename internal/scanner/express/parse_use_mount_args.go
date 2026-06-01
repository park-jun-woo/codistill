//ff:func feature=scan type=extract control=selection topic=express
//ff:what app.use()/lazyUse() 인자에서 prefix(string/식별자상수)와 마운트 대상(라우터 변수 or 인라인 require)을 파싱한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func parseUseMountArgs(args *sitter.Node, fi *fileInfo, imports map[string]string, absRoot string, aliases map[string]string) *useMount {
	argNodes := collectArgNodes(args)
	if len(argNodes) < 2 {
		return nil
	}
	prefix, ok := resolveMountPrefix(argNodes[0], fi, absRoot, aliases)
	if !ok {
		return nil
	}
	routerNode := argNodes[len(argNodes)-1]
	switch routerNode.Type() {
	case "identifier":
		varName := nodeText(routerNode, fi.Src)
		return &useMount{
			Prefix:   prefix,
			VarName:  varName,
			FilePath: imports[varName],
		}
	case "call_expression":
		// 인라인 require('...') / import('...') 마운트: 대상 파일을 직접 해소한다.
		filePath := resolveInlineRequirePath(routerNode, fi.Src, fi, absRoot, aliases)
		if filePath == "" {
			return nil
		}
		return &useMount{
			Prefix:   prefix,
			VarName:  "",
			FilePath: filePath,
		}
	}
	return nil
}
