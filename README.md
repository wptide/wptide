# ![Tide](docs/images/logo.png)

[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
> A rising tide lifts all boats.
> -- United States President, John F. Kennedy (borrowed from the New England Council)

An automated tool to provide insight into WordPress code and highlight areas to improve the quality of plugins and themes.

We believe the web can be better. With Tide, the code which underpins every website can be more standardized, faster, and more secure. Tide is focused on WordPress, because no other platform has as large an impact on the state of the web. Tide raises the quality of code one plugin or theme at a time, by elevating the importance of code quality in the developer consciousness. **Because a rising Tide lifts all boats.**

## [Table of Contents](#table-of-contents)
   + [Introduction](#introduction)
   + [Setup](#setup)
   + [Usage](#usage)
   + [PHP_CodeSniffer (PHPCS) Server](#php-codesniffer--phpcs--server)
     - [WordPress Coding Standards (WPCS)](#wordpress-coding-standards--wpcs-)
     - [PHP Compatibility](#php-compatibility)
   + [Lighthouse Server](#lighthouse-server)
     - [Status](#status)
     - [Components](#components)
     - [Process](#process)
     - [Lighthouse Results in Tide API](#lighthouse-results-in-tide-api)
       * [Checksum Endpoint (Summary Only)](#checksum-endpoint--summary-only-)
       * [Checksum Endpoint (Details Included)](#checksum-endpoint--details-included-)
   + [Contributing](#contributing)
   + [Contact Us](#contact-us)
   + [Credits](#credits)
   + [License](#license)

### Introduction
Tide services are responsible for the following:
- The Sync Server polls the WordPress.org API's for themes and plugins to process and writes them to a queue.
- The PHPCS Server reads messages from a queue and runs reports against both plugins and themes, then sends the results back to the Tide API.
- The Lighthouse Server reads messages from a queue and runs Google Lighthouse reports against the themes only, then sends the results back to the Tide API.

### Setup
@todo

### Usage

### PHP_CodeSniffer (PHPCS) Server
@todo

#### WordPress Coding Standards (WPCS)
@todo

#### PHP Compatibility
@todo

### Lighthouse Server
This section outlines the components included in the Lighthouse integration with Tide, demonstrates how the integration processes Lighthouse audits of WordPress.org themes, and provides examples of the Tide API with Lighthouse audit results.

#### Status
Lighthouse auditing of WordPress themes has been integrated by running each theme from wp-themes.com through the Lighthouse CLI, stores the full report, and includes the summary results and a link to the full report in the Tide API. This currently functions locally, XWP is working to get this running on GCP.

#### Components
The following outlines the components added to Tide in order to integrate Lighthouse in overall Tide auditing of WordPress themes.
1. Docker Container
   - Lighthouse CLI
   - Lighthouse Server binary (lighthouse-server)
   **Note:** Uses an Alpine Image with working Chromium version. Produces consistent results with Lighthouse Chrome extension. Allows reduced image size of 432MB.
2. Go Lighthouse Server Source Code
   - `cmd/lighthouse-server` for binary
   - `src` for packages
   - `vendor` for imported packages
     - [github.com/nanobox-io/golang-scribble](https://github.com/nanobox-io/golang-scribble)
     - [github.com/wptide/pkg](https://github.com/wptide/pkg)
   - Running Tide cluster

#### Process
The following demonstrates how a WordPress theme is run through a Lighthouse audit and has its results stored and returned via the Tide API.
1. `tide-cluster`
   - Starts all Tide services and listens...

   ![](docs/images/starts-all-tide-services-and-listens.png)
2. `lighthouse-server`
   - authenticate with Tide API    

   ![](docs/images/authenticate-with-tide-api.png)
   - polls a job queue (SQS) for messages to process

   ![](docs/images/polls-a-job-queue-for-messages-to-process.png)
   - downloads theme and calculates checksum

   ![](docs/images/downloads-theme-and-calculates-checksum.png)
   - runs source through `gocloc` to get code information
   - scans source for Theme header information (required for Lighthouse Report)

   ![](docs/images/scans-source-for-theme-header-info.png)
   - runs theme through `lighthouse` at `https://wp-themes.com/<theme-slug>` and keeps polling for next job.

   ![](docs/images/runs-theme-through-lighthouse-and-keeps-polling-for-next-job.png)
   - take full report and generate detailed report
   - saves full report to a file store (S3)

   ![](docs/images/saves-full-report-to-a-file-store.png)
   - grabs subset of results `reportCategories` with only:
     - category name
     - score
     - Description

   ![](docs/images/grabs-subset-of-results.png)
   - bundles summary result and reference to full report as a message payload
   - payload sent to Tide API instance

   ![](docs/images/payload-sent-to-tide-api-instance.png)

#### Lighthouse Results in Tide API
The following are example responses from the Tide API showing a summary of a Lighthouse report and a detailed result of a Lighthouse report.

##### Checksum Endpoint (Summary Only)
https://tide.local/api/tide/v1/audit/ab38727534cbeeef043faf1e54a786e22e4e7c6a172a5ccccf23fe6b2f3d28bd?standards=lighthouse

![](docs/images/checksum-endpoint-summary-only.png)

##### Checksum Endpoint (Details Included)
https://tide.local/api/tide/v1/audit/ab38727534cbeeef043faf1e54a786e22e4e7c6a172a5ccccf23fe6b2f3d28bd?standards=lighthouse&details=all

![](docs/images/checksum-endpoint-details-included.png)

**Note:** Results for details is exactly the same as the output from a Lighthouse CI report.

### Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Contact Us
Have questions? Don't open an Issue, come join us in the [`#tide` channel](https://wordpress.slack.com/messages/C7TK8FBUJ/) in [WordPress Slack](https://make.wordpress.org/chat/). Even though Slack is a chat service, sometimes it takes several hours for community members to respond — please be patient.

### Credits
Props: [@rheinardkorf](https://github.com/rheinardkorf), [@valendesigns](https://github.com/valendesigns), [@danlouw](https://github.com/danlouw), [@jeffpaul](https://github.com/jeffpaul)

### License
Tide utilizes an [MIT license](https://github.com/xwp/go-tide/blob/master/LICENSE).
