package service

type Service interface {
	CheckGpu(string) (string, error)
}
