package rest_errors

import "net/http"

//criando um objeto "RestErr" do tipo struct
type RestErr struct {
	Message string `json:"message"` 
	Err string `json:"err"`
	Code int `json:"code"`
	Causes []Causes `json:"causes,omitempty"`
}
//`json:""` deixa o objeto mapeado pra json
//omitempty: a propriedade não irá aparecer caso a mesma estiver vazia

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

//cria uma função chamada "Error", que implementa RestErr e retorna uma string
func (r *RestErr) Error() string {
	return r.Message
}

//criando uma função que recebe os parâmetros message, err, code e causes e retorna um ponteiro de RestErr
func NewRestError(message string, err string, code int, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "Bad Request",
		Code: http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "Not Found",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "Forbidden",
		Code: http.StatusForbidden,
	}
}

func NewUnprocessableEntityError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "Unprocessable Entity",
		Code: http.StatusUnprocessableEntity,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "Internal Server Error",
		Code: http.StatusInternalServerError,
	}
}