package request

type CreateBrandsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}
