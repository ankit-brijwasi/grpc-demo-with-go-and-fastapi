# Golang & FastAPI Demo

A simple boilerplate application, which uses gRPC to establish communication between fastapi and golang servers

## Prerequisites

1. Go v1.19
2. Python v3.8 or higher

## How to use?

1. Install the `Protobuf` compiler
```bash
$ sudo apt install -y protobuf-compiler
```

2. Install the `gen-go` and `gen-go-grpc` plugins
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. Install the dependencies for the `golang` service
```bash
$ go mod tidy
```

4. Install the dependencies for the `fastapi` service
```bash
$ pip install -r fastapi/app/requirements.txt
```

5. Run the `golang` service
```bash
$ go run main.go -port 8080
```

6. Run the `fastapi` service
```bash
$ cd ./fastapi/app && uvicorn main:app --host 127.0.0.1 --port 8000 --reload && cd ../../
```

Once the servers are running visit, [http://127.0.0.1:8000/docs](http://127.0.0.1:8000/docs) and make a get
request to the `/token/verify/` api route. This route makes an internal gRPC call to the `golang`
service which is responsible for handling the authentication

## Re-generate gRPC code

1. Generate `golang` gRPC code
```bash
$ protoc --go_out=. --go-grpc_out=. protobufs/auth.proto
```

2. Generate `python` gRPC code
```bash
$ cd ./fastapi/app && python -m grpc_tools.protoc -I ../protobufs --python_out=. --grpc_python_out=. ../protobufs/auth.proto cd ../../
```

## Common Error

1. Import error on `auth_pub2_grpc.py`

Open the `auth_pub2_grpc.py` file and change the import from
```python
import auth_pb2 as auth__pb2
```

to

```python
from . import auth_pb2 as auth__pb2
```

Cheers!