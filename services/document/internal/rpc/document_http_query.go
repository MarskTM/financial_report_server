package rpc

import (
	"io"
	"os"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/golang/glog"
)

func (d *DocumentService) UploadFile(stream pb.Document_UploadFileServer) error {

	glog.V(1).Info("Document_service::UploadFile request")

	// Mở file để lưu dữ liệu từ client
	f, err := os.Create("../../cdn/public/upload_file")
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// Gửi phản hồi về khi upload hoàn tất
			break
		}
		if err != nil {
			return err
		}

		// Ghi dữ liệu vào file
		if _, err := f.Write(chunk.Content); err != nil {
			return err
		}
	}

	// Trả về response
	stream.SendAndClose(&pb.UploadStatus{
		Filename: "",
		Message:  "Upload successful",
		Success:  true,
	})
	return nil
}
