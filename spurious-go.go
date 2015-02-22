package main

import (
  "os"
    "github.com/codegangsta/cli"
    )

func main() {
  app := cli.NewApp()
  app.Name = "spurious"
  app.Usage = "Spurious is a toolset allowing development against a subset of AWS resource, locally."
  app.Action = func(c *cli.Context) {
    println("boom! I say!")
  }
  app.Run(os.Args)
}
