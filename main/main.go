package main

import (
	"github.com/chirino/swag-bug-example/lib-go"
	"net/http"
)

// @BasePath /v2
func main() {
	http.HandleFunc("/testapi/get-foo", GetFoo)
	http.ListenAndServe(":8080", nil)
}

// @Description get Foo
// @ID get-foo
// @Accept json
// @Produce json
// @Success 200 {object} Foo
// @Router /testapi/get-foo [get]
func GetFoo(w http.ResponseWriter, r *http.Request) {
	var _ = Foo{}
}

type Foo struct {
	Field lib.Bar
}
