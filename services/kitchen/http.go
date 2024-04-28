package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/AchillesLe/kitchen/services/common/genproto/orders"
	"github.com/gorilla/schema"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (h *HttpServer) Run() error {
	router := http.NewServeMux()
	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		var decoder = schema.NewDecoder()
		var input orders.GetOrdersRequest
		err := decoder.Decode(&input, r.URL.Query())

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: input.GetCustomerID(),
		})

		if err != nil {
			fmt.Printf("Client error %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			fmt.Printf("Template error: %v", err)
		}

		// json.NewEncoder(w).Encode(orderDb)
	})

	fmt.Println("Starting http server on ", h.addr)
	return http.ListenAndServe(h.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Orders</title>
</head>
<body>
    <table border="1">
		<tr>
			<th>Order ID</th>
			<th>Customer ID</th>
			<th>Quantity</th>
		</tr>
		{{ range . }}
		<tr>
			<td>{{.OrderID}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
		{{ end }}
    </table>
</body>
</html>
`
