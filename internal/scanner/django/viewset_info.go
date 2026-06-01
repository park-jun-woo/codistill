//ff:type feature=scan type=model topic=django
//ff:what ViewSet 정보 구조체
package django

// viewsetInfo holds information about a DRF ViewSet class.
type viewsetInfo struct {
	name            string
	parents         []string       // parent class names
	methods         []actionMethod // CRUD methods resolved transitively through the class index
	serializerClass string         // serializer_class attribute value
	actions         []actionInfo
	file            string // relative file path
	line            int
}
