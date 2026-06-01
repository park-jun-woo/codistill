//ff:func feature=scan type=test control=sequence topic=django
//ff:what 커스텀 베이스를 통한 CRUD 메서드 전이 해소를 검증한다
package django

import "testing"

func TestResolveMethodsTransitive(t *testing.T) {
	idx := classIndex{
		"ProjectViewSet":  {"ModelCrudViewSet"},
		"ModelCrudViewSet": {"ModelViewSet"},
		"LocalesViewSet":  {"ReadOnlyListViewSet"},
		"ReadOnlyListViewSet": {"ReadOnlyModelViewSet"},
	}

	// Transitive: ProjectViewSet -> ModelCrudViewSet -> ModelViewSet => 6 CRUD methods.
	got := resolveMethodsTransitive([]string{"ModelCrudViewSet"}, idx)
	if len(got) != 6 {
		t.Errorf("ProjectViewSet transitive: got %d methods, want 6", len(got))
	}

	// ReadOnly transitive => 2 methods.
	ro := resolveMethodsTransitive([]string{"ReadOnlyListViewSet"}, idx)
	if len(ro) != 2 {
		t.Errorf("ReadOnly transitive: got %d methods, want 2", len(ro))
	}

	// Direct DRF base still resolves (no regression).
	direct := resolveMethodsTransitive([]string{"ModelViewSet"}, idx)
	if len(direct) != 6 {
		t.Errorf("direct ModelViewSet: got %d methods, want 6", len(direct))
	}

	// nil idx degrades to direct-parent check.
	none := resolveMethodsTransitive([]string{"ModelCrudViewSet"}, nil)
	if len(none) != 0 {
		t.Errorf("nil idx custom base: got %d methods, want 0", len(none))
	}
}
