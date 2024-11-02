package rpc

import (
	"io"
	"os"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

type DocumentServerImpl struct {
	pb.UnimplementedDocumentServer
}

func (d *DocumentServerImpl) UploadFile(stream pb.Document_UploadFileServer) error {
	// Mở file để lưu dữ liệu từ client
	f, err := os.Create("uploaded_file")
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// Gửi phản hồi về khi upload hoàn tất
			return stream.SendAndClose(&pb.UploadStatus{Message: "Upload successful"})
		}
		if err != nil {
			return err
		}

		// Ghi dữ liệu vào file
		if _, err := f.Write(chunk.Content); err != nil {
			return err
		}
	}
}
