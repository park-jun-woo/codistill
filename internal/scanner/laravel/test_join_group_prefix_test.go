//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJoinGroupPrefix 테스트
package laravel

import "testing"

func TestJoinGroupPrefix(t *testing.T) {
	if joinGroupPrefix("", "v1") != "v1" {
		t.Fatal("empty outer")
	}
	if joinGroupPrefix("api", "") != "api" {
		t.Fatal("empty inner")
	}
	if joinGroupPrefix("api", "v1") != "api/v1" {
		t.Fatal("both")
	}
	if joinGroupPrefix("api", "/v1") != "api/v1" {
		t.Fatal("leading slash inner")
	}
	if joinGroupPrefix("api", "/currencies/") != "api/currencies" {
		t.Fatal("both-end slash inner")
	}
	if joinGroupPrefix("", "/v1") != "v1" {
		t.Fatal("empty outer leading slash")
	}
}
