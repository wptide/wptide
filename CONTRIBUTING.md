# Tide Contributing Guide

Before submitting your contribution, please make sure to take a moment and read 
through the following guidelines.

## Code of Conduct

This project and everyone participating in it is governed by the 
[Tide Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected 
to uphold this code. Please report unacceptable behavior to engage@xwp.co.

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
    - Provide a convincing reason to add this feature. Ideally, you should open a suggestion issue first and have it green-lit before working on it.
- If fixing a bug:
    - Provide a detailed description of the bug in the PR. Live demo preferred.
    - Add appropriate test coverage if applicable.

Travis CI will run the unit tests whenever you push changes to your PR. Tests are required to pass successfully for a merge to be considered.
