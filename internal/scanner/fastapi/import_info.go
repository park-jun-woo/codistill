//ff:type feature=scan type=model topic=fastapi
//ff:what import 정보 구조체
package fastapi

// importInfo maps imported names to their source module path.
type importInfo struct {
	name     string // imported (local) name (e.g., "UserCreate" or alias "foo")
	module   string // module path (e.g., ".models" or "app.models")
	origName string // original name in the source module (e.g., "router" for `import router as foo`); defaults to name
}
