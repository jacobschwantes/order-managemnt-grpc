package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jacobschwantes/order-managemnt-grpc/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		_, err := client.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 123,
			ProductID:  123,
			Quantity:   2,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orders, err := client.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 123,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		t.Execute(w, orders)
	})

	log.Printf("Kitchen service listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<html>
<head>
	<title>Orders List</title>
</head>
<body>
	<h1>Orders</h1>
	<table>
		<tr>
			<th>Order ID</th>
			<th>Customer ID</th>
			<th>Product ID</th>
			<th>Quantity</th>
		</tr>
		{{range .Orders}}
		<tr>
			<td>{{.OrderID}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.ProductID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`
