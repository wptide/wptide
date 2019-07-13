# Tide Contributing Guide

There are many ways to contribute to Tide. You can help us champion the adoption of quality testing results in the WordPress project. You can also help by contributing code or documentation to Tide itself.

Note that you can also contribute to [WordPress meta](https://make.wordpress.org/meta/handbook/documentation/contributing-with-git/) to improve the WordPress plugin directory.

Before submitting your contribution, please make sure to take a moment and read through the following guidelines.

## Code of Conduct

This project and everyone participating in it is governed by the [Tide Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to derek@wptide.org.

## Issue Reporting Guidelines

- The Issue list of this repo is **exclusively** for bug reports and feature requests.
- Try to search for your issue, it may have already been answered or even fixed in the `develop` branch.
- Check if the issue is reproducible with the latest stable version. If you are using a pre-release, please indicate the specific version you are using.
- It is **required** that you clearly describe the steps necessary to reproduce the issue you are running into. Issues without clear reproducible steps will be closed immediately.
- If your issue is resolved but still open, don't hesitate to close it. In case you found a solution by yourself, it could be helpful to explain how you fixed it.

## Pull Request Guidelines

- Checkout a topic branch from `develop` and merge back against `develop`.
    - If you are not familiar with branching please read [_A successful Git branching model_](http://nvie.com/posts/a-successful-git-branching-model/) before you go any further.
- If adding a new feature:
    - Add accompanying test case.
    - Provide convincing reason to add this feature. Ideally you should open a suggestion issue first and have it green-lit before working on it.
- If fixing a bug:
    - Provide detailed description of the bug in the PR. Live demo preferred.
    - Add appropriate test coverage if applicable.

Travis CI will run the unit tests whenever you push changes to your PR. Tests are required to pass successfully for a merge to be considered.

## Profile Badges

As outlined in [the announcement post](https://make.wordpress.org/tide/2019/06/20/tide-profile-badges/), badges related to work on Tide are awarded as follows:

### Tide Team

![](service/api/wp-content/themes/docs/docs/images/Tide-Team.png)

The Tide Team badge will be manually assigned to all active Tide maintainers – i.e those who are listed as “Maintainers” on [this page](https://github.com/wptide/wptide#maintainers) (also [here](https://github.com/wptide/docs#maintainers)).

### Tide Contributor

![](service/api/wp-content/themes/docs/docs/images/Tide-Contributor.png)

The Tide Contributor badge will be manually assigned to those who provide valuable contributions to Tide -- i.e. those who are listed as “Contributors” on [this page](https://github.com/wptide/wptide#contributors) (also [here](https://github.com/wptide/docs#credits), [here](https://github.com/wptide/pkg#props), and [here](https://github.com/wptide/wp-tide-api#props)).

*The easiest way to have the Tide Team or Tide Contributor badges assigned to you is for you to request them (the system doesn’t allow us to add the badge to your profile until you submit a request). To make this request please go the [Tide Team](https://profiles.wordpress.org/associations/tide-team/) or [Tide Contributor](https://profiles.wordpress.org/associations/tide-contributor/) pages and request membership for the group.*

