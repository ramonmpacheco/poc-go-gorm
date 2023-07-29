package main

import (
	"fmt"

	gormdataprovider "github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/gorm_internal/gorm_dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/rest"
)

func main() {
	fmt.Println("About to start app")
	rest.Start(
		rest.InitRoutes(
			repository.NewPastelRepository(gormdataprovider.NewPostgres()),
		),
	)
}
