//ff:type feature=scan type=model topic=fastapi
//ff:what 파싱된 Python 파일 정보 구조체
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// fileInfo holds parsed information for a single Python source file.
type fileInfo struct {
	absPath    string
	relPath    string
	src        []byte
	root       *sitter.Node
	imports    []importInfo
	prefixes   map[string]string   // router variable -> resolved prefix
	routerDeps map[string][]string // router variable -> middleware from APIRouter(dependencies=[...])
	hidden     map[string]bool     // router variable -> true if include_in_schema=False (router/include level)
	models     map[string][]pydanticField
}
