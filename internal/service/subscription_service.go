package service

import (
	"zadaie/internal/model"
	"zadaie/internal/repository"
)

type SubscriptionService struct {
	Repo *repository.SubscriptionRepository
}

func (s *SubscriptionService) Create(sub *model.Subscription) error {
	return s.Repo.Create(sub)
}

func (s *SubscriptionService) GetAll() ([]model.Subscription, error) {
	return s.Repo.GetAll()
}

func (s *SubscriptionService) GetByID(id uint) (model.Subscription, error) {
	return s.Repo.GetByID(id)
}

func (s *SubscriptionService) Update(id uint, sub model.Subscription) error {
	return s.Repo.Update(id, sub)
}

func (s *SubscriptionService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *SubscriptionService) GetTotalCost(
	userID string,
	serviceName string,
	start string,
	end string,
) (int, error) {
	return s.Repo.GetTotalCost(userID, serviceName, start, end)
}
