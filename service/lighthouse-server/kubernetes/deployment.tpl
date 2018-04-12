apiVersion: v1
kind: Secret
metadata:
  name: ${LH_GKE_CLUSTER}-secret
type: Opaque
stringData:
  TIDE_API_KEY: ${TIDE_API_KEY}
  TIDE_API_SECRET: ${TIDE_API_SECRET}
  LH_SQS_KEY: ${LH_SQS_KEY}
  LH_SQS_SECRET: ${LH_SQS_SECRET}
  LH_S3_KEY: ${LH_S3_KEY}
  LH_S3_SECRET: ${LH_S3_SECRET}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${LH_GKE_CLUSTER}-deployment
  labels:
    app: ${LH_GKE_CLUSTER}
spec:
  replicas: ${LH_GKE_REPLICAS}
  selector:
    matchLabels:
      app: ${LH_GKE_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${LH_GKE_CLUSTER}
    spec:
      containers:
      - image: ${LH_GCR_IMAGE_TAG}
        name: ${LH_GKE_CLUSTER}
        env:
        - name: TIDE_API_AUTH_URL
          value: "${TIDE_API_AUTH_URL}"
        - name: TIDE_API_HOST
          value: "${TIDE_API_HOST}"
        - name: TIDE_API_PROTOCOL
          value: "${TIDE_API_PROTOCOL}"
        - name: TIDE_API_KEY
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: TIDE_API_KEY
        - name: TIDE_API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: TIDE_API_SECRET
        - name: TIDE_API_VERSION
          value: "${TIDE_API_VERSION}"
        - name: LH_SQS_VERSION
          value: "${LH_SQS_VERSION}"
        - name: LH_SQS_REGION
          value: "${LH_SQS_REGION}"
        - name: LH_SQS_KEY
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: LH_SQS_KEY
        - name: LH_SQS_SECRET
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: LH_SQS_SECRET
        - name: LH_SQS_QUEUE_NAME
          value: "${LH_SQS_QUEUE_NAME}"
        - name: LH_S3_REGION
          value: "${LH_S3_REGION}"
        - name: LH_S3_KEY
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: LH_S3_KEY
        - name: LH_S3_SECRET
          valueFrom:
            secretKeyRef:
              name: ${LH_GKE_CLUSTER}-secret
              key: LH_S3_SECRET
        - name: LH_S3_BUCKET_NAME
          value: "${LH_S3_BUCKET_NAME}"
        - name: LH_CONCURRENT_AUDITS
          value: "${LH_CONCURRENT_AUDITS}"
---
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: ${LH_GKE_CLUSTER}-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: ${LH_GKE_CLUSTER}-deployment
  minReplicas: ${LH_GKE_MIN_PODS}
  maxReplicas: ${LH_GKE_MAX_PODS}
  targetCPUUtilizationPercentage: ${LH_GKE_CPU_PERCENT}