# Go Application with New Relic Monitoring

This repository contains a Go application integrated with New Relic for performance monitoring and observability.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Viewing New Relic Data](#viewing-new-relic-data)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.20 or later installed on your machine
- A New Relic account
- A New Relic license key

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/balasl342/go-newrelic-example.git
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Install the New Relic Go agent**:

    ```bash
    go get github.com/go-newrelic-example/go-agent/v3/newrelic
    ```

## Configuration

1. **Set up New Relic environment variables**:

    Update the config file in the {root}/config directory of your project and add your New Relic license key:

    ```ini
    LicenseKey=your_new_relic_license_key
    UserKey=your_relic_user_key
    AppName=YourAppName

## Running the Application

- To run the application, use the following command:

    ```bash
    go run main.go

## Viewing New Relic Data

- To view the data collected by New Relic:

    1. Log in to your New Relic account.
    2. Navigate to the APM section.
    3. Select your application from the list.
    4. Explore the various dashboards and metrics to monitor your application's performance.
