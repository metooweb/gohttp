package gohttp

import (
	"testing"
	"net/http"
	"reflect"
	"context"
	"fmt"
)

type FuncHandler struct{}

func (h FuncHandler) ParseArgs(options *OptionParseArgs) ([]reflect.Value, error) {

	return DefaultArgsParser(options)
}

func (h FuncHandler) ParseArgsErrorHandler(options *OptionParseArgsErrorHandler) {

	fmt.Println("3",options)

}

func (h FuncHandler) CalledHandler(options *OptionCalledHandler) {
	fmt.Println("2",options)
}

func (h FuncHandler) RecoverHandler(options *OptionRecoverHandler) {
	fmt.Println("1",options)
}

type OptionTest struct {
	A string `form:"a"`
}

func TestRun(t *testing.T) {

	r := NewRouter()

	r.HandleFunc("/", func(ctx context.Context, options *OptionTest) {

		fmt.Println(options)

	}, &FuncHandler{})

	http.ListenAndServe(":8082", r)

}
