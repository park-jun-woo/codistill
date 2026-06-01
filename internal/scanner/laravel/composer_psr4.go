//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what composer.json의 autoload(+autoload-dev) psr-4 매핑(네임스페이스 접두→디렉터리)을 읽는다
package laravel

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// composerPSR4 reads composer.json at absRoot and returns its PSR-4 autoload
// map (namespace prefix -> directory), merging both autoload and autoload-dev.
// A psr-4 value may be a string or an array of strings; only the first
// directory of an array form is kept (the conventional primary location). The
// returned namespace prefixes retain their trailing "\\" exactly as composer
// declares them (e.g. "App\\" -> "app/"). Missing/invalid composer.json yields
// an empty map, which simply means PSR-4 resolution finds nothing and callers
// fall back to the hardcoded candidate paths.
func composerPSR4(absRoot string) map[string]string {
	data, err := os.ReadFile(filepath.Join(absRoot, "composer.json"))
	if err != nil {
		return map[string]string{}
	}
	var doc struct {
		Autoload struct {
			PSR4 map[string]json.RawMessage `json:"psr-4"`
		} `json:"autoload"`
		AutoloadDev struct {
			PSR4 map[string]json.RawMessage `json:"psr-4"`
		} `json:"autoload-dev"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return map[string]string{}
	}
	out := make(map[string]string)
	add := func(m map[string]json.RawMessage) {
		for ns, raw := range m {
			if dir := firstPSR4Dir(raw); dir != "" {
				if _, exists := out[ns]; !exists {
					out[ns] = dir
				}
			}
		}
	}
	add(doc.Autoload.PSR4)
	add(doc.AutoloadDev.PSR4)
	return out
}
