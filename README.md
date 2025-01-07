## **README.md**

### **How to Run the Project**

1. **Set Environment Variables**  

```bash

export USER_REQUEST_LIMIT=3
export GLOBAL_REQUEST_LIMIT=10

go run cmd/server/main.go

```


2. **execute curl**  


`curl "http://localhost:8088/request?client_id=user1"`

`curl "http://localhost:8088/request?client_id=user2"`
