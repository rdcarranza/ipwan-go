package login

import "encoding/xml"

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
