package download

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Item user defined type
type Item struct {
	Input       string
	Output      string
	DivisionNum int
	FileLen     int
}

// NewItem is initiate NewItem
func NewItem(input string, output string, divisionNum int) (*Item, error) {
	return &Item{input, output, divisionNum, 0}, nil
}

// GetFileLen is getting length of file
func (item *Item) GetFileLen() error {
	if item == nil {
		return errors.New("Not define structure")
	}

	resp, err := http.Get(item.Input)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileLenStr, ext := resp.Header["Content-Length"]
	if len(fileLenStr) == 0 || !ext {
		return errors.New("Not know file length")
	}

	fileLenInt, err := strconv.Atoi(fileLenStr[0])
	if err != nil {
		return err
	}

	item.FileLen = fileLenInt
	return err
}

// GetFileSize is getting size of file
func GetFileSize(fileLen int) (string, error) {
	if fileLen < 1024 {
		return fmt.Sprintf("%d Byte", fileLen), nil
	}
	kb := float32(fileLen) / 1024
	if kb < 1024 {
		return fmt.Sprintf("%f KB", kb), nil
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%f MB", mb), nil
	}
	gb := mb / 1024
	if gb < 1024 {
		return fmt.Sprintf("%f GB", gb), nil
	}

	return fmt.Sprintf("Your file more than PB"), errors.New("You should not download because file size is PB")
}

// Start is getting size of file
func (item *Item) Start(ctx context.Context) error {
	if item == nil {
		return errors.New("Not define structure")
	}

	fmt.Println("Start download. Time Limit 5 minutes")
	now := time.Now()
	file, err := os.Create(item.Output)
	if err != nil {
		return err
	}
	defer file.Close()

	divisionSize := item.FileLen / item.DivisionNum
	fmt.Println(divisionSize)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	ch := make(chan int, 1)
	for i := 0; i < item.DivisionNum; i++ {
		ch <- i
		select {
		case <-ctx.Done():
			fmt.Println("time over...")
			return nil
		case <-ch:
			err := item.Download(i, divisionSize, file)
			if err != nil {
				return err
			}

			fmt.Printf("経過: %vs\n", time.Since(now).Seconds())
		}
	}
	close(ch)
	return nil
}

// Download is getting length of file
func (item *Item) Download(i int, size int, file *os.File) error {
	if item == nil {
		return errors.New("Not define structure")
	}

	var start, end int
	fmt.Printf("%v/%v downloading\n", i+1, item.DivisionNum)
	var mutex sync.Mutex
	mutex.Lock()

	client := new(http.Client)
	req, err := http.NewRequest("GET", item.Input, nil)
	if err != nil {
		return err
	}

	switch {
	case i != 0 && i+1 == item.DivisionNum:
		start = i*size + 1
		end = item.FileLen
	case i == 0:
		start = 0
		end = (i + 1) * size
	default:
		start = i*size + 1
		end = (i + 1) * size
	}

	rangeBytes := "bytes=" + strconv.FormatInt(int64(start), 10) + "-" + strconv.FormatInt(int64(end), 10)
	req.Header.Set("Range", rangeBytes)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}

	op, err := file.Seek(int64(end), 0)
	if err != nil {
		return err
	}
	fmt.Println(op, "/", item.FileLen, "downloaded")
	mutex.Unlock()
	return nil
}
