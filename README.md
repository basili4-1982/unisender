```markdown
# Unisender Go API Client

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/basili4-1982/unisender-go-api)](https://goreportcard.com/report/github.com/basili4-1982/unisender-go-api)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Go клиент для работы с API Unisender Go (транзакционные email-рассылки).

## 📦 Установка

```bash
go get github.com/basili4-1982/unisender
```

## 🚀 Быстрый старт

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/basili4-1982/unisender-go-api"
    "github.com/basili4-1982/unisender-go-api/models"
)

func main() {
    // Создание клиента
    client, err := unisender.NewClient(
        "your-api-key",
        unisender.WithTimeout(30*time.Second),
        unisender.WithDebug(true),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    // Отправка email
    resp, err := client.Emails().Send(context.Background(), &models.SendEmailRequest{
        Message: models.EmailMessage{
            Recipients: []models.Recipient{
                {Email: "user@example.com"},
            },
            Subject:   "Привет из Go SDK!",
            FromEmail: "sender@example.com",
            FromName:  "Unisender Go",
            Body: models.EmailBody{
                HTML: "<h1>Hello, {{to_name}}!</h1>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Email sent! Task ID: %s\n", resp.Result.TaskID)
}
```

## 📚 Возможности

- ✅ Полная поддержка всех методов API Unisender Go
- ✅ Типизированные модели запросов и ответов
- ✅ Валидация входных данных
- ✅ Автоматическая обработка ошибок
- ✅ Поддержка retry при временных сбоях
- ✅ Контекст для отмены операций
- ✅ Удобные опции конфигурации
- ✅ Поддержка пагинации

## 📖 Документация API

### Ресурсы

| Ресурс | Описание | Методы |
|--------|----------|--------|
| **Emails** | Отправка email и подписка | `Send`, `Subscribe` |
| **Templates** | Управление шаблонами | `Set`, `Get`, `List`, `Delete` |
| **Webhooks** | Управление вебхуками | `Set`, `Get`, `List`, `Delete` |
| **Projects** | Управление проектами | `Create`, `Update`, `List` |
| **Suppression** | Стоп-лист | `Set`, `List` |
| **Tags** | Управление тегами | `List`, `Delete` |
| **Domains** | Управление доменами | `Delete` |
| **Events** | Выгрузка событий | `Create`, `List` |
| **System** | Системная информация | `Info` |

## 💡 Примеры использования

### Отправка email с подстановками

```go
resp, err := client.Emails().Send(ctx, &models.SendEmailRequest{
    Message: models.EmailMessage{
        Recipients: []models.Recipient{
            {
                Email: "user@example.com",
                Substitutions: map[string]string{
                    "name": "Иван",
                    "order_id": "12345",
                },
            },
        },
        Subject:   "Ваш заказ #{{order_id}}",
        FromEmail: "shop@example.com",
        FromName:  "Интернет-магазин",
        Body: models.EmailBody{
            HTML: "<h2>Здравствуйте, {{name}}!</h2><p>Ваш заказ #{{order_id}} оформлен.</p>",
        },
        TrackLinks: 1,
        TrackRead:  1,
    },
})
```

### Отправка email с вложениями

```go
import "encoding/base64"

// Читаем файл и кодируем в base64
fileContent, _ := os.ReadFile("document.pdf")
encodedContent := base64.StdEncoding.EncodeToString(fileContent)

resp, err := client.Emails().Send(ctx, &models.SendEmailRequest{
    Message: models.EmailMessage{
        Recipients: []models.Recipient{
            {Email: "user@example.com"},
        },
        Subject:   "Документы",
        FromEmail: "docs@example.com",
        Body: models.EmailBody{
            HTML: "<p>Ваши документы во вложении</p>",
        },
        Attachments: []models.Attachment{
            {
                Type:    "file",
                Name:    "document.pdf",
                Content: encodedContent,
            },
        },
    },
})
```

### Использование шаблонов

```go
// Создание шаблона
template, err := client.Templates().Set(ctx, &models.SetTemplateRequest{
    Name:      "Welcome Email",
    Subject:   "Welcome to our service!",
    FromEmail: "welcome@example.com",
    Type:      "code",
    Body: models.TemplateBody{
        HTML: "<h1>Hello {{name}}!</h1><p>Welcome aboard!</p>",
    },
})

// Отправка с использованием шаблона
resp, err := client.Emails().Send(ctx, &models.SendEmailRequest{
    Message: models.EmailMessage{
        Recipients: []models.Recipient{
            {
                Email: "user@example.com",
                Substitutions: map[string]string{
                    "name": "John",
                },
            },
        },
        TemplateID: template.ID,
        Subject:    "Welcome!",
        FromEmail:  "welcome@example.com",
    },
})
```

### Получение системной информации

```go
info, err := client.System().Info(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("API Version: %s\n", info.Version)
fmt.Printf("Service: %s (%s)\n", info.Service.Name, info.Service.Status)
fmt.Printf("Max Recipients per request: %d\n", info.Limits.MaxRecipients)
```

### Управление стоп-листом

