package main

import (
	"ECDICT-GO/ent"
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func t() {
	file, err := os.Open("./ecdict.mini.csv")
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

	fileName := "./stardict.csv"

	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	ff, err := os.Open(fileName)
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
	/*
		file, err := os.Open("./ecdict.mini.csv")
		if err != nil {
			fmt.Printf("Error: %s \n", err.Error())
			return
		}

		defer file.Close()

			br := bufio.NewReader(file)
			count := 0
			for {
				count = count + 1
				if count == 1 {
					continue
				}
				a, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}
				_, err := CreateEntry(ctx, client, string(a))
				if err != nil {
					fmt.Println(err)
				}
			}
	*/

}

func CreateEntry(ctx context.Context, client *ent.Client, entry_slice []string) (*ent.Ecdict, error) {
	//	entry_slice := strings.Split(entryStr, ",")
	//	if len(entry_slice) != 13 {
	//	return nil, nil
	//}
	//fmt.Println(len(entry_slice))
	//fmt.Println(entry_slice)
	//fmt.Println(entry_slice)
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
