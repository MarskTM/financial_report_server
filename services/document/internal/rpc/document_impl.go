package rpc

import (
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

type DocumentService struct {
	DocsModel model.DocumentModel
	pb.UnimplementedDocumentServer
}

func NewDocsService(model model.DocumentModel) *DocumentService {
	return &DocumentService{
		DocsModel: model,
	}
}
