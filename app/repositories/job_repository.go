package repositories

import (
	"errors"
	"github.com/vantutran2k1-commits-collector/producer/app/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetLatestJob() (*models.CollectionJob, error)
	CreateJob(job *models.CollectionJob) error
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{
		db: db,
	}
}

type jobRepository struct {
	db *gorm.DB
}

func (r *jobRepository) GetLatestJob() (*models.CollectionJob, error) {
	var job models.CollectionJob
	if err := r.db.Order("collected_from DESC").First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &job, nil
}

func (r *jobRepository) CreateJob(job *models.CollectionJob) error {
	return r.db.Create(job).Error
}
