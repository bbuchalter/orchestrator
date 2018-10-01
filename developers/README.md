# Orchestrator for Developers

# Requirements
* `orchestrator` is built on Linux and OS/X.

# Quick Start Guide
* Install sqlite
* `script/build` # builds the orchestrator binary and installs the desired go version if needed
* `cp conf/orchestrator-sample-sqlite.conf.json conf/orchestrator.conf.json` # use a simple, default config
* `bin/orchestrator http` # start orchestrator
* Open browser to http://localhost:3000
* `developers/scripts/frontend_db_setup.sh` and follow instructions # build the mysql nodes for orchestrator to manage
* Then visit http://localhost:3000/web/clusters/ and see your cluster!

# Tests
## Unit tests
Simply run: `go test ./go/...`

## Integration tests
The integration test suite assumes you have `sqlite` and `wget` installed.
Simply run: `./tests/integration/test.sh`

If you have any trouble with the integration tests, you can get detailed command output by including `DEBUG=1` in the command.

The integration test suite will install [dbdeployer](https://github.com/datacharmer/dbdeployer) to install and run an isolated instance of MySQL.
Basic stop/start commands for these sandboxed installations are available from the `dbdeployer/sandbox/integration_tests` directory.

All variables which control the MySQL sandbox are configured at the top of [`tests/integration/test.sh`](/tests/integration/test.sh).

# Customization
There are some hooks in the Orchestrator web frontend which can be used to add customizations via CSS and JavaScript.
The corresponding files to edit are `resources/public/css/custom.css` and `resources/public/js/custom.js`.
You can find available hooks via `grep -r 'orchestrator:' resources/public/js`.
Please note that all APIs and structures are bound to change and any customizations are unsupported. Please file issues against uncustomized versions.

# Forking and Pull-Requesting
If you want to submit [pull-requests](https://help.github.com/articles/using-pull-requests/) you should first fork `http://github.com/github/orchestrator`.

Setting up the environment is basically the same, except you don't want to

	go get github.com/github/orchestrator/...

But instead clone your own repository.

Assume you fork onto `github.com/you-are-awesome/orchestrator`. _Golang_ has tight coupling between source code import paths and actual URIs. This leads to much confusion. Please consult [Forking Golang repositories on GitHub and managing the import path](http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path) as for ways to solve
that coupling.

Very briefly, you will either want to:

	go get github.com/github/orchestrator/...
	git remote add awesome-fork https://github.com/you-are-awesome/orchestrator.git

Or you will workaround as follows:

	cd $GOPATH
	mkdir -p {src,bin,pkg}
	mkdir -p src/github.com/github/
	cd src/github.com/github/
	git clone git@github.com:you-are-awesome/orchestrator.git # OR: git clone https://github.com/you-are-awesome/orchestrator.git
	cd orchestrator/


You will have a fork of `orchestrator` to which you can push your changes and from which you can send pull requests.
It is best that you first consult (use the [project issues](https://github.com/github/orchestrator/issues)) whether some kind of development would indeed be merged.

You will need to license your code in [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0) or compatible.

Thank you for considering contributions to `orchestrator`!