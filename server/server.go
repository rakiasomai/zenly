package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"strconv"

	pb "github.com/rakiasomai/Zenly/proto"

	"google.golang.org/grpc"
)

var ErrNotfound = fmt.Errorf("name not found")

type server struct{}

type Ob struct {
	Key string
	Value string
}
// our fake data
var ObList = []*Ob{
	&Ob{
		Key: "Khalil",
		Value: "Ariana",
	},
	&Ob{
		Key: "Rakia",
		Value: "Tunis",
	},
	&Ob{
		Key: "Steeve",
		Value: "Paris",
	},
	&Ob{
		Key: "Julien",
		Value: "San Francisco",
	},
}

func (s server) Add(in *pb.KeyValue, srv pb.KV_AddServer) error {
	log.Printf("fetch response for key : %s", in.Key)
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)

		log.Printf("%s  %s", in.Key, in.Value)

		newKV :=  Ob{Key: in.Key, Value: in.Value}

		ObList = append(ObList, &newKV)
		log.Printf("%v", ObList)

		go func() {
			resp := pb.KeyValue{Value: fmt.Sprintf("Your Key and Value have been saved")}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}()
	}
	wg.Wait()
	return nil
}

func find(name string) string {
	for i, p := range ObList {
		if p.Key == name {
			return  strconv.Itoa(i) 
		}
	}
	return  strconv.Itoa(-1)
}

func getValueByKey(name string) string {
	for _, p := range ObList {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}

func (s server) Get(in *pb.KeyRequest, srv pb.KV_GetServer) error {
	log.Printf("fetch response for key : %s", in.Key)
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)

		userKey := in.Key

		if find(userKey) != "-1" {
			value := getValueByKey(userKey)
			log.Printf("The value of %s is %s", userKey, value)
		} else {
			log.Printf("The %s does not exist", userKey)
		}

		go func() {
			userKey := in.Key

			var resMsg string
			if find(userKey) != "-1" {
				value := getValueByKey(userKey)
				resMsg = "The value of " + userKey + " is " + value
			} else {
				resMsg = "The " + userKey + " does not exist"
			}

			resp := pb.KeyResponse{Value: fmt.Sprintf(resMsg)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}()
	}
	
	wg.Wait()
	return nil
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterKVServer(s, server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}