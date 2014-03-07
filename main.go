package main

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/ec2"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	prog := filepath.Base(os.Args[0])
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	if len(os.Args[1:]) < 1 {
		log.Printf("usage: %s ADDRESS1 [ADDRESS2 ...]", prog)
		os.Exit(2)
	}

	addresses := os.Args[1:]

	auth, err := aws.GetAuth("", "", "", time.Time{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("getting instance-id from metadata")
	instanceId, err := aws.GetMetaData("instance-id")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found instance-id: %s\n", instanceId)

	log.Println("connecting to ec2 API endpoint")
	conn := ec2.New(auth, aws.USEast)

	log.Println("getting address information")
	resp, err := conn.DescribeAddresses(addresses, []string{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range resp.Addresses {
		if addr.InstanceId == "" {
			log.Printf("claiming unassociated address: %s\n", addr.PublicIp)
			options := ec2.AssociateAddressOptions{
				AllocationId: addr.AllocationId,
				InstanceId:   string(instanceId),
			}
			resp, err := conn.AssociateAddress(&options)
			if err != nil {
				log.Fatal(err)
			}
			if resp.Return != true {
				log.Fatal(resp.Return)
			}
			log.Printf("successfully associated: %s\n", addr.PublicIp)
			break
		}
	}
}
