package main

import (
	"fmt"
	"log"
	"time"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// Dapatkan Account SID dan Auth Token dari Twilio Console
	accountSid := "Your SID"
	authToken := "Your Token"
	client := twilio.NewClient(accountSid, authToken, nil)

	// Nomor pengirim Twilio dan penerima
	from := "Provided Number" // Misalnya: +1234567890
	to := "abdi.txt"          // Misalnya: nomor.txt

	// Jumlah pengiriman OTP yang diinginkan
	jumlahOTP := 10

	// Perulangan untuk mengirim OTP
	for i := 1; i <= jumlahOTP; i++ {
		// OTP sementara
		otp := fmt.Sprintf("%06d", i*123456) // Generate OTP sederhana untuk uji coba

		// Kirim SMS
		_, err := client.ApiV2010.CreateMessage(&api.CreateMessageParams{
			To:   &to,
			From: &from,
			Body: &otp,
		})

		if err != nil {
			log.Fatalf("Error mengirim OTP: %s\n", err)
		}

		fmt.Printf("OTP ke %s: %s (Pengiriman ke-%d)\n", to, otp, i)

		// Delay antara setiap permintaan (misalnya 2 detik)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Selesai mengirim OTP.")
}
