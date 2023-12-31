## Общие требования
Необходимо реализовать полноценный проект ежедневника с инфраструктурой построенной на docker-compose.

Пользователю будет доступна HTTP API со следующими бизнес требованиями:
- Регистрация и аутентификация пользователей;
- Добавление, изменение, завершение, удаление и получение списка дел;
- Получать уведомления по задачам.

<br><br>

## Абстрактная схема взаимодействия с системой
Ниже представлена абстрактная бизнес-логика взаимодействия пользователя с системой:
1. Пользователь регистрируется в системе;
2. При необходимости пользователь может сменить пароль;
3. Пользователь добавляет задачу в планировщике с возможностью настроить несколько уведомлений;
4. Пользователь может посмотреть свои задачи на сегодня, на неделю, просроченные и бессрочные и завершенные.
5. По достижению необходимого времени, пользователь получает запланированное уведомление;
6. При необходимости пользователь может изменить задачу или перенести уведомления;
7. Пользователь может завершить задачу, которая попадёт в архив.

<br><br>

## Технологии
- **Nginx**. Базовый HTTP роутинг;
- **Golang**. Основной язык разработки;
- **Redis**. Кэш для хранения валидированных токенов;
- **RabbitMQ**. Брокер сообщений для внутреннего меж-сервисного взаимодействия;
- **PosrgreSQL**. База данных.

<br><br>

## Архитектуры проекта
Было принято решение выбрать Event-Driven микросервисную архитектуру без использования API Gateway.

```mermaid
flowchart LR
    user[User]
    nginx[NGINX]
    authService[Auth Service]
    notificationService[Notification Service]
    plannerService[Planner Service]
    rmq(RabbitMQ)
    redis(Redis)
    
    user --> nginx
    
    nginx --> authService
    nginx --> notificationService
    nginx --> plannerService

    authService --> redis
    authService <-.-> rmq
    plannerService <-.-> rmq
    notificationService <-.-> rmq
```
Для общения с клиентом, сервисы будут предоставлять http ручки. В качестве роутера выступает `nginx`.

Для общения между микросервисами используется брокер сообщений `RabbitMQ`.

<br><br>

## Auth Service
Сервис регистрации, аутентификации пользователя.
В этом сервисе также реализован механизм валидации JWT токена для других сервисов.

#### Регистрация
После успешной валидации отправленных данных, в базу сохраняется пользователь.
Пароль предварительно хэшируется\шифруется, а не хранится в базе в открытом виде.

#### Аутентификация
Отправляется запрос для получения JWT токена, сроком жизни в одни сутки.
По истечению срока, токен перестает действовать и требуется повторная авторизация.

#### Задачи со звездочкой
* Реализовать пару access-token / refresh-token, для обновления краткосрочного токена.
* Реализовать кэш ранее валидированных токенов, для ускорения работы. 

#### Архитектура
В качестве архитектуры, была выбрана вариация гексоганальной.

```mermaid
flowchart LR

    subgraph controller[Controller]
        direction TB
        rest(REST)
        rmqConsumer(RabbitMQ)
    end

    subgraph service[Service]
        direction TB
        userService(User)
        authService(Auth)
    end

    subgraph repository[Repository]
        direction TB
        psql[(PostgreSQL)]
    end

    subgraph adapter[Adapter]
        direction TB
        redis[(Redis)]
    end
    
    rest & rmqConsumer --> service
    service -.-> repository & adapter
```

- В качестве входных портов используется слой контроллеров.
    - REST для получения внешних запросов от клиента;
    - RabbitMQ для получения внутренних запросов от других микросервисов.

- В качестве инфраструктуры используется слой сервисов.

- В качестве выходных портов используется слой репозиториев и адаптеров.
    - Репозиторий PostgreSQL для хранения постоянного хранения данных
    - Адаптер Redis для кеширования валидированных ранее JWT токенов

#### REST эндпоинты
- `post` `/auth/register` Регистрация аккаунта с хешированием\шифрованием пароля
- `post` `/auth/change_password` Смена пароля
- `post` `/auth/login` Аутентификация пары логин-пароль и генерация JWT токена.

#### RabbitMQ подписки
- `auth.jwt.validation` Валидация токена JWT.



