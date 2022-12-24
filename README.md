# Weather Service Project

## Overview
This project provides an HTTP Server that exposes an endpoint for gathering weather related information for a particular 
area. This endpoint recieves latitude and longitude coordinates and should return what the weather conditions are outside
in the area specified by the coordinates. The weather conditions will indicate whether there is snow or rain in the area
or whether it's hot, cold or moderate outside. Any weather alerts for a particular area will be returned by API as well.

## Content
- [Getting Started](#getting-started)
- [Application Architecture](#application-architecture)
- [Project Structure](#project-structure)
- [Api Endpoints](#api-endpoints)

## Getting Started
- Build application using Docker Compose:

```sh
$ make build
```

- Build & run the application using Docker Compose:

```sh
$ make up
```

- View the application logs while the Docker container is running:

```sh
$ make logs
```

- Destroy application container and clean up:

```sh
make down
```

## Application Architecture
This application makes usage of the principles of the Clean Architecture (similar to Onion architecture, Hexagonal, Ports
and Adapters, etc) model for structuring how an application can be layered in a manner that separates application concerns
into distinct swappable units. The application utilizes dependency inversion techniques where the direction of the dependencies
goes from the outer layers of the application to the innermost layer (business logic). In doing so, the application is
primarily divided into two layers:
1. **Business logic** (Only standard Go libraries are used here).
2. **Tools** (databases, servers, message brokers, any other packages and frameworks).

This allows the inner layer with the business logic to remain clean and:
- Not have packages imported from the outer layer
- Use only the capabilities of the standard Go libraries
- Make all calls to the outer layer via interfaces

Therfore the business logic doesn't know anything about any web API or database implementations. It has an interface for working
with an abstraction of the other layers of the application.

The outer layer of the application has its own limitations in this model such as:
- All components of this layer remain unaware of each other.
- All calls to the inner layer are made via an interface
- Data is always transferred in a form that is convenient for consumption by the business logic.

## Project Structure
### `/main.go`
The main application file initializes the applications configuration and executes a start command for the HTTP Server
in `infrastructure/httpserver.go`. 

### `/adapter/`
This folder contains the handlers, presenters, controllers (interface that holds methods needed by the application which 
is implemented by Web, Devices and/or External devices). Interface adapters depend on the application business rules or 
use cases.

### `/config/`
The configuration is initialized by reading the config.yml to set up the default environment variables for the application. 
The [viper](https://github.com/spf13/viper) library has been used for loading the configuration file. Please note that 
only the non-securirty-sensitive default values are stored in the config.yml file. All security-sensitive values should
stored in a vault solution or injected as secrets during the application deployment.

### `/domain/`
All business domain concerns are stored in the domain folder specifically within the entities subfolder.

### `infrastructure/`
For simplicity of the application there is a single _Start_ function in the `httpserver.go` file, which is called from the _main_ 
function. Most of the applications primary objects are created here including the starting of the main web server. 
The application relies on dependency injection via "New" constructors. This allows the application to be loosely coupled 
with layers of inter-connecting interfaces, which ultimately keeps the business (domain specific) logic independent of 
the other layers of the application. The benefit is that the layers of the application can be instrumented with different 
implementations which would not cause a rewrite of other parts of the application as a result of a change in instrumentation 
or implementation.

#### `/infrastructure/http`
- HTTP client and request abstractions used by the entire application 
- Client retry instrumentation and implementation
- These implementation can be substituted for other implementations

The server router constructor is implemented in the same pattern:
- Handlers are grouped by commonality
- For each grouping, a separate router structure is created with methods that process the associated path(s)
- The business logic interfaces and structures are injected into the router which will be called by the handlers

#### `/infrastructure/logger`
 - Logger interface for the entire application.
 - Utilizes the `logrus` package.
 - Implements field logging 

#### `/infrastructure/router`
Entities represent the business logic (models, domain, etc) and can be used in any layer. The specific business logic is
instrumented as usecases.

### `/usecase`
The business logic is grouped by what is termed `usecase` each with its own respective methods, structure, files, etc. For 
this application, t business logic works with an abstraction of the web API. If the instrumentation or implementation of 
the web API undergoes changes, the business logic does not need to modified to accommodate this change. We can always override 
the implementation of the interface without making changes to the usecase package.

In general:
- Methods are grouped by area of application (on a common basis)
- Each group has its own structure
- One file - one structure

**Please note: Repositories, webapi, rpc, and other business logic structures are injected into business logic structures.**


## API Endpoints

|    Endpoint    | HTTP Method |        Description         |
|:--------------:|:-----------:|:--------------------------:|
| `/v1/weather`  |   `POST`    | `Gets weather information` |
| `/healthCheck` |    `GET`    |  `Server heartbeat/ping`   |



### Test endpoints API

#### Curl Post Request

`Request`
```bash
curl --location --request POST 'http://127.0.0.1:3001/v1/weather' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude": "54.0307",
    "longitude": "47.2127"
}'
```

`Response`
```bash
HTTP/1.1 200 Ok
Content-Type: application/json
Date: Fri, 23 Dec 2022 15:09:27 GMT
Content-Length: 532
```
```json
{
  "Weather":{
    "base":"stations",
    "clouds":{
      "all":100
    },
    "cod":200,
    "coord":{
      "lat":54.0307,
      "lon":47.2127
    },
    "dt":1671845319,
    "id":473994,
    "main":{
      "feels_like":267.93,
      "grnd_level":996,
      "humidity":79,
      "pressure":1018,
      "sea_level":1018,
      "temp":273.68,
      "temp_max":273.68,
      "temp_min":273.68
    },
    "name":"Veshkayma",
    "sys":{
      "country":"RU",
      "sunrise":1671858571,
      "sunset":1671885090
    },
    "timezone":14400,
    "visibility":10000,
    "weather":[
      {
        "description":"overcast clouds",
        "icon":"04n",
        "id":804,
        "main":"Clouds"
      }
    ],
    "wind":{
      "deg":206,
      "gust":15.73,
      "speed":6.98
    }
  }
}
```

#### PostMan Collection
- Please see the weather_service_project.postman_collection.json file.

