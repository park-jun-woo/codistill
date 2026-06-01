//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestCollectFromAugmentedAssignments_I18n — urlpatterns += i18n_patterns(path(...)) 후미 라우트 수집 검증
package django

import "testing"

func TestCollectFromAugmentedAssignments_I18n(t *testing.T) {
	fi := newTestFileInfo(t,
		"urlpatterns = [path('favicon.ico', v)]\n"+
			"urlpatterns += i18n_patterns(\n"+
			"    path(settings.SECRET, include('crm.urls')),\n"+
			"    path('contact-form/<uuid:uuid>/', contact_form, name='contact_form'),\n"+
			")\n")

	entries := collectFromAugmentedAssignments(fi)
	var found bool
	for _, e := range entries {
		if e.pattern == "contact-form/<uuid:uuid>/" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected contact-form route collected from augmented i18n_patterns, got %+v", entries)
	}
}
