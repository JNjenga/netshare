package filesystem

import (
    "io/fs"
    "os"
    "log"
)

// TODO: Build a dir cache for every call to this
func ListDir(root string, path string) ([] fs.DirEntry, error) {
    fsys := os.DirFS(root)
    return fs.ReadDir(fsys, path)
}

func IsDir(path string) (bool, error) {
    stat, err := os.Stat(path)

    if err == nil { return stat.IsDir(), nil }

    return false, err
}

func Exists(path string) (bool, error) {
    _, err := os.Stat(path)

    if err == nil { return true, nil }

    if os.IsNotExist(err) { return false, nil }

    return false, err
}

func ReadFile(path string) ([] byte, error) {
    log.Printf("Reading file path: '%s'\n", path)
    return os.ReadFile(path)
}

func WriteFile(file_name string, data []byte){
    f, err := os.Create(file_name)
    checkErr(err)

    defer f.Close()

    f.Write(data)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
