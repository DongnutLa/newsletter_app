package services

import (
	"bytes"
	"context"
	"errors"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rs/zerolog"
)

type FilesService struct {
	logger *zerolog.Logger
	cloud  *cloudinary.Cloudinary
}

var _ ports.FilesService = (*FilesService)(nil)

func NewFilesService(ctx context.Context, logger *zerolog.Logger) *FilesService {
	cloudName := utils.GetConfig("cloudinary_cloudname")
	apiKey := utils.GetConfig("cloudinary_api_key")
	apiSecret := utils.GetConfig("cloudinary_api_secret")
	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	return &FilesService{
		logger: logger,
		cloud:  cld,
	}
}

func (s *FilesService) SaveFile(ctx context.Context, buf *bytes.Buffer, fileName, folder string) (string, *domain.ApiError) {
	p := uploader.UploadParams{PublicID: folder + "/" + fileName}

	result, err := s.cloud.Upload.Upload(ctx, buf, p)

	if err != nil || len(result.Error.Message) > 0 {
		s.logger.Err(err).Errs("err", []error{err, errors.New(result.Error.Message)}).Msg("ERROR | Upload Files")
		return "", domain.ErrUploadFile
	}

	return result.SecureURL, nil
}
