package controller

import (
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
		r.Index()
	case "create":
		r.Create()
	case "show":
		r.Show()
	case "update":
		r.Update()
	case "delete":
		r.Delete()
	}
}

func (r *Route) Index() {
	users, err := r.userCase.Index()
	if err != nil {
		r.l.Fatal("fail to index: %v", err)
	}

	for _, user := range users {
		r.l.Info("%#v", user)
	}
}

func (r *Route) Create() {
	user, err := r.userCase.Create(r.cfg.Name, r.cfg.Email)
	if err != nil {
		r.l.Fatal("fail to create: %v", err)
	}

	r.l.Info("%#v", user.Name)
}

func (r *Route) Show() {
	user, err := r.userCase.Read(r.cfg.ID)
	if err != nil {
		r.l.Fatal("fail to show: %v", err)
	}

	r.l.Info("%#v", user.Name)
}

func (r *Route) Update() {
	user, err := r.userCase.Update(r.cfg.ID, r.cfg.Name, r.cfg.Email)
	if err != nil {
		r.l.Fatal("fail to update: %v", err)
	}

	r.l.Info("%#v", user.Name)
}

func (r *Route) Delete() {
	err := r.userCase.Delete(r.cfg.ID)
	if err != nil {
		r.l.Fatal("fail to delete: %v", err)
	}

	r.l.Info("deleted")
}
