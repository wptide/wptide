apiVersion: v1
kind: Secret
metadata:
  name: ${GKE_SYNC_CLUSTER}-secret
type: Opaque
stringData:
  AWS_API_KEY: $AWS_API_KEY
  AWS_API_SECRET: $AWS_API_SECRET
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ${GKE_SYNC_CLUSTER}-pv-claim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi
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
      - image: ${GCR_SYNC_IMAGE_TAG}
        name: ${GKE_SYNC_CLUSTER}
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
        - name: SYNC_ACTIVE
          value: "${SYNC_ACTIVE}"
        - name: SYNC_API_BROWSE_CATEGORY
          value: "${SYNC_API_BROWSE_CATEGORY}"
        - name: SYNC_DATA
          value: "${SYNC_DATA}"
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
        - name: SYNC_PHPCS_ACTIVE
          value: "${SYNC_PHPCS_ACTIVE}"
        - name: SYNC_POOL_DELAY
          value: "${SYNC_POOL_DELAY}"
        - name: SYNC_POOL_WORKERS
          value: "${SYNC_POOL_WORKERS}"
        volumeMounts:
        - name: ${GKE_SYNC_CLUSTER}-persistent-storage
          mountPath: $SYNC_DATA
      volumes:
      - name: ${GKE_SYNC_CLUSTER}-persistent-storage
        persistentVolumeClaim:
          claimName: ${GKE_SYNC_CLUSTER}-pv-claim