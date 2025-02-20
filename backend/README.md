# NutBooks/backend

[![Backend - API server Main](https://github.com/NutBooks/NutBooks/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/NutBooks/NutBooks/actions/workflows/go.yml)
[![Backend - API server Dev](https://github.com/NutBooks/NutBooks/actions/workflows/go.yml/badge.svg?branch=develop)](https://github.com/NutBooks/NutBooks/actions/workflows/go.yml)

## Env Variables

| 서비스 |         변수명         |   값(예시)   |                         비고                          |
|:---:|:-------------------:|:---------:|:---------------------------------------------------:|
| DB  |     MYSQL_HOST      | 127.0.0.1 | ipv4 또는 도커 컴포즈 서비스 명<br>localhost인 경우는 127.0.0.1 사용 |
| DB  |     MYSQL_PORT      |   3306    |               외부 접속 시 DB 컨테이너로 포트 매핑                |
| DB  |     MYSQL_USER      |   user    |                                                     |
| DB  |   MYSQL_PASSWORD    |   1234    |                                                     |
| DB  | MYSQL_ROOT_PASSWORD |   5678    |                    root 계정 비밀번호                     |
| DB  |   MYSQL_DATABASE    | nutbooks  |                      데이터베이스 이름                      |
| API |      API_PORT       |   8081    |                      API 서버 포트                      |

## API 서버 빌드

```bash
go build -o ./bin/main main.go
```

## DB 실행

### MySQL 컨테이너 실행

```bash
docker compose up -d db; docker compose logs -f --tail=1000 db
```

### Migrate DB

- DB 설치 후 최초 1회 실행

```bash
$ ./bin/main migrate
```

## 서버 실행

```bash
./bin/main run
```

### Swagger API 문서

- <http://localhost:8081/docs>

```mermaid
graph LR
  subgraph "EndpointLayer('/api/v1')"
		direction TB
		subgraph "/auth"
			E.SignUp["/signup/ POST"]
			E.LogIn["/login/ POST"]
		end
		subgraph "/bookmark"
			E.AddBookmark["/ POST"]
			E.GetBookMarkById["/{id}/ GET"]
			E.GetAllBookmarks["/ GET"]
		end
		subgraph "/user"
			E.AddUser["/ POST"]
			E.GetUserById["/{id}/ GET"]
			E.GetAllUsers["/ GET"]
		end
	end 
	subgraph ControllerLayer
		subgraph "Auth Handler"
			E.SignUp --> C.SignUpHandler[SignUpHandler]
			E.LogIn --> C.LogInHandler[LogInHandler]		
		end
		subgraph "Bookmark Handler"
			E.AddBookmark --> C.AddBookmarkHandler[AddBookmarkHandler]
			E.GetBookMarkById --> C.GetBookmarkHandler[GetBookmarkHandler]
			E.GetAllBookmarks --> C.GetAllBookmarksHandler[GetAllBookmarksHandler]
		end
		subgraph "User Handler"
			E.AddUser --> C.AddUserHandler[AddUserHandler]
			E.GetUserById --> C.GetUserByIdHandler[GetUserByIdHandler]
			E.GetAllUsers --> C.GetAllUsersHandler[GetAllUsersHandler]
		end
	end
	subgraph CRUDLayer
		subgraph UserCRUD
			C.SignUpHandler --> GetUserById
			C.SignUpHandler --> AddUser
			C.LogInHandler --> GetUserById
			C.AddUserHandler --> AddUser
			C.GetUserByIdHandler --> GetUserById
			C.GetAllUsersHandler --> GetUsers
		end
		subgraph BookmarkCRUD
			C.AddBookmarkHandler --> AddBookmark
			C.GetBookmarkHandler --> GetBookmarkById
			C.GetAllBookmarksHandler --> GetAllBookmarks
		end
	end
	UserCRUD --> DB[("`**DB**
											MySQL`")]
	BookmarkCRUD --> DB
```

## 개발 환경

### 테스팅

```bash
go test -v -cover ./... | tee TestResults
```

## 참고 자료

- Effective Go
    - <https://go.dev/doc/effective_go>
- Go 코딩 스타일: Uber의 Go 스타일 가이드 참고
    - <https://github.com/TangoEnSkai/uber-go-style-guide-kr>
- 주석 형식 가이드
    - <https://go.dev/doc/comment>
- DB behind traefik
    - <https://dev.to/kanzitelli/deploying-postgresql-and-redis-behind-traefik-in-the-cloud-5an2>
- traefik tls
    - <https://doc.traefik.io/traefik/user-guides/docker-compose/acme-tls/>
