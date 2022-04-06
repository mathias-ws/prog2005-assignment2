# Assignment 2

Project for the second assignment in PROG2005-Cloud Technologies 2022. This rest api is created for getting information
about covid policies and cases in different countries around the world. With this api you can se any past or present
covid policies and the stringency of the policies in a country at a given time. With the cases endpoint the most recent
data regarding confirmed cases, deaths, growth rate and some other information.

This api relies on the covidtracker api from the university of Oxford: https://covidtracker.bsg.ox.ac.uk for the policy
information. For the covid cases this api relies on https://covid19-graphql.now.sh from rlindskog on GitHub. It uses
data from the Johns Hopkins University.

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

# How to deploy

## Docker

## Go build

# Design choices

# Known bugs

# Extras