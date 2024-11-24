package rpc

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/go-chi/render"
	"github.com/golang/glog"
)

func (c *gatewayController) UploadFile(w http.ResponseWriter, r *http.Request) {
	var response utils.Response

	// Giới hạn kích thước tối đa của form, ví dụ 10 MB
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Form quá lớn", http.StatusBadRequest)
		return
	}

	// Lấy file "profile_picture" từ form
	file, fileHeader, err := r.FormFile("profile")
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("Không thể lấy file từ form: %v", err))
		return
	}
	defer file.Close()


	glog.V(1).Infoln("FileName:", fileHeader.Filename)
	// ------------------------------------------------------------------------------------------------

	// Gọi hàm upload file qua gRPC
	resp, errReps := uploadFileToDocumentService(c.GateModel.DocsClient, file, fileHeader.Filename)
	if errReps != nil {
		utils.InternalServerErrorResponse(w, r, errReps)
		return
	}

	// reply
	response = utils.Response{
		Data:    resp,
		Success: true,
		Message: "Authenticated",
	}
	render.JSON(w, r, response)
}

func (c *gatewayController) Delete(w http.ResponseWriter, r *http.Request) {}

// ------------------------------------------- Docunment Utils Function -----------------------------------------------------
func uploadFileToDocumentService(client pb.DocumentClient, file io.Reader, fileName string) (*pb.UploadStatus, error) {
	// Tạo stream để upload file
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not upload file: %v", err)
	}

	// Đọc file theo từng chunk và gửi qua gRPC stream
	buffer := make([]byte, 1024) // 1KB buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read chunk: %v", err)
		}

		// Gửi chunk qua gRPC
		req := &pb.FileChunk{
			FileName: fileName,
			Content: buffer[:n],
		}
		if err := stream.Send(req); err != nil {
			return nil, fmt.Errorf("could not send chunk: %v", err)
		}
	}

	// Báo hiệu không còn dữ liệu nào để gửi
	if err := stream.CloseSend(); err != nil {
		return nil, fmt.Errorf("error close send chunk file: %v", err)
	}

	// Nhận phản hồi từ server sau khi đã đóng stream
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("could not receive response: %v", err)
	}

	glog.V(1).Infof("File uploaded successfully: %v", resp.Message)
	return resp, nil
}
