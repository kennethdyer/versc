package search
import (
    "fmt"
    "os"
    "strings"
    "text/tabwriter"

    "github.com/kennethdyer/versc/file"
)

func RunSearch(target string, products []string){
    files := file.GetFiles(target)
    rets := collectData(files, products)
    
    w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprint(w, strings.Join(rets, "\n"))
    w.Flush()
}
