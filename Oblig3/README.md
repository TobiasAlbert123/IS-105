# Obligatorisk oppgave 3 Midgets

##### _Tobias Albert_, _Espen Thorsen Frank_, _Benjamin Vraspilai_, _Fredrik Svartvatn_

----------------------------------------------------------------------------------

### Oppgave 1  

##### I den første oppgaven skulle vi opprette en lokal webserver. Vi fulgte en ganske bra tutorial for å hjelpe oss med å fullføre oppgaven. Den gikk igjennom step by step hvordan man kunne sette opp en webserver som lytter på 8080, og senere skrive inn "Hello Client!".
https://golang.org/doc/articles/wiki/

#### Vi løste oppgven med følgende kode:
#### [/oppgave1/main.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave1/main.go)

#### I browser:  


----------------------------------------------------------------------------------

### Oppgave 2

##### I denne oppgaven skulle vi bruke 5 ulike json APIer. Det å finne 5 forstålige API var litt vanskelig men etter veldig mye leting på https://www.ssb.no/ fant vi noen vi kunne bruke
##### Disse 5 ble brukt
#### page1: http://api.open-notify.org/astros.json  
![page1](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page1.png?raw=true)  

#### page2: https://hotell.difi.no/api/json/difi/geo/kommune  
![page2](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page2.png?raw=true)  

#### page3: https://hotell.difi.no/api/json/fad/reise/utland?  
![page3](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page3.png?raw=true)  

#### page4: https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8  
![page4](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page4.png?raw=true)  

#### page5: https://data.norge.no/api/dcat/data.json?page=1  
![page5](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/images/2/page5.png?raw=true)  


#### Vi løste oppgaven med følgende kode:
#### [/oppgave2/main.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig3/src/oppgave2/main.go)

------------------------------------------------------------------------------------

### Oppgave 3

##### I oppgave 3 skulle vi implementere et serverprogram i henhold til RFC 865, som er en Quote Of The Day (QOTD) protocol. Serveren bruker port 17, og svarer på både UDP og TCP. UDP- og TCP-clienter er vedlagt i /oppgave3.
