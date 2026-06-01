//ff:type feature=scan type=model topic=quarkus
//ff:what 엔드포인트 추출 중간 결과 구조체
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

type endpointInfo struct {
	method        string
	path          string
	handler       string
	file          string
	line          int
	roles         []string
	params        []scanner.Param
	query         []scanner.Param
	headers       []scanner.Param
	bodyType      string
	bodyVarName   string
	bodyIsArray   bool
	formType      string
	formParams    []scanner.Param
	files         []scanner.Param
	statusCode    string
	returnType    string
	returnIsArray bool
}
