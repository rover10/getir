# Getir

This web project is based on pure golang native packages(http). 

# Builing project
Clone the project from this repo
Run ``go build`` 
Build is creates a binary `getir` in the current location

# Running 
The application can be run with enviromnet specific varaibales. If theses values are not set then default values are provided.
``
MONGO_USER
MONGO_PASSWORD
MONGO_HOST
PORT
MONGO_DATABASE
``
After setting desired env varaiables run the binary ``./getir``

# APIs
Implements two APIs
1. GET  -  ``/in-memory`` with query param ``key``
2. POST -  ``/in-memory`` which accepts body ``{"key": "Key1", "value": "Value1"}``

# Code
The struct ``Server`` holds in-memory datastructure, MongoDb client and it's context.
```
type Server struct {
	InmemoryDatabase map[string]string
	MongoClient      *mongo.Client
	Context          context.Context
}
```
