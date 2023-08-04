


# Event-Driven Architecture in Golang

#### 👨‍💻 Full list what has been used:
[EventStoreDB](https://www.eventstore.com) The database built for Event Sourcing<br/>
[gRPC](https://github.com/grpc/grpc-go) Go implementation of gRPC<br/>
[PostgreSQL](https://www.postgresql.org) Application database<br/>
[gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) gRPC to JSON proxy generator following the gRPC HTTP spec<br/>
[Chi](https://github.com/go-chi/chi) A lightweight, idiomatic and composable router for building Go HTTP services<br/>


## Build and run with Docker compose
1. `make up`
2. `make db`
3. `make migrate`

### Software and Hardware List
| Software required      | OS required                        |
|------------------------|------------------------------------|
| Go 1.18+               | Windows, Mac OS X, and Linux (Any) |
| Docker 20.10.x         | Windows, Mac OS X, and Linux (Any) |
| EventStoreDB 22.10.x   | Windows, Mac OS X, and Linux (Any) |
| NATS 2.9               | Windows, Mac OS X, and Linux (Any) |

### Swagger UI:

http://localhost:8080

### EventStoreDB UI:

http://localhost:2113