# Test 예시

이 폴더는 Test를 작성하는 예시를 보여주기 위한 폴더입니다. 참고해서 돌려보시고 본인 프로젝트에 맞게 수정하시면 됩니다. 저도 아직도 Mocking/Testing 이해가 어려워서 잘못된 부분이 있을 수도 있으니까 편하게 커멘트/수정 부탁드립니다.

## 테스트 라이브러리

- [Testify](https://github.com/stretchr/testify): assert, mock 용 라이브러리
- [mockery](https://vektra.github.io/mockery/latest/): mock 자동 생성 라이브러리

## example 구성

- `server.go`: `pkg`의 `Server`를 참고해 구현한 예시용 서버 및 CRUD 함수
- `request.go`: Request 구조체 값들을 하나의 string으로 만들어주는 함수(예시:`{"id": 1, "name": "test"}` -> `?id=1&name=test`)
- `main.go`: `server.go`의 서버를 실행하는 테스트용 `Main()` 함수. 실제로 테스트 해보려면 프로젝트 root 폴더의 `main.go` 에서 `example.Main()` 추가 후 실행.(현재 token/header 등이 설정되어 있지 않아서 404 리턴됨.)

### server_test.go 설명

- TestExampleSeverCreate: Mockery를 이용하여 `pkg`의 `Server`를 Mocking한 후 `Create()` 함수를 테스트하는 예시
- 총 두 가지의 테스트
  - `Success - 성공`: 정상적인 request & response 테스트
  - `Failed - 필수 파라미터 누락`: request 전송 시 필수 파라미터 누락 시 에러 테스트
- 각 Test code를 `t.Run()`으로 묶어서 테스트 실행 시 한 번에 실행되도록 함.
- `t.Run()`의 첫 번째 인자는 테스트 이름, 두 번째 인자는 테스트 함수
- `t.Run()`의 테스트 이름은 테스트 결과에 표시됨.
- error 값이 nil이 아닐 경우 `The error should be nil` 출력
- Response 값이 예상 값과 다를 경우 `The response should be equal` 출력

## `mocks`

- mock 자동 생성 라이브러리를 이용하여 생성한 mock 파일

## `.mockery.yml`

- mock 자동 생성 라이브러리 설정 파일
- example 패키지와 pkg ㅍ패키지의 모든 파일을 대상으로 mock 파일 생성

## 테스트 파일 구성법

- `brew install mockery`로 `mockery` 설치
- `.mockery.yml` 파일이 있는 root 폴더에서 명령어 실행

```bash
mockery
```

## 테스트 방법

```bash
go test -v ./...
```

### 테스트 커버리지 확인

```bash
go test -v ./... -cover
```
