package repository

import (
	"zadaie/internal/model"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DB *gorm.DB
}

func (r *SubscriptionRepository) Create(sub *model.Subscription) error {
	return r.DB.Create(sub).Error
}

func (r *SubscriptionRepository) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.DB.Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepository) GetByID(id uint) (model.Subscription, error) {
	var sub model.Subscription
	err := r.DB.First(&sub, id).Error
	return sub, err
}

func (r *SubscriptionRepository) Update(id uint, sub model.Subscription) error {
	return r.DB.Model(&model.Subscription{}).
		Where("id = ?", id).
		Updates(sub).Error
}

func (r *SubscriptionRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Subscription{}, id).Error
}

func (r *SubscriptionRepository) GetTotalCost(
	userID string,
	serviceName string,
	start string,
	end string,
) (int, error) {

	var total int

	query := r.DB.Model(&model.Subscription{}).
		Where("user_id = ?", userID)

	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	if start != "" {
		query = query.Where("start_date >= ?", start)
	}

	if end != "" {
		query = query.Where("start_date <= ?", end)
	}

	err := query.Select("COALESCE(SUM(price),0)").Scan(&total).Error

	return total, err
}
