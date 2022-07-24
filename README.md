# restful-api-go
In this project, our goal is to build a RESTful API that performs simple CURD operations using Golang.

    "REST": Stands for "Representational State Transfer"
    "Restful": The service that use REST API to communicate
    
    CURD:               HTTP Method/Operations:
    Create --------->   POST
    Read   --------->   GET
    Update --------->   PUT
    Delete --------->   DELETE
    
    Pros of a REST API:
        1.Simple/Standardized Approach
        2.Scalable/Stateless(Easy to modify/Do not need to worry about which data is in which state ad keep track of that acrossing client&server)
        3.High Performance/Has Cache

    Here is the external packages/libararies we use in this project: 
        1. gorilla/mux:= For creating routes and HTTP handlers for our endpoints
        2. jinzhu/gorm:= An ORM tool for MySQL.
        3. go-sql-driver/mysql:= MYSQL driver.

