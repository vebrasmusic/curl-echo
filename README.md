# curl-echo

`curl-echo` is a command-line interface (CLI) tool 
designed to help developers manage and test API routes 
by echoing their responses. 
It provides a simple way to initialize, add, list, and 
remove API routes, as well as execute them to see their responses.

## Features

- **Initialize**: Set up `curl-echo` in your project with default configurations.
- **Add**: Add new API routes to be echoed.
- **List**: Display all available routes or filter by a specific group.
- **Echo**: Run API routes and save their responses to files.
- **Remove**: Delete the `curl-echo` folder and all its contents.

## Installation

To install `curl-echo`, clone the repository and build the CLI using Go:

```
git clone https://github.com/vebrasmusic/curl-echo.git
cd curl-echo
go build -o curl-echo
```

## Usage

### Initialize

Initialize `curl-echo` in your project:

```
./curl-echo init
```

### Add

Add a new API route:

```
./curl-echo add
```

### List

List all available routes or filter by a specific group:

```
./curl-echo list
./curl-echo list -g <group>
```

### Echo

Run API routes and save their responses:

```
./curl-echo echo
./curl-echo echo -r <route>
./curl-echo echo -n <nickname>
./curl-echo echo -g <group>
```

### Remove

Delete the `curl-echo` folder and all its contents:

```
./curl-echo rm
```

## Configuration

The configuration file `config.json` is created during initialization and contains settings such as the root API path and maximum echo timeout.

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.