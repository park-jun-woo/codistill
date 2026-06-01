//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what Flask 프로젝트를 스캔하여 엔드포인트를 추출한다
package flask

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a Flask project root and extracts endpoints.
// Pass 1: collect Blueprint structure (Blueprint instances, register_blueprint chains).
// Pass 2: extract routes from decorated handler functions.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	pyFiles, err := findPyFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding py files: %w", err)
	}
	if len(pyFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, pyFiles)
	if len(files) == 0 {
		return &scanner.ScanResult{}, nil
	}

	bpPrefixes := resolveBlueprintPrefixes(files)
	nsPrefixes := collectNamespacePrefix(files)
	catalog := collectResourceClasses(files)

	var endpoints []scanner.Endpoint
	for _, fi := range files {
		routes := extractRoutes(fi.root, fi.src, bpPrefixes, fi.relPath)
		for _, ri := range routes {
			endpoints = append(endpoints, buildEndpoint(ri))
		}
	}

	// Flask-RESTful class-based routes (api.add_resource / configure_api tuples).
	for _, ri := range extractResourceRoutes(files, catalog, bpPrefixes, nsPrefixes) {
		endpoints = append(endpoints, buildEndpoint(ri))
	}

	// flask_restx class-based routes (@ns.route on Resource subclasses).
	for _, ri := range extractClassRoutes(files, nsPrefixes) {
		endpoints = append(endpoints, buildEndpoint(ri))
	}

	// Flask-AppBuilder API routes (@expose + ModelRestApi CRUD on BaseApi/*RestApi).
	for _, ri := range extractAppbuilderRoutes(files) {
		endpoints = append(endpoints, buildEndpoint(ri))
	}

	// add_url_rule(rule, endpoint, view, methods=...) registrations
	// (Indico IndicoBlueprint + standard Flask function views).
	for _, ri := range extractAddURLRuleRoutes(files, bpPrefixes) {
		endpoints = append(endpoints, buildEndpoint(ri))
	}

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
