# Api-Workbench

Api-Workbench is a Terminal User Interface (TUI) application
that allows you to create and run mock APIs from the commandline.
It provides a convenient way to simulate API endpoints during
development or testing without relying on a production environment.

## Features

- **Create Mock APIs**: Define new API configurations with custom
endpoints, responses, and HTTP methods.
- **Loading existing Configurations**: Start an HTTP server from
pre-existing API configurations defined in JSON.
- **Mock API Endpoints**: Easily mock production API endpoints by
specifying the host and path.
- **TUI Interface**: Interact with the application using a
user-friendly terminal interface.

## Installation

To install Api-Workbench, make sure you have Go installed on
your system. Then, run the following command:

```shell
go install github.com/christianhturner/api-workbench
```

## Usage

To start the Api-Workbench application, run:

```shell
api-workbench
```

## Debugging

To debug the application when developing do the following from the
source directory:

1. Build the code, and use the following build instructions:

```shell
go build -gcflags=all="-N -l"
```

2. Once, built, from a separate terminal, go to the source directory
and run the binary: `./api-workbench`

3. Set your breakpoints, and use the `dlv attach [PID]` command. 
If using Neovim, set your breakpoints, press continue in DAP
and search for api-workbench and run the process

4. Your now running the sourcecode!

## ROADMAP

### Release Ready TODOs

- Support Complete API creation and definition generation within the TUI, and store that within our database.
- Provide 1:1 feature support and compatibility with the HTTP Spec.
- Provide 1:1 feature support for RESTAPIs
- Provide the option to import API creation and definitions through importing JSON files rather than strictly throughthe TUI interface.
- Support exporting our API definitions into our JSON format, so that your data can be migrated between instances of api-workbench

### Post Release TODOs
- Add support for websocket spec 1:1
- Create tooling that would allow users to incorporate API-Workbench as an integration testing suite with their codebase (preferably language agnostic) and interactions should mirror interacting with the HTTP protocol as closely as possible.
