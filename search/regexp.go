package search

import (
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

// Product represents a product mapped in the configuration file.  Every product must have a name and
// an array of valid version numbers to use to determining whether a version descriptor match is
// legitimate or obsolete.
type Product struct {
	Name     string   `mapstructure:"name"`
	Versions []string `mapstructure:"versions"`
}

func generateRegexp(key string) []*regexp.Regexp {
	prods := make(map[string]Product)
	viper.UnmarshalKey("products", &prods)
	name := prods[key].Name
	return []*regexp.Regexp{
		regexp.MustCompile("^\\s*\\.\\. version\\S+:: ([0-9]+\\.[0-9]+)"),
		regexp.MustCompile(fmt.Sprintf("^.*? %s ([0-9]+\\.[0-9]+)", name)),
	}
}
func generateRegexpUnmatch(key string, vers []string) []*regexp.Regexp {
	prods := make(map[string]Product)
	viper.UnmarshalKey("products", &prods)
	name := prods[key].Name
	res := []*regexp.Regexp{}

	for _, ver := range vers {
		res = append(res, regexp.MustCompile(fmt.Sprintf("^\\s*\\.\\. version\\S+:: %s", ver)))
		res = append(res, regexp.MustCompile(fmt.Sprintf("^.*? %s %s", name, ver)))
	}

	return res
}

func findKey(keys []string) string {
	if len(keys) > 0 {
		return keys[0]
	}
	prods := make(map[string]Product)
	viper.UnmarshalKey("products", &prods)
	key := ""
	for k := range prods {
		key = k
	}

	return key
}
