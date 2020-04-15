package main

import (
	"./controller"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-stack/stack"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	newLogger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger := newLogContext(newLogger, "API")

	// Download and load config files
	workingdir, err := os.Getwd()
	configRepo := os.Getenv("CONFIG_REPO")
	if err != nil {
		logger.Log("FATAL", "unable to get current working directory")
		os.Exit(1)
	}

	envName := "dev" // Default to dev
	if os.Getenv("ENV") != "" {
		envName = os.Getenv("ENV")
	}

	filename := fmt.Sprintf("crypto-currency-%s.json", envName)

	data, err := ioutil.ReadFile(path.Join(workingdir, filename))
	if err != nil {
		panic(err)
	}

	if data == nil {
		if err = DownloadFile(path.Join(workingdir, filename), fmt.Sprintf("%s/%s", configRepo, filename)); err != nil {
			logger.Log("FATAL", "unable to download config file")
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	//Set up and start http server
	router := mux.NewRouter()

	// Healthcheck and Prometheus routes
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(map[string]string{"status": "up"})
	})

	apiRouter := router.PathPrefix("/v1.0/api").Subrouter()

	apiRouter.HandleFunc("/currency/{symbol}", controller.GetCryptoCurrency).Methods("GET")
	apiRouter.HandleFunc("/currency", controller.GetAllCryptoCurrency).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	logger.Log("event", "serving...")

	err = server.ListenAndServe()

	if err != nil {
		logger.Log("event", "exiting", "err", err)
		os.Exit(1)
	}
}

func newLogContext(logger log.Logger, app string) log.Logger {
	return log.With(logger,
		"time", log.DefaultTimestampUTC,
		"app", app,
		"caller", log.Valuer(func() interface{} {
			return pkgCaller{stack.Caller(3)}
		}),
	)
}

// pkgCaller wraps a stack.Call to make the default string output include the package path.
type pkgCaller struct {
	c stack.Call
}

func (pc pkgCaller) String() string {
	caller := fmt.Sprintf("%+v", pc.c)
	caller = strings.TrimPrefix(caller, "../crypto-currency-service/")
	return caller
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	client := http.Client{}
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	//request.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, response.Body)
	return err
}
