package file
import (
    "io/ioutil"
    "os"
    "path"
    "github.com/kennethdyer/versc/logger"
)
func getPath(arg string) []string{
    return getPaths([]string{arg})
}
func getPaths(args []string) []string {
    fileList := []string{}

    for _, arg := range args {
        fileInfo, infoErr := os.Stat(arg)
        if infoErr != nil {
            logger.Warn("Error stating file path: ", arg, ":\n", infoErr)
            continue
        }
        if fileInfo.IsDir() {
            files, ioErr := ioutil.ReadDir(arg)
            if ioErr != nil {
                logger.Warn("Error reading directory: ", arg, ":\n", ioErr)
            }
            fpaths := []string{}
            for _, f := range files {
                if !excludePattern(f.Name()) {
                    fpaths = append(fpaths, path.Join(arg, f.Name()))
                }
            }
            fileList = append(fileList, getPaths(fpaths)...)
        } else {
            if includeSuffixes(arg){
                fileList = append(fileList, arg)
            }
        }

    }

    return fileList
}
