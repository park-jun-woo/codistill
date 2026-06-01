//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what schema pair의 키에 따라 schemaInfo의 해당 섹션에 값을 할당한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func assignSchemaSection(si *schemaInfo, pair *sitter.Node, src []byte) {
	key := pairKeyName(pair, src)
	val := pairValueNode(pair)
	if val == nil {
		return
	}
	switch key {
	case "body":
		si.Body = val
	case "querystring":
		si.Querystring = val
	case "params":
		si.Params = val
	case "response":
		collectResponseSchemas(si, val, src)
	case "operationId":
		if val.Type() == "string" {
			si.OperationID = unquoteTS(nodeText(val, src))
		}
	case "summary":
		if val.Type() == "string" {
			si.Summary = unquoteTS(nodeText(val, src))
		}
	case "description":
		if val.Type() == "string" {
			si.Description = unquoteTS(nodeText(val, src))
		}
	case "tags":
		if val.Type() == "array" {
			si.Tags = collectStringArray(val, src)
		}
	}
}
