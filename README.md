# Form3 Take Home Exercise

## Overview

This is a client library written in Go that interacts with a fake Form3 Account API. It supports the following operations :
* Create - creates a new organisation account
* Fetch - retrieves a single organisation account if it exists
* Delete - deletes an organisation account given its Id and Version

I am a beginner in GOLang and this API is my first project in this. Although, I have 6 years of experience in software development (primarily using JAVA), therefore, have tried to develop this client library in an extensible way using design principles, such that it is reusable for other entities as well.

## Running the tests

For running the tests against the provided fake account API, the following command can be run :

```
docker-compose up
```

## Example of using the client library

```go
f3 := form3.Init("http://localhost:8080")

// creates an organisation account insidde Form3
org, err = f3.Create(account)

// retrieve a single organisation using the ID below
org, err = f3.Fetch(accountId)

// remove an organisation account with the ID and version below
err = f3.Delete(accountId, 0)

```

## Technical Overview

An overview of library design and decisions :

### Project Structure

I have tried to keep files for similar functionality in same package. The project has a 'core' package in which common files for making a request, parsing response etc are added. This is written in an extensible way and can be reused for other entities.
For the main functionality, I have created another package named functions which holds the functional implementation for each of the requested options. The core functionality is wrapped in a package named core, similarly we have different package for models as well.

### Client library Implementation
I have tried to keep the requests, response and client implementation separate so that each of them could be used independently for other entities if required. The request entity is a wrapper over vanilla http request and exposes all the relevant HTTP Methods and provides extensibility to include different headers, parameters if required.

### Testing
All the modules have been thoroughly tested. Along with this, I have added integration tests by connecting to the FakeAccount API provided.

