# Tide

[![Build Status](https://travis-ci.org/wptide/wptide.svg?branch=develop)](https://travis-ci.org/wptide/wptide) [![Coverage Status](https://coveralls.io/repos/github/wptide/wptide/badge.svg?branch=develop)](https://coveralls.io/github/wptide/wptide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](CONTRIBUTING.md) [![Shipping faster with ZenHub.io](https://img.shields.io/badge/Shipping_faster_with-ZenHub.io-6567bd.svg?style=flat)](https://www.zenhub.com/)

> A rising tide lifts all boats.
> -- United States President, John F. Kennedy (borrowed from the New England Council)

Tide is an automated tool to provide insight into WordPress code and highlight areas to improve the quality of plugins and themes.

We believe the web can be better. With Tide, the code which underpins every website can be more standardized, faster, and more secure. Tide is focused on WordPress, because no other platform has as large an impact on the state of the web. Tide raises the quality of code one plugin or theme at a time, by elevating the importance of code quality in the developer consciousness. **Because a rising Tide lifts all boats.**

## Table of Contents

* [Introduction](#introduction)
* [Dependencies](#dependencies)
* [Cloning](#cloning)
* [Environment Variables](#environment-variables)
* [API](#api)
* [Build Images](#build-images)
* [Start Servers](#start-servers)
* [Running Audits](#running-audits)
* [Contributing](#contributing)
* [Contact Us](#contact-us)
* [Credits](#credits)
* [License](#license)

### Introduction

The main focus of this documentation is to setup a local development environment. If you want to deploy to a cloud environment, please read our [full documentation](https://www.wptide.org/).

### Dependencies

* Install [Composer](https://getcomposer.org/) and test if it works by running `composer --version`.
* Install [Docker](https://docs.docker.com/install/). _Note: you may not be able to setup and run Tide properly with legacy Docker Toolbox._
* Install [Go](https://golang.org/doc/install) and test if your installation works by following the instructions on the installation page.
* Install [Glide](https://glide.readthedocs.io/en/latest/#installing-glide), a package manager for Go. There are a few ways to install Glide:
  - Use the shell script to try an automatically install it. `curl https://glide.sh/get | sh`
  - Download a [versioned release](https://github.com/Masterminds/glide/releases). Glide releases are semantically versioned.
  - Use a system package manager to install Glide. For example, `brew install glide` can be used if you're using [Homebrew](http://brew.sh/) on Mac.
  - The latest development snapshot can be installed with go get. For example, `go get -u github.com/Masterminds/glide`. This is not a release version.

### Cloning

Tide needs to be cloned to a directory inside your Go workspace specified by the [`$GOPATH`](https://golang.org/doc/code#GOPATH) environment variable. Your `$GOPATH` defaults to a directory named `go` inside your home directory, so `$HOME/go` on Mac/Unix and `%USERPROFILE%\go` (usually `C:\Users\YourName\go`) on Windows.

Create the following directory inside your Go workspace:

```
src/github.com/xwp
```

Open a shell and change into the directory:

```
cd $GOPATH/src/github.com/wptide
```

Clone Tide:

```
git clone -b develop --recursive https://github.com/wptide/wptide.git
```

Change to Tide working directory:  

```
cd wptide
```

Update submodules:

```
git submodule update --init --recursive
```

### Environment Variables

Copy the `.env.dist` file to `.env`.

```
cp .env.dist .env
```
_If you are running Tide locally, you do not need to change any of these `.env` values. If you are deploying this into the cloud, make sure to look at the full documentation for which variables you should update and their values._ 

Create an empty `.env.gcp` file.

```
touch .env.gcp
```

---

### API

Install the dependencies as follows:

```
$ make api.composer
```

Then start the API Docker images in isolation:

```
make api.up
```

Open a new shell and run the setup script:

```
make api.setup
```

_You can run the setup script to initialize WordPress for the first time or if you would like a convenient way to update the default values when you change relevant environment variables._

### Build Images

Install Glide dependencies and build images:

```
make build.images
```

### Start Servers

It is recommended that you run these Docker containers separately and start the servers in new shells in order to isolate the output. However, you can start all services (API, Lighthouse Server, PHPCS Server, Sync Server) with the following command:

```
make up
```

#### Isolated Start

Lighthouse Server:

```
make lighthouse.up
```

PHPCS Server:

```
make phpcs.up
```

Sync Server:

```
make sync.up
```

### Hosts File

Add the following entry to your hosts file to make sure `tide.local` is pointed at your local Tide instance:

```
127.0.0.1 tide.local
```
_You can change the `tide.local` URL to some other value by modifying the `API_AUTH_URL` and `API_HTTP_HOST` variables inside the `.env` file._

---

### Running Audits

Now that all the Tide services are up and running, you can run audits locally. If we want to run an audit against the `twentyseventeen` theme for example, we would use this endpoint:

```
http://tide.local/api/tide/v1/audit/wporg/theme/twentyseventeen/2.1
```

Or for the `akismet` plugin:

```
http://tide.local/api/tide/v1/audit/wporg/plugin/akismet/4.1.1
```

When you request an audit you will receive a JSON object back that indicated the audit is pending. If the audit has previously been requested and is complete then you will receive a JSON object with information about the theme/plugin and summary reports with links to the full report.

If the audit is pending, your shell should have some output to indicate that the audit is running. Once this output stops and all your services go back to the `polling` status, you can refresh the API request in the browser and you should see the updated JSON object with completed Tide reports.

For a full list of API Endpoints that can be used with Tide, see the [API Endpoints section](https://www.wptide.org/services/api#endpoints) of the documentation.

### Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Contact Us
Have questions? Don't open an Issue, come join us in the [`#tide` channel](https://wordpress.slack.com/messages/C7TK8FBUJ/) in [WordPress Slack](https://make.wordpress.org/chat/). Even though Slack is a chat service, sometimes it takes several hours for community members to respond — please be patient.

### Credits
Props: [Bartek Makoś (@MakiBM)](https://github.com/MakiBM), 
[Brendan Woods (@brendanwoods-xwp)](https://github.com/brendanwoods-xwp), 
[Daniel Louw (@danlouw)](https://github.com/danlouw), 
[David Cramer (@davidcramer)](https://github.com/davidcramer), 
[David Lonjon (@davidlonjon)](https://github.com/davidlonjon), 
[Derek Herman (@valendesigns)](https://github.com/valendesigns), 
[Jeffrey Paul (@jeffpaul)](https://github.com/jeffpaul), 
[Justin Kopepasah (@kopepasah)](https://github.com/kopepasah), 
[Leo Postovoit (@postphotos)](https://github.com/postphotos), 
[Lubos Kmetko (@luboskmetko)](https://github.com/luboskmetko), 
[Luke Carbis (@lukecarbis)](https://github.com/lukecarbis), 
[Miina Sikk (@miina)](https://github.com/miina), 
[Mike Crantea (@mehigh)](https://github.com/mehigh), 
[Rheinard Korf (@rheinardkorf)](https://github.com/rheinardkorf), 
[Rob Stinson (@robstino)](https://github.com/robstino), 
[Sayed Taqui (@sayedtaqui)](https://github.com/sayedtaqui), 
[Utkarsh Patel (@PatelUtkarsh)](https://github.com/PatelUtkarsh)

### License
Tide utilizes an [MIT license](LICENSE).
