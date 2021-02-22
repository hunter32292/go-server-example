# Go Server Example
[![Build Status](https://github.com/hunter32292/go-server-example/actions/workflows/build-status.yml/badge.svg)](https://github.com/hunter32292/go-server-example/actions/workflows/build-status.yml)
[![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fhunter32292%2Fgo-server-example.git)](https://jessemillar.com/r/man-hours)

This is an example project to show how to setup a simple go server that produces either a webpage or an API endpoint of data. Currently data is generated with [mockaroo](https://www.mockaroo.com/) and that data is stored in a CSV file under [data](./data), this information is stored in a slice for quick access of in memory data.

## Recent Additions

- Added Metrics Exposure and scrapping via prometheus server
- Added Logging via file and injection pipeline with FileBeat, Elastic, and Kibana
- Added Tracing with jaeger and opentracing
- Added Kubernetes deployment example
- Added Routing on 404 and home redirect
- Added CRUD actions for mock user data

### Tech

The specific goal of this project was to teach a basic structure for a golang application and the surrounding components to productionalize the service. Using as many native golang libraries and very few 3rd party libraries the project is meant to be easily understood and clear from the start.

* Golang
* Docker
* TravisCI

### Installation

The server requires [Go](https://golang.org/) to run.
Once golang has been installed, run `make` command to start service.

```sh
$ make
```

For production environments...

```sh
$ make docker
$ docker run -e LOG_FILE=service-name --rm -d -p 8080:8080 IMAGE_TAG
```


### Todos

 - Add End to End testing Example
 - Streamline Main Server and Router setup
 - Add config file based startup sequence


License
----

[MIT](LICENSE)
