package system

import (
	"context"
	"fmt"
	"os"
	"time"

	"meetplan/api/v1/types"

	"github.com/cloudwego/hertz/pkg/app"
)

var (
	MaxUploadSize  int64 = 4 << 20 // 4MB
	FileUploadPath       = "./file/upload/"
)

func init() {
	// Create the file upload directory if it doesn't exist
	if err := os.MkdirAll(FileUploadPath, os.ModePerm); err != nil {
		panic(err)
	}
}

func UploadFile(ctx context.Context, c *app.RequestContext, _ *struct{}) (string, *types.PageInfo, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", nil, err
	}
	if file.Size > MaxUploadSize {
		return "", nil, fmt.Errorf("file size is too big")
	}

	filePath := fmt.Sprintf("%s%d-%s", FileUploadPath, time.Now().Unix(), file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", nil, err
	}
	return filePath, nil, nil
}
