// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	client "github.com/heptio/sonobuoy/pkg/client"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyPtrToClientSonobuoyClient() *client.SonobuoyClient {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*client.SonobuoyClient))(nil)).Elem()))
	var nullValue *client.SonobuoyClient
	return nullValue
}

func EqPtrToClientSonobuoyClient(value *client.SonobuoyClient) *client.SonobuoyClient {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *client.SonobuoyClient
	return nullValue
}