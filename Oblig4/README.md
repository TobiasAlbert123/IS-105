oblig_4

I denne oppgaven har vi hatt litt friere tøyler enn fra de tidligere oppgavene. Her skulle vi lage en web applikasjon som blir hostet på en lokal server og skal gi useren nyttig informasjon om f.eks vær som vi selv kunne velge. Etter endel research og forslag kom vi fram til at vi skulle bruke en ISS plassering API.

## Systemarkitektur

Vårt systems hovedfunksjon er å vise plasseringen til ISS(The International Space Station), ved hjelp av en API som gir oss altitude,
longitude og timestamp.  
Applikasjonen vi har utviklet er en Single-page applikasjon. Den har ingen user-input og self oppdaterer hvert 15. sekund.   
Mer konkret er applikasjonen en Thin Server Architecture hvor kompleksiteten blir flyttet fra server sia over til klient side.   
Dette vil overall reduserer kompleksiteten til systemet.  
I vår applikasjon bruker vi Go, HTML og CSS. 
Vi bruker Go til 
HTML--->
CSS--->
