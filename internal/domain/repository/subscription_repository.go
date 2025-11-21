package repository

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"gorm.io/gorm"
)

// const (
// 	queryGetSub                 = "SELECT id, service_id, msisdn, keyword, adnet, latest_subject, latest_status, amount, renewal_at, purge_at, unsub_at, charge_at, retry_at, success, ip_address, is_retry, is_purge, is_active, created_at, updated_at FROM subscriptions WHERE service_id = ? AND msisdn = ? AND deleted_at IS NULL LIMIT 1"
// 	queryCountRetrySub          = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ? AND is_retry = true AND is_active = true"
// 	queryCountSub               = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ?"
// 	queryCountNotActiveSub      = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ? AND is_active = false"
// 	queryCountActiveSubByMsisdn = "SELECT COUNT(*) as count FROM subscriptions WHERE msisdn = ? AND is_active = true"
// 	queryUpdateKeyword          = "UPDATE subscriptions SET keyword = ?, adnet = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
// 	queryUpdateLatest           = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
// 	queryUpdateSuccess          = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, amount = amount + ?, renewal_at = ?, charge_at = ?, success = success + ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
// 	queryUpdateFailed           = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, renewal_at = ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"
// 	querySubUpdateDisable       = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, unsub_at = ?, is_retry = false, is_active = false, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
// 	queryInsertSub              = "INSERT INTO subscriptions(service_id, msisdn, keyword, adnet, latest_subject, latest_status, ip_address, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
// 	queryRenewalAll             = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
// 	queryRenewalOdd             = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 1 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
// 	queryRenewalEven            = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 0 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
// 	queryRetry                  = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_retry = true AND is_active = true AND deleted_at IS null ORDER BY success DESC"
// 	queryPurge                  = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(purge_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
// )

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

type ISubscriptionRepository interface {
	CountActive(int, string) (int64, error)
	Count(int, string) (int64, error)
	CountRenewal(int, string) (int64, error)
	CountRetry(int, string) (int64, error)
	Get(int, string) (*entity.Subscription, error)
	Save(*entity.Subscription) error
	Update(*entity.Subscription) error
}

func (r *SubscriptionRepository) CountActive(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).
		Where("service_id = ? AND msisdn = ? AND is_active = true", serviceId, msisdn).
		Count(&count).Error
	return count, err
}

func (r *SubscriptionRepository) Count(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).
		Where("service_id = ? AND msisdn = ?", serviceId, msisdn).
		Count(&count).Error
	return count, err
}

func (r *SubscriptionRepository) CountRenewal(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).
		Where("service_id = ? AND msisdn = ? AND is_active = true", serviceId, msisdn).
		Count(&count).Error
	return count, err
}

func (r *SubscriptionRepository) CountRetry(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).
		Where("service_id = ? AND msisdn = ? AND is_retry = true AND is_active = true", serviceId, msisdn).
		Count(&count).Error
	return count, err
}

func (r *SubscriptionRepository) Get(serviceId int, msisdn string) (*entity.Subscription, error) {
	var e entity.Subscription
	err := r.db.
		Where("service_id = ? AND msisdn = ?", serviceId, msisdn).
		First(&e).Error
	return &e, err
}

func (r *SubscriptionRepository) Save(e *entity.Subscription) error {
	return r.db.Create(e).Error
}

func (r *SubscriptionRepository) Update(e *entity.Subscription) error {
	return r.db.Save(e).Error
}
