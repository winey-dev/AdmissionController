## ca.crt의 파일위치와 해당 스크립트를 수행하는 경로가 같으면 ca.crt 
## ca.crt의 파일위치가 해당 스크립트를 수행하는 경로와 다르면 경로를 입력 
## ex ) ca.crt 위치가 /home/key 이면 $(car /home/key/ca.crt | base64 | tr -d '\n') 로 입력
cat > validate.yaml << EOF
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: yiaw-validator
webhooks:
  - name: rev.validation.yiaw.io
    namespaceSelector:
      matchExpressions:
      - key: yiaw-org-webhook
        operator: In
        values:
        - "true"
    admissionReviewVersions:
    - v1beta1
    - v1
    clientConfig:
      caBundle:  $(cat ca.crt | base64 | tr -d '\n')
      service:
        name: webhook
        namespace: yiaw
        path: /validate
        port: 443
    rules:
    - apiGroups:
      - '*'
      apiVersions:
      - v1
      operations:
      - CREATE
      resources:
      - pods
      scope: '*'
    sideEffects: None
EOF
