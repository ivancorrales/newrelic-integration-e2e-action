package executor

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/newrelic/newrelic-integration-e2e/pkg/agent"
	"github.com/newrelic/newrelic-integration-e2e/pkg/settings"
	"github.com/sirupsen/logrus"
)

func Exec(ag agent.Agent, settings settings.Settings) error{
	spec := settings.Spec()
	for i := range spec.Scenarios {
		scenario := spec.Scenarios[i]
		settings.Logger().Debugf("[scenario]: %s", scenario.Description)
		if err:=ag.SetUp(scenario);err!=nil{
			return err
		}
		if err := executeOSCommands(settings.RootDir(),scenario.Before); err != nil {
			return err
		}
		if err:=ag.Launch();err!=nil{
			return err
		}

		/**
		This block will be used to run the tests
		 */




		time.Sleep(1 * time.Minute)



		if err := executeOSCommands(settings.RootDir(),scenario.After); err != nil {
			println(err.Error())
		}
		if err:=ag.Stop();err!=nil{
			return err
		}
	}

	return nil
}

func executeOSCommands(rootDir string, statements []string) error {
	for i := range statements {
		stmt := statements[i]
		fmt.Println(stmt)
		cmd := exec.Command("bash", "-c", stmt)
		cmd.Dir=rootDir
		stdout, err := cmd.Output()
		if err != nil {
			logrus.Error(stdout)
			return err
		}
	}
	return nil
}

