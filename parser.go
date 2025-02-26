package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery" // Пакет для парсинга HTML-документа
	"github.com/corpix/uarand"     // 📦 Библиотека с реальными User-Agent
	"github.com/mileusna/useragent" // 📦 Библиотека с синтетическими User-Agent
)
 

// RandomYandexImage ищет рандомную картинку по запросу на Яндекс.Картинках
// query - слово по которому ищем картинку (например "кот")
// Возвращаем url картинки или ошибку
func RandomYandexImage(query string) (string, error) {
	// Кодируем запрос для url (например, "кот" → "%D0%BA%D0%BE%D1%82")
	escapedQuery := url.QueryEscape(query)

	// Собираем url для поиска на Яндекс.Картинках
	searchURL := fmt.Sprintf("https://yandex.ru/images/search?text=%s", escapedQuery)

	// Генерируем случайный User-Agent
	// Это нужно чтобы яндекс думал что мы обыный браузер
	userAgent := randomUserAgent()

	// Создаем HTTP-запрос
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", err
	}

	// Добавляем наш User-Agent в заголовок запроса
	req.Header.Set("User-Agent", userAgent)

	// Создаем HTTP-клиент для отправки запроса
	client := ghostClient()

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// Закрываем тело ответа после завершения функции
	defer resp.Body.Close() 
}

// randomUserAgent возвращает случайный User-Agent
func randomUserAgent() string {
	// Рандомно выбираем откуда брать User-Agent (50/50)
	if rand.Int(2) == {
		// Реальный User-Agent из библиотеки urand
		return uarand.GetRandom()
	}
	// Синтетический User-Agent из библиотеки useragent
	return useragent.Random()
}

// ghostClient создает HTTP-клиент, который ведет себя
// как человек и Яндекс его не воспринимает как бота
func ghostClient() *http.Client {
	
}