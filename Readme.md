### Postgres
Установка Postgres

```bash
$ sudo apt-get install postgresql17-server
# Для инициализации схемы
$ sudo /etc/init.d/postgresql initdb
```

Для запуска службы Postgres
```bash
$ sudo service postgresql start
```

Для запуска интерактивной оболочки
```bash
$ sudo psql -U postgres
```

Создание пользователя
```bash
$ sudo createuser -U postgres --no-superuser --no-createdb --no-createrole --encrypted --pwprompt [ПОЛЬЗОВАТЕЛЬ]
```

Создание базы данных:
```bash
$ createdb -U postgres -O [ПОЛЬЗОВАТЕЛЬ] [БАЗА]
```

Просмотр всех пользователей:
```psql
\du 
```

Просмотр всех бд:
```psql
\l
```

Просмотр всех таблиц в базе:
```psql
\dt
```

Переключение на определённую базу
```psql
\c [database_name]
```

Изменение пароля пользователя:
```psql
ALTER USER [username] PASSWORD ['newpassword'];
```

Выдача админских прав пользователю на базу:
```psql
GRANT ALL PRIVILEGES ON DATABASE ["dbname"] to [username];
```

Создание роли для пользователя:
```psql
CREATE ROLE [username] WITH LOGIN PASSWORD ['passwordsequence'];
```
**Примечание**: В PostgreSQL "пользователь" = "роль с правами логина".

#### Назначение прав
Подключаться к БД, создавать новые объекты и временные таблицы

- **CONNECT** — позволяет пользователю подключаться к базе petrdb;

- **CREATE** — позволяет пользователю создавать новые объекты внутри базы данных (например, новые схемы);

- **TEMPORARY** — позволяет создавать временные таблицы в базе данных.

```psql
GRANT CREATE, CONNECT, TEMPORARY ON DATABASE [namedb] to [username];
```

#### Просмотр схемы, обращение к объектам в схеме:
```psql
GRANT USAGE ON SCHEMA public TO [username];
```

Чтобы пользователь базы мог ещё создавать схему:
```psql
GRANT CREATE, USAGE ON SCHEMA public TO [username];
```

USAGE на схему public означает:

- пользователь может видеть схему, может обращаться к объектам в схеме (но без права изменять объекты, пока отдельно не
выдано разрешение).

CREATE на схему public означает:
- пользователь может создавть схему.

**Примечание**: Без права USAGE на схему пользователь не сможет использовать даже существующие таблицы.

Права на все существующие таблицы:

```psql
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON ALL TABLES IN SCHEMA public TO [username];
```

Данная команда даёт права на все существующие таблицы в схеме public:

- **SELECT** — читать данные;

- **INSERT** — вставлять данные;

- **UPDATE** — изменять данные;

- **DELETE** — удалять записи;

- **REFERENCES** — создавать внешние ключи (FOREIGN KEY) на эти таблицы.

**Примечание**: ⚡ Но только для уже существующих таблиц. Новые таблицы пока не затрагиваются.

Права на все существующие последовательности:

```psql
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO [username];
```
Даёт права на все существующие последовательности (SEQUENCE, используются для генерации автоинкрементных ID) в схеме public:

- **USAGE** — право использовать последовательность (например, получать её текущее значение);

- **SELECT** — читать текущее значение последовательности;

- **UPDATE** — изменять значение (например, устанавливать nextval вручную).

**Примечание**: В PostgreSQL автонумерация часто завязана на последовательности.

#### Права поумолчанию на таблицы

```psql
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLES TO [username];
```

Устанавливает права по умолчанию:

- На все будущие таблицы, которые будут созданы в схеме public.

- Пользователь name сразу будет иметь права SELECT, INSERT, UPDATE, DELETE, REFERENCES на новые таблицы без необходимости
вручную давать права каждый раз.

Примечание: Это очень важно, иначе после создания новых таблиц права придётся назначать вручную.

#### Права по умолчанию для новых последовательностей

```psql
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO [username];
```

- Аналогично предыдущему пункту, но для новых последовательностей;

- Устанавливаются права на USAGE, SELECT, UPDATE для всех последовательностей, созданных в будущем в схеме public.


### Подключение пользователя к БД

```bash
sudo psql -U [username] \[databasename]
```