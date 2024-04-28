# Simple golang GRPC

### Run:  `make run-orders`
### Run:  `make run-kitchen`

- use htttp -> to make orders
    + POST:  http://localhost:8000/orders
    + Payload: {"CustomerID" : 1,"ProductID": 2,"Quantity": 3}

- user http -> to get order from kitchen
    + GET: http://localhost:1000
    + Payload: {"CustomerID": 2}