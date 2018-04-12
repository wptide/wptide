# Tide

[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

> A rising tide lifts all boats.
> -- United States President, John F. Kennedy (borrowed from the New England Council)

Tide is an automated tool to provide insight into WordPress code and highlight areas to improve the quality of plugins and themes.

We believe the web can be better. With Tide, the code which underpins every website can be more standardized, faster, and more secure. Tide is focused on WordPress, because no other platform has as large an impact on the state of the web. Tide raises the quality of code one plugin or theme at a time, by elevating the importance of code quality in the developer consciousness. **Because a rising Tide lifts all boats.**

## [Table of Contents](#table-of-contents)
   + [Introduction](#introduction)
   + [Dependencies](#dependencies)
   + [Cloning](#cloning)
   + [Setup]($setup)
   + [API](#api)
   + [Lighthouse Server](#lighthouse-server)
   + [PHPCS Server](#phpcs-server)
   + [Sync Server](#sync-server)
   + [Deployment](#deployment)
       - [Google App Engine](#google-app-engine)
       - [Google Kubernetes Engine](#google-kubernetes-engine)
   + [Contributing](#contributing)
   + [Contact Us](#contact-us)
   + [Credits](#credits)
   + [License](#license)

### Introduction
Tide services are responsible for the following:
- The API implements MySQL, PHP-FPM, and an Nginx web server with WordPress installed serving a theme and a REST API.
- The Sync Server is a Go binary that polls the WordPress.org API's for themes and plugins to process and writes them to a queue.
- The PHPCS Server is a Go binary that reads messages from a queue and runs PHPCS reports against both plugins and themes, then sends the results back to the Tide API.
- The Lighthouse Server is a Go binary that reads messages from a queue and runs Google Lighthouse reports against themes only, then sends the results back to the Tide API.

### Dependencies

* Install [Composer][composer]
* Create a new Cloud Project using the [Cloud Console][cloud-console]
* Enable Billing on that project
* [Enable Cloud SQL API][cloud-sql-api-enable]
* Install [Google Cloud SDK][gcloud-sdk]

### Cloning

Ensure you're in the directory where you would like to install Tide:

```
git clone -b develop --recursive https://github.com/xwp/go-tide.git tide
```

Change to Tide working directory:  

```
cd tide
```

Update submodules:

```
git submodule update --init --recursive
```

### Setup

Copy the `.env.dist` file to `.env`.

```
cp .env.dist .env

```

_Update placeholder values in `.env` to reflect your environment. For example, Tide API specific details; AWS key pairs; SQS queues and S3 buckets._

Configure Google Cloud SDK with your account and the appropriate project ID:

```
$ gcloud init
```

Create an App Engine application within your new project:

```
$ gcloud app create
```

Then configure the App Engine default GCS bucket for later use. The default App Engine bucket is named YOUR_PROJECT_ID.appspot.com. Change the default Access Control List (ACL) of that bucket as follows:

```
$ gsutil defacl ch -u AllUsers:R gs://YOUR_PROJECT_ID.appspot.com
```

Finally, go to the [the Credentials section][credentials-section] of your project in the Console. Click 'Create credentials' and then click 'Service account key.' For the Service account, select 'App Engine app default service account.' Then click 'Create' to create and download the JSON service account key to your local machine. Save it as `service-account.json` in the `service/api` directory for use with connecting to both Cloud Storage and Cloud SQL.

### API

First generate the API templates:

```
$ make api.tpl
```

Then install the dependencies as follows:

```
$ make api.composer
```

Next start the API Docker images in isolation:

```
make api.up
```

Last run the setup script:

```
make api.setup
```

Run the setup script to initialize WordPress for the first time or if you would like a convenient way to update the default values when you change certain environment variables.

The local database is stored in the `data/api/mysql` directory. If you ever need to start from scratch delete that directory and run `make api.setup` again. Be sure to stop the API with `make api.down` or `make down` and then start it again with `make api.up`.

Note: Running `make down` will stop all Docker services.

### Lighthouse Server

First build the Lighthouse Server Docker image:

```
$ make lighthouse.build.image
```

Next start the Lighthouse Server in isolation:

```
$ make lighthouse.up
```

You can combine the previous two steps and simply run:

```
$ make lighthouse.build.up
```

Take the isolated Lighthouse Server down:

```
$ make lighthouse.down
```

### PHPCS Server

First build the PHPCS Server Docker image:

```
$ make phpcs.build.image
```

Next start the PHPCS Server in isolation:

```
$ make phpcs.up
```

You can combine the previous two steps and simply run:

```
$ make phpcs.build.up
```

Take the isolated PHPCS Server down:

```
$ make phpcs.down
```

### Sync Server

First build the Sync Server Docker image:

```
$ make sync.build.image
```

Next start the Sync Server in isolation:

```
$ make sync.up
```

You can combine the previous two steps and simply run:

```
$ make sync.build.up
```

Take the isolated Sync Server down:

```
$ make sync.down
```

### Deployment

To make deploying services easier you can create an `.env.gcp` file in the 
project root that will override values in `.env`. However, this only effects 
the `.tpl` files that get converted, which are used to deploy to your project 
on GCP.

#### Google App Engine

Deploying the API to App Engine is fairly straight forward. You'll need to 
provision a database and then deploy the app. With one caveat. The first time 
you deploy the API you **must** remove the `readiness_check` section from your 
`service/api/app.yaml` or your app will never become healthy. The section looks 
like this:

```yaml
readiness_check:
  path: "/health-check.php"
  timeout_sec: 4
  check_interval_sec: 5
  failure_threshold: 2
  success_threshold: 2
  app_start_timeout_sec: 60
```

_Once you've deployed and installed WordPress add the `readiness_check` 
section back to `app.yaml` and re-deploy. The health check should be working 
and ensuring the health of your containers on App Engine._

First create the Cloud SQL database:

```
make api.deploy.sql
```

_This command will create a database instance and failover instance, set the 
root password, create a database for WordPress, and create a user for that 
database._

Then deploy the API:

```
make api.deploy.app
```

_You will need to activate the plugins, create the necessary API user accounts, 
and setup permalinks manually._

#### Google Kubernetes Engine

All the Go routines are deployed with the same basic steps. Push the image to 
Google Container Registry (GCR), create the Kubernetes cluster, and then create 
a Kubernetes deployment.

I'll demonstrate with the PHPCS Server. Other than specific `make` commands for 
each service, these step are all the same.

Push the image to GCR:

```
make phpcs.push.image
```

Create the GKE cluster:

```
make phpcs.build.cluster
```

Deploy the GKE cluster:

```
make phpcs.deploy.cluster
```

Delete the GKE cluster:

```
make phpcs.clean.cluster
```

### Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Contact Us
Have questions? Don't open an Issue, come join us in the [`#tide` channel](https://wordpress.slack.com/messages/C7TK8FBUJ/) in [WordPress Slack](https://make.wordpress.org/chat/). Even though Slack is a chat service, sometimes it takes several hours for community members to respond — please be patient.

### Credits
Props: [@danlouw](https://github.com/danlouw), [@jeffpaul](https://github.com/jeffpaul), [@rheinardkorf](https://github.com/rheinardkorf), [@valendesigns](https://github.com/valendesigns)

### License
Tide utilizes an [MIT license](https://github.com/xwp/go-tide/blob/master/LICENSE).

[composer]: https://getcomposer.org/
[cloud-console]: https://console.cloud.google.com/
[cloud-sql-api-enable]: https://console.cloud.google.com/flows/enableapi?apiid=sqladmin
[gcloud-sdk]: https://cloud.google.com/sdk/
[credentials-section]: https://console.cloud.google.com/apis/credentials/
