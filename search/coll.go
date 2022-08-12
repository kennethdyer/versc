package search
import (
    "fmt"
    "time"

    "github.com/spf13/viper"
)

func getParallel(count int) int {
    parallel := viper.GetInt("parallel")
    if count < parallel {
        parallel = count
    }
    if parallel == 0 {
        parallel = 1
    }
    return parallel
}

func collectData(files []string, keys []string) []string{

    // Initialize Variables
    lenFiles := len(files)
    parallel := getParallel(lenFiles)
    key := findKey(keys)
    res := generateRegexp(key)
    vers := []string{}
    viper.UnmarshalKey(fmt.Sprintf("products.%s.versions", key), &vers)

    rets := []string{}
    var ret string

    inputChan := make(chan string, lenFiles)
    medChan := make(chan *Line)
    outChan := make(chan string)

    // Initialize File Worker Pool
    for w := 0; w < parallel; w++ {
        go FileWorker(w, inputChan, medChan)
    }

    // Initialize Line Worker Pool
    for w := 0; w < parallel * 3; w++ {
        go LineWorker(w, res, vers, medChan, outChan)
    }

    // Send Files to File Workers
    for _, f := range files {
        inputChan <- f
    }
    close(inputChan)

    // Collect Return Values
    timeout := time.Duration(int64(viper.GetInt("timeout")))
    for {
        select {
        case ret = <- outChan:
            rets = append(rets, ret)
        case <-time.After(time.Second * timeout):
            return sortStrings(rets)
        }
    }
}

