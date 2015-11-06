package brasil

import (
	"bytes"
	"errors"
	"strings"
)

var (
	uf_estados = map[string]string{
		"AC": "Acre",
		"AL": "Alagoas",
		"AP": "Amapá",
		"AM": "Amazonas",
		"BA": "Bahia",
		"CE": "Ceará",
		"DF": "Distrito Federal",
		"ES": "Espírito Santo",
		"GO": "Goiás",
		"MA": "Maranhão",
		"MT": "Mato Grosso",
		"MS": "Mato Grosso do Sul",
		"MG": "Minas Gerais",
		"PA": "Pará",
		"PB": "Paraíba",
		"PR": "Paraná",
		"PE": "Pernambuco",
		"PI": "Piauí",
		"RJ": "Rio de Janeiro",
		"RN": "Rio Grande do Norte",
		"RS": "Rio Grande do Sul",
		"RO": "Rondônia",
		"RR": "Roraima",
		"SC": "Santa Catarina",
		"SP": "São Paulo",
		"SE": "Sergipe",
		"TO": "Tocantins",
	}
	estados_uf = map[string]string{
		"acre":                "AC",
		"alagoas":             "AL",
		"amapa":               "AP",
		"amazonas":            "AM",
		"bahia":               "BA",
		"ceara":               "CE",
		"distrito federal":    "DF",
		"espirito santo":      "ES",
		"goias":               "GO",
		"maranhao":            "MA",
		"mato grosso":         "MT",
		"mato grosso do sul":  "MS",
		"minas gerais":        "MG",
		"para":                "PA",
		"paraiba":             "PB",
		"parana":              "PR",
		"pernambuco":          "PE",
		"piaui":               "PI",
		"rio de janeiro":      "RJ",
		"rio grande do norte": "RN",
		"rio grande do sul":   "RS",
		"rondonia":            "RO",
		"roraima":             "RR",
		"santa catarina":      "SC",
		"sao paulo":           "SP",
		"sergipe":             "SE",
		"tocantins":           "TO",
	}
	norm = map[rune]rune{
		'ã': 'a',
		'á': 'a',
		'â': 'a',
		'Ã': 'a',
		'Á': 'a',
		'Â': 'a',
		'ê': 'e',
		'é': 'e',
		'Ê': 'e',
		'É': 'e',
		'î': 'i',
		'í': 'i',
		'Î': 'i',
		'Í': 'i',
		'õ': 'o',
		'ó': 'o',
		'ô': 'o',
		'Õ': 'o',
		'Ó': 'o',
		'Ô': 'o',
		'û': 'u',
		'ú': 'u',
		'Û': 'u',
		'Ú': 'u',
	}
)

func normalize(str string) string {
	str = strings.ToLower(str)
	buffer := new(bytes.Buffer)
	for _, v := range str {
		if repl, ok := norm[v]; ok {
			buffer.WriteRune(repl)
		} else {
			buffer.WriteRune(v)
		}
	}
	return buffer.String()
}

func UF(estado string) (string, error) {
	estado = normalize(estado)
	if uf, ok := estados_uf[estado]; ok {
		return uf, nil
	}
	return "", errors.New("not found")
}

func ValidUF(uf string) bool {
	_, ok := uf_estados[uf]
	return ok
}

func Estado(uf string) (string, error) {
	uf = strings.ToUpper(uf)
	if estado, ok := uf_estados[uf]; ok {
		return estado, nil
	}
	return "", errors.New("not found")
}
