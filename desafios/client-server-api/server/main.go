package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Registro struct {
	ID    string
	Valor float64
}

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", GetCotacao)
	http.ListenAndServe(":8080", nil)
}

func GetCotacao(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// contexto HTTP
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	// requisição
	req, err := http.NewRequestWithContext(ctx, "get", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// leitura
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// conversão
	var CotacaoDolar Cotacao
	err = json.Unmarshal(res, &CotacaoDolar)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.Write([]byte(CotacaoDolar.Usdbrl.Bid))

	// registro
	floatValue, _ := strconv.ParseFloat(CotacaoDolar.Usdbrl.Bid, 64)
	registro := NewRegistro(floatValue)
	_ = insertRegistro(registro)

}

func NewRegistro(valor float64) *Registro {
	return &Registro{
		ID:    uuid.New().String(),
		Valor: valor,
	}
}

func insertRegistro(registro *Registro) error {
	db, err := sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query("create table if not exists cotacoes (id string NOT NULL PRIMARY KEY, valor float NOT NULL)")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("insert into cotacoes (id, valor) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(registro.ID, registro.Valor)
	if err != nil {
		return err
	}
	return nil

}
