# Практическое задание № 5 Борисов Денис Александрович ЭФМО-01-25
Тема: Реализация HTTPS (TLS-сертификаты). Защита от SQL-инъекций

Задачи практики:
1.	Изучить назначение HTTPS и роль TLS в защите сетевого взаимодействия.
2.	Понять, чем HTTP отличается от HTTPS на уровне backend-приложения.
3.	Освоить подключение TLS-сертификата к серверу на Go.
4.	Научиться запускать HTTP-сервис в защищённом режиме через HTTPS.
5.	Изучить природу SQL-инъекций и причины их возникновения.
6.	Понять, почему конкатенация SQL-строк с пользовательским вводом опасна.
7.	Освоить безопасную работу с SQL-запросами через параметризованные запросы и prepared statements.
8.	Реализовать учебный пример backend-приложения с HTTPS и безопасным доступом к данным.


Выполнение практического задания.

1.	Структура проекта

<img width="223" height="665" alt="image" src="https://github.com/user-attachments/assets/22c6efc6-c4d8-40c8-8373-5de006201d8f" />

2.	Установка зависимостей Go и инструментов генерации.

Установка зависимостей Go

<img width="737" height="97" alt="1" src="https://github.com/user-attachments/assets/311fefe6-b328-4a7a-ba79-e9d21c7f7939" />


3.	Создание файла docker-compose.yml.

docker-compose.yml, предназначен для развертывания БД Postgres 

<img width="513" height="352" alt="2" src="https://github.com/user-attachments/assets/cc2d8bc7-af61-40db-a5e0-0fbda9c7f772" />

Развертывание сервисов

<img width="867" height="161" alt="3" src="https://github.com/user-attachments/assets/fb261420-3faa-4514-938b-12f31d0c0d0d" />

Рабочие контейнеры в Docker

<img width="1567" height="296" alt="4" src="https://github.com/user-attachments/assets/f93cb82a-a6f6-4c48-b9ee-4b6ff191a7be" />

Создание и заполнение БД

<img width="887" height="293" alt="5" src="https://github.com/user-attachments/assets/2f9454ed-bc20-4a32-9ee8-49e84daf3b87" />

4. Создание SSH ключа

<img width="1462" height="451" alt="6" src="https://github.com/user-attachments/assets/503929da-9919-43a1-98ae-a35445634884" />

Созданные ключи

![Uploading 7.png…]()

5. Тестирование

Запуск сервера

<img width="723" height="93" alt="9" src="https://github.com/user-attachments/assets/c3f8dbf0-d5bd-48bb-b93f-7b8cc9b3f675" />

Выполнение тестов

1. Вывод первого студенты

<img width="957" height="552" alt="11" src="https://github.com/user-attachments/assets/9bb4b668-d790-4923-b83d-993f502b488c" />

2. Ошибка несуществующий студент

<img width="953" height="607" alt="12" src="https://github.com/user-attachments/assets/275ee7e6-1231-421f-9173-eec27f4c3284" />

3. Попытка иньекции

<img width="952" height="592" alt="13" src="https://github.com/user-attachments/assets/67e3fa63-ef11-4353-a5d8-befa02a219a3" />

4. Поиск по почте студента

<img width="952" height="687" alt="14" src="https://github.com/user-attachments/assets/00706056-5af1-48cb-abaa-5326df8dd2b5" />

5. Переадресация

<img width="961" height="650" alt="15" src="https://github.com/user-attachments/assets/d5274f94-6201-47d3-91ee-75a97af21dab" />
