package cachingproxy

import (
	"fmt"
	"hung1299/go-projects/caching-proxy/global"
	"hung1299/go-projects/caching-proxy/internal/cache"
	"hung1299/go-projects/caching-proxy/internal/routes"
	"hung1299/go-projects/caching-proxy/pkg/util/regex"
	"hung1299/go-projects/caching-proxy/pkg/util/text"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

const (
	Port       = "port"
	Origin     = "origin"
	ClearCache = "clear-cache"
)

func checkErr(log string, err error) {
	if err != nil {
		fmt.Println(log, err)
		os.Exit(1)
	}
}

func clearCache() {
	cache.Storage = cache.Items{}
	cache.Storage.Save()
}

func InitializeCommand() {
	c := &cobra.Command{
		Use:   "caching-proxy",
		Short: "caching-proxy is a caching proxy server",
		Long:  "caching-proxy is a caching proxy server, it will forward requests to the actual server and cache the response",
		Run:   run,
	}

	c.PersistentFlags().Int32P(Port, "P", 3000, "port is the port on which the caching proxy server will run")
	c.PersistentFlags().StringP(Origin, "O", "", "origin is the URL of the server to which the requests will be forwarded")
	c.PersistentFlags().BoolP(ClearCache, "C", true, "use this to clear all cache")
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	b := cmd.Flags().Changed(ClearCache)
	if b {
		clearCache()
		return
	}

	p, err := cmd.Flags().GetInt32(Port)
	checkErr("error occurred!!", err)

	o, err := cmd.Flags().GetString(Origin)
	checkErr("error occurred!!", err)

	if text.IsEmpty(o) {
		fmt.Println("origin can not be empty!!")
		os.Exit(1)
	}

	if b, _ := regex.IsUrl(o); !b {
		fmt.Println("origin must be url format!!")
		os.Exit(1)
	}

	u, err := url.Parse(o)
	if err != nil {
		fmt.Println("Something happen when parsing url", err)
		os.Exit(1)
	}

	// TODO: handle case user not enter right url format
	if u.Scheme == "" {
		fmt.Println("origin is not a valid url!!")
		os.Exit(1)
	}

	global.Config.Port = p
	global.Config.Origin = fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	cache.InitCache()
	routes.InitRouter()
}
