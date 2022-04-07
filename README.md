# Assignment 2

Project for the second assignment in PROG2005-Cloud Technologies 2022. This rest api is created for getting information
about covid policies and cases in different countries around the world. With this api you can se any past or present
covid policies and the stringency of the policies in a country at a given time. With the cases endpoint the most recent
data regarding confirmed cases, deaths, growth rate and some other information.

This api relies on the covidtracker api from the university of Oxford: https://covidtracker.bsg.ox.ac.uk for the policy
information. For the covid cases this api relies on https://covid19-graphql.now.sh from rlindskog on GitHub. It uses
data from the Johns Hopkins University.

The api also supports webhooks. Webhooks can be registered with custom triggering parameters. When a webhook is
triggered information about how many times a given country has been search for is provided.

[TOC]

# Endpoints

The api has four endpoint:

    /corona/v1/notifications
    /corona/v1/policy
    /corona/v1/cases
    corona/v1/status

If no endpoint is found the error code 404 not found is returned. This indicates that the user entered a wrong path.

## Policy

This endpoint can be used to get information about policies in a given country at a given date. If no date is provided,
today's date is used.

### Request

The policy endpoint can be used with both a scope and country parameter or just the country parameter.

#### Parameters

`country` is the country that the stringency and policy information is returned for. It must be the alpha three code for
the country and not the full name.  
`scope` is the date that you want to see the stringency and policy information for. The date must be provided on the
form yyyy-mm-dd.

Example request:

    /corona/v1/policy?country=nor
    /corona/v1/policy?country=nor&scope=2021-10-20

### Response

A response will have the content type `application/json`.

Status codes:

* 200: Everything is ok.
* 400: Client side error, wrong parameter/other.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

Example body:

    {
        "country_code": "nor",
        "date_value": "2022-01-02",
        "stringency": 51.85,
        "policies": 20
    }

## Cases

This endpoint returns information about the most recent covid cases in a given country. The data includes number of
cases, number of deaths and growth rate.

### Request

A request to this endpoint must be a get request and contain the country parameter.

#### Parameters

`country` is the country that to retrieve the cases information for. It can be a full name or an alpha three code. Note
that not all alpha three codes will work correctly, please see the known bugs section.

Example request:

    /corona/v1/cases?country=norway
    /corona/v1/cases?country=nor

### Response

A response will have the content type `application/json`.

Status codes:

* 200: Everything is ok.
* 400: Client side error, wrong parameter/other.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

An invalid country parameter will yield an empty json structure.

Example body:

    {
        "country": "Norway",
        "date": "2022-03-28",
        "confirmed": 1399714,
        "recovered": 0,
        "deaths": 2339,
        "growth_rate": 0.0014853631627073677
    }

## Notifications

The notification endpoint is for retrieving, adding and fetching webhooks. Webhooks can be registered and set to be
triggered when a specified country has been called x number of times in any of the endpoints.

### Request

The endpoint supports the three rest methods of: get, post and delete.

#### Get request

When the request is of method get, it can be used with and without a parameter. When no parameter is used, all
registered webhooks will be in the response. If the id parameter is used with a webhook id (64 characters) of a
registered webhook, information of that webhook is in the response.

Example request:

    /corona/v1/notifications?id=4b10588a43e6b4658097114e5e8a9d0fd6ffb576e43904aace08b8f3a43f2cc1
    /corona/v1/notifications

#### Post request

When registering a new webhook the rest method post is used. The body of a valid post request must have a json body
following the following pattern:

    {
        "url": "https://cool.webhook/safag3wwgwefsoij9",
        "country": "Norway",
        "calls": 10
    }

The url field should contain the url that is to be triggered when the webhook is triggered. The country field should
contain the full name of a country. The country must be in either the cases api or the countries api, please see the
know bugs section for more information. The calls field should contain the number of times the country should be called
before the webhook is triggered.

The body should be sent to the following url:

    /corona/v1/notification

#### Delete request

For deleting a webhook the rest method delete should be used followed by the id parameter containing the id of the
webhook to be deleted.

Example request:

    /corona/v1/notifications?id=4b10588a43e6b4658097114e5e8a9d0fd6ffb576e43904aace08b8f3a43f2cc1

#### Parameters

`id` the webhook id to delete or get information about.

### Response

A response will have the content type `application/json`.

Status codes:

* 200: Everything is ok.
* 400: Client side error, wrong parameter/other.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

#### Get response

##### One webhook

When using the id parameter to get information about a given webhook the following body is returned:

    {
        "url": "https://cool.webhook/safag3wwgwefsoij9",
        "country": "Norway",
        "calls": 10
    }

It contains the same information as the one being posted at registration.

##### All webhooks

When fetching all webhooks, they will be returned as a list. All registered webhooks, with their ids will be in the
list. Example response:

    [
    {
        "id": "21222154417c758332ab2764adccf61d809bcab49f92ddde8796debd2c18b20b",
        "url": "https://cool.webhook/safag3wwgwefsoij9",
        "country": "Norway",
        "calls": 10
    },
    {
        "id": "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470",
        "url": "https://webhook.site/bf1c45f9-7ca4-4867-a17b-d7a10a49cea6",
        "country": "Sweden",
        "calls": 2
    }
    ]

