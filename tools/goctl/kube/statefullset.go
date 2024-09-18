package kube

var statefullsetTemplate = `apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: {{.Name}}
  name: {{.Name}}
  namespace: {{.Namespace}}
spec:
  podManagementPolicy: OrderedReady
  replicas: {{.MinReplicas}}
  revisionHistoryLimit: {{.Revisions}}
  selector:
    matchLabels:
      app: {{.Name}}
  serviceName: {{.Name}}
  template:
    metadata:
      labels:
        app: {{.Name}}
    spec:
      containers:
      - env:
        - name: SET_NAME
          value: {{.Name}}
        envFrom:
        - configMapRef:
            name: flyele-dev-config
        - secretRef:
            name: flyele-dev-secrets
        image: {{.Image}}
        imagePullPolicy: Always
        name: {{.Name}}
        ports:
        - containerPort: 8080
          name: http
          protocol: TCP
        - containerPort: 8084
          name: grpc
          protocol: TCP
        - containerPort: 5000
          name: headth
          protocol: TCP
        resources:
          requests:
            cpu: {{.RequestCpu}}m
            memory: {{.RequestMem}}Mi
          limits:
            cpu: {{.LimitCpu}}m
            memory: {{.LimitMem}}Mi
        volumeMounts:
        - name: timezone
          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai
  updateStrategy:
    type: RollingUpdate
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{.Name}}
  name: {{.Name}}-svc
  namespace: {{.Namespace}}
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: grpc
    port: 8084
    protocol: TCP
    targetPort: 8084
  selector:
    app: {{.Name}}
  type: ClusterIP

---

# http ingress
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/enable-cors: "true"
    nginx.ingress.kubernetes.io/rewrite-target: /$2
    nginx.ingress.kubernetes.io/auth-response-headers: x-auth-user,x-auth-platform,X-AUTH-Version,x-auth-device
    nginx.ingress.kubernetes.io/proxy-body-size: 10m
    kubernetes.io/ingress.class: nginx
{{- if .IsNotUserCenter}}
    nginx.ingress.kubernetes.io/auth-url: "http://usercenter-svc.{{.Namespace}}.svc.cluster.local:8080/v1/user/verify"
{{- end}}
  name: {{.Name}}-ingress
  namespace: {{.Namespace}}
spec:
  rules:
    - host: {{.Domain}}
      http:
        paths:
          - backend:
              service:
                name: {{.Name}}-svc
                port:
                  number: 8080 
            path: /{{.IngressPrefix}}(/|$)(.*)
            pathType: Prefix

    - host: {{getProfileDomain .Domain}}
      http:
        paths:
          - backend:
              service:
                name: {{.Name}}-svc
                port:
                  number: 5000 
            path: /{{.IngressPrefix}}(/|$)(.*)
            pathType: Prefix
{{- if .EnableTls}}
  tls:
    - hosts:
        - {{.Domain}}
      secretName: tls-secret
    - hosts:
        - {{getProfileDomain .Domain}}
      secretName: tls-secret
{{end}}
`
