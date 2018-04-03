apiVersion: v1
kind: Secret
metadata:
  name: sync-server-secret
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
  name: sync-server-pv-claim
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
  name: sync-server-deployment
  labels:
    app: sync-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sync-server
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: sync-server
    spec:
      containers:
      - image: ${REPO}/${SYNC_TAG}
        name: sync-server
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
              name: sync-server-secret
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
              name: sync-server-secret
              key: PHPCS_SQS_SECRET
        - name: LH_SQS_KEY
          valueFrom:
            secretKeyRef:
              name: sync-server-secret
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
              name: sync-server-secret
              key: LH_SQS_SECRET
        volumeMounts:
        - name: sync-server-persistent-storage
          mountPath: $SYNC_DATA
      volumes:
      - name: sync-server-persistent-storage
        persistentVolumeClaim:
          claimName: sync-server-pv-claim