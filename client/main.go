package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	pb "github.com/vietwow/test/pb"
)

func initConfig() error {
	viper.SetConfigType("yaml")
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	err := viper.ReadConfig(bytes.NewReader(file))
	if err != nil {
		return err
	}

	viper.SetDefault("SERVER", "localhost:50051")

	return nil
}
func main() {
	initConfig()
	endpoint := viper.GetString("SERVER")

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()
	
	c := pb.
}
