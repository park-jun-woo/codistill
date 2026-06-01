//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseWrapperRegisterCall 테스트
package django

import "testing"

func TestParseWrapperRegisterCall(t *testing.T) {
	src := "" +
		"def reg(prefix, viewset, **kwargs):\n" +
		"    router.register(prefix, viewset)\n" +
		"\n" +
		"reg(r'insights', EnterpriseInsightsViewSet)\n"
	fi := newTestFileInfo(t, src)

	wrappers := collectRegisterWrappersFromFile(fi)
	if !wrappers["reg"] {
		t.Fatalf("expected reg to be detected as wrapper, got %v", wrappers)
	}

	regs := extractWrapperRegisterCalls([]fileInfo{fi}, wrappers)
	if len(regs) != 1 {
		t.Fatalf("expected 1 wrapper registration, got %d", len(regs))
	}
	if regs[0].prefix != "insights" {
		t.Errorf("prefix: %q", regs[0].prefix)
	}
	if regs[0].viewsetName != "EnterpriseInsightsViewSet" {
		t.Errorf("viewsetName: %q", regs[0].viewsetName)
	}
}
