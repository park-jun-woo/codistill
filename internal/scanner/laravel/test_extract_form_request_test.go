//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what FormRequest rules() 추출 테스트
package laravel

import "testing"

func TestExtractFormRequest_Basic(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Requests/StoreUserRequest.php", `<?php
namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreUserRequest extends FormRequest {
    public function rules(): array {
        return [
            'name' => 'required|string|max:255',
            'email' => 'required|email|unique:users',
            'age' => 'nullable|integer|min:0|max:150',
            'role' => ['required', 'in:admin,user,editor'],
        ];
    }
}
`)
	parsedFiles := make(map[string]*fileInfo)
	fields := extractFormRequest(dir, "StoreUserRequest", nil, parsedFiles)
	if len(fields) != 4 {
		t.Fatalf("expected 4 fields, got %d", len(fields))
	}

	// name field
	if fields[0].Name != "name" {
		t.Errorf("field[0].Name = %q, want %q", fields[0].Name, "name")
	}
	if fields[0].Type != "string" {
		t.Errorf("field[0].Type = %q, want %q", fields[0].Type, "string")
	}
	if fields[0].MaxLength == nil || *fields[0].MaxLength != 255 {
		t.Errorf("field[0].MaxLength = %v, want 255", fields[0].MaxLength)
	}

	// age field
	if fields[2].Name != "age" {
		t.Errorf("field[2].Name = %q, want %q", fields[2].Name, "age")
	}
	if fields[2].Type != "integer" {
		t.Errorf("field[2].Type = %q, want %q", fields[2].Type, "integer")
	}
	if !fields[2].Nullable {
		t.Error("expected field[2].Nullable = true")
	}

	// role field (enum)
	if fields[3].Name != "role" {
		t.Errorf("field[3].Name = %q, want %q", fields[3].Name, "role")
	}
	if len(fields[3].Enum) != 3 {
		t.Errorf("field[3].Enum = %v, want 3 values", fields[3].Enum)
	}
}
