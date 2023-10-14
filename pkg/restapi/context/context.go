package restapi_context

import "mime/multipart"

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	TransactionID() string
	Audience() string
	GetToken() string
	GetParam(string) (int, error)
	GetQuery(string) string
	FormFile(string) (*multipart.FileHeader, error)
}
