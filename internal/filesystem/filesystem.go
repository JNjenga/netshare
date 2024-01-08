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
