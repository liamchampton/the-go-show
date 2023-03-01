# The Go Show EP2

## Creating your first web server in Go

Hello everyone and welcome to today's episode where we'll be creating your first web server in Go! This session is tailored for new developers in the tech industry and those who want to learn about the language.

In this session, we'll be utilizing the standard library in Go and showcasing the power of GitHub Copilot, a powerful tool that can help you write code faster and more efficiently.

By the end of this session, you'll have a solid understanding of how to build a basic web server in Go and be equipped with the tools to continue exploring this exciting language. So, grab your coffee and join us for a fun and informative session!

Recording link: https://www.youtube.com/live/uhhxPZNKRWM?feature=share

## To run this in Docker or Podman locally
    
```bash
// navigate to the root of the project file
cd web-server

// build the image
docker build --tag docker.io/goshow:v1 .   

// run the image
docker run -dt -p 8080:8080/tcp docker.io/library/goshow:v1
```