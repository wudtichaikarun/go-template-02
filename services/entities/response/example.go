package response

type ExampleRes struct {
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleListRes struct {
	ExampleList []ExampleRes `json:"examples"`
	// pagination metadata
}
