apiVersion: v1
kind: Secret
metadata:
  name: ${GKE_PHPCS_CLUSTER}-secret
type: Opaque
stringData:
  API_KEY: ${API_KEY}
  API_SECRET: ${API_SECRET}
  AWS_API_KEY: ${AWS_API_KEY}
  AWS_API_SECRET: ${AWS_API_SECRET}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${GKE_PHPCS_CLUSTER}-deployment
  labels:
    app: ${GKE_PHPCS_CLUSTER}
spec:
  replicas: ${GKE_PHPCS_REPLICAS}
  selector:
    matchLabels:
      app: ${GKE_PHPCS_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${GKE_PHPCS_CLUSTER}
    spec:
      containers:
      - image: ${GCR_PHPCS_IMAGE_TAG}
        name: ${GKE_PHPCS_CLUSTER}
        env:
        - name: API_AUTH_URL
          value: "${API_AUTH_URL}"
        - name: API_HTTP_HOST
          value: "${API_HTTP_HOST}"
        - name: API_KEY
          valueFrom:
            secretKeyRef:
              name: ${GKE_PHPCS_CLUSTER}-secret
              key: API_KEY
        - name: API_PROTOCOL
          value: "${API_PROTOCOL}"
        - name: API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${GKE_PHPCS_CLUSTER}-secret
              key: API_SECRET
        - name: API_VERSION
          value: "${API_VERSION}"
        - name: AWS_API_KEY
          valueFrom:
            secretKeyRef:
              name: ${GKE_PHPCS_CLUSTER}-secret
              key: AWS_API_KEY
        - name: AWS_API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${GKE_PHPCS_CLUSTER}-secret
              key: AWS_API_SECRET
        - name: AWS_S3_BUCKET_NAME
          value: "${AWS_S3_BUCKET_NAME}"
        - name: AWS_S3_REGION
          value: "${AWS_S3_REGION}"
        - name: AWS_S3_VERSION
          value: "${AWS_S3_VERSION}"
        - name: AWS_SQS_QUEUE_PHPCS
          value: "${AWS_SQS_QUEUE_PHPCS}"
        - name: AWS_SQS_REGION
          value: "${AWS_SQS_REGION}"
        - name: AWS_SQS_VERSION
          value: "${AWS_SQS_VERSION}"
        - name: PHPCS_CONCURRENT_AUDITS
          value: "${PHPCS_CONCURRENT_AUDITS}"
        - name: PHPCS_TEMP_FOLDER
          value: "${PHPCS_TEMP_FOLDER}"
---
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: ${GKE_PHPCS_CLUSTER}-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: ${GKE_PHPCS_CLUSTER}-deployment
  minReplicas: ${GKE_PHPCS_MIN_PODS}
  maxReplicas: ${GKE_PHPCS_MAX_PODS}
  targetCPUUtilizationPercentage: ${GKE_PHPCS_CPU_PERCENT}