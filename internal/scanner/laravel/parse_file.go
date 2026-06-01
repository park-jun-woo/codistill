//ff:func feature=scan type=parse control=sequence topic=laravel
//ff:what PHP 소스를 tree-sitter로 파싱하여 fileInfo를 반환한다 (과대 파일은 스킵)
package laravel

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
	php "github.com/smacker/go-tree-sitter/php"
)

// maxParseFileBytes caps the size of a single PHP file that parseFile will hand
// to tree-sitter. Abnormally large generated PHP files (e.g. a multi-megabyte
// data fixture) can dominate a whole scan; skipping them keeps one pathological
// file from stalling the run. The limit is generous — real route files,
// providers, controllers, requests and resources are far below it.
const maxParseFileBytes = 4 << 20 // 4 MiB

func parseFile(absRoot, absPath string) (*fileInfo, error) {
	src, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", absPath, err)
	}
	if len(src) > maxParseFileBytes {
		return nil, fmt.Errorf("skipping oversized file %s (%d bytes > %d cap)", absPath, len(src), maxParseFileBytes)
	}
	parser := sitter.NewParser()
	parser.SetLanguage(php.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse %s: %w", absPath, err)
	}
	relPath, _ := filepath.Rel(absRoot, absPath)
	return &fileInfo{
		absPath: absPath,
		relPath: relPath,
		src:     src,
		root:    tree.RootNode(),
	}, nil
}
