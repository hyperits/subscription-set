package main

import (
	"context"
	"flag"
	pb "github.com/hyperits/protocol/subscription_set"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	//res, err := http.Get("https://services.nvd.nist.gov/rest/json/cves/2.0/?resultsPerPage=100&startIndex=0")
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//
	//defer res.Body.Close()
	//
	//content, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//var cveList dto.CveDTO
	//err = json.Unmarshal(content, &cveList)
	//if err != nil {
	//	fmt.Println("marshal error ", err.Error())
	//}
	//for _, v := range cveList.Vulnerabilities {
	//	fmt.Println(v.Info.ID)
	//}
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
