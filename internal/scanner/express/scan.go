//ff:func feature=scan type=extract control=sequence topic=express
//ff:what Express 프로젝트를 스캔하여 엔드포인트를 추출한다 (2-pass)
package express

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	tsFiles, err := findTSFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding ts files: %w", err)
	}
	if len(tsFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}
	ctx := scanPass1(tsFiles, absRoot)
	endpoints := scanPass2(ctx, absRoot)
	// 데코레이터 라우팅(@RestController + @Get/@Post 등, @n8n/decorators 스타일)을
	// 별도 패스로 추출해 raw express 결과와 (method,path) 기준 합집합으로 반환한다.
	decoratorEndpoints := scanDecoratorPass(ctx, absRoot)
	endpoints = mergeDedupEndpoints(endpoints, decoratorEndpoints)
	// Medusa v2 파일기반 라우팅(src/api/**/route.ts + export const VERB)을 별도
	// 패스로 추출해 (method,path) 기준 합집합으로 반환한다.
	filebasedEndpoints := scanFilebasedPass(ctx, absRoot)
	endpoints = mergeDedupEndpoints(endpoints, filebasedEndpoints)
	// 커스텀 Controller 베이스 클래스(Unleash 스타일)의 this.route({...}) /
	// this.<method>("...") 라우트 등록을 별도 패스로 추출해 합집합으로 반환한다.
	controllerEndpoints := scanControllerPass(ctx, absRoot)
	endpoints = mergeDedupEndpoints(endpoints, controllerEndpoints)
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
