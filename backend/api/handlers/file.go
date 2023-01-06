package handler

import (
	"strings"
	model "ybigta/bard-backend/models"
	repository "ybigta/bard-backend/repository/db"
	fileservice "ybigta/bard-backend/services/file"
	util "ybigta/bard-backend/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {

	var createdFiles []*model.File

	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]

	for _, file := range files {

		createdFile, err := fileservice.UploadToS3(file)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		res, err := (*util.BLIP).Predict(ctx,
			util.BLIPInput{
				Image: createdFile.Url,
				Task:  "image_captioning",
			})
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		createdFile.Caption = strings.SplitN(res.(string), " ", 2)[1]

		createdFiles = append(createdFiles, createdFile)

	}

	repository.DB.Create(&createdFiles)

	ctx.JSON(201, gin.H{
		"files": createdFiles,
	})

}
