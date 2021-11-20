package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type PathsUrls []struct {
	path string `yaml:"path"`
	url  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest, pr := pathsToUrls[r.URL.Path]
		if pr {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// using struct to unmarshal yaml is not working
func YAMLHandler(bbbbb []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pu []map[string]string
	err := yaml.Unmarshal(bbbbb, &pu)
	fmt.Println(pu)
	if err != nil {
		return nil, err
	}
	pathToURLs := make(map[string]string)
	for _, v := range pu {
		pathToURLs[v["path"]] = v["url"]
	}

	return MapHandler(pathToURLs, fallback), nil
}
