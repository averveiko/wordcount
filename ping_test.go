package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Этот тест делает реальные запросы в сеть, что довольно долго (почти секунду)
func TestPing(t *testing.T) {

	// Этот тест будет пропущен в коротком режиме: $ go test -v -short (так как он слишком долгий)
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	client := &http.Client{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com")
	if !got {
		t.Errorf("Expected example.com to be available")
	}
	got = pinger.Ping("https://example.com/404")
	if got {
		t.Errorf("Expected example.com/404 to be unavailable")
	}
}

// Далее пример использования заглушки
// MockClient - это заглушка http-клиента для тестов.
type MockClient struct{}

// Head возвращает http-ответ со статусом, указанным в url.
// Например:
// url = https://ya.ru/200 -> статус = 200
// url = https://ya.ru/404 -> статус = 404
func (client *MockClient) Head(url string) (resp *http.Response, err error) {
	parts := strings.Split(url, "/")
	last := parts[len(parts)-1]
	statusCode, err := strconv.Atoi(last)
	if err != nil {
		return nil, err
	}
	resp = &http.Response{StatusCode: statusCode}
	return resp, nil
}

func TestPingMocked(t *testing.T) {
	client := &MockClient{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com/200")
	if !got {
		t.Errorf("Expected example.com/200 to be available")
	}
	got = pinger.Ping("https://example.com/404")
	if got {
		t.Errorf("Expected example.com/404 to be unavailable")
	}
}

// Вообще, для http-заглушек в Go есть отдельный пакет net/http/httptest. Здесь общий пример общего подхода:
// 	- Заменяем конкретную зависимость на зависимость от интерфейса.
// 	- Реализуем интерфейс в тестах и используем его.

// setup и teardown (например подготовка данных перед тестами и очистка после)
func TestMain(m *testing.M) {
	fmt.Println("Setup tests...")
	start := time.Now()

	m.Run()

	fmt.Println("Teardown tests...")
	end := time.Now()
	fmt.Println("Tests took", end.Sub(start))
}
