# State Boundries

RESTful API written in Go for finding out which states border which.

Utilizes the [gorilla/mux](https://pkg.go.dev/github.com/gorilla/mux) package for routing and a few built-in Go packages.

## Live Server
This project has been deployed on Heroku. If you use visit `https://state-boundries.herokuapp.com/api/borders/`, it will `GET` the list of all 50 states and the states that they share a border with.


## Installation & Testing locally

Make sure that you have the latest version of [Go](https://golang.org/dl/) installed on your machine.

You can clone this repository, install the requisite dependencies and run it on your machine like so:

```
$ git clone https://github.com/J-Obog/state-boundries.git
$ cd state-boundries
$ go install
$ go run .
```

## Endpoints

`GET` **/api/borders/:state** - *Get the a specific state and the states that border it*

`GET` **/api/borders/** - *Get all states and their borders*