package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type Request struct {
    Expression string `json:"expression"`
}

type Response struct {
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}

func calculate(expression string) (float64, error) {
    // Простой парсер для арифметических выражений
    // Здесь можно использовать более сложный парсер для реальных приложений
    result, err := eval(expression)
    if err != nil {
        return 0, err
    }
    return result, nil
}

func eval(expression string) (float64, error) {
    // Пример простого вычисления (только для демонстрации)
    // В реальном приложении используйте безопасный парсер
    parts := strings.Fields(expression)
    if len(parts) != 3 {
        return 0, errors.New("invalid expression")
    }

    a, err := strconv.ParseFloat(parts[0], 64)
    if err != nil {
        return 0, err
    }
    b, err := strconv.ParseFloat(parts[2], 64)
    if err != nil {
        return 0, err
    }

    switch parts[1] {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        return a / b, nil
    default:
        return 0, errors.New("invalid operator")
    }
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    result, err := calculate(req.Expression)
    if err != nil {
        http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
        return
    }

    response := Response{Result: strconv.FormatFloat(result, 'f', -1, 64)}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/api/v1/calculate", calculateHandler)
    http.ListenAndServe(":8080", nil)
}