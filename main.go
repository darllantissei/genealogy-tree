package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	dbcommon "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/common"
	webserver "github.com/darllantissei/genealogy-tree/adapters/web"
	"github.com/darllantissei/genealogy-tree/application"
	factorycommonservice "github.com/darllantissei/genealogy-tree/factories/common_service"
	factoryperson "github.com/darllantissei/genealogy-tree/factories/person"
	factoryrelationship "github.com/darllantissei/genealogy-tree/factories/relationship"

	_ "github.com/mattn/go-sqlite3"
)

const instructionsService = `
	Type of service undefined. Valid types: 
	http: Will start the application in the mode web, delivering resources of API.
	E.g.: When start the application, set the flag service=<service-type>. Like this: -service=http
	`
const instructionsPort = `Port of application will running`

const instructionsDebug = `Information of debug`

const flagService = "service"
const flagPort = "port"
const flagDebug = "debug"

const typeServiceHTTP = "http"

func init() {
	flag.String(flagService, typeServiceHTTP, instructionsService)
	flag.String(flagPort, "port", instructionsPort)
	flag.String(flagDebug, "debug", instructionsDebug)
}

func main() {
	port := getPortHTTP()
	serviceType := getServiceType()
	inDebugMode := getInDebugMode()

	// sourceDataBase := "./database/genealogy-tree.sqlite"
	sourceDataBase := os.Getenv("DATA_SOURCE_APP")

	if strings.EqualFold(sourceDataBase, "") {
		sourceDataBase = "database/genealogy-tree.sqlite"
	}

	factoryCommonService := factorycommonservice.FactoryCommonService{}

	factoryPerson := factoryperson.FactoryPerson{
		CommonDB:      *dbcommon.NewCommon(sourceDataBase),
		CommonService: factoryCommonService.NewService(),
	}

	factoryRelationship := factoryrelationship.FactoryRelationship{
		CommonDB:      *dbcommon.NewCommon(sourceDataBase),
		CommonService: factoryCommonService.NewService(),
		PersonService: factoryPerson.NewService(),
	}

	application := application.Application{
		CommonService:       factoryCommonService.NewService(),
		PersonService:       factoryPerson.NewService(),
		RelationshipService: factoryRelationship.NewService(),
	}

	switch serviceType {
	case typeServiceHTTP:

		serverWeb := webserver.MakeNewWebServer(application)

		serverWeb.Serve(port, inDebugMode)

	default:
		panic("Type service not setted")
	}

	_ = port

	_ = serviceType
}

func getPortHTTP() (port int) {
	flag.Parse()
	flagPort := flag.Lookup("port")

	if flagPort != nil {
		value := flagPort.Value.String()

		strPort, err := strconv.Unquote(value)

		if err != nil {
			strPort = value
		}

		portEnv, err := strconv.Atoi(strPort)

		if err != nil {
			port = 9000
			return
		}

		port = portEnv
	}

	return
}

func getServiceType() (serviceType string) {

	flag.Parse()

	flagServiceType := flag.Lookup(flagService)

	if flagServiceType == nil {
		panic(instructionsService)
	}

	flagAllowed := []string{
		"http",
		"agent",
	}

	if !strings.Contains(strings.Join(flagAllowed, "|"), flagServiceType.Value.String()) {
		panic(instructionsService)
	}

	serviceType = flagServiceType.Value.String()

	return

}

func getInDebugMode() (inDebugMode bool) {
	flag.Parse()
	flagDebugMode := flag.Lookup("debug")

	if flagDebugMode != nil {
		value := flagDebugMode.Value.String()

		strPort, err := strconv.Unquote(value)

		if err != nil {
			strPort = value
		}

		debugMode, err := strconv.ParseBool(strPort)

		if err != nil {
			inDebugMode = true
			return
		}

		inDebugMode = debugMode
	}

	return
}
