## oblig_4

I denne oppgaven har vi hatt litt friere tøyler enn fra de tidligere oppgavene. Her skulle vi lage en web applikasjon som blir hostet på en lokal server og skal gi useren nyttig informasjon om f.eks vær som vi selv kunne velge. Etter endel research og forslag kom vi fram til at vi skulle bruke en ISS plassering API.

## Systembeskrivelse

Hensikten med systemet vårt er å vise plasserings til ISS (The International Space Station). Vi får informasjon om hvem som er om bord og hva de heter, og hvis man trykker på navnet deres vil man bli tatt til deres Wikipedia-side. I tillegg til dette får vi informasjon om hvilket klokkeslett det er, hvilken tidssone romstasjonen befinner seg i, høyde, breddegrad og lengregrad, og havdybde (dersom det er noe hav under). 

## Systemarkitektur

Vårt systems hovedfunksjon er å vise plasseringen til ISS(The International Space Station), ved hjelp av en API som gir oss latitude,
longitude, elevation og timestamp. http://api.open-notify.org/iss-now.json     
Applikasjonen vi har utviklet er en Single-page applikasjon. Den har ingen user-input og self oppdaterer hvert 15. sekund.   
//Mer konkret er applikasjonen en Thin Server Architecture.  
I vår applikasjon bruker vi Go, HTML og CSS sammen til å gi oss resultatet vi ønsker.  
Vi bruker Go hovedsakelig til å handle json API--->  
For å finne elapsed time til de ulike astronautene bruker vi time.Date func i time packagen. Det som er nyttig med denne funksjonen er at man kan oppgi en start date fra når tickeren skal begynne fram til present time. Dette vil alltid gi oss en oppdatert tid ettersom vi har satt start tiden til deres respektive oppskytning.   
HTML--->  
CSS--->  


## Enhetstester  

Siden applikasjonen ikke inneholder noen user input, er det vanskelig å lage gode tester. Testene vi har laget sjekker at APIen til ISS var brukbar (resten av programmet baserer seg på den), og om det er nok API-nøkler til å bruke de ulike Google API-tjenestene. Ellers er det brukt mye defensiv programmering i koden. Hver gang en API blir hentet blir `status` sjekket, og om den ikke er lik `OK`, vil programmet ikke prøve å bruke data som ikke eksisterer, men heller produsere sin egen (for eksempel: "Country: N/A").
