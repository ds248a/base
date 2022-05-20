# Base

Шаблон построения клиент-серверных приложений, включающий пакеты приложений:
**chi, grpc, cmux, testify, zap, opentelemetry, prometheus**

## Основные компоненты и их версии

**app**  
github.com/satori/go.uuid  
go.uber.org/zap  
github.com/stretchr/testify  


**client**  
google.golang.org/grpc  


**server**  
github.com/go-chi/chi  
google.golang.org/grpc  
github.com/grpc-ecosystem/grpc-gateway/v2/runtime  
github.com/soheilhy/cmux  
github.com/go-openapi/spec  


**server - health**  
google.golang.org/grpc  
google.golang.org/grpc/health/grpc_health_v1  
github.com/grpc-ecosystem/go-grpc-middleware  


**server - interceptors**  
google.golang.org/grpc  


**server - metrics (prometheus)**  
github.com/prometheus/client_golang/prometheus  


**metrics - prometheus**  
github.com/prometheus/common/expfmt  
github.com/prometheus/client_golang/prometheus  
github.com/prometheus/client_golang/prometheus/push  

**server - swagger_ui**  
github.com/go-openapi/spec  
github.com/rakyll/statik/fs  

**server - trace**  
google.golang.org/grpc  
go.opentelemetry.io/otel  


**tracing - otel**  
go.opentelemetry.io/otel  
go.opentelemetry.io/otel/sdk/trace  
go.opentelemetry.io/otel/attribute  
go.opentelemetry.io/otel/exporters/jaeger  
go.opentelemetry.io/otel/semconv/v1.4.0  
github.com/hashicorp/go-retryablehttp  


**logger**  
go.uber.org/zap  
go.uber.org/zap/zapcore  
