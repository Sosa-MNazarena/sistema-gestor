package events

import (
	"fmt"
	"gopkg.in/gomail.v2"
)	

func SendLoadSuccessEmail(email string) error {
	fmt.Println("Enviando correo de éxito de carga a:", email)

	m := gomail.NewMessage()
	m.SetHeader("From", "mariana.sosa@davinci.edu.ar")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Datos cargados con éxito")
	m.SetBody("text/html", "<p>Los datos se han cargado correctamente en el sistema.</p>")
	//clave de aplicación de Gmail, cada persona/cuenta genera la suya
	d := gomail.NewDialer("smtp.gmail.com", 587, "mariana.sosa@davinci.edu.ar", "bpqn njgd kiyg rckk")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
	} else {
		fmt.Println("Correo enviado exitosamente.")
	}
	return nil
}