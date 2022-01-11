openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 100000 -out ca.crt -subj "/CN=admission_ca"

## req_distinguished_name 섹션과 alt_names 섹션에 입력하는 CN, DNS.1 값은 k8s 에서 생성할 svc의 명과 동일하게 설정 
## local 환경에서 테스트는 /etc/hosts 에 127.0.0.1   k8s 에서 생성할 svc 로 작성 
## ex) CN = webhook.yiaw.svc 일 경우 
## :> cat /etc/hosts
## 127.0.0.1 localhost webhook.yiaw.svc 
cat >server.conf <<EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no
[req_distinguished_name]
CN = webhook.yiaw.svc
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = webhook.yiaw.svc
EOF

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config server.conf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 100000 -extensions v3_req -extfile server.conf
