package main

import (
	"context"
	"download/download"
	"flag"
	"fmt"
	"os"
)

// initialize flag
var (
	input       = flag.String("i", "https://people.sc.fsu.edu/~jburkardt/data/csv/hw_200.csv", "input path")
	output      = flag.String("o", "./mlb_players.csv", "output path")
	divisionNum = flag.Int("d", 3, "division number")
)

func init() {
	flag.Usage = func() {
		fmt.Printf(`Usage: -i input path -o output path -d  division number
        Use: download in splits.
        Default input path : https://people.sc.fsu.edu/~jburkardt/data/csv/hw_200.csv
				output path : ./mlb_players.csv
				division number : 3
        `)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	Item, err := download.NewItem(*input, *output, *divisionNum)

	if err != nil {
		fmt.Println(err)
	}
	err = Item.GetFileLen()
	if err != nil {
		fmt.Println(err)
	}
	Size, err := download.GetFileSize(Item.FileLen)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	} else {
		fmt.Println("ファイルサイズ：", Size)
	}
	ctx := context.Background()

	err = Item.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}

}
