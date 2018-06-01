apiVersion: v1
kind: Secret
metadata:
  name: ${GKE_REDIS_CLUSTER}-secret
type: Opaque
stringData:
  API_REDIS_AUTH: ${API_REDIS_AUTH}
---
apiVersion: v1
kind: Service
metadata:
  name: ${GKE_REDIS_CLUSTER}-service
  labels:
    app: redis
spec:
  type: LoadBalancer
  ports:
    - port: 6379
      name: redis
  selector:
    app: redis
---
apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  name: ${GKE_REDIS_CLUSTER}-stateful-set
  labels:
    app: redis
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  serviceName: ${GKE_REDIS_CLUSTER}-service
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
        - name: redis
          image: redis:3.2-alpine
          imagePullPolicy: Always
          args: ["--requirepass", "$(API_REDIS_AUTH)", "--appendonly", "no", "--save", "3600", "1", "--save", "300", "100", "--save", "60", "10000"]
          ports:
            - containerPort: 6379
              name: redis
          env:
            - name: API_REDIS_AUTH
              valueFrom:
                secretKeyRef:
                  name: ${GKE_REDIS_CLUSTER}-secret
                  key: API_REDIS_AUTH
          volumeMounts:
            - name: ${GKE_REDIS_CLUSTER}-volume
              mountPath: /data
      volumes:
        - name: ${GKE_REDIS_CLUSTER}-volume
          gcePersistentDisk:
            pdName: ${GKE_REDIS_CLUSTER}-disk
            fsType: ext4