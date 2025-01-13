package daemon

import (
	"time"

	"github.com/dezh-tech/panda/config"
	"github.com/dezh-tech/panda/deliveries/grpc"
	"github.com/dezh-tech/panda/deliveries/http"
	"github.com/dezh-tech/panda/infrastructures/database"
	grpcClient "github.com/dezh-tech/panda/infrastructures/grpc_client"
	"github.com/dezh-tech/panda/infrastructures/redis"
	"github.com/dezh-tech/panda/pkg/logger"
	"github.com/dezh-tech/panda/repositories"
	service "github.com/dezh-tech/panda/services/domain"
)

type Daemon struct {
	config     config.Config
	httpServer http.Server
	grpcServer *grpc.Server
	database   *database.Database
	redis      *redis.Redis
}

func New(cfg *config.Config) (*Daemon, error) {
	db, err := database.Connect(cfg.Database)
	if err != nil {
		return nil, err
	}

	r, err := redis.New(cfg.RedisConf)
	if err != nil {
		return nil, err
	}

	_, err = grpcClient.New(cfg.GRPCClient.Endpoint)
	if err != nil {
		return nil, err
	}

	domainRepo := repositories.NewDomainRepository(db.Client, cfg.Database.DBName,
		time.Duration(cfg.Database.QueryTimeout)*time.Millisecond)

	hs := http.New(cfg.HTTPServer, service.NewDomainService(domainRepo))
	gs := grpc.New(&cfg.GRPCServer, r, db, time.Now())

	return &Daemon{
		config:     *cfg,
		httpServer: hs,
		database:   db,
		redis:      r,
		grpcServer: gs,
	}, nil
}

func (d *Daemon) Start() chan error {
	errCh := make(chan error, 2)

	logger.Info("starting daemon.")

	go func() {
		if err := d.httpServer.Start(); err != nil {
			errCh <- err
		}
	}()

	go func() {
		if err := d.grpcServer.Start(); err != nil {
			errCh <- err
		}
	}()

	logger.Info("daemon started successfully.")

	return errCh
}

func (d *Daemon) Stop() error {
	logger.Info("stopping the server.")

	if err := d.httpServer.Stop(); err != nil {
		return err
	}

	if err := d.grpcServer.Stop(); err != nil {
		return err
	}

	if err := d.database.Stop(); err != nil {
		return err
	}

	logger.Info("daemon stopped successfully.")

	return nil
}
