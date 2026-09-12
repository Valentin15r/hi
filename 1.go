package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	// Раздаем статические файлы из папки web
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	// Определяем локальный IP-адрес автоматически
	localIP := getLocalIP()

	addr := ":80"
	log.Printf("Запуск HTTP сервера на http://%s%s", localIP, addr)
	log.Printf("Локально: http://localhost%s", addr)
	log.Printf("В сети:   http://%s%s", localIP, addr)

	// Запускаем обычный HTTP сервер (без TLS)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

// getLocalIP возвращает первый не-loopback IPv4 адрес
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return ip4.String()
			}
		}
	}
	return "127.0.0.1"
}