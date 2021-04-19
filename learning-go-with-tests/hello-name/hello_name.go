package hello_name

const prefixPortuguese = "Olá, "
const prefixEnglish = "Hello, "
const prefixEspanish = "Hola, "

func HelloName(name string) string {
	if name != "" {
		return "Hello, " + name + "!"
	}

	return "Hello, World!"
}
