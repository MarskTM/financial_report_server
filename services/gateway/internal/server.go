package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/go-chi/jwtauth"
)

var (
	config         env.GatewayConfig
	configDocument env.DocumentConfig
	managerDao     database.ManagerDBDao
	DecodeAuth     *jwtauth.JWTAuth
	EncodeAuth     *jwtauth.JWTAuth
)

// ----------------------------------------------------------------
type GatewayService struct {
	apiServer *http.Server
	server    *grpc.Server

	// ----------------------------------------------------------------
	clientConnection map[string]*grpc.ClientConn
}

// Constructor creates a new GatewayServer
func NewGatewayService() system.ServicesInterface {
	return &GatewayService{
		clientConnection: make(map[string]*grpc.ClientConn),
	}
}

// -----------------------------------------------------------------------------
// service interface
func (s *GatewayService) Install() error {
	glog.V(1).Infof(">>> gateway::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		glog.V(1).Infof("(-) gateway::Initialize - Error: %+v", err)
		return err
	}

	glog.V(1).Infof("(+) load configuration for gateway successfully!")

	// 2. Install DAO
	managerDao.ConnectDB(config.DBConfig, system.PostgresDB)

	// 3. Install gRPC client
	conn, err := grpc.Dial(configDocument.URL, grpc.WithInsecure())
	if err != nil {
		glog.Error("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 4. Install gRPC client
	// clientConnection["document"] = pb.NewDocumentClient(document.)
	// clientConnection["biz_server"] = pb.New

	// 4. Install gRPC server

	// 5. Install HTTP server
	s.apiServer = &http.Server{
		Addr:         config.Addr,
		Handler:      Router(s.clientConnection),
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}

	return nil
}

func (s *GatewayService) Start() {
	go func() {
		err := s.apiServer.ListenAndServe()
		if err != nil {
			glog.Fatalf("gateway::ListenAndServer - Error: %+v", err)
		}
	}()
}

func (s *GatewayService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
