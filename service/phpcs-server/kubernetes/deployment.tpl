apiVersion: v1
kind: Secret
metadata:
  name: ${PHPCS_GKE_CLUSTER}-secret
type: Opaque
stringData:
  TIDE_API_KEY: ${TIDE_API_KEY}
  TIDE_API_SECRET: ${TIDE_API_SECRET}
  PHPCS_SQS_KEY: ${PHPCS_SQS_KEY}
  PHPCS_SQS_SECRET: ${PHPCS_SQS_SECRET}
  PHPCS_S3_KEY: ${PHPCS_S3_KEY}
  PHPCS_S3_SECRET: ${PHPCS_S3_SECRET}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${PHPCS_GKE_CLUSTER}-deployment
  labels:
    app: ${PHPCS_GKE_CLUSTER}
spec:
  replicas: ${PHPCS_GKE_REPLICAS}
  selector:
    matchLabels:
      app: ${PHPCS_GKE_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${PHPCS_GKE_CLUSTER}
    spec:
      containers:
      - image: ${PHPCS_GCR_IMAGE_TAG}
        name: ${PHPCS_GKE_CLUSTER}
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
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: TIDE_API_KEY
        - name: TIDE_API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: TIDE_API_SECRET
        - name: TIDE_API_VERSION
          value: "${TIDE_API_VERSION}"
        - name: PHPCS_SQS_VERSION
          value: "${PHPCS_SQS_VERSION}"
        - name: PHPCS_SQS_REGION
          value: "${PHPCS_SQS_REGION}"
        - name: PHPCS_SQS_KEY
          valueFrom:
            secretKeyRef:
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: PHPCS_SQS_KEY
        - name: PHPCS_SQS_SECRET
          valueFrom:
            secretKeyRef:
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: PHPCS_SQS_SECRET
        - name: PHPCS_SQS_QUEUE_NAME
          value: "${PHPCS_SQS_QUEUE_NAME}"
        - name: PHPCS_S3_REGION
          value: "${PHPCS_S3_REGION}"
        - name: PHPCS_S3_KEY
          valueFrom:
            secretKeyRef:
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: PHPCS_S3_KEY
        - name: PHPCS_S3_SECRET
          valueFrom:
            secretKeyRef:
              name: ${PHPCS_GKE_CLUSTER}-secret
              key: PHPCS_S3_SECRET
        - name: PHPCS_S3_BUCKET_NAME
          value: "${PHPCS_S3_BUCKET_NAME}"
---
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: ${PHPCS_GKE_CLUSTER}-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: ${PHPCS_GKE_CLUSTER}-deployment
  minReplicas: ${PHPCS_GKE_MIN_PODS}
  maxReplicas: ${PHPCS_GKE_MAX_PODS}
  targetCPUUtilizationPercentage: ${PHPCS_GKE_CPU_PERCENT}