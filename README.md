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

### Request

#### Parameters

### Response

## Notifications

### Request

#### Parameters

### Response

# How to deploy

## Docker

## Go build

# Design choices

# Known bugs

# Extras