<div align="center">

# Simple FTP Client

[![Go](https://img.shields.io/badge/Go-1.19%2B-brightgreen?logo=go&logoColor=white)](https://go.dev/)
[![FTP](https://img.shields.io/badge/FTP-Supported-brightgreen?logo=files&logoColor=white)](https://en.wikipedia.org/wiki/File_Transfer_Protocol)
[![Go FTP](https://img.shields.io/badge/Go%20FTP-jlaffaye%2Fftp-brightgreen?logo=go&logoColor=white)](https://github.com/jlaffaye/ftp)
[![Base64](https://img.shields.io/badge/Base64-Encoding-brightgreen?logo=security&logoColor=white)](https://datatracker.ietf.org/doc/html/rfc4648)
[![Docker](https://img.shields.io/badge/Docker-Ready-brightgreen?logo=docker&logoColor=white)](https://www.docker.com/)

*Микросервис для загрузки файлов с FTP-серверов по логину/паролю или анонимно, реализованный на Go.*

</div> 

### 🚀 Возможности

- 🔐 **Поддержка авторизации**: Загрузка файлов по логину и паролю.
- 👤 **Анонимный доступ**: Возможность скачивать файлы без логина/пароля.
- 💾 **Base64 кодирование**: Файлы возвращаются в JSON с Base64-содержимым.
- 🌐 **REST API**: Удобный интерфейс для скачивания через HTTP.
- 🐳 **Простота развертывания**: Быстрый запуск с Docker и Docker Compose.


### ⚙️ Архитектура


```text
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│  Пользователь/  ├──── > │    REST API     ├──── > │   FTP Сервис    ├──── > │   FTP Сервер    │
│     Система     │ < ────┤   (Gin-Gonic)   │ < ────┤ (jlaffaye/ftp)  │ < ────┤  (binary/ASCII) │
└─────────────────┘       └─────────────────┘       └─────────────────┘       └─────────────────┘   
```

<div align="center">

## 📦 Установка
</div>

### 1. Клонирование репозитория

```bash
git clone https://github.com/Khrllw/FTP_client.git
cd FTP_client
```

### 2. Конфигурация приложения

Откройте файл .env и при необходимости измените его

```dotenv
# App
APP_PORT=8080
GIN_MODE=debug

# Logger
LOGGER_ENABLE=true
LOGGER_LOGS_DIR=./logs
LOGGER_LOG_LEVEL=DEBUG
LOGGER_SAVING_DAYS=7
```

### 3. Запуск приложения

```
# Windows
./build/windows_ftp_client.exe

# Linux
./build/linux_ftp_client

# MacOS
./build/macos_ftp_client

# Golang
go run cmd/app/main.go
```


<div align="center">

## 🗂️ Структура проекта
</div>

```
Ftp_client/
│ 
├── cmd/app/                     
│       └── 📄 main.go                 # Главная точка входа приложения
├── 📁 docs/                           # Документация проекта
├── internal/
│   ├── 📁 app/                        # Сборка и запуск приложения с помощью Fx для DI
│   ├── 📁 config/                     # Логика загрузки конфигурации из .env
│   ├── adapters/ 
│   │   └── 📁 handlers/               # Обработчики HTTP-запросов (слой API на Gin)
│   ├── 📁 domain/                     # DTO, Request/Response модели  
    ├── 📁 interfaces/                 # Go-интерфейсы для всех слоев (контракты)       
│   ├── middleware/
│   │   ├── 📁 logging/                # Логирование
│   │   └── 📁 swagger/                # Swagger/OpenAPI документация
│   ├── services/ 
│   │   └── 📁 ftp_service/            # FTP-клиент
│   └── 📁 usecases/                   # Бизнес-логика
├── logs/ 
│   └── 📄 { _date_ }.log              # Логи приложения по дням
├── pkg/
│   ├── 📁 client/                     # Клиентская библиотека для API
│   └── 📁 errors/                     # Пользовательские ошибки 
├── tools/build/
│       └── 📄 build.go                # Скрипт для сборки исполняемых файлов
├── 📁 build/                          # Папка с готовыми исполняемыми файлами
├── 📄 .env                            # Файл конфигурации
├── 📄 docker-compose.yml              # Файл для запуска Kafka и Kafka-UI
├── 📄 LICENSE
└── 📄 README.md
```

## 🆘 Поддержка

- 🐛 [Создайте issue](https://github.com/Khrllw/FTP_service/issues)
- 📧 Напишите на email: khrllw@gmail.com

## 📝 Лицензия

Проект распространяется под [лицензией MIT](LICENSE)

Copyright (c) 2025 khrllw
