package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" // Easy and flexible package for web development
	"github.com/javier-cortina/goapi/internal/handlers"
	log "github.com/sirupsen/logrus" // Loggin for debuggging (alias as log)
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API")

	fmt.Println(`
       _    __      _______ _  _____            _____ _____ 
      | |  /\ \    / /_   _( )/ ____|     /\   |  __ \_   _|
      | | /  \ \  / /  | | |/| (___      /  \  | |__) || |  
  _   | |/ /\ \ \/ /   | |    \___ \    / /\ \ |  ___/ | |  
 | |__| / ____ \  /   _| |_   ____) |  / ____ \| |    _| |_ 
  \____/_/    \_\/   |_____| |_____/  /_/    \_\_|   |_____|`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
