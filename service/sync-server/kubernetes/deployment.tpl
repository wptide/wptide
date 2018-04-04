apiVersion: v1
kind: Secret
metadata:
  name: ${SYNC_GKE_CLUSTER}-secret
type: Opaque
stringData:
  LH_SQS_KEY: $LH_SQS_KEY
  LH_SQS_SECRET: $LH_SQS_SECRET
  PHPCS_SQS_KEY: $PHPCS_SQS_KEY
  PHPCS_SQS_SECRET: $PHPCS_SQS_SECRET
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ${SYNC_GKE_CLUSTER}-pv-claim
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
  name: ${SYNC_GKE_CLUSTER}-deployment
  labels:
    app: ${SYNC_GKE_CLUSTER}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ${SYNC_GKE_CLUSTER}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: ${SYNC_GKE_CLUSTER}
    spec:
      containers:
      - image: ${SYNC_GCR_IMAGE_TAG}
        name: ${SYNC_GKE_CLUSTER}
        env:
        - name: TIDE_API_HOST
          value: "${TIDE_API_HOST}"
        - name: TIDE_API_PROTOCOL
          value: "${TIDE_API_PROTOCOL}"
        - name: TIDE_API_VERSION
          value: "${TIDE_API_VERSION}"
        - name: SYNC_ACTIVE
          value: "${SYNC_ACTIVE}"
        - name: SYNC_API_BROWSE_CATEGORY
          value: "${SYNC_API_BROWSE_CATEGORY}"
        - name: SYNC_DATA
          value: "${SYNC_DATA}"
        - name: SYNC_DEFAULT_CLIENT
          value: "${SYNC_DEFAULT_CLIENT}"
        - name: SYNC_DEFAULT_VISIBILITY
          value: "${SYNC_DEFAULT_VISIBILITY}"
        - name: SYNC_FORCE_AUDITS
          value: "${SYNC_FORCE_AUDITS}"
        - name: SYNC_ITEMS_PER_PAGE
          value: "${SYNC_ITEMS_PER_PAGE}"
        - name: SYNC_POOL_DELAY
          value: "${SYNC_POOL_DELAY}"
        - name: SYNC_POOL_WORKERS
          value: "${SYNC_POOL_WORKERS}"
        - name: PHPCS_SQS_KEY
          valueFrom:
            secretKeyRef:
              name: ${SYNC_GKE_CLUSTER}-secret
              key: PHPCS_SQS_KEY
        - name: PHPCS_SQS_QUEUE
          value: "${PHPCS_SQS_QUEUE}"
        - name: PHPCS_SQS_QUEUE_NAME
          value: "${PHPCS_SQS_QUEUE_NAME}"
        - name: PHPCS_SQS_REGION
          value: "${PHPCS_SQS_REGION}"
        - name: PHPCS_SQS_SECRET
          valueFrom:
            secretKeyRef:
              name: ${SYNC_GKE_CLUSTER}-secret
              key: PHPCS_SQS_SECRET
        - name: LH_SQS_KEY
          valueFrom:
            secretKeyRef:
              name: ${SYNC_GKE_CLUSTER}-secret
              key: LH_SQS_KEY
        - name: LH_SQS_QUEUE
          value: "${LH_SQS_QUEUE}"
        - name: LH_SQS_QUEUE_NAME
          value: "${LH_SQS_QUEUE_NAME}"
        - name: LH_SQS_REGION
          value: "${LH_SQS_REGION}"
        - name: LH_SQS_SECRET
          valueFrom:
            secretKeyRef:
              name: ${SYNC_GKE_CLUSTER}-secret
              key: LH_SQS_SECRET
        volumeMounts:
        - name: ${SYNC_GKE_CLUSTER}-persistent-storage
          mountPath: $SYNC_DATA
      volumes:
      - name: ${SYNC_GKE_CLUSTER}-persistent-storage
        persistentVolumeClaim:
          claimName: ${SYNC_GKE_CLUSTER}-pv-claim