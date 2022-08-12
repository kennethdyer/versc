package search

import (
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

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
	for k, _ := range prods {
		key = k
	}

	return key
}
