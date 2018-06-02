apiVersion: v1
kind: Secret
metadata:
  name: ${GKE_LH_CLUSTER}-secret
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
  name: ${GKE_LH_CLUSTER}-deployment
  labels:
    app: ${GKE_LH_CLUSTER}
spec:
  replicas: ${GKE_LH_REPLICAS}
  selector:
    matchLabels:
      app: ${GKE_LH_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${GKE_LH_CLUSTER}
    spec:
      containers:
      - image: ${GCR_LH_IMAGE_TAG}
        name: ${GKE_LH_CLUSTER}
        terminationMessagePolicy: FallbackToLogsOnError
        env:
        - name: API_AUTH_URL
          value: "${API_AUTH_URL}"
        - name: API_HTTP_HOST
          value: "${API_HTTP_HOST}"
        - name: API_KEY
          valueFrom:
            secretKeyRef:
              name: ${GKE_LH_CLUSTER}-secret
              key: API_KEY
        - name: API_PROTOCOL
          value: "${API_PROTOCOL}"
        - name: API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${GKE_LH_CLUSTER}-secret
              key: API_SECRET
        - name: API_VERSION
          value: "${API_VERSION}"
        - name: AWS_API_KEY
          valueFrom:
            secretKeyRef:
              name: ${GKE_LH_CLUSTER}-secret
              key: AWS_API_KEY
        - name: AWS_API_SECRET
          valueFrom:
            secretKeyRef:
              name: ${GKE_LH_CLUSTER}-secret
              key: AWS_API_SECRET
        - name: AWS_S3_BUCKET_NAME
          value: "${AWS_S3_BUCKET_NAME}"
        - name: AWS_S3_REGION
          value: "${AWS_S3_REGION}"
        - name: AWS_S3_VERSION
          value: "${AWS_S3_VERSION}"
        - name: AWS_SQS_QUEUE_LH
          value: "${AWS_SQS_QUEUE_LH}"
        - name: AWS_SQS_REGION
          value: "${AWS_SQS_REGION}"
        - name: AWS_SQS_VERSION
          value: "${AWS_SQS_VERSION}"
        - name: LH_CONCURRENT_AUDITS
          value: "${LH_CONCURRENT_AUDITS}"
        - name: LH_MESSAGE_PROVIDER
          value: "${LH_MESSAGE_PROVIDER}"
        - name: LH_STORAGE_PROVIDER
          value: "${LH_STORAGE_PROVIDER}"
        - name: LH_TEMP_FOLDER
          value: "${LH_TEMP_FOLDER}"
        - name: GCF_QUEUE_LH
          value: "${GCF_QUEUE_LH}"
        - name: GCP_PROJECT
          value: "${GCP_PROJECT}"
        - name: GCS_BUCKET_NAME
          value: "${GCS_BUCKET_NAME}"
---
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: ${GKE_LH_CLUSTER}-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: ${GKE_LH_CLUSTER}-deployment
  minReplicas: ${GKE_LH_MIN_PODS}
  maxReplicas: ${GKE_LH_MAX_PODS}
  targetCPUUtilizationPercentage: ${GKE_LH_CPU_PERCENT}