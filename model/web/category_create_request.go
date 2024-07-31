package web

type CategoryCreateRequest struct {
	// json=name untuk binding front ngirim name bukan Name
	Name string `validate:"required,min=1,max=100", json="name"`
}
