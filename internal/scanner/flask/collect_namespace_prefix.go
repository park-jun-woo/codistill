//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 전 파일에서 api.add_namespace(ns, path=) 호출로 ns 변수→prefix 맵을 만든다
package flask

// collectNamespacePrefix walks every parsed file for api.add_namespace(ns, path="/x")
// calls and builds a flask_restx Namespace variable -> prefix map. The first
// positional identifier argument is the namespace variable; the prefix is taken
// from a string positional argument or the path= keyword (e.g.
// api.add_namespace(ns, "/x") or api.add_namespace(ns, path="/x")). Namespaces
// registered without a path carry an empty prefix. Used by the @ns.route
// class-based route scanner to compose namespace_prefix + route_path.
func collectNamespacePrefix(files []fileInfo) namespacePrefix {
	prefixes := make(namespacePrefix)
	for _, fi := range files {
		mergeNamespacePrefixes(prefixes, fileNamespacePrefixes(fi.root, fi.src))
	}
	return prefixes
}
