package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func forTestGetIdxArg(t *testing.T, nameArg string) int {
	idxPort := -1
	for idx, arg := range os.Args {

		if strings.Contains(arg, nameArg) {
			idxPort = idx
		}
	}

	return idxPort
}

func Test_getPortHTTP(t *testing.T) {

	const flagPort = "port"

	os.Args = append(os.Args, `-port=9100`)

	port := getPortHTTP()

	require.GreaterOrEqual(t, port, 9100, "the port of application informed as integer must be returned as integer - as must be")

	idxPort := forTestGetIdxArg(t, flagPort)

	os.Args[idxPort] = `-port="9200"`

	port = getPortHTTP()

	require.GreaterOrEqual(t, port, 9200, "the port of application informed as string must be returned as integer - as must be")

}

func Test_getServiceType(t *testing.T) {

	typeService := getServiceType()

	require.Equal(t, "http", typeService, "Type of service undefined, then the default will the service http")

	os.Args = append(os.Args, "-service=undefined")

	require.Panics(t, func() {
		typeService = getServiceType()

		errRecover := recover()

		require.NotNil(t, errRecover, "Panic occurred because flag service not setted")
		require.Empty(t, typeService, "Will not return any type service")
	})

}

func Test_getInDebugMode(t *testing.T) {
	const flagDebug = "debug"

	os.Args = append(os.Args, `-debug=true`)

	isDebug := getInDebugMode()

	require.True(t, isDebug, "the debug mode of application informed as boolean must be returned as boolean - as must be")

	idxDebug := forTestGetIdxArg(t, flagDebug)

	os.Args[idxDebug] = `-debug="false"`

	isDebug = getInDebugMode()

	require.False(t, isDebug, "the debug mode of application informed as string must be returned as boolean - as must be")
}
