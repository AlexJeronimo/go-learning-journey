# Тиждень 1, День 2. Інтерфейси (поглиблено) та Обробка Помилок (кращі практики)

## Наступні кроки: Тиждень 1, День 2: Інтерфейси (поглиблено) та Обробка Помилок (кращі практики)

Сьогодні ми зосередимося на двох важливих концепціях Go: **інтерфейсах** та **кращих практиках обробки помилок**.

### Теоретична частина:

1. **Інтерфейси (поглиблено):**
    - **Природа інтерфейсів Go:** Інтерфейси в Go відрізняються від інтерфейсів в інших мовах. Вони є **неявними (implicit)**. Це означає, що тип реалізує інтерфейс автоматично, якщо він має всі методи, оголошені в інтерфейсі. Вам не потрібно явно вказувати, що тип "реалізує" інтерфейс.
    - **Принцип "менше інтерфейсів, більше щастя" (interface segregation principle):** У Go поширений підхід створення малих, сфокусованих інтерфейсів. Наприклад, замість одного великого інтерфейсу `FullReaderWriter` краще мати `Reader` та `Writer`.
    - **"Empty Interface" (`interface{}` / `any`):** `interface{}` може утримувати значення будь-якого типу. В Go 1.18 і пізніших версіях для цього ж використовується аліас `any`. Хоча це потужний інструмент, його слід використовувати обережно, оскільки він втрачає інформацію про тип і вимагає перевірок типу (type assertions) під час використання.
    - **Type Assertions та Type Switches:** Як перевірити базовий тип значення, яке зберігається в інтерфейсі.
2. **Обробка Помилок (кращі практики):**
    - **Ідіома `if err != nil`:** Це стандартний спосіб обробки помилок у Go. Помилки повертаються як останнє значення з функції.
    - **Створення власних типів помилок:** Крім `errors.New`, ви можете створювати власні типи помилок (наприклад, структури, що реалізують інтерфейс `error`). Це дозволяє додавати більше контексту до помилок.
    - **Обгортання помилок (`fmt.Errorf` з `%w`):** У Go 1.13 з'явилася можливість "обгортати" помилки, що дозволяє зберігати ланцюжок причин помилки. Це важливо для відлагодження. Ви можете перевіряти обгорнуті помилки за допомогою `errors.Is` та отримувати доступ до них за допомогою `errors.As`.
    - **`defer`, `panic`, `recover` (повторення/поглиблення):** Хоча `panic` використовується для виняткових, невідновлюваних ситуацій, `defer` і `recover` дозволяють перехопити паніку та відновити виконання, якщо це необхідно (наприклад, для закриття ресурсів). `defer` – це потужний інструмент для відкладеного виконання функцій, часто використовується для очищення ресурсів (`file.Close()`, `db.Close()`).

### Ресурси для вивчення:

