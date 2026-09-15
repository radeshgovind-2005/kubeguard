package main

import(
	"fmt"
	"log"
	"github.com/radeshgovind-2005/kubeguard/config"
)

func main(){
	fmt.Println("Kubeguard starting...")

	// Load Config file
	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
        // If there's an error, log it and exit the program.
        log.Fatalf("Failed to load config: %v", err)
    }

	fmt.Printf("Watching resources: %v in namespaces: %v\n", cfg.Watch.Resources, cfg.Watch.Namespaces)

}