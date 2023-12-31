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