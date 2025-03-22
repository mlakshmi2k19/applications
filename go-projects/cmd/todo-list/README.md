curl -X GET http://localhost:8080/tasks | jq
curl -X POST http://localhost:8080/create-new-task -H "Content-Type: application/json" -d '{"Title": "write API", "Content": "In-progress"}' 
