dev:
	go run *.go

testpost:
	curl -d '{"name": "indomie", "price": 20000}' http://localhost:8080/api/v1/belonjoan