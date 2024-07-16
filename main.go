package main

import (
	"net/http"

	"github.com/namelessaffirmation/na-web/config"
	"github.com/namelessaffirmation/na-web/database"
	"github.com/namelessaffirmation/na-web/routers"
)

func main() {

	config := config.LoadEnvVars()

	database.InitializeDBConnection(config.DBConnectionString)

	r := routers.LoadRouters()

	http.ListenAndServe(":3000", r)
}
