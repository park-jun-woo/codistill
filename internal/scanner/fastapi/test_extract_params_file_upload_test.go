//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractParams_FileUpload 테스트
package fastapi

import "testing"

func TestExtractParams_FileUpload(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, UploadFile

app = FastAPI()

@app.post("/upload")
async def upload(file: UploadFile):
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routes := extractRoutes(root, src, prefixes, nil, nil, "main.py", nil)

	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if len(routes[0].files) != 1 {
		t.Fatalf("expected 1 file param, got %d", len(routes[0].files))
	}
	if routes[0].files[0].Name != "file" {
		t.Fatalf("expected file, got %s", routes[0].files[0].Name)
	}
}
