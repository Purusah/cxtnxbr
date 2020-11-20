Test project to implement uri visits counter and check default project layout

Requirements:
* `docker`

Start:
1. `make build` - build docker container with app
2. `make redis` - (optional) start redis instance (or change config at `scripts/default.env`)
3. `make run` - start app server

Additional commands:
* `make test` - run test
* refer to `Makefile`
