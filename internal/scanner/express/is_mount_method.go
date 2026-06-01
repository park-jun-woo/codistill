//ff:func feature=scan type=extract control=selection topic=express
//ff:what 메서드명이 라우터 마운트성 메서드(use/lazyUse)인지 판정한다 (보수적 화이트리스트)
package express

func isMountMethod(name string) bool {
	switch name {
	case "use", "lazyUse":
		return true
	}
	return false
}
