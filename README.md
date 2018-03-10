# GO Tide

[![Build Status](https://travis-ci.org/xwp/go-tide.svg?branch=develop)](https://travis-ci.org/xwp/go-tide) [![Coverage Status](https://coveralls.io/repos/github/xwp/go-tide/badge.svg?branch=develop)](https://coveralls.io/github/xwp/go-tide?branch=develop) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Tide services responsible for polling the WordPress.org APIs for themes and plugins to audit ("Sync Server"), for running audits and sending the results to the Tide API service ("Audit Server"), and for running Google Lighthouse audits against themes in the WordPress.org repository ("Lighthouse Server").

## Table of Contents
@todo

### Setup
@todo

### Usage

#### WPCS Auditing
@todo

#### PHP Compatibility Auditing
@todo

#### Google Lighthouse Auditing
This section outlines the components included in the Lighthouse integration with Tide, demonstrates how the integration processes Lighthouse audits of WordPress.org themes, and provides examples of the Tide API with Lighthouse audit results.

Status:
Lighthouse auditing of WordPress themes has been integrated by running each theme from wp-themes.com through the Lighthouse CLI, stores the full report, and includes the summary results and a link to the full report in the Tide API. This currently functions locally, XWP is working to get this running on GCP.

Components:
The following outlines the components added to Tide in order to integrate Lighthouse in overall Tide auditing of WordPress themes.
@todo

Process:
The following demonstrates how a WordPress theme is run through a Lighthouse audit and has its results stored and returned via the Tide API.
@todo

Lighthouse Results in Tide API:
The following are example responses from the Tide API showing a summary of a Lighthouse audit and a detailed result of a Lighthouse audit.
@todo

### Contributing
@todo

### Credits
Props: [@rheinardkorf](https://github.com/rheinardkorf), [@valendesigns](https://github.com/valendesigns)

### License
Tide utilizes an [MIT license](https://github.com/xwp/go-tide/blob/master/LICENSE).
