package file

type File struct {
    Path string `json:"path"`
}

func initFiles(paths []string) []*File {
    files := []*File{}

    var f *File
    for _, p := range paths {
        f = &File{
            Path: p,
        }
        files = append(files, f)
    }

    return files
}

