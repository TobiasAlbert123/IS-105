# Obligatorisk oppgave 2 Midgets

##### _Tobias Albert_, _Espen Thorsen Frank_, _Benjamin Vraspi_, _Fredrik Svartvatn_

----------------------------------------------------------------------------------

### Oppgave 1  

##### I denne oppgaven skulle vi lage et program som skal kunne hente inn en fil og returnere filstørrelsen i bytes, KB, MB og GB. I tillegg til det skal programmet kunne sjekke en rekke fler detaljert informasjon om en fil, for eksempel om filen er et directory.
#### Vi løste oppgven med følgende kode: 
##### [filecount.go](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/src/oppgave2/filecount.go)


##### Her er filen kjørt i cmd på text.txt:  

##### ![Bilde av oppgave 1 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave1.png "Oppgave 1 resultat i cmd")

##### Filen kan eventuelt bli kjørt i cmd direkte fra [fileinfo.bat](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave1/fileinfo.bat)

#### b)

##### [fileinfo.exe](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin/fileinfo.exe)-filen ligger i [bin](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin)-mappen på git. Filstien for å finne text.txt er relativ, og programmet vil lete i [files](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/files)-mappen på git etter filen du vil bruke. For å kjøre .exe filen skriver man kun 'fileinfo text.txt'.

------------------------------------------------------------------------------------

### Oppgave 2

##### Her var oppgaven å finne antall linjer i en fil, samt finne de fem mest brukte runes i en .txt fil. Utfordringen her var å få programmet til å kjøre uten at det tok veldig lang tid.
#### Vi løste oppgaven med følgende kode:
##### Vi tar et argument når programmet blir kjørt som blir satt lik filnavnet, i en funksjon getFile(). Deretter bruker to funksjoner getFile(). Den første funksjonen, countLines(), bruker os.Open() på filen og bufio.NewScanner() og teller hver gang den gjør en Scan() på filen, da finner man antall linjer i filen. Funksjon nummer 2, MapMaker(), blir også kjørt på filen og bruker os.ReadFile(). Filen blir dermed konvertert til en string, og så splittet opp i en string slice med strings.Split. Hver posisjon i slicen tilsvarer én karakter / rune i filen. Deretter legger vi runer vi finner i filen inn et map, og legger til 1 til int-verdien for en rune hver gang vi ser den en gang til. Eksempel: vi finner 'a' i filen, 'a' i mapet : [a:1] blir til [a:2]. I xMostUsed går vi gjennom (med en for-loop) alle runene vi har funnet, og hver gang vi finner en som har blitt skrevet flere ganger enn 'highestCount', setter vi 'highestCount' lik den verdien. Når vi er ferdig printer vi ut 'highestCount' og tilsvarende rune. Runen blir deretter slettet fra mapet, og den for-loopen x ganger (x er satt til 5 i main())
##### Se [filecount.go](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/src/oppgave2/filecount.go) for koden.  


##### Her er filen kjørt i cmd på text.txt:  

##### ![Bilde av oppgave 2 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave2.png "Oppgave 2 resultat i cmd")

##### [filecount.exe](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin/filecount.exe) filen ligger i [bin](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin)-mappen på git. Filstien for å finne text.txt er relativ, og programmet vil lete i [files]()-mappen på git etter filen du vil bruke. For å kjøre .exe filen skriver man kun 'fileinfo text.txt'.

------------------------------------------------------------------------------------

### Oppgave 3

##### I oppgave 3


#### a)

##### Se [addup.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/addup.go) for koden.  

##### ![Bilde av oppgave 3a ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave3a.png "Oppgave 3a resultat i cmd")

#### b)

##### Se [3b.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/3b_run/3b.go), [addtofile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/addtofile.go) og [sumfromfile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/sumfromfile.go) for .go-filer    

##### ![Bilde av oppgave 3b ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave3b.png "Oppgave 3b resultat i cmd")

##### c)




#### d)

##### Eksempler på SIGINT-signalet (CTRL+C) brukt i oppgave 3a og 3b:  

##### ![Bilde av oppgave 3c1 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/sigint3a.png "Oppgave 3c resultat i cmd")  
_ _  
##### ![Bilde av oppgave 3c2 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/sigint3b.png "Oppgave 3c resultat i cmd")

#### e)

##### Se [bin-mappen](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin). Disse filene ble laget ved å skrive 'go build <filnavn>' i cmd i riktig directory.