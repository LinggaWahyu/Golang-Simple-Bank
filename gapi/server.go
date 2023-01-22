package gapi

import (
	"fmt"

	db "github.com/LinggaWahyu/Golang-Simple-Bank/db/sqlc"
	"github.com/LinggaWahyu/Golang-Simple-Bank/pb"
	"github.com/LinggaWahyu/Golang-Simple-Bank/token"
	"github.com/LinggaWahyu/Golang-Simple-Bank/util"
	"github.com/LinggaWahyu/Golang-Simple-Bank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaksDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaksDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor ,
	}

	return server, nil
}
