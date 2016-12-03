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

/*
Endpoint Structure:
http://127.0.0.1:8080/plugins
http://127.0.0.1:8080/plugins/
http://127.0.0.1:8080/plugins/type/:type [eg collector, publisher, publisher]
http://127.0.0.1:8080/plugin/:name [for a specific plugin or ListPlugins]
http://127.0.0.1:8080/plugin/




*/

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

// Filter takes in an array of plugins, a condition, and returns
// a filtered array of plugins
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

// ListPlugins works for a specific plugin or all of them
func ListPlugins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := getPluginData("http://staging.webapi.snap-telemetry.io/plugin")
	pluginNames := make([]Plugin, 0)
	err := json.Unmarshal(body, &pluginNames)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		pluginName := strings.ToLower(ps.ByName("name"))
		if pluginName != "" {
			pluginNames = Filter(pluginNames, func(v Plugin) bool {
				return strings.Contains(v.FullName, pluginName)
			})
		}

		output, _ := json.MarshalIndent(pluginNames, "", "    ")
		fmt.Fprint(w, string(output))
	}
}

func getPluginData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/plugin", ListPlugins)
	router.GET("/plugin/:name", ListPlugins)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