```go
// Добавление в стоп-лист
_, err := client.Suppression().Set(ctx, &models.SetSuppressionRequest{
    Items: []models.SuppressionItem{
        {
            Value:  "spammer@example.com",
            Type:   "email",
            Reason: "Spam complaint",
        },
        {
            Value:  "spam-domain.com",
            Type:   "domain",
            Reason: "Known spam domain",
        },
    },
})

// Получение списка стоп-листа
suppressions, err := client.Suppression().List(ctx, &models.ListOptions{
    Limit:  50,
    Offset: 0,
})
```

### Работа с вебхуками

```go
// Создание вебхука
webhook, err := client.Webhooks().Set(ctx, &models.SetWebhookRequest{
    Name: "Email Events",
    URL:  "https://example.com/webhook",
    Events: []models.WebhookEventInput{
        {Type: models.WebhookEventEmailSent},
        {Type: models.WebhookEventEmailOpened},
        {Type: models.WebhookEventEmailClicked},
    },
    IsActive: boolPtr(true),
})

// Получение списка вебхуков
webhooks, err := client.Webhooks().List(ctx)
```

### Управление проектами

```go
// Создание проекта
project, err := client.Projects().Create(ctx, &models.CreateProjectRequest{
    Name:            "My New Project",
    DefaultLanguage: "ru",
    Timezone:        "Europe/Moscow",
})

// Обновление проекта
updated, err := client.Projects().Update(ctx, project.ID, &models.UpdateProjectRequest{
    Name: "Updated Project Name",
})
```

### Выгрузка событий

```go
// Создание выгрузки
dump, err := client.Events().Create(ctx, &models.CreateEventDumpRequest{
    EventTypes: []string{
        models.EventTypeEmailSent,
        models.EventTypeEmailOpened,
        models.EventTypeEmailClicked,
    },
    DateFrom: time.Now().AddDate(0, -1, 0),
    DateTo:   time.Now(),
    Format:   models.ExportFormatJSON,
})

// Получение списка выгрузок
dumps, err := client.Events().List(ctx)
```

## 🔧 Конфигурация

### Опции клиента

```go
client, err := unisender.NewClient(
    "your-api-key",
    
    // Базовый URL (опционально)
    unisender.WithBaseURL("https://goapi.unisender.ru/ru/transactional/api/v1"),
    
    // Таймаут запросов
    unisender.WithTimeout(30*time.Second),
    
    // Настройка retry
    unisender.WithRetry(3, 2*time.Second),
    
    // Включение debug режима
    unisender.WithDebug(true),
    
    // Кастомный User-Agent
    unisender.WithUserAgent("my-app/1.0"),
)
```

## 🐛 Обработка ошибок

Клиент предоставляет типизированные ошибки для удобной обработки:

```go
resp, err := client.Emails().Send(ctx, req)
if err != nil {
    switch e := err.(type) {
    case *unisender.ValidationErrors:
        fmt.Println("Validation errors:")
        for _, ve := range e.Errors {
            fmt.Printf("  - %s: %s\n", ve.Field, ve.Message)
        }
    case *unisender.AuthenticationError:
        fmt.Println("Auth error:", e.Message)
    case *unisender.RateLimitError:
        fmt.Printf("Rate limit! Retry after %d seconds\n", e.RetryAfter)
    case *unisender.NotFoundError:
        fmt.Println("Not found:", e.Message)
    case *unisender.InternalError:
        fmt.Println("Internal server error:", e.Message)
    default:
        fmt.Println("Unknown error:", err)
    }
    return
}
```

## 🧪 Тестирование

```bash
# Запуск всех тестов
go test -v ./...

# Запуск с покрытием
go test -cover ./...

# Запуск конкретного теста
go test -v -run TestSendEmail ./resources
```

## 📁 Структура проекта

```
unisender-go-api/
├── client.go          # Основной клиент
├── config.go          # Конфигурация и опции
├── errors.go          # Типы ошибок
├── interfaces.go      # Интерфейсы ресурсов
├── models/            # Модели данных
│   ├── common.go
│   ├── email.go
│   ├── template.go
│   ├── webhook.go
│   ├── project.go
│   ├── suppression.go
│   ├── tag.go
│   ├── event.go
│   ├── system.go
│   ├── domain.go
│   ├── validation.go
│   └── constants.go
├── resources/         # Реализации ресурсов
│   ├── base.go
│   ├── emails.go
│   ├── templates.go
│   ├── webhooks.go
│   ├── projects.go
│   ├── suppression.go
│   ├── tags.go
│   ├── domains.go
│   ├── events.go
│   └── system.go
├── examples/          # Примеры использования
│   ├── basic/
│   ├── templates/
│   └── webhooks/
└── README.md
```

## 🤝 Внесение вклада

1. Форкните репозиторий
2. Создайте ветку для фичи (`git checkout -b feature/amazing-feature`)
3. Зафиксируйте изменения (`git commit -m 'Add amazing feature'`)
4. Отправьте в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📄 Лицензия

MIT License. См. файл [LICENSE](LICENSE) для деталей.

## 🔗 Ссылки

- [Документация Unisender Go API](https://godocs.unisender.ru/web-api-ref)
- [История изменений API](https://godocs.unisender.ru/api-changelog)
- [GitHub репозиторий](https://github.com/basili4-1982/unisender-go-api)

## 📞 Поддержка

Если у вас возникли вопросы или проблемы:
- Создайте [Issue](https://github.com/basili4-1982/unisender-go-api/issues)
- Напишите на email: your-email@example.com
```