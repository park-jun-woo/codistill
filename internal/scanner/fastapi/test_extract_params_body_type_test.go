//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractParams_BodyType 테스트
package fastapi

import "testing"

func TestExtractParams_BodyType(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI
from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
    email: str

app = FastAPI()

@app.post("/users", status_code=201)
async def create_user(user: UserCreate):
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
	if routes[0].bodyType != "UserCreate" {
		t.Fatalf("expected UserCreate, got %s", routes[0].bodyType)
	}
}
