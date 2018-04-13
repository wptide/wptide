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
       - [API Settings](#api-settings)
   + [Lighthouse Server](#lighthouse-server)
       - [Lighthouse Server Settings](#lighthouse-server-settings)
   + [PHPCS Server](#phpcs-server)
       - [PHPCS Server Settings](#phpcs-server-settings)
   + [Sync Server](#sync-server)
       - [Sync Server Settings](#sync-server-settings)
   + [Deployments to Google Cloud Platform (GCP)](#deployments-to-google-cloud-platform-(gcp))
       - [GCP Settings](#gcp-settings)
       - [Google Cloud SQL (GCSQL)](#google-cloud-sql-(gcsql))
           * [GCSQL API Settings](#gcsql-api-settings)
       - [Google App Engine (GAE)](#google-app-engine-(gae))
           * [GAE API Settings](#gae-api-settings)
       - [Google Kubernetes Engine (GKE)](#google-kubernetes-engine-(gke))
           * [GKE Lighthouse Server Settings](#gke-lighthouse-server-settings)
           * [GKE PHPCS Server Settings](#gke-phpcs-server-settings)
           * [GKE Sync Server Settings](#gke-sync-server-settings)
   + [AWS](#aws)
       - [AWS Settings](#aws-settings)
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

_Update the value for each [environment variable](#environment-variables) in your custom `.env` file._

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

#### API Settings

| Variable | Description |
| :--- | :--- |
| `API_ADMIN_EMAIL` | @todo |
| `API_ADMIN_PASSWORD` | @todo |
| `API_ADMIN_USER` | @todo |
| `API_AUTH_URL` | @todo Default is `http://tide.local/api/tide/v1/auth` |
| `API_BLOG_DESCRIPTION` | @todo Default is `Automated insight into your WordPress code`. |
| `API_BLOG_NAME` | @todo Default is `Tide`. |
| `API_DB_HOST` | @todo Default is `api-mysql`. |
| `API_DB_NAME` | @todo Default is `wordpress`. |
| `API_DB_PASSWORD` | @todo Default is `wordpress`. |
| `API_DB_ROOT_PASSWORD` | @todo Default is `wordpress`. |
| `API_DB_USER` | @todo Default is `wordpress`. |
| `API_HTTP_HOST` | @todo Default is `tide.local`. |
| `API_KEY` | @todo |
| `API_PROTOCOL` | @todo Default is `http`. |
| `API_SECRET` | @todo |
| `API_THEME` | @todo Default is `twentyseventeen`. |
| `API_VERSION` | @todo Default is `v1`. |

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

#### Lighthouse Server Settings

| Variable | Description |
| :--- | :--- |
| `LH_CONCURRENT_AUDITS` | @todo Default is `5` |

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

#### PHPCS Server Settings

| Variable | Description |
| :--- | :--- |
| `PHPCS_CONCURRENT_AUDITS` | @todo Default is `5` |
| `PHPCS_TEMP_FOLDER` | @todo Default is `/tmp` |

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

#### Sync Server Settings

| Variable | Description |
| :--- | :--- |
| `SYNC_ACTIVE` | @todo Default is `off` |
| `SYNC_API_BROWSE_CATEGORY` | @todo Default is `updated` |
| `SYNC_DATA` | @todo Default is `/srv/data` |
| `SYNC_DEFAULT_CLIENT` | @todo Default is `wporg` |
| `SYNC_DEFAULT_VISIBILITY` | @todo Default is `public` |
| `SYNC_FORCE_AUDITS` | @todo Default is `no` |
| `SYNC_ITEMS_PER_PAGE` | @todo Default is `250` |
| `SYNC_LH_ACTIVE` | @todo Default is `off` |
| `SYNC_PHPCS_ACTIVE` | @todo Default is `on` |
| `SYNC_POOL_DELAY` | @todo Default is `600` |
| `SYNC_POOL_WORKERS` | @todo Default is `125` |

### Deployments to Google Cloud Platform (GCP)

To make deploying services easier you can create an `.env.gcp` file in the 
project root that will override values in `.env`. However, this only effects 
the `.tpl` files that get converted, which are used to deploy to your project 
on GCP.

#### GCP Settings

| Variable | Description |
| :--- | :--- |
| `GCP_PROJECT` | @todo |
| `GCP_REGION` | @todo |
| `GCP_ZONE` | @todo |

#### Google Cloud SQL (GCSQL)

@todo

##### GCSQL API Settings

| Variable | Description |
| :--- | :--- |
| `GCSQL_API_BACKUP_START_TIME` | @todo |
| `GCSQL_API_DATABASE_VERSION` | @todo |
| `GCSQL_API_DB_NAME` | @todo |
| `GCSQL_API_DB_PASSWORD` | @todo |
| `GCSQL_API_DB_ROOT_PASSWORD` | @todo |
| `GCSQL_API_DB_USER` | @todo |
| `GCSQL_API_INSTANCE` | @todo |
| `GCSQL_API_FAILOVER_REPLICA_NAME` | @todo |
| `GCSQL_API_TIER` | @todo |
| `GCSQL_API_MAINTENANCE_RELEASE_CHANNEL` | @todo |
| `GCSQL_API_MAINTENANCE_WINDOW_DAY` | @todo |
| `GCSQL_API_MAINTENANCE_WINDOW_HOUR` | @todo |
| `GCSQL_API_STORAGE_SIZE` | @todo |

#### Google App Engine (GAE)

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

##### GAE API Settings

| Variable | Description |
| :--- | :--- |
| `GAE_API_AS_COOL_DOWN_PERIOD_SEC` | @todo |
| `GAE_API_AS_CPU_TARGET_UTILIZATION` | @todo |
| `GAE_API_AS_MAX_NUM_INSTANCES` | @todo |
| `GAE_API_AS_MIN_NUM_INSTANCES` | @todo |
| `GAE_API_CRON_SCHEDULE_MINS` | @todo |
| `GAE_API_RC_APP_START_TIMEOUT_SEC` | @todo |
| `GAE_API_RC_CHECK_INTERVAL_SEC` | @todo |
| `GAE_API_RC_FAILURE_THRESHOLD` | @todo |
| `GAE_API_RC_SUCCESS_THRESHOLD` | @todo |
| `GAE_API_RC_TIMEOUT_SEC` | @todo |

#### Google Kubernetes Engine (GKE)

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

##### GKE Lighthouse Server Settings

| Variable | Description |
| :--- | :--- |
| `GKE_LH_CLUSTER` | @todo |
| `GKE_LH_CLUSTER_VERSION` | @todo |
| `GKE_LH_CPU_PERCENT` | @todo |
| `GKE_LH_DISK_SIZE` | @todo |
| `GKE_LH_IMAGE` | Default is`lighthouse-server`. |
| `GKE_LH_MACHINE_TYPE` | @todo |
| `GKE_LH_MAX_NODES` | @todo |
| `GKE_LH_MAX_PODS` | @todo |
| `GKE_LH_MIN_NODES` | @todo |
| `GKE_LH_MIN_PODS` | @todo |
| `GKE_LH_NUM_NODES` | @todo |
| `GKE_LH_REPLICAS` | @todo |

##### GKE PHPCS Server Settings

| Variable | Description |
| :--- | :--- |
| `GKE_PHPCS_CLUSTER` | @todo |
| `GKE_PHPCS_CLUSTER_VERSION` | @todo |
| `GKE_PHPCS_CPU_PERCENT` | @todo |
| `GKE_PHPCS_DISK_SIZE` | @todo |
| `GKE_PHPCS_IMAGE` | Default is `phpcs-server`. |
| `GKE_PHPCS_MACHINE_TYPE` | @todo |
| `GKE_PHPCS_MAX_NODES` | @todo |
| `GKE_PHPCS_MAX_PODS` | @todo |
| `GKE_PHPCS_MIN_NODES` | @todo |
| `GKE_PHPCS_MIN_PODS` | @todo |
| `GKE_PHPCS_NUM_NODES` | @todo |
| `GKE_PHPCS_REPLICAS` | @todo |

##### GKE Sync Server Settings

| Variable | Description |
| :--- | :--- |
| `GKE_SYNC_CLUSTER` | @todo |
| `GKE_SYNC_CLUSTER_VERSION` | @todo |
| `GKE_SYNC_DISK_SIZE` | @todo |
| `GKE_SYNC_IMAGE` | Default is `sync-server`. |
| `GKE_SYNC_MACHINE_TYPE` | @todo |

### AWS

@todo

#### AWS Settings

| Variable | Description |
| :--- | :--- |
| `AWS_API_KEY` | The AWS API key. |
| `AWS_API_SECRET` | The AWS API secret. |
| `AWS_S3_BUCKET_NAME` | The name of the S3 bucket.  |
| `AWS_S3_REGION` | The region of the S3 bucket. Default is `us-west-2`. See a list of available [AWS Regions and Enpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region).  |
| `AWS_S3_VERSION` | The S3 API version. Default is `2006-03-01` |
| `AWS_SQS_QUEUE_LH` | The name of the SQS queue for the Lighthouse Server. |
| `AWS_SQS_QUEUE_PHPCS` | The name of the SQS queue for the PHPCS Server. |
| `AWS_SQS_REGION` | The region of the SQS queue. Default is `us-west-2`. See a list of available [AWS Regions and Enpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html#sqs_region).  |
| `AWS_SQS_VERSION` | The SQS API version. Default is `2012-11-05` |

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
