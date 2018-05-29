# Tide

[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](CONTRIBUTING.md) [![Shipping faster with ZenHub.io](https://img.shields.io/badge/Shipping_faster_with-ZenHub.io-6567bd.svg?style=flat)](https://www.zenhub.com/)

> A rising tide lifts all boats.
> -- United States President, John F. Kennedy (borrowed from the New England Council)

Tide is an automated tool to provide insight into WordPress code and highlight areas 
to improve the quality of plugins and themes.

We believe the web can be better. With Tide, the code which underpins every website 
can be more standardized, faster, and more secure. Tide is focused on WordPress, 
because no other platform has as large an impact on the state of the web. Tide raises 
the quality of code one plugin or theme at a time, by elevating the importance of code 
quality in the developer consciousness. **Because a rising Tide lifts all boats.**

## [Table of Contents](#table-of-contents)
   + [Introduction](#introduction)
   + [Dependencies](#dependencies)
   + [Cloning](#cloning)
   + [Setup]($setup)
       - [Environment Variables](#environment-variables)
       - [Google Cloud SDK](#google-cloud-sdk)
       - [Google App Engine](#google-app-engine)
       - [Google Cloud Storage for App Engine](#google-cloud-storage-for-app-engine)
       - [Service Account](#service-account)
   + [API](#api)
       - [API Notes](#api-notes)
       - [API Settings](#api-settings)
   + [Lighthouse Server](#lighthouse-server)
       - [Lighthouse Server Notes](#lighthouse-server-notes)
       - [Lighthouse Server Settings](#lighthouse-server-settings)
       - [Running Lighthouse audits](#running-lighthouse-audits)
   + [PHPCS Server](#phpcs-server)
       - [PHPCS Server Notes](#phpcs-server-notes)
       - [PHPCS Server Settings](#phpcs-server-settings)
   + [Sync Server](#sync-server)
       - [Sync Server Notes](#sync-server-notes)
       - [Sync Server Settings](#sync-server-settings)
   + [Deployments to Google Cloud Platform (GCP)](#deployments-to-google-cloud-platform-gcp)
       - [GCP Settings](#gcp-settings)
       - [Google Cloud SQL (GCSQL)](#google-cloud-sql-gcsql)
           * [GCSQL API Settings](#gcsql-api-settings)
       - [Google App Engine (GAE)](#google-app-engine-gae)
           * [GAE API Settings](#gae-api-settings)
       - [Google Kubernetes Engine (GKE)](#google-kubernetes-engine-gke)
           * [GKE Lighthouse Server Settings](#gke-lighthouse-server-settings)
           * [GKE PHPCS Server Settings](#gke-phpcs-server-settings)
           * [GKE Sync Server Settings](#gke-sync-server-settings)
   + [Google Cloud Storage (GCS)](#google-cloud-storage-gcs)
       - [GCS Settings](#gcs-settings)
   + [Amazon Web Services (AWS)](#amazon-web-services-aws)
       - [AWS Settings](#aws-settings)
   + [Contributing](#contributing)
   + [Contact Us](#contact-us)
   + [Credits](#credits)
   + [License](#license)

### Introduction

The main focus of this documentation is to setup a local development environment. 
Cloud deployments will be covered, though many assumptions will be made about your 
level of understanding not only with Google Cloud Platform (GCP) but also the Tide project and it's manual 
setup process in relation to WordPress on App Engine. We will not be going into 
great detail when it comes to deploying Tide on GCP.

### Dependencies

* Install [Composer][composer]
* Install [Docker][docker]
* Install [Go][go]
* Install [Google Cloud SDK][gcloud-sdk] & [gsutil][gsutil]
* Create a new Cloud Project using the [Cloud Console][cloud-console]
* Enable Billing on that project
* Enable the [Cloud SQL API][cloud-sql-api-enable] (_optional_)

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

Tide is a complicated system of services that we've tried to make easy to build 
and deploy both locally and on GCP. The following directions are by no means 
exhaustive, but should get you there without too much trouble if you follow along 
and do each relevant step correctly and in the following order.

#### Environment Variables

Copy the `.env.dist` file to `.env`.

```
cp .env.dist .env

```

_Update the value for each environment variable in your custom `.env` file. 
The variables and their descriptions can be found at the end of each relevant 
section. You must do this before setting up any of the services._

Additionally, you can create an `.env.gcp` file in the project root. This will make 
deploying services to GCP a bit easier since the `.env.gcp` file will override 
values in the `.env` file. The `.env.gcp` is optional for local development. If 
you see warnings in the console about the file missing when running certain `make` 
commands, that's ok. The `.env.gcp` file will only affect the `.tpl` files and you 
should only add overrides for GCP specific resources. 

The `.tpl` files are template files that through variable interpolation are 
converted and used to deploy your project to GCP and even setup the local API. So 
these files play a critical role in getting Tide setup. If these files are not 
generating the correct output, please [contact us](#contact-us) to troubleshoot and 
figure out a solution for your OS.

So far we've only tested on OS X with and without the `envsubst` command available. 
Other systems may not work correctly and we want to resolve that quickly.

#### Google Cloud SDK 

Configure Google Cloud SDK with your account and the appropriate project ID:

```
$ gcloud init
```

#### Google App Engine

Create an App Engine application within your new project:

```
$ gcloud app create
```

#### Google Cloud Storage for App Engine

Configure the App Engine default GCS bucket for later use. The default App Engine 
bucket is named YOUR_PROJECT_ID.appspot.com. Change the default Access Control 
List (ACL) of that bucket as follows:

```
$ gsutil defacl ch -u AllUsers:R gs://YOUR_PROJECT_ID.appspot.com
```

_**Note**: The previous step is optional if you are setting up Tide for local 
development only._

If you want to upload images to WordPress then you'll need to create a bucket for 
those images to live (unless you opt to use the local file system). Open the 
[Cloud Storage Browser][cloud-storage-browser] and click **Create Bucket**.

Run the following command to change the ACL's of your new bucket:

```
$ gsutil defacl ch -u AllUsers:R gs://YOUR_BUCKET_NAME
```

When the API is up and running, log into the WordPress admin and go to the plugin 
settings page for `GCS Upload` then add your bucket name to the form field and 
click **Save Settings**.

#### Service Account

Finally, go to the [the Credentials section][credentials-section] of your project 
in the Console. Click 'Create credentials' and then click 'Service account key.' 
For the Service account, select 'App Engine app default service account.' Then 
click 'Create' to create and download the JSON service account key to your local 
machine. Save it as `service-account.json` in the the projects root directory for 
use with connecting to both Cloud Storage and Cloud SQL.

---

### API

First generate the API templates:

```
$ make api.tpl
```

_Technically this first step can be skipped since the second command will 
automatically run the first command before installing Composer dependencies._

Next install the dependencies as follows:

```
$ make api.composer
```

Then start the API Docker images in isolation:

```
make api.up
```

Last run the setup script:

```
make api.setup
```

Run the setup script to initialize WordPress for the first time or if you would 
like a convenient way to update the default values when you change relevant 
environment variables.

#### API Notes

The API implements MySQL, PHP-FPM, and an Nginx web server with WordPress installed 
serving a theme and a REST API.

The local database is stored in the `data/api/mysql` directory. If you ever need 
to start from scratch delete that directory and run `make api.setup` again. Be 
sure to stop the API with `make api.down` or `make down` and then start it again 
with `make api.up`.

Running `make down` will stop all Docker services.

If you see an error like this on OS X when bringing up the API you need to add the 
directory to the `Preferences -> File Sharing` section of the Docker for Mac app.

```
ERROR: for gotide_api-mysql_1  Cannot start service api-mysql: b'Mounts denied: ...'
```

For local development you can manually set the `API_KEY` and `API_SECRET` for the 
`audit-server` user, which will automatically update the user meta values when 
`make api.setup` is ran. If you do not set those environment variables, or are 
running Tide in production, then you can access the auto generated key and secret 
from the `audit-server` user profile. 

#### API Settings

| Variable | Description |
| :--- | :--- |
| `API_ADMIN_EMAIL` | The email associated with the local admin account |
| `API_ADMIN_PASSWORD` | The password associated with the local admin account |
| `API_ADMIN_USER` | The username associated with the local admin account |
| `API_AUTH_URL` | The [`wp-tide-api`][wp-tide-api] authentication REST endpoint, used both locally and on GCP. Default is `http://tide.local/api/tide/v1/auth` |
| `API_BLOG_DESCRIPTION` | Site tagline (set in Settings > General). Default is `Automated insight into your WordPress code`. |
| `API_BLOG_NAME` | Site title (set in Settings > General). Default is `Tide`. |
| `API_CACHE` | Whether caching should be active or not. Must be one of: `true`, `false`. Default is `true`. |
| `API_CACHE_DEBUG` | Whether or not to display the caching headers for debugging. Must be one of: `true`, `false`. Default is `false`. |
| `API_CACHE_TTL` | Sets the numeric time to live (TTL) in seconds for page caching. Default is `300`. |
| `API_DB_HOST` | The host of the local database, which connects to a Docker container. Default is `api-mysql`. |
| `API_DB_NAME` | Name of the local database. Default is `wordpress`. |
| `API_DB_PASSWORD` | Password used to access the local database. Default is `wordpress`. |
| `API_DB_ROOT_PASSWORD` | The local database root password. Default is `wordpress`. |
| `API_DB_USER` | Username used to access the local database. Default is `wordpress`. |
| `API_HTTP_HOST` | The API domain name, used both locally and on GCP. Default is `tide.local`. |
| `API_KEY` | The API key used locally to authenticate the `audit-server` user. |
| `API_PROTOCOL` | The API protocol, used both locally and on GCP Default is `http`. |
| `API_REDIS_AUTH` | The Redis database password. Default is `redis`. |
| `API_REDIS_DATABASE` | Use a specific numeric Redis database. Default is `0`. |
| `API_REDIS_HOST` | The host where the Redis database can be reached. Default is `api-redis`. |
| `API_REDIS_PORT` | The port where the Redis database can be reached. Default is `6379`. |
| `API_SECRET` | The API secret used locally to authenticate the `audit-server` user. |
| `API_THEME` | The slug of the local WordPress theme. Default is `twentyseventeen`. |
| `API_UPLOAD_HANDLER` | Tells WordPress how media upload is handled. Uses either the local file system or Google Cloud Storage. Must be one of: `local`, `gcs`. Default is `local`. |
| `API_VERSION` | The API version found in the Tide API REST url, used both locally and on GCP. Default is `v1`. |

---

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

#### Lighthouse Server Notes

The Lighthouse Server is a Go binary that reads messages from a queue and runs Google 
Lighthouse reports against themes, then sends the results back to the Tide API.

#### Lighthouse Server Settings

| Variable | Description |
| :--- | :--- |
| `LH_CONCURRENT_AUDITS` | Sets the number of goroutines the server will perform concurrently. Default is `5` |
| `LH_TEMP_FOLDER` | Sets the temporary folder inside the container used to store downloaded files. Default is `/tmp` |
| `LH_STORAGE_PROVIDER` | Upload reports to the local file system, Google Cloud Storage, or AWS S3. Must be one of: `local`, `gcs`, `s3`. Default is `local`. |
| `LH_MESSAGE_PROVIDER` | Tells the Lighthouse Server which message/queue provider to use; either the AWS SQS or Google Cloud Firestore. Must be one of: `aws`, `firestore`. |

#### Running Lighthouse audits

Details on running Lighthouse audits are [available in the Tide wiki](https://github.com/xwp/go-tide/wiki/Running-Lighthouse-audits).

---

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

#### PHPCS Server Notes

The PHPCS Server is a Go binary that reads messages from a queue and runs PHPCS 
reports against both plugins and themes, then sends the results back to the Tide API.

#### PHPCS Server Settings

| Variable | Description |
| :--- | :--- |
| `PHPCS_CONCURRENT_AUDITS` | Sets the number of goroutines the server will perform concurrently. Default is `5` |
| `PHPCS_TEMP_FOLDER` | Sets the temporary folder inside the container used to store downloaded files. Default is `/tmp` |
| `PHPCS_STORAGE_PROVIDER` | Upload reports to the local file system, Google Cloud Storage, or AWS S3. Must be one of: `local`, `gcs`, `s3`. Default is `local`. |
| `PHPCS_MESSAGE_PROVIDER` | Tells the PHPCS Server which message/queue provider to use; either the AWS SQS or Google Cloud Firestore. Must be one of: `aws`, `firestore`. |

---

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

#### Sync Server Notes

The Sync Server is a Go binary that polls the WordPress.org API's for themes and 
plugins to process and writes them to a queue.

#### Sync Server Settings

| Variable | Description |
| :--- | :--- |
| `SYNC_ACTIVE` | Whether the Sync Server is active or not. Must be one of: `on`, `off`. Default is `off` |
| `SYNC_API_BROWSE_CATEGORY` | The API category used to ingest the wp.org themes and plugins. Must be one of: `popular`, `featured`, `updated`, `new`. Default is `updated` |
| `SYNC_DATA` | When the database provider is set to `local` this will be where the data is stored relative to the `/srv/data` working directory. Default is `./db` |
| `SYNC_DATABASE_DOCUMENT_PATH` |  When the database provider is set to `firestore` this value is the path to the document in Cloud Firestore. Must be in the form of `<collection>/<document>`. Default is `sync-server/wporg`. |
| `SYNC_DATABASE_PROVIDER` | Tells the Sync Server which database provider to use; either the local file system or Google Cloud Firestore. Must be one of: `local`, `firestore`. Default is `local`. |
| `SYNC_MESSAGE_PROVIDER` | Tells the Sync Server which message/queue provider to use; either the AWS SQS or Google Cloud Firestore. Must be one of: `aws`, `firestore`. Default is `aws`. |
| `SYNC_DEFAULT_CLIENT` | The API client used to make requests by the audit servers; also associated with the key and secret those server use. Default is `wporg` |
| `SYNC_DEFAULT_VISIBILITY` | The audit and report visibility. Must be one of: `public`, `private`. Default is `public` |
| `SYNC_FORCE_AUDITS` | Forces audit reports to be generated even if a report exists for the checksum and standard. Must be one of: `yes`, `no`. Default is `no` |
| `SYNC_ITEMS_PER_PAGE` | The number of plugins or themes per page in the API request. Default is `250` |
| `SYNC_LH_ACTIVE` | Send messages to the Lighthouse SQS queue. Must be one of: `on`, `off`. Default is `on` |
| `SYNC_PHPCS_ACTIVE` | Send messages to the PHPCS SQS queue. Must be one of: `on`, `off`. Default is `on` |
| `SYNC_POOL_DELAY` | The wait time in seconds between the wp.org theme and plugin ingests. Default is `600` |
| `SYNC_POOL_WORKERS` | The number of workers (concurrent goroutines) the server will create to ingest the wp.org API. Default is `125` |
| `FIRESTORE_QUEUE_LH` | Specifies which collection in Firestore to use for the Lighthouse message queue. This is a Firestore collection **path**. |
| `FIRESTORE_QUEUE_PHPCS` | Specifies which collection in Firestore to use for the PHPCS message queue. This is a Firestore collection **path**. |

---

### Deployments to Google Cloud Platform (GCP)

Deploying to GCP is optional and not required for local development. In this 
section we've included some of the basic steps required to get setup on GCP.

#### GCP Settings

These three settings are required for both local development and deployments to 
GCP. Be sure to use the same region you chose during the earlier 
`gcloud app create` step.

| Variable | Description |
| :--- | :--- |
| `GCP_PROJECT` | The unique ID of you Google project. |
| `GCP_REGION` | The [region][regions-and-zones] where all your resources will be created. For example, `us-west1`. |
| `GCP_ZONE` | The preferred [zone][regions-and-zones] in your region that resources will be created, For example, `us-west1-a` |

#### Google Cloud SQL (GCSQL)

Deploying a database to Cloud SQL for the WordPress API only requires a bit of 
configuration to the environment variable and then to run a single `make` command.

Create the Cloud SQL database:

```
make api.deploy.sql
```

_This command will create a database instance and failover instance, set the 
root password, create a database for WordPress, and create a user for that 
database._

##### GCSQL API Settings

| Variable | Description |
| :--- | :--- |
| `GCSQL_API_BACKUP_START_TIME` | he start time of daily backups, specified in the 24 hour format - HH:MM, in the UTC timezone. This is the window of time when you would like backups to start. [Learn more](https://cloud.google.com/sql/docs/mysql/instance-settings#backups-and-binary-logging-2ndgen). |
| `GCSQL_API_DATABASE_VERSION` | The database engine type and version. Must be one of: `MYSQL_5_6`, `MYSQL_5_7`. |
| `GCSQL_API_DB_NAME` | Name of the database. |
| `GCSQL_API_DB_PASSWORD` | Password used to access the database. |
| `GCSQL_API_DB_ROOT_PASSWORD` | The database root password. |
| `GCSQL_API_DB_USER` | Username used to access the database. |
| `GCSQL_API_INSTANCE` | Second Generation instance name. Do not include sensitive or personally identifiable information in your instance name; it is externally visible. |
| `GCSQL_API_FAILOVER_REPLICA_NAME` | The name of the failover replica to be created for the new instance. |
| `GCSQL_API_TIER` | The tier for this instance. For Second Generation instances, this is the instance's machine type (e.g., db-n1-standard-1). The machine type determines the number of CPUs and the amount of memory your instance has. [See valid values](https://cloud.google.com/sql/docs/mysql/instance-settings#tier-values). [Learn more](https://cloud.google.com/sql/docs/mysql/instance-settings#machine-type-2ndgen). |
| `GCSQL_API_MAINTENANCE_RELEASE_CHANNEL` | Your preferred timing for instance updates, relative to other instances in the same project. Use `preview` for earlier updates, and `production` for later updates. [Learn more](https://cloud.google.com/sql/docs/mysql/instance-settings#maintenance-timing-2ndgen). |
| `GCSQL_API_MAINTENANCE_WINDOW_DAY` | Day of week for maintenance window, in UTC time zone. Must be one of: `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`. |
| `GCSQL_API_MAINTENANCE_WINDOW_HOUR` | Hour of day - `0` to `23`. Determines a one-hour window when Cloud SQL can perform disruptive maintenance on your instance. |
| `GCSQL_API_STORAGE_SIZE` | Amount of storage allocated to the instance. Must be an integer number of `GB` between `10GB` and `10230GB` inclusive. |

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

Deploy the API to App Engine:

```
make api.deploy.app
```

_You will need to activate the plugins, create the necessary API user accounts, 
and setup permalinks manually._

##### GAE API Settings

| Variable | Description |
| :--- | :--- |
| Auto Scaling: |
| `GAE_API_AS_COOL_DOWN_PERIOD_SEC` | The number of seconds that the autoscaler should wait before it starts collecting information from a new instance. This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable. The cool-down period must be greater than or equal to 60 seconds. An example value is `120`. |
| `GAE_API_AS_CPU_TARGET_UTILIZATION` | Target CPU utilization. CPU use is averaged across all running instances and is used to decide when to reduce or increase the number of instances. Note that instances are downscaled irrespective of in-flight requests 25 seconds after an instance receives the shutdown signal. An example value is `0.5`. |
| `GAE_API_AS_MAX_NUM_INSTANCES` | The maximum number of instances that your service can scale up to. The maximum number of instances in your project is limited by your project's [resource quota](https://cloud.google.com/compute/docs/resource-quotas). |
| `GAE_API_AS_MIN_NUM_INSTANCES` | The minimum number of instances given to your service. When a service is deployed, it is given this many instances and scales according to traffic. Must be `1` or greater, use a minimum of `2` to reduce latency. |
| Cron Schedule: |
| `GAE_API_CRON_SCHEDULE_MINS` | The number of minutes between cron intervals. An example value is `1`. |
| Readiness Check: |
| `GAE_API_RC_APP_START_TIMEOUT_SEC` | The maximum time, in seconds, an instance has to become ready after the VM and other infrastructure are provisioned. After this period, the deployment fails and is rolled back. You might want to increase this setting if your application requires significant initialization tasks, such as downloading a large file, before it is ready to serve. Must be within range of: `1-3600`. |
| `GAE_API_RC_CHECK_INTERVAL_SEC` | Time interval between checks, in seconds. Must be within range of: `1-300`. |
| `GAE_API_RC_FAILURE_THRESHOLD` | An instance is unhealthy after failing this number of consecutive checks. Must be within range of: `1-10`. |
| `GAE_API_RC_SUCCESS_THRESHOLD` | An unhealthy instance becomes healthy after successfully responding to this number of consecutive checks. Must be within range of: `1-10`. |
| `GAE_API_RC_TIMEOUT_SEC` | Timeout interval for each request, in seconds. Must be within range of: `1-300`. |

#### Google Kubernetes Engine (GKE)

All the goroutines are deployed with the same basic steps. Push the image to 
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
| `GKE_LH_CLUSTER` | The name of the cluster. Default is`lighthouse-server`. |
| `GKE_LH_CLUSTER_VERSION` | The Kubernetes version to use for the master and nodes. You can check which Kubernetes versions are default and available in a given zone by running the following command: |
| | `gcloud container get-server-config --zone [COMPUTE-ZONE]` |
| `GKE_LH_CPU_PERCENT` | The average percent CPU utilization across all pods. Must be a range of `1-100`. |
| `GKE_LH_DISK_SIZE` | Size in GB for node VM boot disks. An example value is `100GB`. |
| `GKE_LH_IMAGE` | The name of the Docker image. Default is`lighthouse-server`. |
| `GKE_LH_MACHINE_TYPE` | The type of machine to use for nodes. An example value is `n1-standard-1`. |
| `GKE_LH_MAX_NODES` | Maximum number of nodes to which the node pool can scale. |
| `GKE_LH_MAX_PODS` |  Maximum number of Pods you want to run based on the CPU utilization of your existing Pods. |
| `GKE_LH_MIN_NODES` | Minimum number of nodes to which the node pool can scale. |
| `GKE_LH_MIN_PODS` | Minimum number of Pods you want to run based on the CPU utilization of your existing Pods. |
| `GKE_LH_NUM_NODES` | The number of nodes to be created in each of the cluster's zones. |
| `GKE_LH_REPLICAS` | The number of desired Pods in the initial deployment. |

##### GKE PHPCS Server Settings

| Variable | Description |
| :--- | :--- |
| `GKE_PHPCS_CLUSTER` | The name of the cluster. Default is `phpcs-server`. |
| `GKE_PHPCS_CLUSTER_VERSION` | The Kubernetes version to use for the master and nodes. You can check which Kubernetes versions are default and available in a given zone by running the following command: |
| | `gcloud container get-server-config --zone [COMPUTE-ZONE]` |
| `GKE_PHPCS_CPU_PERCENT` | The average percent CPU utilization across all pods. Must be a range of `1-100`. |
| `GKE_PHPCS_DISK_SIZE` | Size in GB for node VM boot disks. An example value is `100GB`. |
| `GKE_PHPCS_IMAGE` | The name of the Docker image. Default is `phpcs-server`. |
| `GKE_PHPCS_MACHINE_TYPE` | The type of machine to use for nodes. An example value is `n1-standard-1`. |
| `GKE_PHPCS_MAX_NODES` | Maximum number of nodes to which the node pool can scale. |
| `GKE_PHPCS_MAX_PODS` | Maximum number of Pods you want to run based on the CPU utilization of your existing Pods. |
| `GKE_PHPCS_MIN_NODES` | Minimum number of nodes to which the node pool can scale. |
| `GKE_PHPCS_MIN_PODS` | Minimum number of Pods you want to run based on the CPU utilization of your existing Pods. |
| `GKE_PHPCS_NUM_NODES` | The number of nodes to be created in each of the cluster's zones. |
| `GKE_PHPCS_REPLICAS` | The number of desired Pods in the initial deployment. |

##### GKE Sync Server Settings

| Variable | Description |
| :--- | :--- |
| `GKE_SYNC_CLUSTER` | The name of the cluster. Default is `sync-server`. |
| `GKE_SYNC_CLUSTER_VERSION` | The Kubernetes version to use for the master and nodes. You can check which Kubernetes versions are default and available in a given zone by running the following command: |
| | `gcloud container get-server-config --zone [COMPUTE-ZONE]` |
| `GKE_SYNC_DISK_SIZE` | Size in GB for node VM boot disks. An example value is `100GB`. |
| `GKE_SYNC_IMAGE` | The name of the Docker image. Default is `sync-server`. |
| `GKE_SYNC_MACHINE_TYPE` | The type of machine to use for nodes. An example value is `n1-standard-1`. |

---

### Google Cloud Storage (GCS)

If you want to upload Tide audit reports to Google Cloud Storage then you'll need 
to create a bucket for those reports. Open the [Cloud Storage Browser][cloud-storage-browser] 
and click **Create Bucket**.

#### GCS Settings

| Variable | Description |
| :--- | :--- |
| `GCS_BUCKET_NAME` | The name of the GCS bucket. |

---

### Amazon Web Services (AWS)

@todo add docs for AWS.

#### AWS Settings

| Variable | Description |
| :--- | :--- |
| `AWS_API_KEY` | The AWS API key. |
| `AWS_API_SECRET` | The AWS API secret. |
| `AWS_S3_BUCKET_NAME` | The name of the S3 bucket.  |
| `AWS_S3_REGION` | The region of the S3 bucket. Default is `us-west-2`. See a list of available [AWS Regions and Enpoints][s3-region].  |
| `AWS_S3_VERSION` | The S3 API version. Default is `2006-03-01` |
| `AWS_SQS_QUEUE_LH` | The name of the SQS queue for the Lighthouse Server. |
| `AWS_SQS_QUEUE_PHPCS` | The name of the SQS queue for the PHPCS Server. |
| `AWS_SQS_REGION` | The region of the SQS queue. Default is `us-west-2`. See a list of available [AWS Regions and Enpoints][sqs-region].  |
| `AWS_SQS_VERSION` | The SQS API version. Default is `2012-11-05` |

---

### Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, 
and the process for submitting pull requests to us.

### Contact Us
Have questions? Don't open an Issue, come join us in the 
[`#tide` channel][tide-slack] in [WordPress Slack][wp-slack]. Even though Slack is 
a chat service, sometimes it takes several hours for community members to respond 
— please be patient.

### Credits
Props: [Brendan Woods (@brendanwoods-xwp)](https://github.com/brendanwoods-xwp), 
[Daniel Louw (@danlouw)](https://github.com/danlouw), 
[David Cramer (@davidcramer)](https://github.com/davidcramer), 
[David Lonjon (@davidlonjon)](https://github.com/davidlonjon), 
[Derek Herman (@valendesigns)](https://github.com/valendesigns), 
[Jeffrey Paul (@jeffpaul)](https://github.com/jeffpaul), 
[Justin Kopepasah (@kopepasah)](https://github.com/kopepasah), 
[Luke Carbis (@lukecarbis)](https://github.com/lukecarbis), 
[Miina Sikk (@miina)](https://github.com/miina), 
[Mike Crantea (@mehigh)](https://github.com/mehigh), 
[Rheinard Korf (@rheinardkorf)](https://github.com/rheinardkorf), 
[Rob Stinson (@robstino)](https://github.com/robstino), 
[Sayed Taqui (@sayedtaqui)](https://github.com/sayedtaqui), 
[Utkarsh Patel (@PatelUtkarsh)](https://github.com/PatelUtkarsh)

### License
Tide utilizes an [MIT license](LICENSE).

[composer]: https://getcomposer.org/
[docker]: https://docs.docker.com/install/
[go]: https://golang.org/doc/install
[cloud-console]: https://console.cloud.google.com/
[cloud-sql-api-enable]: https://console.cloud.google.com/flows/enableapi?apiid=sqladmin
[gcloud-sdk]: https://cloud.google.com/sdk/
[gsutil]: https://cloud.google.com/storage/docs/gsutil_install
[cloud-storage-browser]: https://console.cloud.google.com/storage/browser
[credentials-section]: https://console.cloud.google.com/apis/credentials/
[wp-tide-api]: https://github.com/wptide/wp-tide-api
[regions-and-zones]: https://cloud.google.com/compute/docs/regions-zones/
[s3-region]: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
[sqs-region]: https://docs.aws.amazon.com/general/latest/gr/rande.html#sqs_region
[tide-slack]: https://wordpress.slack.com/messages/C7TK8FBUJ/
[wp-slack]: https://make.wordpress.org/chat/
