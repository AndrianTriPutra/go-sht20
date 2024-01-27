package main

import (
	"atp-sht20/pkg/sht20"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("======== SHT20 ========", "")

	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Println("usage:")
		log.Println("\t go run . read/write port baudrate idPrev idNext")
	}

	metode := flag.Args()[0]
	port := flag.Args()[1]
	baud := flag.Args()[2]
	pID := flag.Args()[3]
	nID := flag.Args()[4]

	ibaud, err := strconv.Atoi(baud)
	if err != nil {
		log.Fatalf("ibaud:%s", err)
	}
	baudrate := uint(ibaud)

	ipID, err := strconv.Atoi(pID)
	if err != nil {
		log.Fatalf("ipID:%s", err)
	}
	idPrev := uint8(ipID)

	inID, err := strconv.Atoi(nID)
	if err != nil {
		log.Fatalf("inID:%s", err)
	}
	idNext := uint8(inID)

	setting := sht20.Setting{
		Port:     port,
		Baudrate: baudrate,
		Timeout:  5 * time.Second,
	}

	log.Printf("port      :%s", port)
	log.Printf("baudrate  :%v", baudrate)
	fmt.Println()

	ctx := context.Background()
	repo := sht20.NewRepository(setting)

	if metode == "read" {
		log.Printf("READ ID   :%v", idPrev)
		data, err := repo.Read(ctx, idPrev)
		if err != nil {
			log.Fatalf("Read:%s", err)
		}

		jw, err := json.MarshalIndent(data, " ", " ")
		if err == nil {
			msg := string(jw)
			log.Printf("sht20 :\n%s" + msg)
			fmt.Printf("\n\n")
		}
	} else if metode == "write" {
		log.Printf("CHANGE ID [%v] to [%v] ", idPrev, idNext)
		err := repo.Write(ctx, idPrev, idNext)
		if err != nil {
			log.Fatalf("Write:%s", err)
		}
		log.Println("Write succes change id->sht20")
		log.Println("Write unplog power and wait 10s")
		log.Println("Write plug power sht and try read")
	}

}
