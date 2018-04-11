[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

> A rising tide lifts all boats.
> -- United States President, John F. Kennedy (borrowed from the New England Council)

An automated tool to provide insight into WordPress code and highlight areas to improve the quality of plugins and themes.

We believe the web can be better. With Tide, the code which underpins every website can be more standardized, faster, and more secure. Tide is focused on WordPress, because no other platform has as large an impact on the state of the web. Tide raises the quality of code one plugin or theme at a time, by elevating the importance of code quality in the developer consciousness. **Because a rising Tide lifts all boats.**

## [Table of Contents](#table-of-contents)
   + [Introduction](#introduction)
   + [Dependencies](#dependencies)
   + [Setup]($setup)
   + [API](#api)
   + [Contributing](#contributing)
   + [Contact Us](#contact-us)
   + [Credits](#credits)
   + [License](#license)

### Introduction
Tide services are responsible for the following:
- The API is responsible for implementing a web server and REST API.
- The Sync Server polls the WordPress.org API's for themes and plugins 
to process and writes them to a queue.
- The PHPCS Server reads messages from a queue and runs reports against 
both plugins and themes, then sends the results back to the Tide API.
- The Lighthouse Server reads messages from a queue and runs Google Lighthouse 
reports against the themes only, then sends the results back to the Tide API.

### Dependencies

* Install [Composer][composer]
* Create a new Cloud Project using the [Cloud Console][cloud-console]
* Enable Billing on that project
* [Enable Cloud SQL API][cloud-sql-api-enable]
* Install [Google Cloud SDK][gcloud-sdk]

### Setup

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

Run the setup script to initialize WordPress for the first time or if you would 
like a convenient way to update the default values when you change certain 
environment variables.

The local database is stored in the `data/api/mysql` directory. If you 
ever need to start from scratch delete that directory and run `make api.setup` 
again. Be sure to stop the API with `make down` and then start it again with 
`make api.up`.

Note: Running `make down` will stop all Docker services.

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
