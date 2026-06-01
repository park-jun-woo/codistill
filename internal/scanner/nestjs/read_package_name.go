//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what package.json 파일에서 name 필드를 읽는다
package nestjs

import (
	"encoding/json"
	"os"
)

// readPackageName reads the "name" field from the package.json at path.
// Returns "" when the file is missing, unreadable, invalid JSON, or has no
// name. Used to map a workspace package directory to its import name.
func readPackageName(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var pkg struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return ""
	}
	return pkg.Name
}
