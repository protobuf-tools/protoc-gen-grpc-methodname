// SPDX-FileCopyrightText: Copyright 2022 The protobuf-tools Authors
// SPDX-License-Identifier: BSD-3-Clause

// Command protoc-gen-grpc-methodname generates gRPC full method name constants.
package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/protobuf-tools/protoc-gen-grpc-methodname/methodname"
)

func main() {
	flags := flag.NewFlagSet("protoc-gen-grpc-methodname", flag.ExitOnError)

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if f.Generate {
				methodname.GenerateFile(p, f)
			}
		}

		return nil
	})
}
