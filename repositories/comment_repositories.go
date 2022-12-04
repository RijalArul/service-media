package repositories

import (
	models "service-media/models/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository interface {
	Create(comment models.Comment) (models.Comment, error)
	GetCommentsUser(userId uint) ([]models.Comment, error)
	Update(comment models.Comment, commentId uint) (models.Comment, error)
	Delete(commentId uint) error
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (r *CommentRepositoryImpl) Create(comment models.Comment) (models.Comment, error) {
	err := r.DB.Preload(clause.Associations).Create(&comment).First(&comment).Error
	return comment, err
}

func (r *CommentRepositoryImpl) GetCommentsUser(userId uint) ([]models.Comment, error) {
	comments := []models.Comment{}
	err := r.DB.Preload(clause.Associations).Find(&comments, "user_id = ?", userId).Error

	return comments, err
}

func (r *CommentRepositoryImpl) Update(comment models.Comment, commentId uint) (models.Comment, error) {
	err := r.DB.Preload(clause.Associations).Where("id = ?", commentId).Updates(&comment).First(&comment).Error
	return comment, err
}

func (r *CommentRepositoryImpl) Delete(commentId uint) error {
	var comment models.Comment
	err := r.DB.Model(&comment).Delete(&comment, commentId).Error
	return err
}
