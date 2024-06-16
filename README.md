# metrics
Yandex Metrics

/metrics
    /cmd
        /server
            main.go         # Основной исполняемый файл сервера
    /internal
        /handlers          # Обработчики HTTP запросов
            handlers.go    # Файл с функциями обработки HTTP запросов
    /pkg
        /storage           # Логика хранения данных
            storage.go     # Файл для работы с хранилищем метрик
    go.mod                 # Файл управления зависимостями
    go.sum                 # Файл с хешами зависимостей
    README.md              # Описание проекта
