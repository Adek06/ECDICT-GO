package util

import (
	"ECDICT-GO/ent"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func WriteCSVtoDB(csvFilePath string) {

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

	fs, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	ff, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer ff.Close()
	t := csv.NewReader(ff)
	t_, _ := t.ReadAll()
	all_row := len(t_)

	r := csv.NewReader(fs)
	count := 0
	for {
		count += 1
		if count == 1 {
			continue
		}
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		_, err = CreateEntry(ctx, client, row)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(float64(count) / float64(all_row))

	}

}

func CreateEntry(ctx context.Context, client *ent.Client, entry_slice []string) (*ent.Ecdict, error) {
	w, err := client.Ecdict.Create().
		SetWord(entry_slice[0]).
		SetSw("").
		SetPhonetic(entry_slice[1]).
		SetDefinition(entry_slice[2]).
		SetTranslation(entry_slice[3]).
		SetPos(entry_slice[4]).
		SetCollins(StrToInt(entry_slice[5])).
		SetOxford(StrToInt(entry_slice[6])).
		SetTag(entry_slice[7]).
		SetBnc(StrToInt(entry_slice[8])).
		SetFrq(StrToInt(entry_slice[9])).
		SetExchange(entry_slice[10]).
		SetDetail(entry_slice[11]).
		SetAudio(entry_slice[12]).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed create entry: %v", err)
	}
	return w, nil
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("failed translate %s: %v", s, err)
		return -1
	}
	return i
}
