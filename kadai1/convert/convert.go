package convert

import (
    "io/ioutil"
    "path/filepath"
    "os"
    "image"
    "image/jpeg"
    "image/png"
)

// user defined type
type Conv struct {
    from string
    to   string
    dir  string
    paths []string
}

// create Conv
func NewConv(from string, to string, dir string)(*Conv, error){
    return &Conv{from, to, dir, nil}, nil
}

// search directory
func (conv *Conv)FileSearch(dir string, from string)([]string, error){
        var paths []string
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            return nil, err
        }
        for _, file := range files {
            if file.IsDir() {
                conv.FileSearch(filepath.Join(dir, file.Name()), from)
                continue
            }
            ext := filepath.Ext(file.Name())
            if ext[1:] == from {
                fullpath := filepath.Join(dir, file.Name())
                paths = append(paths, fullpath)
            }
        }
        return paths, err
}

// replace filepath
func (conv *Conv) ReplaceExt(path, from, to string) (error) {
    input_file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer input_file.Close()

    img, _, err := image.Decode(input_file)
    if err != nil {
        return err
    }

    out_file, err := os.Create(path[:len(path)-len(filepath.Ext(path))+1] + to)
    if err != nil {
        return err
    }
    defer out_file.Close()

    if to == "jpg" || to == "jpeg" {
            err := jpeg.Encode(out_file, img, &jpeg.Options{})
            if err != nil {
                return err
            }
    } else if to == "png" {
            err := png.Encode(out_file, img)
            if err != nil {
                return err
            }
    }

    err = os.Remove(path)
    if err != nil {
        return err
    }

    return nil
}
