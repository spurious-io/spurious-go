package main

import (
  "os"
    "github.com/codegangsta/cli"
    )

func main() {
  app := cli.NewApp()
  app.Name = "spurious"
  app.Usage = "Spurious is a toolset allowing development against a subset of AWS resource, locally."

  app.Commands = []cli.Command{
    {
      Name:      "init",
      Usage:     "Pulls down the images for and creates the containers.",
      Action: func(c *cli.Context) {
        println("Intialized containers", c.Args().First())
      },
    },
    {
      Name:      "start",
      ShortName: "s",
      Usage:     "Starts the containers for running the services.",
      Action: func(c *cli.Context) {
        println("Starts containers", c.Args().First())
      },
    },
    {
      Name:      "stop",
      ShortName: "st",
      Usage:     "Stops the containers for the spurious services.",
      Action: func(c *cli.Context) {
        println("Stops containers", c.Args().First())
      },
    },     
  }

  app.Run(os.Args)
}
