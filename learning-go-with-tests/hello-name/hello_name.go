package hello_name

const prefixEnglish = "Hello, "
const prefixPortuguese = "Olá, "
const prefixEspanish = "Hola, "
const prefixFrench = "Bonjour, "

func HelloName(name, idiom string) string {

	if name == "" {
		return "Hello, World"
	}

	switch idiom {
	case "portuguese":
		return prefixPortuguese + name
	case "spanish":
		return prefixEspanish + name
	case "french":
		return prefixFrench + name
	}

	return prefixEnglish + name
}
