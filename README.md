# base
Base library for microservices

## import

### app > app_identity
github.com/satori/go.uuid 

### app > jobs
go.uber.org/zap  
github.com/stretchr/testify/assert  

### client > gRPC
google.golang.org/grpc  

### server
github.com/go-chi/chi  
google.golang.org/grpc  
github.com/grpc-ecosystem/grpc-gateway/v2/runtime  
github.com/soheilhy/cmux  
github.com/go-openapi/spec  

### server > health_check
google.golang.org/grpc
google.golang.org/grpc/health/grpc_health_v1
github.com/grpc-ecosystem/go-grpc-middleware

### server > interceptors
google.golang.org/grpc

### server > internal
github.com/soheilhy/cmux
github.com/go-openapi/spec

### server > metrics (prometheus)
github.com/prometheus/client_golang/prometheus

### server > swagger_ui
github.com/go-openapi/spec
github.com/rakyll/statik/fs

### metrics > prometheus
github.com/prometheus/common/expfmt
github.com/prometheus/client_golang/prometheus
github.com/prometheus/client_golang/prometheus/push

### server > trace
go.opentelemetry.io/otel
google.golang.org/grpc

### tracing / otel
go.opentelemetry.io/otel
go.opentelemetry.io/otel/sdk/trace
go.opentelemetry.io/otel/attribute
go.opentelemetry.io/otel/exporters/jaeger
go.opentelemetry.io/otel/semconv/v1.4.0
github.com/hashicorp/go-retryablehttp

### logger
go.uber.org/zap
go.uber.org/zap/zapcore


// ------------------------
//   IMPL
// ------------------------
