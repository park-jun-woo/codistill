//ff:func feature=scan type=test control=sequence
//ff:what TestFormatPathConflictReport 테스트 (빈 입력은 빈 문자열, 충돌은 건수+상세 집계)
package scanner

import (
	"strings"
	"testing"
)

func TestFormatPathConflictReport(t *testing.T) {
	// 회귀 없음: 충돌 0건이면 빈 리포트.
	if got := formatPathConflictReport(nil); got != "" {
		t.Fatalf("expected empty report for no conflicts, got %q", got)
	}

	conflicts := []pathConflict{
		{Path: "/users", Method: "get", KeptHandler: "k1", KeptFile: "a.rs", KeptLine: 1, DropHandler: "d1", DropFile: "b.rs", DropLine: 2},
		{Path: "/posts", Method: "post", KeptHandler: "k2", KeptFile: "c.rs", KeptLine: 3, DropHandler: "d2", DropFile: "d.rs", DropLine: 4},
	}
	report := formatPathConflictReport(conflicts)
	if !strings.Contains(report, "2 operation(s) lost") {
		t.Fatalf("expected total count in report, got %q", report)
	}
	if !strings.Contains(report, "dropped \"d1\"") || !strings.Contains(report, "kept \"k2\"") {
		t.Fatalf("expected dropped/kept handlers in report, got %q", report)
	}
	// path 정렬 결정성: /posts 가 /users 보다 먼저 등장.
	if strings.Index(report, "/posts") > strings.Index(report, "/users") {
		t.Fatalf("expected paths sorted in report, got %q", report)
	}
}
