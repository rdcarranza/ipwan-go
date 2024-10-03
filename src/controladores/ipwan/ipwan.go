package ipwan

import "encoding/xml"

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