- **Обов'язково:**
    - **A Tour of Go: Interfaces:** [https://go.dev/tour/methods/9](https://go.dev/tour/methods/9) (перегляньте ще раз, зосереджуючись на тому, як типи _неявно_ реалізують інтерфейси).
    - **Go by Example: Interfaces:** [https://gobyexample.com/interfaces](https://gobyexample.com/interfaces)
    - **Go Blog: Error Handling and Go:** [https://go.dev/blog/error-handling-and-go](https://go.dev/blog/error-handling-and-go)
    - **Go Blog: Working with Errors in Go 1.13:** [https://go.dev/blog/go1.13-errors](https://go.dev/blog/go1.13-errors) (зосередьтеся на `errors.Is` та `errors.As`).
    - Ваша книга **"Мова програмування Go"** – знайдіть розділи про інтерфейси та обробку помилок.
    - Ваша книга **"Ідіоми та паттерни проектування"** – розділи, що стосуються інтерфейсів та композиції.
- **Додатково:**
    - [Go by Example: Defer](https://gobyexample.com/defer)
    - [Go by Example: Panics](https://www.google.com/search?q=https://gobyexample.com/panics)
    - [Go Blog: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

### Практична частина:

Ми будемо модифікувати ваш поточний проект.

1. **Створіть інтерфейс `Operation`:**
    
    - У файлі `calculator/calculator.go`, вище функції `Calculate`, визначте наступний інтерфейс:
        
        Go
        
        ```
        type Operation interface {
            Apply(num1, num2 float64) (float64, error)
            Symbol() string // Додайте цей метод для отримання символу операції
        }
        ```
        
2. **Реалізуйте `Operation` для різних операцій:**
    
    - Створіть окремі файли в пакеті `calculator` (наприклад, `calculator/addition.go`, `calculator/subtraction.go` і т.д.), або всі в `calculator.go` (якщо вам зручніше для такого маленького прикладу).
    - Визначте для кожної операції **структуру** (наприклад, `Addition`, `Subtraction`, `Multiplication`, `Division`). Ці структури можуть бути порожніми, оскільки їм не потрібно зберігати стан.
    - Кожна з цих структур повинна **реалізовувати інтерфейс `Operation`**:
        - Реалізуйте метод `Apply` для кожної структури. Логіка буде схожою на ту, що вже є у вашому `Calculate` функції.
        - Реалізуйте метод `Symbol()` для кожної структури, який повертає символ операції (наприклад, `"+"`, `"-"`, `"*"` або `/`).
3. **Створіть власний тип помилки для ділення на нуль:**
    
    - У файлі `calculator/calculator.go` (або в окремому файлі `calculator/errors.go`):
        
        Go
        
        ```
        // Власна структура помилки
        type DivideByZeroError struct {
            Message string
        }
        
        func (e *DivideByZeroError) Error() string {
            return e.Message
        }
        
        // Експортована змінна для зручності
        var ErrDivideByZero = &DivideByZeroError{Message: "cannot divide by zero"}
        ```
        
    - У методі `Apply` для `Division` використовуйте `return 0, ErrDivideByZero` замість `errors.New("cannot divide by zero")`.
4. **Модифікуйте `main.go`:**
    
    - Створіть екземпляри ваших структур-операцій: `add := Addition{}` тощо.
    - Викличте їхні методи `Apply`.
    - Використайте **Type Assertion** для перевірки, чи помилка, яку повертає `Division.Apply`, є саме вашим `*DivideByZeroError`. Якщо так, виведіть спеціальне повідомлення.
    - Використайте `errors.Is` для перевірки, чи отримана помилка є `ErrDivideByZero` (це важлива сучасна Go-ідіома).
    - Використайте `defer` для імітації закриття ресурсу (наприклад, виведіть повідомлення "Closing resources" у кінці `main` функції).
5. **Приклад використання в `main.go` (орієнтир):**
    
    Go
    
    ```
    package main
    
    import (
    	"fmt"
    	"glp/calculator"
    	"errors" // Додайте імпорт errors
    )
    
    func main() {
    	// defer-заява
    	defer fmt.Println("Exiting main function. Cleaning up resources.")
    
    	// Приклад використання інтерфейсу
    	var op calculator.Operation
    
    	op = calculator.Addition{} // Припускаємо, що у вас є struct Addition{}
    	res, err := op.Apply(10, 5)
    	if err != nil {
    		fmt.Println("Error:", err)
    	} else {
    		fmt.Printf("Operation %s: %f\n", op.Symbol(), res)
    	}
    
    	op = calculator.Division{} // Припускаємо, що у вас є struct Division{}
    	res, err = op.Apply(10, 0) // Спроба ділення на нуль
    	if err != nil {
    		fmt.Printf("Error during division: %v\n", err)
    
    		// Перевірка власного типу помилки
    		var divErr *calculator.DivideByZeroError
    		if errors.As(err, &divErr) {
    			fmt.Println("Caught a specific DivideByZeroError:", divErr.Message)
    		}
    
    		// Перевірка за допомогою errors.Is
    		if errors.Is(err, calculator.ErrDivideByZero) {
    			fmt.Println("Error is recognized as ErrDivideByZero using errors.Is.")
    		}
    
    	} else {
    		fmt.Printf("Operation %s: %f\n", op.Symbol(), res)
    	}
    
    	// ... інші операції
    }
    ```
    

---

Виконайте ці завдання, і коли будете готові, надішліть оновлені файли (можливо, `calculator.go` та `main.go`, а також нові файли операцій, якщо ви їх створили).