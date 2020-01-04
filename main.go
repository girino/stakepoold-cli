/*
 * Copyright 2019 Girino Vey.
 * Licensed under the Girino's Anarchist License:
 * https://girino.org/license
 */

// Package main implements a client for stakepoold service.
package main

import (
	"context"
	"log"
	"os"
	"time"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"encoding/json"

	pb "./stakepoolrpc"
)

const (
	 TIMEOUT = 10 * time.Second
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] command\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "Valid commands are:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  blockheight\n    Returns a single integer\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  walletinfo\n    Returns a json\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  stakeinfo\n    Returns a json\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  version\n    Returns a json\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  ping\n    Returns 'Alive' or 'Dead'\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Flags are:\n")
    flag.PrintDefaults()
}

func main() {

	flag.Usage = usage

	var cert = flag.String("c", "rpc.cert", "server cert file")
	var host = flag.String("h", "localhost", "server host")
	var port = flag.Int("p", 9113, "server port")
	flag.Parse()

	if (flag.NArg() != 1) {
		usage()
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "blockheight":
	case "walletinfo":
	case "stakeinfo":
	case "ping":
	case "version":
		break;
	default :
		usage()
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)
	// Set up a connection to the server.
	creds, _ := credentials.NewClientTLSFromFile(*cert, "")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStakepooldServiceClient(conn)

	switch flag.Arg(0) {
	case "blockheight":
		r, err := c.GetStakeInfo(ctx, &pb.GetStakeInfoRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Printf("%d\n", r.BlockHeight)
		break
	case "walletinfo":
		r, err := c.WalletInfo(ctx, &pb.WalletInfoRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		j,_ := json.MarshalIndent(r, "", "   ")
		fmt.Printf("%s\n", j)
		break
	case "stakeinfo":
		r, err := c.GetStakeInfo(ctx, &pb.GetStakeInfoRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		j,_ := json.MarshalIndent(r, "", "   ")
		fmt.Printf("%s\n", j)
		break
	case "ping":
		r, err := c.WalletInfo(ctx, &pb.WalletInfoRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if (r.DaemonConnected && r.Unlocked) {
			fmt.Printf("Alive\n")
		} else {
			fmt.Printf("Dead\n")
			os.Exit(2)
		}
		break
	case "version":
		v := pb.NewVersionServiceClient(conn)
		r, err := v.Version(ctx, &pb.VersionRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		j,_ := json.MarshalIndent(r, "", "   ")
		fmt.Printf("%s\n", j)
		break
	default :
		usage()
		os.Exit(-1)
	}
}
