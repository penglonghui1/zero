package kube

var k3sDeploymentTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}-svc
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-svc
spec:
  replicas: {{.Replicas}}
  revisionHistoryLimit: {{.Revisions}}
  selector:
    matchLabels:
      app: {{.Name}}-svc
  template:
    metadata:
      labels:
        app: {{.Name}}-svc
    spec:{{- if .ServiceAccount}}
      serviceAccountName: {{.ServiceAccount}}{{end}}
      restartPolicy: Always
      containers:
      - envFrom:
        - configMapRef:
            name: flyele-dev-config
        - secretRef:
            name: flyele-dev-secrets
        name: {{.Name}}-svc
        image: {{.Image}}
        lifecycle:
          preStop:
            exec:
              command: ["sh","-c","sleep 5"]
        ports:
        - containerPort: 8080
          name: http
          protocol: TCP
        - containerPort: 5000
          name: health
          protocol: TCP
        - containerPort: 8084
          name: grpc
          protocol: TCP
{{- if .EnableWebsocket}}
        - containerPort: 8081
          name: websocket
          protocol: TCP
{{end}}
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

---

apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}-svc
  namespace: {{.Namespace}}
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: healthz
    port: 5000
    protocol: TCP
    targetPort: 5000
  - name: grpc
    port: 8084
    protocol: TCP
    targetPort: 8084
  selector:
    app: {{.Name}}-svc

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{.Name}}-hpa-c
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-hpa-c
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{.Name}}-svc
  minReplicas: {{.MinReplicas}}
  maxReplicas: {{.MaxReplicas}}
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{.Name}}-hpa-m
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-hpa-m
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{.Name}}-svc
  minReplicas: {{.MinReplicas}}
  maxReplicas: {{.MaxReplicas}}
  metrics:
  - type: Resource
    resource:
      name: memory
      targetAverageUtilization: 80

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

var deploymentTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}-svc
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-svc
spec:
  replicas: {{.Replicas}}
  revisionHistoryLimit: {{.Revisions}}
  selector:
    matchLabels:
      app: {{.Name}}-svc
  template:
    metadata:
      labels:
        app: {{.Name}}-svc
    spec:{{- if .ServiceAccount}}
      serviceAccountName: {{.ServiceAccount}}{{end}}
      restartPolicy: Always
      containers:
      - env:
        - name: RELEASE_MODE
          value: {{.Env}}
        envFrom:
        - configMapRef:
            name: flyele-dev-config
        - secretRef:
            name: flyele-dev-secrets
        name: {{.Name}}-svc
        image: {{.Image}}
        lifecycle:
          preStop:
            exec:
              command: ["sh","-c","sleep 5"]
        ports:
        - containerPort: 8080
          name: http
          protocol: TCP
        - containerPort: 5000
          name: health
          protocol: TCP
        - containerPort: 8084
          name: grpc
          protocol: TCP
{{- if .EnableWebsocket}}
        - containerPort: 8081
          name: websocket
          protocol: TCP
{{end}}
        readinessProbe:
          tcpSocket:
            port: 5000
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 5000
          initialDelaySeconds: 15
          periodSeconds: 20
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

---

apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}-svc
  namespace: {{.Namespace}}
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: healthz
    port: 5000
    protocol: TCP
    targetPort: 5000
  - name: grpc
    port: 8084
    protocol: TCP
    targetPort: 8084
  selector:
    app: {{.Name}}-svc

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{.Name}}-hpa-c
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-hpa-c
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{.Name}}-svc
  minReplicas: {{.MinReplicas}}
  maxReplicas: {{.MaxReplicas}}
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{.Name}}-hpa-m
  namespace: {{.Namespace}}
  labels:
    app: {{.Name}}-hpa-m
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{.Name}}-svc
  minReplicas: {{.MinReplicas}}
  maxReplicas: {{.MaxReplicas}}
  metrics:
  - type: Resource
    resource:
      name: memory
      targetAverageUtilization: 80

---

# http ingress
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/enable-cors: "true"
    nginx.ingress.kubernetes.io/rewrite-target: /$2
  name: {{.Name}}-ingress
  namespace: {{.Namespace}}
spec:
  rules:
    - host: {{.Domain}}
      http:
        paths:
          - backend:
              serviceName: {{.Name}}-svc
              servicePort: 8080
            path: /{{.IngressPrefix}}(/|$)(.*)
{{- if .EnableGRPCIngress}}
    - host: {{.Name}}-svc.flyele.vip
      http:
        paths:
          - backend:
              serviceName: {{.Name}}-svc
              servicePort: 8084
            path: /
{{end}}
    - host: {{getProfileDomain .Domain}}
      http:
        paths:
          - backend:
              serviceName: {{.Name}}-svc
              servicePort: 5000
            path: /{{.IngressPrefix}}(/|$)(.*)
{{- if .EnableTls}}
  tls:
    - hosts:
        - {{.Domain}}
      secretName: tls-secret
    - hosts:
        - {{getProfileDomain .Domain}}
      secretName: tls-secret
{{- if .EnableGRPCIngress}}
    - hosts:
        - {{.Name}}-svc.flyele.vip
      secretName: tls-secret
{{end}}
{{end}}
`
