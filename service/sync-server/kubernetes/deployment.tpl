apiVersion: v1
kind: Secret
metadata:
  name: ${GKE_SYNC_CLUSTER}-secret
type: Opaque
stringData:
  AWS_API_KEY: "${AWS_API_KEY}"
  AWS_API_SECRET: "${AWS_API_SECRET}"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${GKE_SYNC_CLUSTER}-deployment
  labels:
    app: ${GKE_SYNC_CLUSTER}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ${GKE_SYNC_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${GKE_SYNC_CLUSTER}
    spec:
      containers:
        - name: ${GKE_SYNC_CLUSTER}
          image: ${GCR_SYNC_IMAGE_TAG}
          terminationMessagePolicy: FallbackToLogsOnError
          env:
            - name: API_HTTP_HOST
              value: "${API_HTTP_HOST}"
            - name: API_PROTOCOL
              value: "${API_PROTOCOL}"
            - name: API_VERSION
              value: "${API_VERSION}"
            - name: AWS_API_KEY
              valueFrom:
                secretKeyRef:
                  name: ${GKE_SYNC_CLUSTER}-secret
                  key: AWS_API_KEY
            - name: AWS_API_SECRET
              valueFrom:
                secretKeyRef:
                  name: ${GKE_SYNC_CLUSTER}-secret
                  key: AWS_API_SECRET
            - name: AWS_SQS_QUEUE_LH
              value: "${AWS_SQS_QUEUE_LH}"
            - name: AWS_SQS_QUEUE_PHPCS
              value: "${AWS_SQS_QUEUE_PHPCS}"
            - name: AWS_SQS_REGION
              value: "${AWS_SQS_REGION}"
            - name: AWS_SQS_VERSION
              value: "${AWS_SQS_VERSION}"
            - name: GCF_QUEUE_LH
              value: "${GCF_QUEUE_LH}"
            - name: GCF_QUEUE_PHPCS
              value: "${GCF_QUEUE_PHPCS}"
            - name: GCP_PROJECT
              value: "${GCP_PROJECT}"
            - name: SYNC_ACTIVE
              value: "${SYNC_ACTIVE}"
            - name: SYNC_API_BROWSE_CATEGORY
              value: "${SYNC_API_BROWSE_CATEGORY}"
            - name: SYNC_DATA
              value: "${SYNC_DATA}"
            - name: SYNC_DATABASE_DOCUMENT_PATH
              value: "${SYNC_DATABASE_DOCUMENT_PATH}"
            - name: SYNC_DATABASE_PROVIDER
              value: "${SYNC_DATABASE_PROVIDER}"
            - name: SYNC_DEFAULT_CLIENT
              value: "${SYNC_DEFAULT_CLIENT}"
            - name: SYNC_DEFAULT_VISIBILITY
              value: "${SYNC_DEFAULT_VISIBILITY}"
            - name: SYNC_FORCE_AUDITS
              value: "${SYNC_FORCE_AUDITS}"
            - name: SYNC_ITEMS_PER_PAGE
              value: "${SYNC_ITEMS_PER_PAGE}"
            - name: SYNC_LH_ACTIVE
              value: "${SYNC_LH_ACTIVE}"
            - name: SYNC_MESSAGE_PROVIDER
              value: "${SYNC_MESSAGE_PROVIDER}"
            - name: SYNC_PHPCS_ACTIVE
              value: "${SYNC_PHPCS_ACTIVE}"
            - name: SYNC_POOL_DELAY
              value: "${SYNC_POOL_DELAY}"
            - name: SYNC_POOL_WORKERS
              value: "${SYNC_POOL_WORKERS}"
          volumeMounts:
            - name: ${GKE_SYNC_CLUSTER}-volume
              mountPath: $SYNC_DATA
      volumes:
        - name: ${GKE_SYNC_CLUSTER}-volume
          gcePersistentDisk:
            pdName: ${GKE_SYNC_CLUSTER}-disk
            fsType: ext4