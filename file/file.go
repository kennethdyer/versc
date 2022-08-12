package file

import "github.com/kennethdyer/versc/logger"

func GetFiles(p string) []string {
    paths := getPath(p)
    logger.Trace("Found ", len(paths), " paths for File list")

    return paths
}

