# web-generator
Massive SEO-optimized static WEB generation tool. This tool will allow you to generate optimized pages in quantity with minimal effort.
Expects an XML input like 
```
<WebData>
    <Title>Mi página de prueba</Title>
    <Body> Bienvenido a "Encuentra el Regalo Perfecto", tu destino definitivo para encontrar ideas de regalos únicas y personalizadas para parejas. Entendemos que cada relación es única, y encontrar el regalo perfecto para tu ser querido puede ser un desafío. Por eso, hemos creado una plataforma intuitiva y fácil de usar que te ayuda a descubrir el regalo ideal que refleje tu amor y aprecio. </Body>
    <Products>
        <Product>
          <Title> Smartbox - Caja Regalo SPA y Relax para Dos </Title>
          <ImageName>imagen1.jpg</ImageName>
          <Url> https://www.amazon.es/Smartbox-Mujer-Momentos-inolvidables-Ideas-Originales-1-Unisex-Adult/dp/B0BL9T9DQC/ref=sr_1_1_sspa?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-1-spons&ufe=app_do%3Aamzn1.fos.5e544547-1f8e-4072-8c08-ed563e39fc7d&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&psc=1 </Url>
        </Product>
        <Product>
          <Title> Smartbox - Caja Regalo SPA y Relax para Dos</Title>
          <ImageName>imagen2.jpg</ImageName>
          <Url> https://www.amazon.es/Smartbox-regalo-actividad-bienestar-personas/dp/B07Z5LKT1Q/ref=sr_1_3_sspa?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2K31EQHS5C13O&keywords=smart%2Bbox&qid=1704302305&sprefix=smart%2Bbox%2Caps%2C169&sr=8-3-spons&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&th=1</Url>
        </Product>
        <Product>
          <Title> ÍNTIMOOS - El Mejor Juego para Parejas- Aniversario Regalos Originales, para 2 jugadores </Title>
          <ImageName>imagen3.jpg</ImageName>
          <Url> https://www.amazon.es/GUATAFAC-Juegos-Pareja-%C3%8DNTIMOOS-Aniversario/dp/B0BG8X2W2M/ref=sr_1_5?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-5</Url>
        </Product>
        <Product>
          <Title> Sin Rechistar - Juego de Parejas - Juego de Parejas mas Divertido de España para Vivir Momentos Inolvidables - Regalos Originales - Regalo Aniversario Pareja </Title>
          <ImageName>imagen4.jpg</ImageName>
          <Url> https://www.amazon.es/SIN-RECHISTAR-Inolvidables-Originales-Aniversario/dp/B0B84LGPGT/ref=sr_1_9?__mk_es_ES=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2IWUDWAQB0J5&keywords=pareja&qid=1704302331&sprefix=pareja%2Caps%2C150&sr=8-9</Url>
        </Product>
    </Products>
</WebData>
```
