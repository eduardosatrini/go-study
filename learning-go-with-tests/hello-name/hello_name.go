package hello_name

const prefixEnglish = "Hello, "
const prefixPortuguese = "Olá, "
const prefixEspanish = "Hola, "
const prefixFrench = "Bonjour, "

func prefixGreetings(idiom string) (prefix string) {

	switch idiom {
	case "portuguese":
		prefix = prefixPortuguese
	case "spanish":
		prefix = prefixEspanish
	case "french":
		prefix = prefixFrench
	default:
		prefix = prefixEnglish
	}

	return
}

func HelloName(name, idiom string) string {

	if name == "" {
		return "Hello, World"
	}

	return prefixGreetings(idiom) + name
}
