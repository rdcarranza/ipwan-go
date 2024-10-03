package login

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rdcarranza/ipwan-go/src/controladores/env"
)

func login() {
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

	pass := Base64encode(pass_host)

	//CONSULTA

	// Datos de la solicitud
	reqBody := Request{
		Username:      user_host,
		Password:      pass,
		Password_type: 4,
	}

	//jsonData, err := json.Marshal(reqBody)
	data, err := Object2xml(reqBody)
	if err != nil {
		panic(err)
	}
	log.Println("xmldata = \n", data)

	urlLogin := "http://" + dir_host + "/api/user/login"
	log.Println("URL: " + urlLogin)

	client := &http.Client{}

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

// Función para codificar una cadena en base64 manualmente
func Base64encode(str string) string {
	// Definir el alfabeto Base64
	const base64chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	var out string
	var c1, c2, c3 int
	i := 0
	lenStr := len(str)

	for i < lenStr {
		// Obtener el primer carácter
		c1 = int(str[i]) & 0xff
		i++

		// Si llegamos al final de la cadena
		if i == lenStr {
			out += string(base64chars[c1>>2])
			out += string(base64chars[(c1&0x3)<<4])
			out += "=="
			break
		}

		// Obtener el segundo carácter
		c2 = int(str[i]) & 0xff
		i++

		// Si llegamos al final de la cadena
		if i == lenStr {
			out += string(base64chars[c1>>2])
			out += string(base64chars[((c1&0x3)<<4)|((c2&0xF0)>>4)])
			out += string(base64chars[(c2&0xF)<<2])
			out += "="
			break
		}

		// Obtener el tercer carácter
		c3 = int(str[i]) & 0xff
		i++

		// Realizar las operaciones bit a bit para obtener los caracteres codificados en base64
		out += string(base64chars[c1>>2])
		out += string(base64chars[((c1&0x3)<<4)|((c2&0xF0)>>4)])
		out += string(base64chars[((c2&0xF)<<2)|((c3&0xC0)>>6)])
		out += string(base64chars[c3&0x3F])
	}

	return out
}

// Convierte una estructura a XML en forma de cadena
func Object2xml(request Request) (string, error) {
	output, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Función que maneja la respuesta XML y la convierte en una estructura
func Xml2object(data []byte) (*Request, error) {
	var ret Request
	err := xml.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

//func login(name string,pw string) string{}
