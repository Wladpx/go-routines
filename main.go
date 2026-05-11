package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.chatgpt.com",
		"https://www.github.com",
		"https://www.linkedin.com",
		"https://www.microsoft.com",
	}

	var wg sync.WaitGroup
	inicio := time.Now()

	fmt.Println("Iniciando Processamento\n")

	for _, url := range urls {
		wg.Add(1)
		go verifyServiceStatus(url, &wg)
	}

	wg.Wait()

	fmt.Printf("Processamento finalizado. Tempo: %.2f\n", time.Since(inicio).Seconds())

}

func verifyServiceStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	resp, err := http.Get(url)
	periodo := time.Since(start).Seconds()

	if err != nil {
		fmt.Printf("Erro ao consultar serviço [%s] [%.2fs]: %v\n", url, periodo, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Sucesso ao consultar serviço [%s] [%.2fs] - status %d\n", url, periodo, resp.StatusCode)
	} else {
		fmt.Printf("Falha ao consultar serviço [%s] [%.2fs] - status %d\n", url, periodo, resp.StatusCode)
	}
}
