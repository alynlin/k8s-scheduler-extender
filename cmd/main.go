package main

import (
	"github.com/k8s-scheduler-extender/pkg/informers"
	"github.com/k8s-scheduler-extender/pkg/routes"
	"github.com/k8s-scheduler-extender/pkg/scheduler"
	"github.com/k8s-scheduler-extender/pkg/utils/signals"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/comail/colog"
	"github.com/julienschmidt/httprouter"
)

func main() {
	flag.CommandLine.Parse([]string{})

	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	level := StringToLevel(os.Getenv("LOG_LEVEL"))
	log.Print("Log level was set to ", strings.ToUpper(level.String()))
	colog.SetMinLevel(level)

	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		port = "39999"
	}

	// Set up signals so we handle the first shutdown signal gracefully.
	stopCh := signals.SetupSignalHandler()

	informerFactory := informers.SharedInformerFactory()

	informerFactory.Core().V1().Pods().Lister()
	informerFactory.Start(stopCh)
	informerFactory.WaitForCacheSync(stopCh)

	predicate := scheduler.NewSingleappPredicate()
	bind := scheduler.NewNoSingleapBind()

	router := httprouter.New()

	routes.AddPProf(router)
	routes.AddVersion(router)
	routes.AddPredicate(router, predicate)
	routes.AddBind(router, bind)

	log.Printf("info: server starting on the port :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func StringToLevel(levelStr string) colog.Level {
	switch level := strings.ToUpper(levelStr); level {
	case "TRACE":
		return colog.LTrace
	case "DEBUG":
		return colog.LDebug
	case "INFO":
		return colog.LInfo
	case "WARNING":
		return colog.LWarning
	case "ERROR":
		return colog.LError
	case "ALERT":
		return colog.LAlert
	default:
		log.Printf("warning: LOG_LEVEL=\"%s\" is empty or invalid, fallling back to \"INFO\".\n", level)
		return colog.LInfo
	}
}

func StringToInt(sThread string) int {
	thread := 1

	return thread
}
