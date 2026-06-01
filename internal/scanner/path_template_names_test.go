//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestPathTemplateNames 테스트
package scanner

import (
	"reflect"
	"testing"
)

func TestPathTemplateNames(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"/users/{id}", []string{"id"}},
		{"/orgs/{org}/users/{id}", []string{"org", "id"}},
		{"/static/{wildcard}", []string{"wildcard"}},
		{"/health", nil},
		{"/a/{}/b", nil},
		// Phase145: `${PREFIX}` interpolation leftover is not a path param.
		{"/${PREFIX_APIV3_DATA}/{modelId}/records", []string{"modelId"}},
		{"/${PREFIX}", nil},
	}
	for _, c := range cases {
		if got := pathTemplateNames(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("pathTemplateNames(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
