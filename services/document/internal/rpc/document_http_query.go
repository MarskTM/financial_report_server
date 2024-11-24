package rpc

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/golang/glog"
)

func (d *DocumentService) UploadFile(stream pb.Document_UploadFileServer) error {

	glog.V(1).Info("Document_service::UploadFile request")
	var fileName string
	var file *os.File

	// // Mở file để lưu dữ liệu từ client
	// f, err := os.Create("../../cdn/public/upload_file")
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// Gửi phản hồi về khi upload hoàn tất
			break
		}
		if err != nil {
			return err
		}

		// Tạo file khi nhận chunk đầu tiên
		if file == nil {
			fileName = chunk.FileName
			if fileName == "" {
				return fmt.Errorf("UploadFile() - err: filename is empty")
			}

			// Tạo đường dẫn lưu file
			outputPath := filepath.Join("../../cdn/public/", fileName)
			err := os.MkdirAll("uploads", os.ModePerm) // Đảm bảo thư mục tồn tại
			if err != nil {
				return fmt.Errorf("UploadFile() - err: failed to create directory: %v", err)
			}

			// Tạo file với tên được truyền lên
			file, err = os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("UploadFile() - err: failed to create file: %v", err)
			}
		}

		// Ghi dữ liệu vào file
		if _, err := file.Write(chunk.Content); err != nil {
			return err
		}
	}

	// Lưu thông tin vào db
	

	// Trả về response
	stream.SendAndClose(&pb.UploadStatus{
		DocId:    0,
		FilePath: "",
		Message:  "Upload successful",
		Success:  true,
	})
	return nil
}
