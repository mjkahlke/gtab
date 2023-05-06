# Architecture
This document describes the code layout and how to build and run **gtab**.

## Code Layout
The layout is fairly simple. Shared implementation code is kept in its own directory while applications that run **gtab** are kept in separate directories. Unit tests are kept in the same directory as the module they test.

### `/`
The toplevel directory contains markdown and build files:
* `build.bat` - Batch file for building **gtab** on Windows
* `Dockerfile` - pulled in by docker compose
* `docker-compose.yml` - for building the REST service in a container

### `impl`
The implementation directory contains global variables, constants, utilities, and algorithms.

### `cmd`
This directory contains subdirectories for each application:
* `cli` - the command line version. Run `cli -usage` to see how to use it.
* `rest` - the REST API version. Use curl, postman, or your browser to query `http://localhost:7777/tabs/{chord}`

## Build

### Docker
I use Rancher Desktop on my PC, just use docker instead of nerdctl if you use Docker desktop.

From the toplevel directory of the repository, execute `nerdctl compose build --file docker-compose.yml`. Verify that your image was built via `nerdctl image ls`

Start up the container with `nerdctl run -d -p 7777:7777`

Find the name of the container using `nerdctl ps` and note the container name. When you're finished, execute `nerdctl kill <container_name>` to terminate the REST service.

### Windows
`build.bat` should be run from the toplevel directory of the repository. It cleans up then regenerates go.mod files, runs unit tests, and builds (but does not install) standalone executable applications in subdirectories under `cmd`.

## Run

### CLI

| parameter | default | description |
|:----------|:-------:|:------------|
| `-chord`  | C       | C major     |
| `-list`   | false   | list of supported chord suffixes |

### REST

| parameter       | description |
|:----------------|:------------|
| `/tabs/{chord}` | returns tabs for named chord     |
| `/list`         | list of supported chord suffixes |

## Sample CLI output

The output of several runs are shown below.

### Usage
    # gtab -help
    Usage of gtab:
      -chord string
            set to chord name, b=flat, #=sharp (default "C")
      -list
            list supported chord formulas

### Cmaj7
    # gtab -chord Cmaj7
    8 10 9 9 8 8
    x 3 2 0 0 0
    x 3 5 4 5 3
    x x 10 12 12 12
    x x x 5 5 7

### B7
    # gtab -chord B7
    7 9 7 8 7 7
    x 2 1 2 0 2
    x 2 4 2 4 2
    x x 9 11 10 11
    x x x 4 4 5

### Abmaj7
    # gtab -chord Abmaj7
    4 3 1 0 1 3
    4 6 5 5 4 4
    x 11 13 12 13 11
    x x 6 8 8 8
    x x x 1 1 3

### D
    # gtab -chord D
    10 12 12 11 10 10
    x 5 7 7 7 5
    x x 0 2 3 2
    x x x 7 7 10
    
### F#sus4
    # gtab -chord F#sus4
    2 2 4 4 2 2
    2 2 4 x 4 1
    x 9 9 11 12 9
    x x 4 4 7 7
    x x x 11 12 14

## Sample REST API output

Sharp chords have a hash-tag symbol `#` which must be escaped before it can be used as a REST API endpoint. Replace `#` with either `%23` or `H`. All responses set the Content-Type header to `application/json` and return indented JSON.

### Cmaj7

    GET http://localhost:7777/tabs/Cmaj7
    {
        "chord": "Cmaj7",
        "tabs": [
            " 8 10 9 9 8 8",
            " x 3 2 0 0 0",
            " x 3 5 4 5 3",
            " x x 10 12 12 12",
            " x x x 5 5 7"
        ]
    }

### B7

    GET http://localhost:7777/tabs/B7
    {
        "chord": "B7",
        "tabs": [
            " 7 9 7 8 7 7",
            " x 2 1 2 0 2",
            " x 2 4 2 4 2",
            " x x 9 11 10 11",
            " x x x 4 4 5"
        ]
    }

### Abmaj7

    GET http://localhost:7777/tabs/Abmaj7
    {
        "chord": "Abmaj7",
        "tabs": [
            " 4 3 1 0 1 3",
            " 4 6 5 5 4 4",
            " x 11 13 12 13 11",
            " x x 6 8 8 8",
            " x x x 1 1 3"
        ]
    }

### D

    GET http://localhost:7777/tabs/D
    {
        "chord": "D",
        "tabs": [
            " 10 12 12 11 10 10",
            " x 5 7 7 7 5",
            " x x 0 2 3 2",
            " x x x 7 7 10"
        ]
    }

### F#sus4

    GET http://localhost:7777/tabs/F%23sus4
    {
        "chord": "F#sus4",
        "tabs": [
            " 2 2 4 4 2 2",
            " 2 2 4 x 4 1",
            " x 9 9 11 12 9",
            " x x 4 4 7 7",
            " x x x 11 12 14"
        ]
    }

    GET http://localhost:7777/tabs/FHsus4
    {
        "chord": "F#sus4",
        "tabs": [
            " 2 2 4 4 2 2",
            " 2 2 4 x 4 1",
            " x 9 9 11 12 9",
            " x x 4 4 7 7",
            " x x x 11 12 14"
        ]
    }

