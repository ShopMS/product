package api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shopms/product/config"
	"github.com/shopms/product/ent"
	"github.com/shopms/product/internal/repository"
	"github.com/shopms/product/internal/service"
	pb "github.com/shopms/product/rpc/product"

	"github.com/shopms/common/grpc"
	log "github.com/shopms/common/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx = context.Background()
)

func Run() {
	config.Load()

	log.New(&log.Config{
		Level:  config.Values.LoggerLevel,
		Format: "json",
		Name:   "shopms-product",
	})

	log.Infof("Start running product service")

	dbLink := config.Values.DB.User + ":" +
		config.Values.DB.Password + "@tcp(" +
		config.Values.DB.Host + ":" +
		strconv.FormatUint(uint64(config.Values.DB.Port), 10) + ")/" +
		config.Values.DB.DBName + "?parseTime=True"

	fmt.Println(dbLink)

	clientDB, err := ent.Open("mysql", dbLink)

	if err != nil {
		log.Panicf("db open connection on failure ", err)
	}

	defer func() {
		if err := clientDB.Close(); err != nil {
			panic(err)
		}
	}()

	productApi := initProductApi(clientDB)

	grpcService := grpc.StandardRunner{
		Address: config.Values.ServerAddress,
		Server:  grpc.New(ctx),
	}

	pb.RegisterProductAPIServer(grpcService.Server, productApi)

	if err := grpcService.Run(ctx); err != nil {
		log.Panicf("failed to serve the schedule service", err)
	}

}

func initProductApi(client *ent.Client) *ProductApi {
	productRepository := repository.NewProductRepository(client)
	productService := service.NewProductService(productRepository)
	productEndpoint := &ProductApi{
		ProductService: productService,
	}
	return productEndpoint
}
