package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Plugin struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Forks       int    `json:"fork_count"`
	Stars       int    `json:"star_count"`
	Watchers    int    `json:"watch_count"`
	Issues      int    `json:"issues_count"`
}

func Filter(vs []Plugin, f func(Plugin) bool) []Plugin {
	vsf := make([]Plugin, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, `Snap Plugin API Server:

/plugin
/plugin/collector
/plugin/processor
/plugin/publisher
/plugin/:name`)
}

func ListPlugin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, e := ioutil.ReadFile("./plugins.json")
	if e != nil {
		fmt.Fprintf(w, "File error: %v\n", e)
	}
	pluginNames := make([]Plugin, 0)
	err := json.Unmarshal(file, &pluginNames)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		plugin_name := strings.ToLower(ps.ByName("full_name"))
		if plugin_name != "" {
			pluginNames = Filter(pluginNames, func(v Plugin) bool {
				return strings.Contains(v.FullName, plugin_name)
			})
		}

		output, _ := json.MarshalIndent(pluginNames, "", "    ")
		fmt.Fprint(w, string(output))
	}

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/plugin", ListPlugin)
	router.GET("/plugin/:full_name", ListPlugin)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
