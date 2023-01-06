package fileservice

import (
	"mime/multipart"
	"strings"
	model "ybigta/bard-backend/models"
	util "ybigta/bard-backend/utils"

	"github.com/google/uuid"
)

func UploadToS3(file *multipart.FileHeader) (*model.File, error) {

	var createdFile model.File

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	file_key := uuid.New().String()
	file_ext := strings.Split(file.Filename, ".")[1]
	object_name := file_key + "." + file_ext

	err = util.S3.UploadFile(src, object_name)
	if err != nil {
		return nil, err
	}

	createdFile.Name = file.Filename
	createdFile.Url = util.S3.GetFileUrl(object_name)
	createdFile.Size = file.Size

	return &createdFile, nil

}
