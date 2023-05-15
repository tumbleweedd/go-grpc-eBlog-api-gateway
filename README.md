 # go-grpc-eBlog-api-gateway
## Repositories
* https://github.com/tumbleweedd/go-grpc-eBlog-comment - Comment SVC (gRPC)
* https://github.com/tumbleweedd/go-grpc-eBlog-post - Post SVC (gRPC)
* https://github.com/tumbleweedd/go-grpc-eBlog-user - User SVC (gRPC)
* https://github.com/tumbleweedd/go-grpc-eBlog-auth - Authentication SVC (gRPC)
* https://github.com/tumbleweedd/go-grpc-eBlog-api-gateway - API Gateway (HTTP)

# ОПИСАНИЕ:
Сервис, предоставляющий работу с API интернет-блога, построенный на микросервисной архитектуре с использованием grpc-протокола.

# Описание конечных точек:

### AuthService
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------: |
| POST          |/auth/sign-up                          |Регистрация                       |
| POST          |/auth/sign-in                          |Вход в систему                    |

### UserService
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------------: |
| GET           |/api/users/userList                    |Получить список всех пользователей      |
| GET           |/api/users/me                          |Зайти в свой профиль                    |
| GET           |/api/users/:username/profile           |Зайти в профиль пользователя            |
| GET           |/api/users/:username/posts             |Просмотр постов пользователя            |
| POST          |/api/users/                            |Добавить пользователя (only for admin)  |
| PUT           |/api/users/:username                   |Обновить профиль (only for admin)       |
| DELETE        |/api/users/:username                   |Удалить профиль (only for admin)        |

### PostService
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------------: |
| GET           |/api/posts                             |Получить список всех постов             |
| GET           |/api/posts/:id                         |Получить пост по id                     |
| POST          |/api/posts/                            |Добавить новый пост                     |
| PUT           |/api/posts/:id                         |Обновить пост                           |
| DELETE        |/api/posts/:id                         |Удалить пост                            |

### CommentService
| Метод         | URL                 |Описание|
| ------------- |:-------------------------------------:| -------------------------------------: |
| POST          |/api/posts/:id/comments/               |Добавить комментарий к посту            |
| GET           |/api/posts/:id/comments/               |Получить комментарии к данному посту    |
| GET           |/api/posts/:id/:commentId              |Получить комментарий по его id          |
| DELETE        |/api/posts/:id/:commentId              |Удалить комментарий                     |
