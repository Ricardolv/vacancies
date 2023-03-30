package handler

import "fmt"

// CreateOptions
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequest(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is requeid", name, typ)
}

func (request *CreateOpeningRequest) Validate() error {

	if request.Role == "" && request.Company == "" && request.Location == "" && request.Remote == nil && request.Link == "" && request.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}

	if request.Role == "" {
		return errParamIsRequest("role", "string")
	}

	if request.Company == "" {
		return errParamIsRequest("company", "string")
	}

	if request.Location == "" {
		return errParamIsRequest("location", "string")
	}

	if request.Link == "" {
		return errParamIsRequest("link", "string")
	}

	if request.Remote == nil {
		return errParamIsRequest("remote", "bool")
	}

	if request.Salary <= 0 {
		return errParamIsRequest("salary", "int64")
	}

	return nil
}
