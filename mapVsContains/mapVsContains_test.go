package mapvscontains

import (
	"testing"
)

var estadosBR = [...]string{
	"Acre (AC)",
	"Alagoas (AL)",
	"Amapá (AP)",
	"Amazonas (AM)",
	"Bahia (BA)",
	"Ceará (CE)",
	"Distrito Federal (DF)",
	"Espírito Santo (ES)",
	"Goiás (GO)",
	"Maranhão (MA)",
	"Mato Grosso (MT)",
	"Mato Grosso do Sul (MS)",
	"Minas Gerais (MG)",
	"Pará (PA)",
	"Paraíba (PB)",
	"Paraná (PR)",
	"Pernambuco (PE)",
	"Piauí (PI)",
	"Rio de Janeiro (RJ)",
	"Rio Grande do Norte (RN)",
	"Rio Grande do Sul (RS)",
	"Rondônia (RO)",
	"Roraima (RR)",
	"Santa Catarina (SC)",
	"São Paulo (SP)",
	"Sergipe (SE)",
	"Tocantins (TO)",
}

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
	"twenty",
	"twenty-one",
	"twenty-two",
	"twenty-three",
	"twenty-four",
	"twenty-five",
	"twenty-six",
	"twenty-seven",
	"twenty-eight",
	"twenty-nine",
	"thirty",
	"thirty-one",
	"thirty-two",
	"thirty-three",
	"thirty-four",
	"thirty-five",
	"thirty-six",
	"thirty-seven",
	"thirty-eight",
	"thirty-nine",
	"forty",
	"forty-one",
	"forty-two",
	"forty-three",
	"forty-four",
	"forty-five",
	"forty-six",
	"forty-seven",
	"forty-eight",
	"forty-nine",
	"fifty",
	"fifty-one",
	"fifty-two",
	"fifty-three",
	"fifty-four",
	"fifty-five",
	"fifty-six",
	"fifty-seven",
	"fifty-eight",
	"fifty-nine",
	"sixty",
	"sixty-one",
	"sixty-two",
	"sixty-three",
	"sixty-four",
	"sixty-five",
	"sixty-six",
	"sixty-seven",
	"sixty-eight",
	"sixty-nine",
	"seventy",
	"seventy-one",
	"seventy-two",
	"seventy-three",
	"seventy-four",
	"seventy-five",
	"seventy-six",
	"seventy-seven",
	"seventy-eight",
	"seventy-nine",
	"eighty",
	"eighty-one",
	"eighty-two",
	"eighty-three",
	"eighty-four",
	"eighty-five",
	"eighty-six",
	"eighty-seven",
	"eighty-eight",
	"eighty-nine",
	"ninety",
	"ninety-one",
	"ninety-two",
	"ninety-three",
	"ninety-four",
	"ninety-five",
	"ninety-six",
	"ninety-seven",
	"ninety-eight",
	"ninety-nine",
}

var result bool

func contains(find string, list []string) bool {
	for _, item := range list {
		if item == find {
			return true
		}
	}
	return false
}

func Benchmark_contains_numbers_firstElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("one", numbers)
	}
	result = found
}

func Benchmark_contains_numbers_lastElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("ninety-nine", numbers)
	}
	result = found
}

func Benchmark_contains_numbers_inexistentElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("zero", numbers)
	}
	result = found
}

func Benchmark_contains_estadosBR_firstElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("Acre (AC)", numbers)
	}
	result = found
}

func Benchmark_contains_estadosBR_lastElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("Tocantins (TO)", numbers)
	}
	result = found
}

func Benchmark_contains_estadosBR_inexistentElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		found = contains("Fernando de Noronha", numbers)
	}
	result = found
}

func Benchmark_map_numbers_firstElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range numbers {
			mapper[s] = exists
		}

		if _, item := mapper["one"]; item {
			found = true
		}
	}
	result = found
}

func Benchmark_map_numbers_lastElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range numbers {
			mapper[s] = exists
		}

		if _, item := mapper["ninety-nine"]; item {
			found = true
		}
	}
	result = found
}

func Benchmark_map_numbers_inexistentElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range numbers {
			mapper[s] = exists
		}

		if _, item := mapper["zero"]; item {
			found = true
		}
	}
	result = found
}

func Benchmark_map_estadosBR_firstElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range estadosBR {
			mapper[s] = exists
		}

		if _, item := mapper["Acre (AC)"]; item {
			found = true
		}
	}
	result = found
}

func Benchmark_map_estadosBR_lastElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range estadosBR {
			mapper[s] = exists
		}

		if _, item := mapper["Tocantins (TO)"]; item {
			found = true
		}
	}
	result = found
}

func Benchmark_map_estadosBR_inexistentElement(b *testing.B) {
	found := false
	for i := 0; i < b.N; i++ {
		mapper := map[string]struct{}{}
		exists := struct{}{}
		for _, s := range estadosBR {
			mapper[s] = exists
		}

		if _, item := mapper["Fernando de Noronha"]; item {
			found = true
		}
	}
	result = found
}
