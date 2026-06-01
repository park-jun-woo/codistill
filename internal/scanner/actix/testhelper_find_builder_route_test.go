//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what findBuilderRoute 테스트 헬퍼: method/path로 builderRoute를 찾는다
package actix

func findBuilderRoute(routes []builderRoute, method, path string) *builderRoute {
	for i := range routes {
		if routes[i].method == method && routes[i].path == path {
			return &routes[i]
		}
	}
	return nil
}
