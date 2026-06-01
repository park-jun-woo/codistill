//ff:func feature=scan type=test control=sequence topic=flask
//ff:what extractMethodsArg가 tuple methods=("POST",) 형태도 처리하는지 검증한다
package flask

import (
	"reflect"
	"testing"
)

func TestExtractMethodsArg_Tuple(t *testing.T) {
	args, src := argListOf(t, `expose('/', methods=("POST",))`+"\n")
	if got := extractMethodsArg(args, src); !reflect.DeepEqual(got, []string{"POST"}) {
		t.Fatalf("tuple methods: got %v, want [POST]", got)
	}

	args, src = argListOf(t, `expose('/', methods=("GET", "POST"))`+"\n")
	if got := extractMethodsArg(args, src); !reflect.DeepEqual(got, []string{"GET", "POST"}) {
		t.Fatalf("tuple multi methods: got %v", got)
	}
}
