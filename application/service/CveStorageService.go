package service

import "github.com/hyperits/subscription-sets/internal/data"

// CveStorageService curd
type CveStorageService struct {
	repo *data.CveRepo
}

// NewCveStorageService init
func NewCveStorageService() *CveStorageService {
	return &CveStorageService{
		repo: data.NewCveRepo(),
	}
}

// TimingInsertion 定时插入
func (c *CveStorageService) TimingInsertion() error {
	return nil
}
