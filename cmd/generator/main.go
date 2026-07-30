//go:build ignore

package main

import (
	"flag"
	"log"
	"strings"

	"github.com/untrustedmodders/go-plugify"
)

type sliceFlags []string

func (i *sliceFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *sliceFlags) Set(value string) error {
	for _, dep := range strings.Split(value, ",") {
		trimmed := strings.TrimSpace(dep)
		if trimmed != "" {
			*i = append(*i, trimmed)
		}
	}
	return nil
}

func main() {
	name := flag.String("name", "menus", "Plugin name")
	version := flag.String("version", "1.0.0", "Plugin version")
	desc := flag.String("desc", "menus", "Plugin description")
	author := flag.String("author", "hdmi1519", "Author")
	website := flag.String("website", "https://github.com/hdmi1519/", "Website")
	license := flag.String("license", "MIT", "License")
	output := flag.String("output", "menus.pplugin", "Output file")
	entry := flag.String("entry", "bin/menus", "Entry binary name")

	var dependencies sliceFlags
	flag.Var(&dependencies, "dependencies", "List of dependencies")

	flag.Parse()

	if len(dependencies) == 0 {
		dependencies = append(dependencies, "s2sdk", "polyhook")
	}

	params := plugify.GenerateParams{
		Version:      *version,
		Description:  *desc,
		Author:       *author,
		Website:      *website,
		License:      *license,
		Platforms:    []string{},
		Dependencies: dependencies,
		Conflicts:    nil,
	}

	err := plugify.Generate(
		*output,
		*name,
		*entry,
		false,
		params,
		nil,
	)

	if err != nil {
		log.Fatalf("Gen error: %v", err)
	}
}
