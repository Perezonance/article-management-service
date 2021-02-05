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
- Now that I've established some order, I'll write a simple "Hello, world." program to test everything
- As I mentioned beforehand, I plan on using gorilla/mux for building this server for just the few convenience features so I'll start by setting up the router, defining the paths with handlers, and starting the server with ListenAndServe on http://localhost:8000
- Now I know that I want to follow a structured design from the beginning with this project so I'm going to develop it based similarly to the MVC model. 
- I've gone ahead and setup a package structure for the internal directory to reflect this design.
- I've started with building out the controller struct which will define the layer responsible for processing the request and response, (input)authentication, and authorization
- While I start working on the controller as well as its methods I've gone ahead and used the subrouting feature from gorilla/mux to assign all the handlers to each path and method defined by the API
- Before continuing with work on the controller methods I need to define the server and its functions as well. This is of course housed in the internal/server pkg
- Defined the Storage interface behavior and will implement the signatures on a MockDynamo struct that will mimic a key-value store, in-memory.
- provide the Article data model within the internal/models package
- while I'm defining the data model for an article, i've also included a model for a NewArticle which is the model we'll use for the request payload when providing an article for a creation operation. Reason for is because we want the server to handle articleID generation.
- I'll also need to create each of these instances within the main.go as part of startup of the server.
- For article creation I've come up with a simple in-memory counter to provide a "unique" id for each article.





