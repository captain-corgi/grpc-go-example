package svc

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	routeGuideService "github.com/captain-corgi/grpc-go-example/internal/svc/route_guide"
	userService "github.com/captain-corgi/grpc-go-example/internal/svc/user"
)

var (
	swaggerDir = flag.String("swagger_dir", "/api", "path to the directory which contains swagger definitions")
)

//StartServices run all services
func StartServices(grpcEndpoint string) {
	// Define server
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sv := grpc.NewServer()

	//Register services
	routeGuideService.RegisterGRPCService(sv)
	userService.RegisterGRPCService(sv)

	// Start server
	log.Printf("GRPC server listening on port %s\n", grpcEndpoint)
	log.Fatalf("failed to start server %v", sv.Serve(lis))
}

//StartGateway run all gRPC gateways
func StartGateway(rpcPort string, grpcEndpoint string, opts ...runtime.ServeMuxOption) {
	defer glog.Flush()

	// Initial context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create runtime server mux
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	userService.RegisterGateway(ctx, mux, grpcEndpoint, dialOpts)

	// Create http server mux to handle independent http function
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	httpMux.HandleFunc("/v1/swagger/", serveSwagger)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("GRPC server started on port %s\n", grpcEndpoint)
	log.Printf("HTTP server listening on port %s\n", rpcPort)
	log.Fatal(http.ListenAndServe(rpcPort, allowCORS(httpMux)))
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// TODO: Not working yet.....
func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Swagger JSON not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/v1/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}
