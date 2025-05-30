
# День 5. Розширені можливості HTTP-сервера та Робота з HTML-шаблонами

Ми продовжуємо поглиблюватись у веб-розробку з Go, переходячи до більш складних аспектів HTTP-серверів та генерації динамічного HTML.

---

### Завдання 3.4: Обробка HTTP-помилок та статус-кодів

- **Мета:** Навчитися повертати відповідні HTTP-статус-коди для різних сценаріїв помилок.
- **У `main`:**
    - Залиште попередні обробники (`homeHandler`, `aboutHandler`, `greetHandler`).
    - Створіть новий обробник `notFoundHandler(w http.ResponseWriter, r *http.Request)` для шляху `/404`.
        - Він повинен повертати HTTP-статус **`404 Not Found`** і повідомлення "Сторінка не знайдена.".
        - Використайте `http.Error(w, message, statusCode)` для цього.
    - Створіть обробник `internalServerErrorHandler(w http.ResponseWriter, r *http.Request)` для шляху `/500`.
        - Він повинен повертати HTTP-статус **`500 Internal Server Error`** і повідомлення "Сталася внутрішня помилка сервера.".
        - Використайте `http.Error` або `w.WriteHeader` та `fmt.Fprint`.
    - Зареєструйте ці обробники.

---

### Завдання 3.5: Робота з HTML-шаблонами (`html/template`)

- **Мета:** Навчитися використовувати стандартний пакет Go `html/template` для генерації динамічного HTML-вмісту.
- **Кроки:**
    1. **Створіть файл шаблону:** У корені вашого проекту створіть файл `templates/index.html` з таким вмістом:

        ```html
        <!DOCTYPE html>
        <html lang="uk">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>{{.Title}}</title>
        </head>
        <body>
            <h1>{{.Greeting}}</h1>
            <p>Ваше ім'я: {{.Name}}</p>
            <p>Повідомлення: {{.Message}}</p>
        
            <h2>Список елементів:</h2>
            <ul>
                {{range .Items}}
                    <li>{{.}}</li>
                {{end}}
            </ul>
        </body>
        </html>
        ```
        
    2. **У `main`:**
        - Створіть структуру `PageData`, яка буде містити дані для шаблону:

            ```go
            type PageData struct {
                Title    string
                Greeting string
                Name     string
                Message  string
                Items    []string
            }
            ```
            
        - Створіть новий обробник `templateHandler(w http.ResponseWriter, r *http.Request)` для шляху `/template`.
            - Всередині обробника:
                1. **Завантажте шаблон:** Використайте `template.ParseFiles("templates/index.html")`. Обробляйте помилки завантаження шаблону (наприклад, `log.Fatal` або `http.Error`).
                2. **Створіть дані для шаблону:** Створіть екземпляр `PageData` і заповніть його будь-якими тестовими даними (наприклад, `Title: "Динамічна сторінка"`, `Greeting: "Привіт, світе шаблонів!"`, `Name: "Go Розробник"`, `Items: []string{"Елемент 1", "Елемент 2", "Елемент 3"}`).
                3. **Виконайте шаблон:** Використайте `tmpl.Execute(w, data)`. Це передасть дані до шаблону і запише згенерований HTML у `http.ResponseWriter`. Також обробляйте можливі помилки виконання шаблону.
        - Зареєструйте `templateHandler` для шляху `/template`.

---

### Завдання 3.6: Обробка HTML-форм

- **Мета:** Приймати дані з HTML-форми (метод `POST`) і відображати їх.
- **Кроки:**
    1. **Оновіть `templates/index.html`:** Додайте до нього просту форму (після списку елементів):

        ```html
        <hr>
        <h2>Надіслати дані:</h2>
        <form action="/submit-form" method="POST">
            <label for="username">Ім'я користувача:</label><br>
            <input type="text" id="username" name="username" value="За замовчуванням"><br><br>
            <label for="comment">Коментар:</label><br>
            <textarea id="comment" name="comment" rows="4" cols="50"></textarea><br><br>
            <input type="submit" value="Надіслати">
        </form>
        ```
        
    2. **У `main`:**
        - Створіть новий обробник `submitFormHandler(w http.ResponseWriter, r *http.Request)` для шляху `/submit-form`.
            - Всередині обробника:
                1. **Перевірка методу:** Переконайтеся, що це `POST`-запит. Якщо ні, поверніть `http.StatusMethodNotAllowed`.
                2. **Парсинг форми:** Використайте `r.ParseForm()`. Це парсить дані форми з тіла запиту (для `POST`) або URL-запиту (для `GET`).
                3. **Отримання даних:** Використайте `r.Form.Get("username")` та `r.Form.Get("comment")` для отримання значень полів.
                4. **Відповідь:** Сформуйте просту HTML-сторінку (можна знову використати шаблон або просто `fmt.Fprintf`), яка відображає отримані ім'я користувача та коментар. Наприклад: "Отримано дані: Ім'я: [ім'я], Коментар: [коментар]".
        - Зареєструйте `submitFormHandler` для шляху `/submit-form`.

---

Після виконання цих завдань ми будемо мати функціональний веб-сервер, який може обслуговувати статичні сторінки, генерувати динамічний HTML та обробляти дані, надіслані користувачем.