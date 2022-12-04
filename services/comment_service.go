package services

import (
	"fmt"
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type CommentService interface {
	Create(commentInput web.CommentRequest, userId uint, photoId uint) (web.CreateCommentResponse, error)
	GetComments(userId uint) ([]web.CreateCommentResponse, error)
	UpdateComment(commentInput web.CommentRequest, userId uint, commentId uint) (web.CreateCommentResponse, error)
}

type CommentServiceImpl struct {
	CommentRepository repositories.CommentRepository
}

func NewCommentService(commentRepository repositories.CommentRepository) CommentService {
	return &CommentServiceImpl{CommentRepository: commentRepository}
}

func ConvertBodyCommentResp(comment models.Comment) web.CreateCommentResponse {
	return web.CreateCommentResponse{
		Id:        comment.ID,
		Message:   comment.Message,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		User:      convertBodyPhotoUser(*comment.User),
		Photo:     convertBodyAssociatePhoto(*comment.Photo),
		CreatedAt: *comment.CreatedAt,
	}
}

func (s *CommentServiceImpl) Create(commentInput web.CommentRequest, userId uint, photoId uint) (web.CreateCommentResponse, error) {
	comment := models.Comment{
		Message: commentInput.Message,
		UserID:  userId,
		PhotoID: photoId,
	}
	newComment, err := s.CommentRepository.Create(comment)

	return ConvertBodyCommentResp(newComment), err
}

func (s *CommentServiceImpl) GetComments(userId uint) ([]web.CreateCommentResponse, error) {
	fmt.Println(">>>>>>>>>>>>>>>>>>", userId)
	comments, err := s.CommentRepository.GetCommentsUser(userId)
	commentResp := []web.CreateCommentResponse{}
	for i := 0; i < len(comments); i++ {
		comment := ConvertBodyCommentResp(comments[i])
		commentResp = append(commentResp, comment)
	}

	return commentResp, err
}

func (s *CommentServiceImpl) UpdateComment(commentInput web.CommentRequest, userId uint, commentId uint) (web.CreateCommentResponse, error) {
	comment := models.Comment{
		Message: commentInput.Message,
		UserID:  userId,
	}

	updateComment, err := s.CommentRepository.Update(comment, commentId)
	return ConvertBodyCommentResp(updateComment), err
}
