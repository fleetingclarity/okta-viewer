# Okta user and group viewer

This CLI tool allows development and support staff to view user details and group memberships from the command line.

## Prerequisites
* Okta API Token with read access to the users and groups APIs
* An Okta domain

## Setup for local development
1. clone the repository 
    ```bash
    git clone https://github.com/fleetingclarity/okta-viewer.git
    ```
2. Configure environment variables
    ```bash
    export OV_OKTA_ORG_URL="https://yourtenant.okta.com"
    export OV_OKTA_API_TOKEN="your-api-token"
    ```
3. Install dependencies
    ```bash
    go mod tidy
    ```