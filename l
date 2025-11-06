[1mdiff --git a/k8s/overlays/prod/patches/backend-deployment.yaml b/k8s/overlays/prod/patches/backend-deployment.yaml[m
[1mindex 7fb88cf..ed4b73b 100644[m
[1m--- a/k8s/overlays/prod/patches/backend-deployment.yaml[m
[1m+++ b/k8s/overlays/prod/patches/backend-deployment.yaml[m
[36m@@ -8,7 +8,7 @@[m [mspec:[m
     spec:[m
       containers:[m
       - name: backend[m
[31m-        image: 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend:latest[m
[32m+[m[32m        image: 907849381252.dkr.ecr.us-east-1.amazonaws.com/globepay-backend:v1.0.1[m
         resources:[m
           requests:[m
             memory: "512Mi"[m
