package main

import (
	"io"

	"fmt"

	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/rdcarranza/ipwan-go/src/controladores/env"
	"github.com/rdcarranza/ipwan-go/src/controladores/ipwan"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	//Probando variables de entorno
	env_ := "./.env"
	env_copia := "./src/controladores/env/env.copia"
	if env.VerificarEnv(env_, env_copia) {
		estado_env, err := env.GetEnv("estado_env", env_)
		if err == nil && estado_env == "1" {
			fmt.Println("Archivo env: " + env_ + " - cargado correctamente.")
		} else {
			if err == nil {
				log.Fatal("Verificar configuración de archivo env")
			} else {
				log.Fatal("Verificar configuración de archivo env: " + err.Error())
			}

		}
	}

	dir_host, _ := env.GetEnv("dir_host", env_)

	// Crear un nuevo cookie jar
	jar, _ := cookiejar.New(nil)

	// Crear un cliente HTTP con el cookie jar
	client := &http.Client{
		Jar: jar,
	}

	r0, _ := client.Get("http://" + dir_host)
	defer r0.Body.Close()

	// Imprimir las cookies obtenidas
	/*for _, cookie := range jar.Cookies(r0.Request.URL) {
		fmt.Printf("Cookie0: %s=%s\n", cookie.Name, cookie.Value)
	}*/

	cookie := jar.Cookies(r0.Request.URL)
	//fmt.Println("\nCookie0: ", cookie) //se imprime la cookie sin formato.
	urlString := "http://" + dir_host + "/api/monitoring/status"
	u1, _ := url.Parse(urlString)
	client.Jar.SetCookies(u1, cookie) //se carga la cookie en el contexto de la siguiente petición.

	/*
		u, _ := url.Parse("http://192.168.253.250/api/monitoring/status")

		ck := &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Domain:   u.Host,
			Path:     "/",
			Secure:   false, // Si la cookie debe ser segura (HTTPS)
			HttpOnly: false, // Si la cookie solo debe ser accesible desde el navegador
		}

		// Agregar la cookie al cookiejar
		client.Jar.SetCookies(u, []*http.Cookie{ck})
	*/

	r1, err := client.Get("http://" + dir_host + "/api/monitoring/status")
	if err != nil {
		log.Fatal(err)
	}
	defer r1.Body.Close()
	body1, _ := io.ReadAll(r1.Body)

	//fmt.Println("\nrespuesta1:\n" + string(body1))

	res1, err := ipwan.Xml2object(body1)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nRespuesta: \n")
	fmt.Println("IPv4 WAN: " + res1.WanIpV4)
	//fmt.Println("IPv6 WAN: " + res1.WanIpV6)
	fmt.Print("\n")

}
