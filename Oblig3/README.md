# Obligatorisk oppgave 3, Group _Midgets_

##### _Tobias Albert_, _Espen Thorsen Frank_, _Benjamin Vraspilai_, _Fredrik Svartvatn_

----------------------------------------------------------------------------------

#### **Merk: Alle oppgavene har blitt testet i windows cmd, og vil ikke alltid fungere riktig i Goland IDE ettersom Goland og cmd tolker relative path litt forskjellig.**

----------------------------------------------------------------------------------

### Oppgave 1  

##### I den første oppgaven skulle vi opprette en lokal webserver. Vi fulgte en ganske bra tutorial for å hjelpe oss med å fullføre oppgaven. Den gikk igjennom step by step hvordan man kunne sette opp en webserver som lytter på 8080, og senere skrive inn "Hello, client".
> ##### Tutorial: https://golang.org/doc/articles/wiki/

#### Vi løste oppgven med følgende kode:
#### [/oppgave1/main.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave1/main.go)

#### I browser:  
![Hello client image](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/1/hello.png?raw=true)



----------------------------------------------------------------------------------

### Oppgave 2

##### I denne oppgaven skulle vi bruke 5 ulike json APIer. Det å finne 5 forstålige API var litt vanskelig men etter en del leting på blant annet https://www.ssb.no/ fant vi noen vi kunne bruke.

##### Vi løste oppgaven med følgende kode:
##### [/oppgave2/main.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave2/main.go)

##### Kort fortalt henter programmet json-data fra de ulike API-ene, hvor hver side har typer definert i '/pages/pageX.go' (X = 1-5), og hver funksjon `PageX` i main.go plasserer json-dataen inn i to slices, `names` og `values`. I funksjon `loadTemplate` blir verdiene i `names` og `values` plassert i en 'type Template', og deretter blir [templaten](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave2/page-template.html) brukt på den nylig genererte type Template. På denne måten kan templaten bli brukt på alle API-ene, selv om de i utgangspunktet krever forskjellige structs (definert i '/pages/pageX.go').

#### **De 5 API-ene vi brukte og hvordan dataen ser ut etter å ha blitt plassert i template:**
##### PS: på noen av sidene er noen verdier '-missing-', fordi disse verdiene var noen ganger tomme i json-dataen.  
#### page1: http://api.open-notify.org/astros.json  
![page1](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page1.png?raw=true)  

#### page2: https://hotell.difi.no/api/json/difi/geo/kommune  
![page2](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page2.png?raw=true)  

#### page3: https://hotell.difi.no/api/json/fad/reise/utland?  
![page3](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page3.png?raw=true)  

#### page4: https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8  
![page4](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page4.png?raw=true)  

#### page5: https://data.norge.no/api/dcat/data.json?page=1  
##### I page 5 ser vi html-kode i 'Beskrivelse', som blir vist i klartekst (i stedet for å bli lest som HTML av browseren) fordi datasettet ikke inneholder tegnene '<' og '>' i klartekst, men en ASCII-representasjon av tegnene i stedet (`\u003C` og `\u003E`).
![page5](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page5.png?raw=true)  



------------------------------------------------------------------------------------

### Oppgave 3

##### I oppgave 3 skulle vi implementere et serverprogram i henhold til [RFC 865], som er en Quote Of The Day (QOTD) protocol. Serveren bruker port 17, og svarer på både UDP og TCP. Ved en forbindelse til UDP eller TCP henter den et quote fra en json-fil i samme mappe (filen er data hentet fra en API 20.04.18, men har blitt lagret lokalt ettersom kilden har en grense på 10 besøk per time, som viste seg å være ganske upraktisk). UDP og TCP gir ulike sitater.  

##### Funksjonene `connTCP` og `connUDP` blir kjørt med 'go', som gjør det mulig for en client å alltid koble seg til serveren (også flere clienter samtidig). Det er dermed også lagt inn 10 minutters ventetid på slutten av `main` for å forhindre serveren til å avslutte umiddelbart.

##### UDP connection:  
![UDP](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/3/udp_done.png?raw=true)  

##### TCP connection:  
![UDP](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/3/tcp_done.png?raw=true)  

##### Merk at TCP-forbindelsen blir avbrutt etter at sitatet er sendt, slik det er beskrevet i [RFC 865], mens UDP-forbindelsen fortsatt er åpen.  

##### Kode: [/oppgave3/main.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave3/main.go)


##### UDP- og TCP-clienter er vedlagt i [/oppgave3](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave3).

[RFC 865]: https://tools.ietf.org/html/rfc865
