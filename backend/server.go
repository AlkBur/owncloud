package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	var port int
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	frontend, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	frontend = filepath.Join(frontend, "../frontend/dist")
	fmt.Println(frontend)

	app.Static("/", frontend)

	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()


	app.Get("/api/v1/law", getRandomLaw)

	log.Fatal(app.Listen(fmt.Sprintf(":%v",port)))
  }
  
  type Law struct {
	Name       string `json:"name,omitempty"`
	Definition string `json:"definition,omitempty"`
  }
  
  var HackerLaws = []Law{
	{
	  Name:       "Amdahl's Law",
	  Definition: "Amdahl's Law is a formula which shows the potential speedup of a computational task which can be achieved by increasing the resources of a system.",
	},
	{
	  Name:       "Conway's Law",
	  Definition: "This law suggests that the technical boundaries of a system will reflect the structure of the organisation.",
	},
	{
	  Name:       "Gall's Law",
	  Definition: "A complex system that works is invariably found to have evolved from a simple system that worked.",
	},
  }
  
  func getRandomLaw(c *fiber.Ctx) error {
	randomLaw := HackerLaws[rand.Intn(len(HackerLaws))]
	// j, err := json.Marshal(randomLaw)
	//  if err != nil {
	// 	return err
	// }
  
	// w.Header().Set("Content-Type", "application/json")
	// io.Copy(w, bytes.NewReader(j))
	return c.JSON(randomLaw)
  }