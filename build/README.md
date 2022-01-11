# 확인 사항
### validate, mutate yaml 생성 Scripts 생성 전 수행 해야할 것 
### k8s-exmaple/AdmissionController/Go/src/webhook-example 에서 바이너리 빌드 후 이미지 생성 당시 생성된 crt 파일을 이용 해야 한다.

~~~bash
cp k8s-exmaple/AdmissionController/Go/src/webhook-example/key/ca.crt k8s-example/AdmissionController/deploy/validate

cp k8s-exmaple/AdmissionController/Go/src/webhook-example/key/ca.crt k8s-example/AdmissionController/deploy/mutate
~~~
