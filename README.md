# SAP History  [![Go Report Card](https://goreportcard.com/badge/github.com/ulranh/saphistory)](https://goreportcard.com/report/github.com/ulranh/saphistory)

The purpose of this app is to help solve problems related to SAP ABAP instances. A snapshot of various transactions is taken every minute. Thus, when analyzing, one gets a more accurate idea of the state that existed at a particular time. Because the situation can change every millisecond, this is of course not exact science but you can get a much better understanding of whats going on at the time a problem occured.

# Prerequisites

## NWRFC SDK
You need the current SAP NWRFC SDK 7.50 library as a prequisite for the installation of this app. To download this library you must have a customer or partner account on the SAP Service Marketplace. Please take a look at SAP note "2573790 - Installation, Support and Availability of the SAP NetWeaver RFC Library 7.50" and the [gorfc](https://github.com/SAP/gorfc) readme.

With the nwrfcsdk zip file unpacked in /usr/sap, the following environment variables are necessary under Linux:

```
LD_LIBRARY_PATH="/usr/sap/nwrfcsdk/lib"
CGO_LDFLAGS="-L /usr/sap/nwrfcsdk/lib"
CGO_CFLAGS="-I /usr/sap/nwrfcsdk/include"
CGO_LDFLAGS_ALLOW=.*
CGO_CFLAGS_ALLOW=.*
```

## SAP User 
A SAP user is necessary for every SAP system with read access for all affected remote function modules. 

# Data
At the moment, the following data is collected:

| Transaction | Function Module  | Description |
| ----------  | ---------------  |------------ |
| SM66        | TH_WPINFO        | Current workprocesses |
| SM37        | RFC_READ_TABLE   | Active jobs and aborted jobs of the current day |
| SM13        | RFC_READ_TABLE   | Update table |
| SM12        | ENQUE_READ       | Current lock table |
| SM58        | RFC_READ_TABLE   | Transactional RFC's |
| WE02        | RFC_READ_TABLE   | Idocs of the last minute|
| AL08        | TH_USER_LIST     | Active users of the last minute |
| SM21        | /SDF/GET_SYS_LOG | Syslog entries of the last minute|

After each scrape, the data is serialized using the [protocol buffers](https://github.com/protocolbuffers/protobuf) method, compressed according to the [Zstdandard](https://github.com/facebook/zstd) algorithm and then written to a [Badger DB](https://github.com/dgraph-io/badger). Web [gRPC](https://github.com/improbable-eng/grpc-web) is responsible for the exchange of data between the [Go](https://golang.org/) backend and the [Vue](https://vuejs.org/) client.

# Installation
## Binary
To build the app you need the [Go](https://golang) programming language, the [make](https://www.gnu.org/software/make/) tool and [npm](https://www.npmjs.com/) for the client.
### Without Tls
```
$ export VUE_APP_PORT=8080

$ make client
$ make generate-build
```
Then the app can be started with
```
$ ./saphistory -dbstore </path/to/database directory>
```
where the user needs write rights for the database directory. Afterwards the app should be accessible in the browser on the port 8080
```
http://localhost:8080
```
### With Tls
For tls you also need a tls directory with the files tls.crt and tls.key. To test it with localhost, you can make the certificates for example with [mkcert](https://github.com/FiloSottile/mkcert).

```
$ export VUE_APP_PORT=8443
$ export VUE_APP_TLS_PATH=</path/to/tls directory>

$ make client
$ make generate-build
$ ./saphistory -dbstore </path/to/database directory>

https://localhost:8443
```
## Docker
### Without Tls
The nwrfcsdk directory has to be copied into the app directory.

```
$ export PORT=8080
$ export VERSION=0.1.0
$ export DBSTORE=</path/to/database directory>

$ make docker-build
$ make docker-run

http://<hostname>:8080
```
### With Tls
The nwrfcsdk directory and the tls directory with the certificates have to be copied into the app directory.

```
$ export PORT=8443
$ export VERSION=0.1.0
$ export TLS_PATH=/app/tls
$ export DBSTORE=</path/to/database directory>

$ make docker-build
$ make docker-run

https://<hostname>:8443
```

# Usage
The first thing you need to do is create a system entry. About a minute later, after the first scrape, you can go to the system page via the 'Goto system data' button and should see the associated system information.

The image below shows the table of the SAP lock entries at a given time:

 ![saphistory](/images/saphistory.png)

# ToDo
* Writing tests
* Possibility to delete old database content
* Adapt transaction data
