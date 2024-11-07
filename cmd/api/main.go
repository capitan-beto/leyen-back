package main

import (
	"fmt"
	"net/http"
	"os"

	"cmd/api/main.go/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
	}

	PORT, exists := os.LookupEnv("PORT")
	if !exists {
		os.Setenv("PORT", "8080")
		PORT = "8080"
	}

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	fmt.Println(("Running on " + PORT + "!"))
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
