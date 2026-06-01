//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 경로 세그먼트 목록에 "api" 디렉터리 세그먼트가 있는지 판정한다
package express

// pathHasAPIDir reports whether any of the given path segments equals "api".
func pathHasAPIDir(segs []string) bool {
	for _, s := range segs {
		if s == "api" {
			return true
		}
	}
	return false
}
