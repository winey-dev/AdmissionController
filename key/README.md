# OpenSSL 스크립트 내용 정리 
Scripts 내용 중 확인하고 수정이 필요한 사항 
**${ServiceName}.${Namespace}.svc** 해당 내용을 k8s에 Deploy 하려는 서비스명으로 수정해서 사용하시기 바랍니다. 

~~~
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 100000 -out ca.crt -subj "/CN=admission_ca"
cat >server.conf <<EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no
[req_distinguished_name]
CN = ${ServiceName}.${Namespace}.svc
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = ${ServiceName}.${Namespace}.svc
EOF
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config server.conf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 100000 -extensions v3_req -extfile server.conf
~~~

# 기타 수정 사항
https server Local 환경에서 테스트를 위해 DNS를 등록해준다.

/etc/hosts 파일에 127.0.0.1 15라인과 22라인에 적은 DNS 이름을 적어주면 된다.

/etc/hosts 파일에 접근하려면 root 권한이 필요하다.
~~~
:> vi /etc/hosts
127.0.0.1 ${ServiceName}.${Namespace}.svc
~~~
