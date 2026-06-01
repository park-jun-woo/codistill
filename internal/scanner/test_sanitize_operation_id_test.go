//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestSanitizeOperationID 테스트
package scanner

import "testing"

func TestSanitizeOperationID(t *testing.T) {
	tests := []struct {
		name string
		ep   Endpoint
		id   string
		want string
	}{
		{
			name: "leading tilde stripped",
			ep:   Endpoint{Method: "GET", Path: "/~/api/v1/invoices"},
			id:   "~GetInvoices",
			want: "GetInvoices",
		},
		{
			name: "leading slash stripped",
			ep:   Endpoint{Method: "POST", Path: "//invoices"},
			id:   "/CreateInvoice",
			want: "CreateInvoice",
		},
		{
			name: "normal camelCase unchanged",
			ep:   Endpoint{Method: "GET", Path: "/api/v1/invoices"},
			id:   "getInvoices",
			want: "getInvoices",
		},
		{
			name: "internal underscore and digits preserved",
			ep:   Endpoint{Method: "GET", Path: "/api/v1/foo"},
			id:   "get_foo2Bar",
			want: "get_foo2Bar",
		},
		{
			name: "leading underscore is identifier, preserved",
			ep:   Endpoint{Method: "GET", Path: "/api/v1/foo"},
			id:   "_foo",
			want: "_foo",
		},
		{
			name: "all non-identifier falls back to method",
			ep:   Endpoint{Method: "GET", Path: "/~/~"},
			id:   "~~",
			want: "get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeOperationID(tt.ep, tt.id)
			if got != tt.want {
				t.Errorf("sanitizeOperationID() = %q, want %q", got, tt.want)
			}
		})
	}
}
