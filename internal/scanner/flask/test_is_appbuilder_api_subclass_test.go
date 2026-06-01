//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what isAppbuilderAPISubclass가 BaseApi/ModelRestApi/*RestApi를 판정하는지 검증한다
package flask

import "testing"

func TestIsAppbuilderAPISubclass(t *testing.T) {
	cases := []struct {
		supers          []string
		aliases         importAlias
		wantAPI, wantMd bool
	}{
		{[]string{"BaseApi"}, nil, true, false},
		{[]string{"ModelRestApi"}, nil, true, true},
		{[]string{"BaseSupersetModelRestApi"}, nil, true, true},
		{[]string{"flask_appbuilder.api.ModelRestApi"}, nil, true, true},
		{[]string{"SomethingApi"}, nil, true, false},
		{[]string{"BaseView"}, nil, false, false},
		{[]string{"Mra"}, importAlias{"Mra": "ModelRestApi"}, true, true},
	}
	for _, c := range cases {
		gotAPI, gotMd := isAppbuilderAPISubclass(c.supers, c.aliases)
		if gotAPI != c.wantAPI || gotMd != c.wantMd {
			t.Fatalf("isAppbuilderAPISubclass(%v)=(%v,%v), want (%v,%v)", c.supers, gotAPI, gotMd, c.wantAPI, c.wantMd)
		}
	}
}
