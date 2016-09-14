# Using truss

## File structure

You start with your definition file.

```
.
└── svc.proto
```

Then: `truss svc.proto`
After:

```
.
├── NAME-service
│   ├── docs
│   │   └── docs.md
│   ├── generated
│   │   └── ...
│   ├── handlers
│   │   ├── client
│   │   │   └── client_handler.go
│   │   └── server
│   │       └── server_handler.go
│   ├── NAME-client
│   │   └── client_main.go
│   ├── NAME-server
│   │   └── server_main.go
│   └── svc.pb.go
├── svc.proto
└── third_party
    └── ...
```

*Our Contact*

1. Only the files in `hanlders` are user modifiable. Only function names defined as rpc's in the definition service, all other functions will be removed. User logic can be imported and executed in the functions in the handlers.
2. Do not create files or directories in NAME-service
3. All user logic should exist outside of NAME-service, up to the users discretion for organization.

 


