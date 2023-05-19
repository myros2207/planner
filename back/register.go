package register

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/planner_mk")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Створюємо таблицю users, якщо вона ще не існує
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(60) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Валідація даних вхідного запиту
		if len(request.Username) < 5 || len(request.Password) < 8 {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}

		// Перевірка, чи не існує вже користувача з таким іменем
		existingUser, err := getUserByUsername(db, request.Username)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if existingUser != nil {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		// Хешування паролю
		hashedPassword, err := hashPassword(request.Password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Генерація унікального ID для користувача
		userID := uuid.New().String()

		// Додавання нового користувача до бази даних
		err = addUser(db, userID, request.Username, hashedPassword)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Створення токена
		token, err := createToken(userID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Формування відповіді з токеном
		response := RegisterResponse{
			Token: token,
		}

		// Відправка відповіді з токеном
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hashPassword(password string) (string, error) {
	// Генерація хешу паролю за допомогою bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func addUser(db *sql.DB, id, username, password string) error {
	// Вставка даних користувача в таблицю users
	_, err := db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, username, password)
	if err != nil {
		return err
	}
	return nil
}

func getUserByUsername(db *sql.DB, username string) (*User, error) {
	// Отримання користувача за іменем користувача з бази даних
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func createToken(userID string) (string, error) {
	// Генерація унікального токена з використанням userID та поточного часу
	token := uuid.New().String() + "_" + userID + "_" + time.Now().Format("20060102150405")

	// Здійснення підпису токена (опційно)
	// signedToken, err := signToken(token)
	// if err != nil {
	//     return "", err
	// }

	return token, nil
}

// func signToken(token string) (string, error) {
//     // Реалізуйте підпис токена, використовуючи алгоритм підпису, такий як HMAC або RSA.
//     // Поверніть підписаний токен
// }
