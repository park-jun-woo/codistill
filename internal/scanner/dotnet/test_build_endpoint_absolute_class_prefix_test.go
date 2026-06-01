//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what 클래스 [Route("~/{cryptoCode}/x")]+메서드 [HttpGet("y")] => /{cryptoCode}/x/y end-to-end 가드
package dotnet

import "testing"

func TestBuildEndpointAbsoluteClassPrefix(t *testing.T) {
	root, src := parseCS(t, `[Route("~/{cryptoCode}/x")] class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	prefix := extractClassRoute(cls, src, "C")
	ci := controllerInfo{prefix: prefix, className: "C"}
	ep := endpointInfo{method: "GET", path: "y", handler: "H"}
	if got := buildEndpoint(ci, ep).Path; got != "/{cryptoCode}/x/y" {
		t.Errorf("absolute class prefix => %q, want /{cryptoCode}/x/y", got)
	}
}