<br><br>



## Planner Service
Сервис работы с задачами внутри планировщика. Доступны следующие функции:
- Добавление задачи;
- Получение списка задач на сегодня;
- Получение списка задач на неделю;
- Получение списка задач без даты выполнения;
- Получение списка просроченных задач;
- Получение списка задач в архиве.

#### Работа с уведомлениями
При добавлении или изменении задачи, у пользователя есть возможность указать несколько уведомлений.
Задачи на создание, изменение или удаление уведомлений отправляются в RabbitMQ, который разбирает соответствующий сервис.


#### Задачи со звездочкой:
* Добавить возможность массово удалять или завершать задачи;
* Добавить возможность поручать задачи другому пользователю и просматривать список порученных задач.

#### Архитектура
В качестве архитектуры, была выбрана вариация гексоганальной.

```mermaid
flowchart LR

    subgraph controller[Controller]
        direction TB
        rest(REST)
    end

    subgraph service[Service]
        direction TB
        plannerService(Planner)
    end

    subgraph repository[Repository]
        direction TB
        psql[(PostgreSQL)]
    end

    subgraph adapter[Adapter]
        direction TB
        rmqPublisher(RabbitMQ)
    end
    
    rest --> service
    service -.-> repository & adapter
```
- В качестве входных портов используется слой контроллеров.
  Для этого сервиса предусмотрен только REST для получения внешних запросов от клиента.

- В качестве инфраструктуры используется слой сервисов.

- В качестве выходных портов используется слой репозиториев и адаптеров.
    - Репозиторий PostgreSQL для хранения постоянного хранения данных
    - Адаптер RabbitMQ для запроса данных из другого микросервиса

#### REST эндпоинты
- `post` `/planner` Добавление задачи
- `get` `/planner/{id}` Получение задачи по ID
- `get` `/planner/today` Получение списка задач на сегодня
- `get` `/planner/week` Получение списка задач на неделю
- `get` `/planner/undated` Получение списка бессрочных задач
- `get` `/planner/overdue` Получение списка просроченных задач
- `get` `/planner/archive` Получение задач в архиве
- `patch` `/planner/{id}` Изменение задачи
- `patch` `/planner/done/{id}` Завершение задачи и добавление её в архив
- `delete` `/planner/{id}` Удаление задачи

<br>

## Notification Service
Сервис для работы с уведомлениями.
Напрямую к нему можно обратиться только за получением списка доступных уведомлений. 
Логика получения уведомлений на стадии MVP ложится на сторону клиента.

Для изменения, удаления и пометки уведомления как прочитанного используется брокер сообщений.

#### Задачи со звездочкой:
* Реализовать возможность автоматической отправки уведомлений в Telegram;
* Реализовать websocket клиент для получения уведомлений.

#### Архитектура
В качестве архитектуры, была выбрана вариация гексоганальной.

```mermaid
flowchart LR

    subgraph controller[Controller]
        direction TB
        rest(REST)
        rmqConsumer(RabbitMQ)
    end

    subgraph service[Service]
        direction TB
        notificationService(Notification)
    end

    subgraph repository[Repository]
        direction TB
        psql[(PostgreSQL)]
    end

    subgraph adapter[Adapter]
        direction TB
        rmqPublisher(RabbitMQ)
    end
    
    rest & rmqConsumer --> service
    service -.-> repository & adapter
```
- В качестве входных портов используется слой контроллеров.
  Для этого сервиса предусмотренно два контроллера:
    - REST для получения внешних запросов от клиента;
    - RabbitMQ для получения внутренних запросов от других микросервисов.

- В качестве инфраструктуры используется слой сервисов.

- В качестве выходных портов используется слой репозиториев и адаптеров.
    - Репозиторий PostgreSQL для хранения постоянного хранения данных
    - Адаптер RabbitMQ для запроса данных из другого микросервиса

#### REST эндпоинты
- `get` `/notification` Получение списка непрочитанных уведомлений
- `patch` `/notification/check/{id}` Пометка уведомления как прочитанного

#### RabbitMQ подписки
- `notification.create` Создание уведомления
- `notification.update` Изменение уведомления
- `notification.delete` Удаление уведомления

