package cli

import (
	"errors"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/smallfish/simpleyaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"reflect"
)

var yamlCommand = &cobra.Command{
	Use:  "yaml",
	RunE: yaml,
}

var yamlVars = make(map[string]string)

func yaml(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("You must specify a yaml file and a template file")
	}
	return applyYaml(args[0], args[1])
}

func applyYaml(yamlFile string, templateFile string) error {
	vars, err := loadYamlFile(yamlFile)
	if err != nil {
		return err
	}
	dict, _ := vars.Map()
	recursiveParse("", dict)
	finalContent, err := replaceVars(templateFile)
	if err != nil {
		return err
	}
	fmt.Println(finalContent)
	return nil
}

func recursiveParse(varName string, dict map[interface{}]interface{}) {
	if varName != "" {
		varName += ":"
	}
	for k := range dict {
		if reflect.TypeOf(dict[k]).Kind() == reflect.Map {
			recursiveParse(varName+k.(string), dict[k].(map[interface{}]interface{}))
		} else {
			yamlVars[varName+k.(string)] = dict[k].(string)
		}
	}
}

func replaceVars(templateFile string) (string, error) {
	return mustache.RenderFile(templateFile, yamlVars)
}

func loadYamlFile(file string) (*simpleyaml.Yaml, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return simpleyaml.NewYaml(data)
}