#### Post response

When sending a post request to register a new webhook the following body with a webhook id will be returned:

    {
        "webhook_id": "21222154417c758332ab2764adccf61d809bcab49f92ddde8796debd2c18b20b"
    }

This is the only time to view the webhook id, so save it if needed.

#### Delete response

When doing a delete request there will be no response, but the status code 200 will be returned. To verify that it has
been deleted it is possible to send a get request with the id, the body should be empty.

### Status

The status endpoint provides information about the service and the availability of the backend apis the service is
reliant on.

#### Request

Only a get request is possible.

Example request:

    /corona/v1/status

#### Response

The response will contain the status code of the backend apis, the number of currently registered webhooks, the version
of the services and the uptime of the service. An example body is:

    {
      "cases_api_status": 200,
      "policy_api_status": 200,
      "country_api_status": 200,
      "number_of_webhooks": 1,
      "version": "v1",
      "uptime": "0h 0m 14s"
    }

# Webhook triggered

When a webhook is triggered the server does a post request to the provided url. The body will contain the webhook id,
the name of the country and the total amount of times the country have been searched for in the api.

Example body:

    {
        "webhook_id": "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470",
        "country": "Sweden",
        "calls": 24
    }

# How to deploy

This project can be deployed either using the provided Dockerfile or building it using `go build`. Either way a
firestore service account credential file must be provided. It must be present in the root directory and be
called `auth.json`.

Both of the following methods have some common steps to be done.

1: clone the repo,
using `git clone https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022-workspace/mathias_ws/assignment-2.git`.   
2: change into the cloned directory: `cd assignment-2`.  
3: follow the firestore documentation to create an empty database and a service account authentication file.  
4: move the file into the current folder (`assignment-2`) and rename it to `auth.json`.  
5: follow the steps for you desired deployment method following:

## Docker

If not already done, please follow the general steps before proceeding. These steps will help you build a docker image
of the project and do a simple deployment. This method requires you to already have a working docker environment on
linux.

1: run `docker build . -t assigmnet2:latest`, this can take some time so be patient.  
2: run the container using `docker run --name assignment-2 -d -p 80:80 --restart unless-stopped assigmnet2:latest`. The
service can be reached on port 80 on the ip of the host.

Please see docker's documentation for more parameters or ways to deploy a docker image. The deployment option in step
two is just a simple way to quickly get it running.

## Go build

If not already done, please follow the general steps before proceeding. These steps will help you compile the project
and run the binary. To build this project go 1.17 or higher must be installed. Se the go docs for how to check you
version and how to install go.

1: to build the binary run the following
command: `go build -a -ldflags '-extldflags "-static"' -o assignment-2 cmd/main.go`, it should provide a self-contained
binary called assignment-2. This file can be moved anywhere you please.  
2: in the location the `assignment-2` file is placed use `./assignment-2` to run the service. The service should be
available on port 80 on the ip of the host.

# Design choices

The feature set of this is api is close to the specification given in the assignment. Some additional features have been
added, see the Extra section. The project aims for high cohesion and loose coupling. This project is created to be
highly modular to ease the maintenance and the re-usability in further projects. The api relies on the three backend
apis heavily, without them this api will not work.

I chose to use parameters when searching instead of adding the search string(s) into the path. This to make it clearer
for the user to see what information is to be added where. It also makes it easier to add additional features like:
searching for multiple names and/or countries. It also makes it easier to reuse the code for other endpoints.

The Dockerfiles ends up building the docker image based on the scratch base image for security and reliability reasons.
Only the necessary packages and dependencies are added to the container image. This minimizes the attack surface and
there will be fewer things to break making it more stable and reliable.

The tests are reliant on having some test data already populated into the database. This will make it impossible to run
the tests without the test database. Some tests might also contain some sleeps or use some functions that is being
tested later to make the tests run. This is because of some concurrency issues when go is running tests in parallel. For
example the test for adding a webhook in the logic package might run at the same time as the database function for
getting all webhooks are tested. This will cause the number of webhooks to maybe be different. This could have also been
solved using `go test -p ./...` as the tests are run sequentially.

The test coverage ended up being 100% of the files and 75,3% of the statements according to goland.

# Known bugs

* The country api does not always return the same name of a country as the cases api uses. This can cause some issues
  causing the service to not return a valid response.
* The cases api uses some inconsistent names of countries. This is not handled by this service. Some examples are:
  South, Korea and US. Please see the documentation of the backend case api to see what country names it is using.
* The default scope of the policy endpoint will almost always return an error because the backend api does not usually
  have data for the current day.
* Country api might sometimes show up with the status code of 502, and therefor not be available. This is probably
  because of a certificate issue with the service. Usually it works again after trying multiple times.

# Extras

* A short term cache of the cases endpoint is implemented to reduce requests to the backend api and to make the api
  respond faster. The cache is short-lived (12 hours) to always give the most up-to-date information.
* Collection of statistics of how many times a given country is called. This is also returned in the webhooks to give
  the user some information.
* Hashing of the database collection and document names, and also the webhook ids.
* Added a pipeline that run tests and automatically builds the image and pushes it to my docker registry.