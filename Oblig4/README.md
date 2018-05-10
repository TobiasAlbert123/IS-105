## oblig_4

I denne oppgaven har vi hatt litt friere tøyler enn fra de tidligere oppgavene. Her skulle vi lage en web applikasjon som blir hostet på en lokal server og skal gi useren nyttig informasjon om f.eks vær som vi selv kunne velge. Etter endel research og forslag kom vi fram til at vi skulle bruke en ISS plassering API.

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
