package api

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"loxilb/api/restapi"
	"loxilb/api/restapi/operations"
	"loxilb/options"
)

type ApiHookInterface interface {
    LoxiLBRuleAdd(int) int
}

var ApiHooks ApiHookInterface

func RegisterApiHooks(hooks ApiHookInterface) {
	ApiHooks = hooks
}

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func RunApiServer() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewLoxilbRestAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Loxilb Rest API"
	parser.LongDescription = "Loxilb REST API for Baremetal Scenarios"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	// API server host list
	server.Host = options.Opts.Host
	server.TLSHost = options.Opts.TLSHost
	server.TLSCertificateKey = options.Opts.TLSCertificateKey
	server.TLSCertificate = options.Opts.TLSCertificate
	server.Port = options.Opts.Port
	server.TLSPort = options.Opts.TLSPort

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
