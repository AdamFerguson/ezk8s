package main

import (
	"fmt"
	"os"

	"github.com/tma1/ezk8s"
	"github.com/tma1/ezk8s/config"
	"github.com/tma1/ezk8s/query"
)

func main() {
	conf, err := config.LoadFromKubeConfig("", "minikube")

	cl := conf.Client(
		ezk8s.QueryOpts(
			query.Host("192.168.99.100:8443"),
			query.Scheme("https"),
		),
	)

	res, err := cl.Query(
		query.Pod(""),
		query.Label("name", "nginx"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//var resourceVersion string
	//var generation float64
	//err = res.Scan(
	//	query.Path{"$.metadata.resourceVersion", &resourceVersion},
	//	query.Path{"$.metadata.generation", &generation},
	//)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//fmt.Printf(
	//	"generation = %v\nresourceVersion = %v\n",
	//	generation, resourceVersion,
	//)

	var names []string
	err = res.Scan(
		query.Path{"$.items[:0].metadata.name", &names},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("pod names:")
	for _, name := range names {
		fmt.Println(name)
	}
}
