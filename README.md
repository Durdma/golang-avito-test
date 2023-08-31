# Тестовое задание для стажёра Backend
# Сервис динамического сегментирования пользователей

### Проблема:

В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие.
На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.

### Задача:

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

**Сценарии использования:**

Хотим провести несколько экспериментов и протестировать новый функционал Авито:
- Голосовые сообщения в чатах
- Новые услуги продвижения
- Скидка 30% на услуги продвижения
- Скидка 50% на услуги продвижения

> Кто из пользователей в какой эксперимент попадет будет решать большой отдел аналитики, а мы лишь дадим им возможность для таких тестов.

Допустим аналитики создали сегменты:
- AVITO_VOICE_MESSAGES
- AVITO_PERFORMANCE_VAS
- AVITO_DISCOUNT_30
- AVITO_DISCOUNT_50

и добавили созданные сегменты нескольким пользователям:

| Пользователь | Сегменты которым он принадлежит |
| --- | --- |
| 1000 | [AVITO_VOICE_MESSAGES, AVITO_PERFORMANCE_VAS, AVITO_DISCOUNT_30, …] |
| 1002 | [AVITO_VOICE_MESSAGES, AVITO_DISCOUNT_50, …] |
| 1004 | нет сегментов |

> Формат хранения данных в базе данных не ограничен - можно выбрать любой удобный

**Получили следующие данные:**

- Пользователь 1000 состоит в 3 сегментах: AVITO_VOICE_MESSAGES, AVITO_PERFORMANCE_VAS, AVITO_DISCOUNT_30

- Пользователь 1002 состоит в 2 сегментах: AVITO_VOICE_MESSAGES, AVITO_DISCOUNT_50

- Пользователь 1004 не состоит ни в одном из сегментов 

Теперь мы хотим через API сервиса по user_id получать список сегментов в которых он состоит.

## Требования и детали по заданию

**Технические требования:**

1. Сервис должен предоставлять HTTP API с форматом JSON как при отправке запроса, так и при получении результата.
2. Язык разработки: Golang.
3. Фреймворки и библиотеки можно использовать любые.
4. Реляционная СУБД: MySQL или PostgreSQL.
5. Использование docker и docker-compose для поднятия и развертывания dev-среды.
6. Весь код должен быть выложен на Github/Gitlab с Readme файлом с инструкцией по запуску и примерами запросов/ответов (можно просто описать в Readme методы, можно через Postman, можно в Readme curl запросы скопировать, и так далее).
7. Если есть потребность в асинхронных сценариях, то использование любых систем очередей - допускается.
8. При возникновении вопросов по ТЗ оставляем принятие решения за кандидатом (в таком случае в Readme файле к проекту должен быть указан список вопросов, с которыми кандидат столкнулся и каким образом он их решил).
9. Разработка интерфейса в браузере НЕ ТРЕБУЕТСЯ. Взаимодействие с API предполагается посредством запросов из кода другого сервиса. Для тестирования можно использовать любой удобный инструмент. Например: в терминале через curl или Postman.

**Будет плюсом:**

