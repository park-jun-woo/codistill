//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what JsonResource toArray() 응답 필드 추출 테스트
package laravel

import "testing"

func TestExtractResourceFields(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Resources/UserResource.php", `<?php
namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource {
    public function toArray($request): array {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'email' => $this->email,
            'created_at' => $this->created_at->toISOString(),
        ];
    }
}
`)
	parsedFiles := make(map[string]*fileInfo)
	fields := extractResourceFields(dir, "UserResource", nil, parsedFiles)
	if len(fields) != 4 {
		t.Fatalf("expected 4 fields, got %d", len(fields))
	}
	names := []string{"id", "name", "email", "created_at"}
	for i, name := range names {
		if fields[i].Name != name {
			t.Errorf("field[%d].Name = %q, want %q", i, fields[i].Name, name)
		}
	}
}
