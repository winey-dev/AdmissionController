# 선행 해야 할 것
~~~bash
:> key/create_tls.sh
~~~

# 생성된 인증 파일을 이동 
~~~bash
:> mv key/server.key ./
:> mv key/server.crt ./
:> mv key/ca.crt ./
~~~

# GOPATH 지정
$topath 진짜로 입력하시지 마시고 k8s-example 앞에 임의 경로가 존재한다면 그 경로를 적어야 합니다.
## csh
~~~bash
:> setenv GOPATH $topath/k8s-example/AdmissionController/Go
~~~
## bash
~~~bash
:> export GOTPATH=$topath/k8s-example/AdmissionController/Go
~~~

# GO BUILD 수행 
## go.mod 확인
## k8s.io/api 옛날 버전을 사용 하는 이유
최신 버전의 kubernetes pkg에서 k8s.io/api 참고하는 영역이 stagging/src 폴더로 Replace 되어있다. 
go get으로 다운로드 받은 kubernetes pkg에는 해당 소스가 없어 빌드가 안되어 옜날 pkg로 사용 했다. 
~~~bash
:> go build
~~~

# Docker 이미지 생성 
~~~bash
:> docker -t ${dockerhub.io/${UserName}.${HubRepoName}:${TAG} .
:> docker push ${dockerhub.io/${UserName}.${HubRepoName}:${TAG} 
~~~
