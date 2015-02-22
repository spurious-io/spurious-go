package main

import (
  "os"
    "github.com/codegangsta/cli"
    )

func main() {
  app := cli.NewApp()
  app.Name = "Spurious"
  app.Usage = "Spurious is a toolset allowing development against a subset of AWS resource, locally."
  app.Run(os.Args)
}
