//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildAllEndpoints_WithController 테스트
package nestjs

import "testing"

func TestBuildAllEndpoints_WithController(t *testing.T) {
	dir := t.TempDir()
	ctrl := `
import { Controller, Get } from '@nestjs/common';
@Controller('users')
export class UsersController {
  @Get()
  findAll() {}
}
`
	writeFile(t, dir, "src/users.controller.ts", ctrl)
	files := []string{dir + "/src/users.controller.ts"}
	controllers := collectControllers(files, dir)
	eps, _ := buildAllEndpoints("api", false, "", controllers, dir)
	if len(eps) != 1 {
		t.Fatalf("expected 1, got %d", len(eps))
	}
}
