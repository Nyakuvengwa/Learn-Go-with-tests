package hello

const shonaLang = "Shona"
const tswanaLang = "Tswana"

const englishHelloPrefix = "Hello, "
const shonaHelloPrefix = "Mhoro "
const tswanaHelloPrefix = "Dumela "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case shonaLang:
		prefix = shonaHelloPrefix
	case tswanaLang:
		prefix = tswanaHelloPrefix
	default:
		prefix = englishHelloPrefix

	}
	return
}
