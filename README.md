# article-management-service
Manage Articles

## Initial Thoughts
    Before beginning I've gone ahead and planned out how I want to structure this service and API. That being said I'm a huge fan of building my go RESTful services using Swagger's Codegen tool since it offers a very clean way to document and define your API along with the convenience of generating not only server but also client code stubs. This avoids a lot of the boilerplate code required for developing these kind of services but for the purpose of this assignment I've gone ahead and built the server from scratch. 
    
    My aim is to use go standard packages only with the exception of maybe some tools like a multiplexer(for avoiding the hastle of parsing urls, subrouting, etc.), unit testing, and DB access.

## Readme Structure
    The following readme serves multiple purposes in regards to the service. 
    The first is documenting and defining the API for this service as this is ultimately the most core task for this assignment. 
    This document also provides insight into how the service works, but more importantly reveals the logos behind the design of the service. To better reflect this, I've decided to take down notes detailing my thought process during the development of this service.


## Pre Planning
    Before beginnind the assignment I want to clearly define the api outlined by the assignment requirements and start brainstorming how I want to structure the server as well as set some goals:
    - This API will only be handling posts which I've decided to rename articles to remove ambiguity. The assignment prompt gives the data model for users but does not set any requirements for managing these users. This does make sense if we're abiding by the microservice single responsibility principal.
    - The endpoints identified from the prompt:
        GET     /articles                           - returns all articles
        GET     /articles/{articleId}               - returns article with given id
        GET     /articles?ids=id1, id2, idn...      - returns []articles with given ids
        POST    /articles                           - creates new article
        PUT     /articles/{articleId}               - updates article with given id
        PATCH   /articles/{articleId}               - skipped(redundant for simple data model)
        DELETE  /articles/{articleId}               - deletes article with given id
        GET     /articles/user/{userId}             - returns articles authored by given user

    - Goals for the project:
        1. Minimum Viable Product(MVP) of working service that meets all requirements
        2. Simple structured logging utility
        3. Interface for storage solution
            a. Key-Value in-memory data storage
            b. DynamoDB data storage
        4. Unit testing of all public exported functions
        5. Concurrency for handling multiple ids
            (I'll look for other opportunities to utilize concurrency as well)
        6. Input Authentication
        7. Containerize the application and run on Docker 
        8. Deploy onto AWS via either ECS or EKS
    
    DISCLAIMER - These goals may be overestimating how much work I can accomplish in 4 hours with breaks but it's worth the attempt.

## Notes
- To begin I'll setup a github repo for an article management service
- Now I need to setup the package/project layout for this project. For this I tend to ascribe to the go standards project layout: https://github.com/golang-standards/project-layout
- 





