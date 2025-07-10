package events

import (
	"fmt"
	"gopkg.in/gomail.v2"
)	

func SendCustomEmail(email, subject, body string) error {
	fmt.Println("Enviando correo personalizado a:", email)

	m := gomail.NewMessage()
	m.SetHeader("From", "mariana.sosa@davinci.edu.ar")
	m.SetHeader("To", "mariana.sosa@davinci.edu.ar") //cambiar por un email particular
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, "mariana.sosa@davinci.edu.ar", "bpqn njgd kiyg rckk")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return err
	} 
	fmt.Println("Correo enviado exitosamente.")
	return nil
}