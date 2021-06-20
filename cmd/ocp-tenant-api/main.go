package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tenant-api/internal/repo"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozoncp/ocp-tenant-api/internal/api"
	desc "github.com/ozoncp/ocp-tenant-api/pkg/ocp-tenant-api"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
	httpPort = ":7000"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
	httpEndpoint = flag.String("http-server-endpoint", "0.0.0.0"+httpPort, "HTTP server endpoint")

	swaggerDir = flag.String("swagger-dir", "swagger", "./swagger/")
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		fmt.Printf("Swagger JSON not Found: %s\n", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	fmt.Printf("Serving %s\n", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}

func runHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpTenantApiHandlerFromEndpoint(ctx, gwmux, *grpcEndpoint, opts)
	if err != nil {
		return err
	}

	mux.HandleFunc("/swagger/", serveSwagger)
	mux.Handle("/", gwmux)

	fmt.Printf("Https server listening on %s\n", *httpEndpoint)
	return http.ListenAndServe(*httpEndpoint, mux)
}

func runGrpc() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"localhost", "5432", "ocp_tenant_api_user", "ocp_tenant_api_password", "ocp_tenant_api", "disable")
	db, err := sqlx.Open("pgx", psqlInfo)
	if err != nil {
		log.Printf("db error: %v", err.Error())
		log.Fatalf("failed to create connect to database")
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping to database")
		return err
	}

	repo := repo.New(*db, 5)

	s := grpc.NewServer()
	desc.RegisterOcpTenantApiServer(s, api.NewOcpTenantApi(&repo))

	fmt.Printf("Grps server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	if err := runGrpc(); err != nil {
		log.Fatal(err)
		return
	}

	if err := runHttp(); err != nil {
		log.Fatal(err)
		return
	}
}
