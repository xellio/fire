package fire

import (
	"testing"
)

// dataprovider for TestIsSupportedMethod
func dpTestIsSupportedMethod() map[string]bool {
	return map[string]bool{
		"GET":    true,
		"POST":   true,
		"PUT":    true,
		"DELETE": true,
		"PATCH":  true,
		"get":    false,
		"post":   false,
		"put":    false,
		"delete": false,
		"path":   false,
		"":       false,
		"fooo":   false,
		"FOO":    false,
	}
}

func TestIsSupportedMethod(t *testing.T) {
	data := dpTestIsSupportedMethod()
	for method, expected := range data {
		result := IsSupportedMethod(method)
		if result != expected {
			t.Errorf("%s should be %t - got %t", method, expected, result)
		}
	}
}

// dataprovider for TestIsValidURL
func dpTestIsValidURL() map[string]bool {
	return map[string]bool{
		"http://www.example.com/":                      true,
		"http://www.example.com":                       true,
		"https://www.example.com":                      true,
		"http://www.example.com/index.html":            true,
		"http://www.example.com/index.html?block=true": true,
		"http://www.example.com/?block=true":           true,
		"http://www.example.com?block=true":            true,
		"//www.example.com":                            true,
		"//www.example.com/":                           true,
		"local://www.example.com/":                     true,
		"ftp://www.example.com":                        true,
		"www.example.com":                              false,
		"www.example.com/index.html":                   false,
		"example.com":                                  false,
		"example.com/index.html":                       false,
		"index.html":                                   false,
		"":                                             false,
	}
}

func TestIsValidURL(t *testing.T) {
	data := dpTestIsValidURL()
	for url, expected := range data {
		result := IsValidURL(url)
		if result != expected {
			t.Errorf("%s should be %t - got %t", url, expected, result)
		}
	}
}
