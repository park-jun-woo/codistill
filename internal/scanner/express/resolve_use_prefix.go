//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what app.use/lazyUse("/prefix", router) 호출에서 prefix→파일 매핑을 수집한다
package express

func resolveUsePrefixes(fi *fileInfo, routers map[string]bool, imports map[string]string, absRoot string, aliases map[string]string) []useMount {
	var mounts []useMount
	calls := findAllByType(fi.Root, "call_expression")
	for _, call := range calls {
		m := extractUseMount(call, fi, routers, imports, absRoot, aliases)
		if m != nil {
			mounts = append(mounts, *m)
		}
	}
	return mounts
}
