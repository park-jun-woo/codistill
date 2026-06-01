//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what relPath가 *ServiceProvider.php 파일인지 판정한다 (packages/**/src/Providers/* 등 모듈형 Provider 포함)
package laravel

import "strings"

// isPackageServiceProvider reports whether relPath is a *ServiceProvider.php
// file, the convention modular Laravel apps (e.g. Bagisto) use for package
// providers under packages/**/src/Providers/. These providers commonly load
// their routes via Route::...->group(__DIR__ . '/../Routes/api.php'), a
// directory-relative load that base_path() resolution does not cover.
func isPackageServiceProvider(relPath string) bool {
	p := strings.ReplaceAll(relPath, "\\", "/")
	return strings.HasSuffix(p, "ServiceProvider.php")
}
