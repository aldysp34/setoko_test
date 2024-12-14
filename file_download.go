package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aldysp34/setoko_test/generator"
)

func downloadReceiptFile(ctx context.Context) ([]byte, error) {
	listModelData := []generator.ListModelData{
		{
			HeaderData: "Detail Pembeli",
			ModelData: []generator.ModelData{
				{
					Key:   "neuralism2",
					Value: "",
				},
				{
					Key:   "6281807362365",
					Value: "",
				},
			},
		},
		{
			HeaderData: "Detail Pesanan",
			ModelData: []generator.ModelData{
				{
					Key:   "laptop",
					Value: "",
				},
				{
					Key:   "1 x Rp.1000",
					Value: "Rp.1000",
				},
				{
					Key:   "Total Harga (1 Produk)",
					Value: "Rp.1000",
				},
			},
		},
		{
			HeaderData: "Detail Pengiriman",
			ModelData: []generator.ModelData{
				{
					Key:   "Metode Pengiriman",
					Value: "Kasir",
				},
				{
					Key:   "Biaya Pengiriman",
					Value: "Rp0",
				},
				{
					Key:   "Asuransi Pengiriman",
					Value: "Rp0",
				},
			},
		},
		{
			HeaderData: "Detail Pembayaran",
			ModelData: []generator.ModelData{
				{
					Key:   "Metode Pembayaran",
					Value: "Tunai (Cash)",
				},
				{
					Key:   "Status",
					Value: "Lunas",
				},
				{
					Key:   "Subtotal Pesanan",
					Value: "Rp1.000",
				},
				{
					Key:   "Subtotal Pengiriman",
					Value: "Rp0",
				},
				{
					Key:   "Total Pembayaran",
					Value: "Rp1.000",
				},
			},
		},
	}

	gen, err := generator.NewReceiptFileGenerator(ctx, generator.Paper_A5)
	if err != nil {
		return nil, err
	}
	buf, err := gen.GenerateReceipt("", "setoko", false, listModelData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func downloadReceiptFileHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := downloadReceiptFile(context.Background())
	if err != nil {
		response := response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		// Encode and write the JSON response
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	fileName := "resi"
	dateNow := time.Now().Format("20060102")
	fileLength := len(buf)

	// Set headers for file download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.pdf\"", dateNow, fileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileLength))

	// Write file data to response
	_, err = w.Write(buf)
	if err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
	}
}
