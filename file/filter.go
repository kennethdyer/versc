package file
import (
    "strings"

    "path/filepath"
    "github.com/spf13/viper"

)

func excludePattern(p string) bool {
    for _, pattern := range strings.Split(viper.GetString("exclude_pattern"), ", ") {
        if p == pattern {
            return true
        }
    }
    return false
}

func includeSuffixes(p string) bool {
    ext := filepath.Ext(p)
    for _, pattern := range strings.Split(viper.GetString("include_extensions"), ",") {
        if ext == pattern {
            return true
        }
    }
    return false
}
