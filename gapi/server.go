package gapi

import (
	"fmt"

	db "github.com/LinggaWahyu/Golang-Simple-Bank/db/sqlc"
	"github.com/LinggaWahyu/Golang-Simple-Bank/pb"
	"github.com/LinggaWahyu/Golang-Simple-Bank/token"
	"github.com/LinggaWahyu/Golang-Simple-Bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