1. Покрытие кода тестами.
2. [Swagger](https://swagger.io/solutions/api-design/) файл для вашего API.

**Основное задание (минимум):**

1. Метод создания сегмента. Принимает slug (название) сегмента. 
2. Метод удаления сегмента. Принимает slug (название) сегмента. 
3. Метод добавления пользователя в сегмент. Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.
4. Метод получения активных сегментов пользователя. Принимает на вход id пользователя.

**Детали по заданию:**

1. По умолчанию сервис не содержит в себе никаких данных о сегментах и пользователях (пустая табличка в БД). Данные появляются при создании сегментов и добавлении их пользователям.
2. Валидацию данных и обработку ошибок оставляем на ваше усмотрение.
3. Список полей к методам не фиксированный. Перечислен лишь необходимый минимум. В рамках выполнения доп. заданий возможны дополнительные поля.
4. Механизм миграции не нужен. Достаточно предоставить конечный SQL файл с созданием всех необходимых таблиц в БД.
5. Сегменты пользователя очень важны - на этих данных в будущем строится аналитика о том насколько продукт востребован. Поэтому нужно следить за тем чтобы сегменты не терялись (не перетирались) и не добавлялись лишним пользователям.
6. Мы можем добавлять сегменты пользователю динамически (он уже состоит в нескольких, мы можем добавить еще несколько, не перетирая существующие).
7. В методе получения сегментов пользователя мы должны получить АКТУАЛЬНУЮ информацию о сегментах пользователя с задержкой не более 1 минуты после добавления сегмента.

## Опциональные задания

Далее перечислены дополнительные задания. 

Они не являются обязательными, но их выполнение даст существенный плюс перед другими кандидатами.

Можно выбрать несколько из представленных

*Доп. задание 1:*

Иногда пользователи приходят в поддержку и спрашивают почему у них пропал/появился какой-то новый функционал. Нужно иметь возможность посмотреть когда точно пользователь попал в конкретный сегмент. 

Задача: реализовать сохранение истории попадания/выбывания пользователя из сегмента с возможностью получения отчета по пользователю за определенный период. На вход: год-месяц. На выходе ссылка на CSV файл.

Пример отчета:

идентификатор пользователя 1;сегмент1;операция (добавление/удаление);дата и время

идентификатор пользователя 1;сегмент2;операция (добавление/удаление);дата и время

идентификатор пользователя 2;сегмент3;операция (добавление/удаление);дата и время

*Доп. задание 2:*

Бывают ситуации когда нам нужно добавить пользователя в эксперимент на ограниченный срок. Например выдать скидку всего на 2 дня. 

Задача: реализовать возможность задавать TTL (время автоматического удаления пользователя из сегмента)

Пример: Хотим чтобы пользователь попал в сегмент на 2 дня - для этого в метод добавления сегментов пользователю передаём время удаления пользователя из сегмента отдельным полем

*Доп. задание 3:*

Мы хотим добавлять пользователя в сегмент не в ручную, а автоматически. В сегмент будет попадать заданный процент пользователей.

Задача: в методе создания сегмента, добавить опцию указания процента пользователей, которые будут попадать в сегмент автоматически. В методе получения сегментов пользователя, добавленный сегмент должен отдаваться у заданного процента пользователей.

Пример: создали сегмент AVITO_VOICE_MESSAGES и указали что 10% пользователей будут попадать в него автоматически. Пользователь 1000 попал в этот сегмент автоматически. При запросе сегментов пользователя 1000, сегмент AVITO_VOICE_MESSAGES должен отдаваться всегда.

## Решение
В ходе работы были разработаны следующие end-points
- "http://localhost:8080/user/" - http метод POST - для занесения пользователя в БД и добавление/удаления его в сегменты
- "http://localhost:8080/user/:id" - http метод GET - для получения информации о пользователе и активных сегментах
- "http://localhost:8080/user/:id/history" - http метод GET - для получения csv-файла с историей операций пользователя
- "http://localhost:8080/slug/" - http метод POST - метод для добавления сегмента в БД
- "http://localhost:8080/slug/:name" - http метод DELETE - метод для удаления сегмента из БД

### "http://localhost:8080/user/" - http метод POST
End point для обработки запросов на добавление пользователя в БД и добавление/удаление его в сегменты.
Структура json:
```json
{
    "user_id": 4,
    "slugs_list_to_add": [{"slug_name": "TEST_SLUG_1", "days": "4"}],
    "slugs_list_to_del": ["TEST_SLUG_DEL"]
}
```
Поле user_id(int) - обязательно, остальные опциональны.
Поле slugs_list_to_add - представляет собой массив имен сегментов для добавления с возможностью указать на сколько дней добавить пользователя в сегмент.
Поле days(int)
Поле slugs_list_to_del - представляет собой собой массив имен сегментов, из которых пользователь должен быть удален.

### "http://localhost:8080/user/:id" - http метод GET - для получения информации о пользователе и активных сегментах
End point для получения информации об активных сегментах пользователя. Id пользователя передается через параметр-url.
Структура json-response
```json
{
    "user_id": 4,
    "ActiveSlugs": [
        {
            "slug_id": 1,
            "slug_name": "TEST_SLUG_1",
            "created_at": "2023-08-31 22:57:31",
            "updated_at": "2023-08-31 22:57:31"
        }
    ],
    "created_at": "2023-08-31 22:57:35",
    "updated_at": "2023-08-31 22:57:35"
}
```
Поле ActiveSlugs является массивом сегментов, в которых пользователь находится в данный момент.

### "http://localhost:8080/user/:id/history" - http метод GET - для получения csv-файла с историей операций пользователя
End point для получения информации об операциях с сегментами. Id пользователя передается через параметр-url. Ответ от сервера представляет собой csv файл, в котором хранятся все операции пользователя.

### "http://localhost:8080/slug/" - http метод POST - метод для добавления сегмента в БД
End point представляет возможность добавить новый сегмент в БД.
Структура JSON:
```json
{
    "slug_name": "TEST_SLUG_1"
}
```

###  "http://localhost:8080/slug/:name" - http метод DELETE - метод для удаления сегмента из БД
End point для удаления сегмента из базы данных. Организовано "мягкое" удаление, то есть в БД ставится метка, что запись удалена и при поиске сегментов, данный сегмент не попадает в выборку или в поле активных сегментов пользователей.

### Для запуска
### requirements
- Posgresql
- GoLang 1.20
В любой СУБД запустить скрипт для создания БД из dbscripts public_schema.sql
Из корневого каталога приложения войти в командную строку и ввести
```bash
go run main.go
```
