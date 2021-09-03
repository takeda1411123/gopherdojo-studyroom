package convert

import (
    "io/ioutil"
    "path/filepath"
    "fmt"
    "os"
)

// user defined type
type Conv struct {
    from string
    to   string
    dir  string
}

// create Conv
func NewConv(from string, to string, dir string)(*Conv, error){
    return &Conv{from, to, dir}, nil
}

// search directory
func (conv *Conv) FileSearch(dir string) []string {
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            fmt.Println(err)
        }
        var paths []string
        for _, file := range files {
            if file.IsDir() {
                paths = append(paths, conv.FileSearch(filepath.Join(dir, file.Name()))...)
                continue
            }
            paths = append(paths, filepath.Join(dir, file.Name()))
        }

        return paths
}

// replace filepath
func (conv *Conv) ReplaceExt(path, from, to string) {
    ext := filepath.Ext(path)
    if len(from) > 0 && ext[1:] == from {
        changed := path[:len(path)-len(ext)+1] + to
        if err := os.Rename(path, changed); err != nil {
            fmt.Println(err)
        }
    }
    return
}
