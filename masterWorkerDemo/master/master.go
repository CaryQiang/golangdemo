package master

import (
	"context"
	gw "golangdemo/masterWorkerDemo/pb"
	"golangdemo/masterWorkerDemo/worker"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Server master server
type Server struct {
	config      *worker.Config
	httpServer  *http.Server
	netCancelFn context.CancelFunc
}

// GetNewServer create new server
func GetNewServer(conf *worker.Config) (*Server, error) {
	ser := Server{
		config: conf,
	}

	return &ser, nil
}

// Start start master server
func (ser *Server) Start() error {

	//http
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ser.netCancelFn = cancel

	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterTargetServiceHandlerFromEndpoint(ctx, mux, ser.config.TargetGrpcpoint, dialOptions); err != nil {
		log.Println("Target Server: register service handler for grpc failed: ", err)
		return err
	}
	/*
		ser.httpServer = &http.Server{Addr: ser.config.TargetEndpoint}
		http.Handle("/", mux)
	*/
	go func() {
		if err := http.ListenAndServe(ser.config.TargetEndpoint, mux); err != nil {
			log.Println("Failed to start http service: ", err)
			panic(err)
		}
	}()
	log.Printf("HTTP Server: server listing at %s...\n", ser.config.TargetEndpoint)

	return nil
}

//Stop stop master server
func (ser *Server) Stop() {
	ser.httpServer.Close()
}
