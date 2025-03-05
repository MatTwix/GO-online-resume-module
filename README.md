# Go-модуль для работы с API VK, CODEFORCES и GITHUB
Данный проект представляет собой go-модуль, предоставляющий информацию из профилей VK, Github и Codeforces. 
Архитектура проекта соответствует специальной структуре модулей, что позволит легко его масштабировать под все необходимые нужды.
Также реализована обработка ошибок на каждом этапе работы программы.
Полученные данные после обрабатываются и отображаются на главной странице сайта.
# Структура
/api                # Пакеты для работы с API  
    github.go       # Модуль для работы с GitHub API  
    telegram.go     # Модуль для работы с Telegram API  
    vk.go           # Модуль для работы с VK API  
/client             # Фронтенд-приложение (React)  
    /public         # Публичные файлы (иконки, изображения)  
    /src            # Исходный код фронтенда  
        /assets     # Статические ресурсы (графика, лого)  
        App.css     # Стили главного компонента  
        App.jsx     # Главный React-компонент  
        index.css   # Глобальные стили приложения  
        main.jsx    # Точка входа фронтенда  
    .env_example    # Пример файла окружения для фронтенда  
    index.html      # Основной HTML-шаблон приложения  
    ...             # Прочие файлы конфигурации фронтенда  
/config             # Конфигурация проекта  
    config.go       # Загрузка данных из .env  
/handlers           # Обработчики HTTP-запросов  
    resumeHandler.go # Логика обработки резюме  
/models             # Определение структур данных  
    Resume.go       # Модель данных резюме  
/routes             # Определение маршрутов API  
    routes.go       # Основные маршруты сервера  
.gitignore          # Игнорирование файлов, не отслеживаемых Git  
.env_example        # Шаблон конфигурационного файла окружения  
README.md           # Документация проекта  
go.mod              # Файл управления зависимостями Go  
main.go             # Основной файл запуска сервера на Go  
# Как запустить? 
Необходимо установить сам Golang на свой пк ([ссылка](https://go.dev/doc/install)) и Node.JS [ссылка](https://nodejs.org/en)

После этого необходимо создать 2 файла с персональными ключами- `.env`. 
Первый находится в главной дирректории проекта и отвечает за конфигурационные данные для получения и отправки API React-приложению.
Данные в этом файле должны выглядеть по аналогии с `.env_example`

Аналогично нужно поступить и с `.env` в папке `client/`, тут нужно указать тот же порт, что в вышеупомянутом файле, но уже по аналогии с `.env_example` в текущей дирректории

Далее необходимо перейти в дирректорию проекта и запустить файл main.go:
```
go run main.go

либо по абсолютной ссылке

go run "/{path}/main.go"
```

После этого нужно в другом окне терминала перейти в дирректорию `client/`.
```
cd client/
```
После этого необходимо запустить сделующий ряд комманд:
```
npm install
npm run dev
```

Чтобы увидеть отображение приложения, нужно ввести в поисковую строку браузера `http://localhost:5173` (указанный порт установлен по умолчанию, актуальная ссылка будет отображаться в терминале после запуска `npm run dev`)

# Использованные дополнительные библиотеки
[Go Fiber](https://github.com/gofiber/fiber)
[React](https://react.dev/)
[GoDotEnv](github.com/joho/godotenv)