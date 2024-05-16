package check

import "net/http"

type CheckService struct {
	cl *http.Client
}

func NewCheckService() *CheckService {
	return &CheckService{
		cl: &http.Client{},
	}
}
