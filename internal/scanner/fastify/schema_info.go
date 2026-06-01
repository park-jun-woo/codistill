//ff:type feature=scan type=model topic=fastify
//ff:what JSON Schema 추출 결과 구조체
package fastify

import sitter "github.com/smacker/go-tree-sitter"

type schemaInfo struct {
	Body        *sitter.Node
	Querystring *sitter.Node
	Params      *sitter.Node
	Response    map[string]*sitter.Node
	OperationID string
	Summary     string
	Description string
	Tags        []string
}
