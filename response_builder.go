package main

type ResponseBuilder struct {
	Response Response
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response: Response{},
	}
}

func (rb *ResponseBuilder) WithBalance(balance float64) *ResponseBuilder {
	rb.Response.Account.Balance = balance
	return rb
}

func (rb *ResponseBuilder) FromCity(city string) *ResponseBuilder {
	rb.Response.Address.City = city
	return rb
}

func (rb *ResponseBuilder) WithEmail(email string) *ResponseBuilder {
	rb.Response.User.Email = email
	return rb
}

func (rb *ResponseBuilder) Build() Response {
	return rb.Response
}
