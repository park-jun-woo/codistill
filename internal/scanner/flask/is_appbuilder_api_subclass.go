//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 상위클래스 목록이 Flask-AppBuilder API(BaseApi/ModelRestApi/*RestApi)인지 판정한다
package flask

import "strings"

// isAppbuilderAPISubclass reports whether any superclass resolves to a
// Flask-AppBuilder API base: BaseApi, ModelRestApi, or any name ending in
// "RestApi" / "Api" (e.g. BaseSupersetModelRestApi, ChartRestApi). Dotted
// attribute paths (flask_appbuilder.api.ModelRestApi) and import aliases are
// resolved by taking the final component. The second return reports whether the
// class derives from a ModelRestApi-family base (name ends in "RestApi"), which
// triggers standard CRUD synthesis.
func isAppbuilderAPISubclass(supers []string, aliases importAlias) (isAPI, isModel bool) {
	for _, s := range supers {
		base := s
		if i := strings.LastIndex(s, "."); i >= 0 {
			base = s[i+1:]
		}
		if orig := aliases[base]; orig != "" {
			base = orig
		}
		if base == "BaseApi" {
			isAPI = true
			continue
		}
		if strings.HasSuffix(base, "RestApi") {
			isAPI = true
			isModel = true
			continue
		}
		if strings.HasSuffix(base, "Api") {
			isAPI = true
		}
	}
	return isAPI, isModel
}
