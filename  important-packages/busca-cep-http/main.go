package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// Registrando o handler para a rota "/"
	http.HandleFunc("/", BuscaCEPHandler)
	// Iniciando o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// como se fosse um controller com a req e res
func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	// se a rota for diferente de "/" retorna 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// pego o valor do parametro cep
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// chamo a função BuscaCep passando o cep como parametro
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convertendo a struct ViaCEP para JSON
	// e escrevendo no corpo da resposta
	//-----------------------------------------
	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	// retorno o cep em formato JSON
	// diretamente para o cliente
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	res, error := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if error != nil {
		return nil, error
	}
	defer res.Body.Close()

	// Lendo o corpo da resposta da requisição
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	// Convertendo o JSON para a struct ViaCEP
	var c ViaCEP
	// deserializa o json para a struct
	// e passo o valor para o endereço de memória da var c
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
