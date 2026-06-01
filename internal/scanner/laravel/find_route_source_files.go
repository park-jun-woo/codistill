//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 프로젝트 루트에서 라우트 소스 PHP 파일(routes/**, Providers/**, *ServiceProvider.php, **/Routes/**)만 수집한다
package laravel

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// findRouteSourceFiles walks the project root and returns only the PHP files
// that may define or load routes (see isRouteSourceFile). This is the stage-1
// input: parsing this narrow set — instead of every PHP file — is what keeps
// large apps (thousands of unrelated PHP files) within the scan time budget.
// Controllers/FormRequests/Resources are parsed lazily on demand in stage 2.
func findRouteSourceFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() {
			if skipDirs[name] || scanner.IsTestDir(name) {
				return filepath.SkipDir
			}
			return nil
		}
		if scanner.IsTestFile(name) {
			return nil
		}
		if !strings.HasSuffix(name, ".php") {
			return nil
		}
		relPath, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return nil
		}
		if isRouteSourceFile(relPath) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
