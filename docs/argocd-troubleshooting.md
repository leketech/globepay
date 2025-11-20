# ArgoCD Troubleshooting Guide

This document provides guidance on troubleshooting issues with ArgoCD deployments in the Globepay application.

## Common Issues and Solutions

### 1. Applications Stuck in "OutOfSync" State

#### Symptoms:
- Applications show as "OutOfSync" in the ArgoCD UI
- Changes made in the Git repository are not reflected in the cluster

#### Solutions:
1. Manually sync the application:
   ```bash
   argocd app sync <application-name>
   ```

2. Check if auto-sync is enabled in the application specification

3. Verify that the Git repository URL and revision are correct

### 2. Image Pull Errors

#### Symptoms:
- Pods stuck in "ImagePullBackOff" or "ErrImagePull" state
- Application fails to start

#### Solutions:
1. Verify that the image tags in the values files match the images built by the CI pipeline
2. Check that images exist in the container registry
3. Ensure proper image pull secrets are configured

### 3. Health Checks Failing

#### Symptoms:
- Applications show as "Progressing" or "Degraded" instead of "Healthy"
- Readiness/liveness probes failing

#### Solutions:
1. Check pod logs:
   ```bash
   kubectl logs -n <namespace> <pod-name>
   ```

2. Verify that the application is listening on the correct port
3. Check that required environment variables and secrets are properly configured

### 4. Resource Quota Issues

#### Symptoms:
- Pods stuck in "Pending" state
- Error messages about resource quotas

#### Solutions:
1. Check resource quotas in the namespace:
   ```bash
   kubectl describe quota -n <namespace>
   ```

2. Adjust resource requests and limits in the Helm values files

## Diagnostic Commands

### Check Application Status
```bash
# List all applications
argocd app list

# Get detailed status of a specific application
argocd app get <application-name>

# View application history
argocd app history <application-name>
```

### Check Kubernetes Resources
```bash
# Check pod status
kubectl get pods -n <namespace>

# Check deployment status
kubectl get deployments -n <namespace>

# Check service status
kubectl get services -n <namespace>

# Check ingress status
kubectl get ingress -n <namespace>
```

### View Logs
```bash
# View pod logs
kubectl logs -n <namespace> <pod-name>

# View logs for previous container instance
kubectl logs -n <namespace> <pod-name> --previous

# Stream logs in real-time
kubectl logs -n <namespace> <pod-name> -f
```

## Best Practices

1. **Consistent Image Tagging**: Ensure that the image tagging strategy is consistent between CI/CD pipelines and ArgoCD configurations.

2. **Proper Sync Policies**: Configure appropriate sync policies in the Application manifests to control how changes are applied.

3. **Resource Management**: Set appropriate resource requests and limits to prevent resource contention.

4. **Monitoring and Alerting**: Implement monitoring and alerting for ArgoCD applications to detect issues early.

5. **Regular Sync**: Regularly sync applications to ensure they reflect the desired state in the Git repository.