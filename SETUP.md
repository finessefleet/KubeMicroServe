# KubeMicroServe - Setup & Usage Guide

## Quick Start

The included Makefile streamlines the entire workflow:

```bash
# Install dependencies and set up K3d cluster
make install

# Build the application (automatically done if needed)
make build

# Deploy to K3d cluster
make cluster-up
```

Once deployed, you can access the application at http://localhost:8080

## Monitoring

This project includes a lightweight monitoring stack using Prometheus and Grafana.

### Metrics Available

The Go application exposes the following metrics via its `/metrics` endpoint:

- `myapp_http_requests_total` - Counter of HTTP requests by path and status code
- `http_request_duration_seconds` - Histogram of HTTP request durations by path

### Setting up Monitoring

```bash
# Deploy Prometheus and Grafana
make monitoring

# Access monitoring dashboards
# Prometheus: http://localhost:9090
# Grafana: http://localhost:3000

# Stop monitoring port-forwarding
make monitoring-stop
```

### Dashboard Setup in Grafana

Once Grafana is running:

1. Navigate to http://localhost:3000 & login using admin / admin
2. Go to "Data Sources" and add Prometheus
   - URL: http://prometheus-service:9090
3. Import a dashboard
   - Go to "+" > "Import"
   - Copy paste the template.json / upload it for the template dashboard

### Prometheus Query Examples

Some useful Prometheus queries:

- Request rate: `rate(myapp_http_requests_total[1m])`
- Error rate: `rate(myapp_http_requests_total{status=~"5.."}[1m])`
- 90th percentile response time: `histogram_quantile(0.9, rate(myapp_http_request_duration_seconds_bucket[5m]))`
- also you can test in cli by running `curl -s http://localhost:9090/metrics | head -n20`

## Project Structure

```
.
├── app/                    # Go application
│   ├── internal/           # Handlers & middleware
│   ├── main.go             # Main application code
│   └── Dockerfile          # Container definition
├── config/                 # Kubernetes manifests
│   ├── deployment.yaml     # Kubernetes deployment
│   ├── namespace.yaml      # Application namespace
│   └── service.yaml        # Service definition
├── hack/                   # Helper scripts
│   ├── build.sh            # Build Docker image
│   ├── install.sh          # Install dependencies
│   └── run.sh              # Deploy to cluster
├── Makefile                # Automation commands
├── LICENSE                 # MIT License
└── README.md               # Project documentation
```

## Commands

The Makefile provides the following commands:

| Command | Description |
|---------|-------------|
| `make` | Checks for Docker image, builds if needed, and deploys to cluster |
| `make install` | Installs dependencies and sets up K3d cluster |
| `make build` | Builds the Docker image |
| `make cluster-up` | Deploys application to K3d cluster |
| `make clean` | Removes application from cluster |
| `make monitoring` | Deploys Prometheus and Grafana |
| `make monitoring-stop` | Stops monitoring port-forwarding |
| `make help` | Shows available commands |

## Installation Details

The installation script (`hack/install.sh`) handles:
- Installing Go 1.22.2 (if not present)
- Installing kubectl (if not present)
- Installing Docker (in local environments, if not present)
- Creating a K3d/K3s cluster

The script adapts to different environments:
- GitHub Codespaces: Uses K3d with a LoadBalancer
- Local Linux: Installs Docker and K3s/K3d as needed

## Development Notes

- The application runs on port 8080 inside the container
- The service exposes this on port 80
- The K3d cluster maps port 80 to host port 80
- The deployment includes proper health checks and resource limits

## Testing the Application

After deployment, the service is available:

```bash
# Using port-forwarding (automatically started in make cluster-up)
# Access at http://localhost:8080

# Check pod status
kubectl get pods -n go-k3-app-ns

# View logs
kubectl logs -f -l app=go-k3-app -n go-k3-app-ns
```

## Troubleshooting

- **Pods not starting**: Check for image pull or readiness probe issues with `kubectl describe pod -n go-k3-app-ns`
- **Service not accessible**: Verify port forwarding is active or check service with `kubectl get svc -n go-k3-app-ns`
- **Build errors**: Ensure Docker is running and Go is properly installed 