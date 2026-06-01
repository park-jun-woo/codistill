//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectNamespacePrefix가 add_namespace의 path 키워드/위치 인자로 ns→prefix 맵을 만드는지 검증한다
package flask

import "testing"

func TestCollectNamespacePrefix(t *testing.T) {
	src := []byte(`api.add_namespace(inner_api_ns, path="/inner/api")
api.add_namespace(positional_ns, "/pos")
api.add_namespace(bare_ns)
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	files := []fileInfo{{relPath: "x.py", src: src, root: root}}
	got := collectNamespacePrefix(files)
	if got["inner_api_ns"] != "/inner/api" {
		t.Errorf("inner_api_ns: got %q want /inner/api", got["inner_api_ns"])
	}
	if got["positional_ns"] != "/pos" {
		t.Errorf("positional_ns: got %q want /pos", got["positional_ns"])
	}
	if v, ok := got["bare_ns"]; !ok || v != "" {
		t.Errorf("bare_ns: got %q ok=%v want empty", v, ok)
	}
}
