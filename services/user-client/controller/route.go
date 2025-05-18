package controller

import (
	"fmt"
	"log"

	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/services/user-client/config"
	"github.com/SHshzik/homework_crud/services/user-client/usecase"
)

type Route struct {
	l        logger.Interface
	userCase usecase.User
	cfg      *config.Config
}

func NewRoute(l logger.Interface, userCase usecase.User, cfg *config.Config) *Route {
	return &Route{l: l, userCase: userCase, cfg: cfg}
}

func (r *Route) Run() {
	switch r.cfg.RequestType {
	case "index":
		users, err := r.userCase.Index()
		if err != nil {
			log.Fatalf("fail to index: %v", err)
		}

		for _, user := range users {
			fmt.Printf("%#v\n", user.Name)
		}
	case "create":
		user, err := r.userCase.Create(r.cfg.Name, r.cfg.Email, r.cfg.Phone)
		if err != nil {
			log.Fatalf("fail to create: %v", err)
		}
		fmt.Printf("%#v\n", user.Name)
	case "show":
		user, err := r.userCase.Read(r.cfg.ID)
		if err != nil {
			log.Fatalf("fail to show: %v", err)
		}
		fmt.Printf("%#v\n", user.Name)
	case "update":
		user, err := r.userCase.Update(r.cfg.ID, r.cfg.Name, r.cfg.Email, r.cfg.Phone)
		if err != nil {
			log.Fatalf("fail to update: %v", err)
		}
		fmt.Printf("%#v\n", user.Name)
	}
}
