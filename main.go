package main

import (
	"ECDICT-GO/ent"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func t() {
	file, err := os.Open("./ecdict.csv")
	if err != nil {
		fmt.Printf("Error: %s \n", err.Error())
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}

func main() {
	client, err := ent.Open("sqlite3", "file:./test.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return
	}

	u, err := client.User.Create().SetAge(20).SetName("adek").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return
	}
	log.Println("user ", u)
}
