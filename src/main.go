package main

import (
	"fmt"

	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider"
	"github.com/ramonmpacheco/poc-go-gorm/app/dataprovider/repository"
	"github.com/ramonmpacheco/poc-go-gorm/app/rest"
)

func main() {
	fmt.Println("About to start app")
	rest.Start(
		rest.InitRoutes(
			repository.NewPastelRepository(dataprovider.NewPostgres()),
		),
	)
}
