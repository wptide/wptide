# GO Tide

[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Common Prerequisites

* Install [Composer][composer]
* Create a new Cloud Project using the [Cloud Console][cloud-console]
* Enable Billing on that project
* [Enable Cloud SQL API][cloud-sql-api-enable]
* Install [Google Cloud SDK][gcloud-sdk]

## Project preparation

Copy the `.env.dist` file to `.env`.

```
cp .env.dist .env

```

Update placeholder values in `.env` to reflect your environment. For example, 
Tide API specific details; AWS key pairs; SQS queues and S3 buckets.

Configure Google Cloud SDK with your account and the appropriate project ID:

```
$ gcloud init
```

Create an App Engine application within your new project:

```
$ gcloud app create
```

Then configure the App Engine default GCS bucket for later use. The default App
Engine bucket is named YOUR_PROJECT_ID.appspot.com. Change the default Access
Control List (ACL) of that bucket as follows:

```
$ gsutil defacl ch -u AllUsers:R gs://YOUR_PROJECT_ID.appspot.com
```

Go to the [the Credentials section][credentials-section] of your project in the
Console. Click 'Create credentials' and then click 'Service account key.' For
the Service account, select 'App Engine app default service account.' Then
click 'Create' to create and download the JSON service account key to your
local machine. Save it as `service-account.json` in the `service/api` directory 
for use with connecting with both Cloud Storage and Cloud SQL.

## API Setup

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

Run the setup script to initialize WordPress for the first time or if you would 
like a convenient way to update the default values when you change certain 
environment variables.

The local database is stored in the `data/api/mysql` directory. If you 
ever need to start from scratch delete that directory and run `make api.setup` 
again. Be sure to stop the API with `make down` and then start it again with 
`make api.up`.

Note: Running `make down` will stop all Docker services.

## Props

[@rheinardkorf](https://github.com/rheinardkorf), [@valendesigns](https://github.com/valendesigns)

[composer]: https://getcomposer.org/
[cloud-console]: https://console.cloud.google.com/
[cloud-sql-api-enable]: https://console.cloud.google.com/flows/enableapi?apiid=sqladmin
[gcloud-sdk]: https://cloud.google.com/sdk/
[credentials-section]: https://console.cloud.google.com/apis/credentials/