package run

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Name() string
	RootPath() string
	Init(context.Context, *gin.RouterGroup) error
	Finalize(context.Context)
}

type Services []Service

func (ss Services) init(ctx context.Context, handler http.Handler) error {
	engine, ok := handler(*gin.Engnine)
	if !ok {
		return errors.Errof("unsupported http handler[%+v]", handler)
	}
	for _, svc := range ss {
		group := engine.Group(svc.RootPath())
		if err := svc.Init(ctx, group); err != nil {
			return errors.Wrapf(err, "init service[%s] with root path...", svc.Name(), svc.RootPath())
		}
	}
	return nil
}

func (ss Services) finalize() {
	var wg sync.WaitGroup
	for _, s := range ss {
		wg.Add(1)
		go func(svc Service) {
			defer wg.Done()
			svc.Finalize(ctx)
		}(s)
	}

	wg.Wait()
	// log here
}
