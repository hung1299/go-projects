# Caching Proxy

The [Caching Proxy](https://roadmap.sh/projects/caching-server) is a caching proxy server.

Idea from [roadmap.sh](https://roadmap.sh/golang/projects)

## How to run

You need to clone this project into your local machine first.

by running these commands:

```bash
git clone https://github.com/hung1299/go-projects.git
cd caching-proxy
```

Run the following command to build and run the project:

```bash
make build
or
make build-window # for window

# To start the caching proxy server
caching-proxy --port <number> --origin <url>
or
caching-proxy -P <number> -O <url>
# Or just define the origin and the default port will be 3000

# To clear all the cache
caching-proxy --clear-cache
or
caching-proxy -C

# To see the list of available commands
caching-proxy --help
```
