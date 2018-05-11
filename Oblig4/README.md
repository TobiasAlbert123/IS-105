## oblig_4

I denne oppgaven har vi hatt litt friere tøyler enn fra de tidligere oppgavene. Her skulle vi lage en web applikasjon som blir hostet på en lokal server og skal gi useren nyttig informasjon om f.eks vær som vi selv kunne velge. Etter endel research og forslag kom vi fram til at vi skulle bruke en ISS plassering API.

### Disclaimers:  
- placeholder

## Systembeskrivelse

Hensikten med systemet vårt er å vise plasserings til ISS (The International Space Station), i tillegg til noen andre detaljer om romstasjonen. Vi får informasjon om hvem som er om bord og hva de heter, og hvis man trykker på navnet deres vil man bli tatt til deres Wikipedia-side. I tillegg til dette får vi informasjon om hvilket klokkeslett det er, hvilken tidssone romstasjonen befinner seg i, høyde, breddegrad og lengregrad, og havdybde (dersom det er noe hav under). For å få den mest oppdaterte informasjonen som er mulig, blir romstasjonens plassert oppdatert hvert 15. sekund.

Personer som vil ha størt nytte av dette systemet er romentusiaster, eller andre, som ønsker å vite hvor den internasjonale romstasjonen befinner seg til enhver tid, og hvem som befinner seg ombord.

## Systemarkitektur

Vårt systems hovedfunksjon er å vise plasseringen til ISS(the International Space Station), ved hjelp av en API som gir oss latitude, longitude og et UNIX timestamp. API-en: http://api.open-notify.org/iss-now.json.  
Mer data om posisjonen, som land, høyde over havet og tidssone blir hentet fra ulike Google API-er:  
- https://developers.google.com/maps/documentation/elevation/intro
- https://developers.google.com/maps/documentation/geocoding/intro#ReverseGeocoding
- https://developers.google.com/maps/documentation/timezone/intro  

Applikasjonen vi har utviklet er en Single-page applikasjon. Den har ingen user-input og self oppdaterer hvert 15. sekund. Mer konkret er applikasjonen en Thin Server Architecture.  
I vår applikasjon bruker vi Go, HTML og CSS sammen til å gi oss resultatet vi ønsker.  
Vi bruker Go hovedsakelig til å handle json API. Json blir hentet i `getJson` og unmarshallet i `formatJson`. `getJson` blir brukt på alle de ulike API-ene, og `formatJson` blir brukt en gang, og setter all data fra de ulike API-ene inn i structen `issData`. Denne structen blir deretter sendt inn i funksjonen `renderTemplate` som gjennom en `http.ResponseWriter` bruker structen på en template og produserer nettsiden som blir vist på 'http://localhost:8080/'.


For å finne elapsed time til de ulike astronautene bruker vi time.Date func i time packagen. Det som er nyttig med denne funksjonen er at man kan oppgi en start date og finne tiden som har gått fra start date til nå. Funksjonen blir kjørt hver gang siden oppdateres, og siden start date aldri fordandres vil elapsed time alltid være oppdatert.  


HTML blir brukt til å sette opp siden, og CSS blir brukt til å kontrolle forskjellige objekter som HTML produserer. CSS-en brukt bestemmer hvor på siden de ulike objektene skal være, og hvordan de skal se ut (størrelse, font, farger osv.).
Siden er delvis responsiv (dvs. sideoppsettet tilpasser seg til skjermstørrelsen), men dette er ikke et stort fokus på oppgaven, og er dermed langt fra perfekt. Dette ble gjort i CSS ved å bruke `@media only screen and` med max-width og max-height.  


Etter at siden har kjørt i 15 sekunder, laster go-serveren opp en html-fil med en linje javascript som fjerner alt innhold på skjermen, og deretter blir hele siden lastet inn på nytt. Dette blir gjort slik at dataen på siden blir hentet fra API-ene hvert 15.sekund og deretter oppdatert på siden. Uten javascriptet risikerer vi at den forrige siden ikke blir borte når den nye printes, og den nye blir bare printet ved slutten av den forrige.

## Enhetstester  

Siden applikasjonen ikke inneholder noen user input, er det vanskelig å lage gode tester. Testene vi har laget sjekker at APIen til ISS var brukbar (resten av programmet baserer seg på den), og om det er nok API-nøkler til å bruke de ulike Google API-tjenestene. Ellers er det brukt mye defensiv programmering i koden. Hver gang en API blir hentet blir `status` sjekket, og om den ikke er lik `OK`, vil programmet ikke prøve å bruke data som ikke eksisterer, men heller produsere sin egen (for eksempel: "Country: N/A").  
Det er også en test som sjekker om `globalError` har innhold, fordi den starter tom og innhold blir bare lagt til når en error forekommer. Programmet feiler og printer ut `globalError`. Det er verdt å merke at alle feilmeldingene som kommer, kommer mange ganger (rundt 9-10 ganger) fordi programmet er kodet til å prøve om igjen rundt 10 ganger

Alle steder hvor en error blir initalisert, har funksjonskall som `log.Fatal()` blitt kommentert vekk i iss.go. Med disse fungerte ikke testene, ettersom programmet ble avsluttet før testene fikk noe data.
