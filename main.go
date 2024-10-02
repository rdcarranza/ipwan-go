package main

import (
	"bytes"
	"io"

	"fmt"

	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/rdcarranza/ipwan-go/src/controladores/env"
	"github.com/rdcarranza/ipwan-go/src/controladores/login"
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
	user_host, _ := env.GetEnv("user_host", env_)
	pass_host, _ := env.GetEnv("pass_host", env_)

	pass := login.Base64encode(pass_host)
	//CONSULTA

	// Datos de la solicitud
	reqBody := login.Request{
		Username:      user_host,
		Password:      pass,
		Password_type: 4,
	}

	//jsonData, err := json.Marshal(reqBody)
	data, err := login.Object2xml(reqBody)
	if err != nil {
		panic(err)
	}
	log.Println("xmldata = \n", data)

	urlLogin := "http://" + dir_host + "/api/user/login"
	log.Println("URL: " + urlLogin)

	// Crear un nuevo cookie jar
	jar, _ := cookiejar.New(nil)

	// Crear un cliente HTTP con el cookie jar
	client := &http.Client{
		Jar: jar,
	}

	r0, _ := client.Get("http://192.168.253.250")
	defer r0.Body.Close()
	// Imprimir las cookies obtenidas

	/*for _, cookie := range jar.Cookies(r0.Request.URL) {
		fmt.Printf("Cookie0: %s=%s\n", cookie.Name, cookie.Value)
	}*/
	cookie := jar.Cookies(r0.Request.URL)
	fmt.Printf("Cookie0:", cookie)
	urlString := "http://192.168.253.250/api/monitoring/status"
	u1, _ := url.Parse(urlString)
	client.Jar.SetCookies(u1, cookie)
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
	r1, _ := client.Get("http://192.168.253.250/api/monitoring/status")
	defer r1.Body.Close()
	for _, cookie := range jar.Cookies(r1.Request.URL) {
		fmt.Printf("Cookie1: %s=%s\n", cookie.Name, cookie.Value)
	}
	body1, _ := io.ReadAll(r1.Body)
	fmt.Println("respuesta1: " + string(body1))

	//client := &http.Client{}

	req, err := http.NewRequest("POST", urlLogin, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("Error creando la solicitud: %v", err)
	}
	req.Header.Set("Content-Type", "application/xml")

	// Realizar la solicitud HTTP

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Imprimir las cookies obtenidas
	for _, cookie := range jar.Cookies(resp.Request.URL) {
		fmt.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error leyendo la respuesta: %v", err)
	}

	/*
		resBody, err := login.Xml2object(body)
		if err != nil {
			log.Fatalf("Error convirtiendo XML a objeto: %v", err)
		}
	*/
	fmt.Println("respuesta: " + string(body))

	/*
		 a borrar
			// Crear la solicitud HTTP
			endpoint := "http://" + dir_host + "/api/user/login"
			req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
			if err != nil {
				panic(err)
			}

			// Agregar el encabezado Content-Type
			req.Header.Set("Content-Type", "application/json")

			// Realizar la solicitud
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			// Leer la respuesta
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println("respuesta: " + string(body))
			// Decodificar la respuesta JSON
			var response Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				panic(err)
			}

			fmt.Println(response.Message)
	*/
}
