package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title    string
	Body     string
	Products []Product
}

type Product struct {
	Title     string
	ImageName string
	Link      string
}

var data = PageData{
	Title: "Mi página de prueba",
	Body:  "Este es el contenido de la página.",
	Products: []Product{
		{
			Title:     "Smartbox - Caja Regalo SPA y Relax para Dos",
			ImageName: "imagen1.jpg",
			Link:      "https://www.amazon.es/Smartbox-Mujer-Momentos-inolvidables-Ideas-Originales-1-Unisex-Adult/dp/B0BL9T9DQC/ref=sr_1_1_sspa?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-1-spons&ufe=app_do%3Aamzn1.fos.5e544547-1f8e-4072-8c08-ed563e39fc7d&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&psc=1",
		},
		{
			Title:     "Smartbox - Caja Regalo SPA y Relax para Dos",
			ImageName: "imagen2.jpg",
			Link:      "https://www.amazon.es/Smartbox-regalo-actividad-bienestar-personas/dp/B07Z5LKT1Q/ref=sr_1_3_sspa?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2K31EQHS5C13O&keywords=smart%2Bbox&qid=1704302305&sprefix=smart%2Bbox%2Caps%2C169&sr=8-3-spons&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&th=1",
		},
		{
			Title:     "ÍNTIMOOS - El Mejor Juego para Parejas- Aniversario Regalos Originales, para 2 jugadores",
			ImageName: "imagen3.jpg",
			Link:      "https://www.amazon.es/GUATAFAC-Juegos-Pareja-%C3%8DNTIMOOS-Aniversario/dp/B0BG8X2W2M/ref=sr_1_5?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-5",
		},
		{
			Title:     "Sin Rechistar - Juego de Parejas - Juego de Parejas mas Divertido de España para Vivir Momentos Inolvidables - Regalos Originales - Regalo Aniversario Pareja",
			ImageName: "imagen4.jpg",
			Link:      "https://www.amazon.es/SIN-RECHISTAR-Inolvidables-Originales-Aniversario/dp/B0B84LGPGT/ref=sr_1_9?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-9",
		},
	},
}

func main() {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./generated/pagina.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
