package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ehazlett/interlock"
)

var (
	name    = "example plugin"
	version = "0.1"
	author  = "@ehazlett"
	url     = "github.com/ehazlett/interlock-plugins/tree/master/example"
)

func handle(in *interlock.PluginInput) (*interlock.PluginOutput, error) {
	var out *interlock.PluginOutput

	switch in.Command {
	case "info":
		info := &interlock.Plugin{
			Name:    name,
			Version: version,
			Author:  author,
			Url:     url,
		}

		out = &interlock.PluginOutput{
			Command: in.Command,
			Output:  nil,
			Error:   nil,
		}

		data, err := json.Marshal(info)
		if err != nil {
			out.Error = bytes.NewBufferString(err.Error()).Bytes()
		}

		if data != nil {
			out.Output = data
		}
	default:
		out = &interlock.PluginOutput{
			Command: in.Command,
			Output:  []byte(name),
			Error:   nil,
		}
	}

	return out, nil
}

func main() {
	in, err := interlock.GetPluginInput()
	if err != nil {
		log.Fatal(err)
	}

	// if not piped input
	if in == nil {
		fmt.Printf("%s %s\n", name, version)
		os.Exit(0)
	}

	output, err := handle(in)
	if err != nil {
		log.Fatal(err)
	}

	if output != nil {
		if err := interlock.SendPluginOutput(output); err != nil {
			log.Fatal(err)
		}
	}

}
