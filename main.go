package main

import (
	"net/http"

	"github.com/squireaj/FirstGoTutorial/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
