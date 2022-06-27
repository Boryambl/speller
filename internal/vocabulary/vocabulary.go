package vocabulary

/*
import (
	"fmt"
	"os"
	"sync"
)

type Vocabulary struct {
	Vocabulary map[string]string `json:"vocabulary"`
	m sync.Mutex
}


func readRates(filename string) (Rates,error) {
	usdRates := make(map[string]float32)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}
	err = json.Unmarshal(file, &usdRates)
	if err != nil {
		return nil, fmt.Errorf("unmarshal file: %v", err)
	}
	for key, value := range usdRates {
		if !IsCurrencyValid(key)  {
			return nil, fmt.Errorf("invalid currency %v", key)
		}
		if value <= 0 {
			return nil, fmt.Errorf("invalid rate for currency %v", key)
		}
	}
	return usdRates, nil
}


 */