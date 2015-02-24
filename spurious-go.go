package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/fsouza/go-dockerclient"
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

        var channel = make(chan ImageStatus)

        images := [...]string{"spurious/sqs", "spurious/s3", "spurious/dynamodb", "spurious/browser"}
        for _, image := range images {
          go getImage(image, channel)
        }

        for {
          entry := <- channel
          print(entry.image + " - " + string(entry.output[:]))
        }

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
    {
      Name:      "delete",
      ShortName: "d",
      Usage:     "Deletes the spurious containers and images.",
      Action: func(c *cli.Context) {
        println("Deletes containers", c.Args().First())
      },
    },     

  }

  app.Run(os.Args)
}

type Output struct {
  image string
  ch chan ImageStatus
}

type ImageStatus struct {
  image string
  output []byte
}

func (o *Output) Write(p []byte) (n int, err error) {
  o.ch <- ImageStatus{image: o.image, output: p}
  return
}

func getImage(image string, channel chan ImageStatus) {
  endpoint := os.Getenv("DOCKER_HOST")
  client, _ := docker.NewTLSClient(endpoint, "/Users/stevenjack/.boot2docker/certs/boot2docker-vm/cert.pem", "/Users/stevenjack/.boot2docker/certs/boot2docker-vm/key.pem", "/Users/stevenjack/.boot2docker/certs/boot2docker-vm/ca.pem")

  output := Output{image: image, ch: channel}

  err := client.PullImage(docker.PullImageOptions{Repository: image, OutputStream: &output}, docker.AuthConfiguration{})

  if err != nil {
    println("Error pulling image: ", err)
  }

  println("Container: " + image + " finished")
}
