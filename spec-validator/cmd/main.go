package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/newrelic/newrelic-integration-e2e-action/spec-validator/pkg"
)

var log = logrus.New()

const (
	flagSpecPath = "spec_path"
)

func main() {
	log.Info("running spec-validator")
	specsPathPtr := flag.String(flagSpecPath, "", "Relative path to the spec file")
	flag.Parse()
	specPath := *specsPathPtr
	if specPath == "" {
		os.Exit(1)
	}
	content,err:=ioutil.ReadFile(specPath)
	if err!=nil{
		log.Error(err)
		os.Exit(1)
	}
	pkg.ParseSpecFile(content)
}

