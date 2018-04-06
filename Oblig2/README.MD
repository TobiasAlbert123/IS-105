# Obligatorisk oppgave 2 Midgets

##### _Tobias Albert_, _Espen Thorsen Frank_, _Fredrik Svartvatn_, _Benjamin Vraspillai_

----

### Oppgave 1  

#### a)

I denne oppgaven skulle vi lage et program som skal kunne hente inn en fil og returnere filstørrelsen i bytes, KB, MB og GB. I tillegg til det skal programmet kunne sjekke en rekke fler detaljert informasjon om en fil, for eksempel om filen er et directory.
##### Vi løste oppgven slik:

Vi tar et argument når programmet blir kjørt som blir satt lik filnavnet, i en funksjon `getFile()`. Deretter bruker to funksjoner `getFile()`. Først har vi `fileSize()`, som bruker `os.Stat().Size()` av filen, og printer ut antall byte som int. Deretter konverterer vi til float64 (for å få desimaltall) og deler antall bytes på 1024, 1024^2 og 1024^3 for å KB, MB og GB og printer. Deretter tar en ny funksjon, `fileInfo()`, `os.Lstat().Mode()` av filen, og finner diverse info om filen ved funksjoner som `IsDir()` og printer informasjonen.

Koden er i [fileinfo.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave1/fileinfo.go) i src/Oppgave1.


##### Her er filen kjørt i cmd på text.txt:  

![Bilde av oppgave 1 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave1.png "Oppgave 1 resultat i cmd")

Filen kan eventuelt bli kjørt i cmd direkte fra [fileinfo.bat](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave1/fileinfo.bat).

#### b)

[fileinfo.exe](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin/fileinfo.exe)-filen ligger i [bin](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin)-mappen på git. Filstien for å finne text.txt er relativ, og programmet vil lete i [files](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/files)-mappen på git etter filen du vil bruke. For å kjøre .exe filen skriver man kun 'fileinfo text.txt'.

----

### Oppgave 2

#### a)

Her var oppgaven å finne antall linjer i en fil, samt finne de fem mest brukte runes i en .txt fil. Utfordringen her var å få programmet til å kjøre uten at det tok veldig lang tid.

##### Vi løste oppgaven slik:
Vi tar et argument når programmet blir kjørt som blir satt lik filnavnet, i en funksjon `getFile()`. Deretter bruker to funksjoner `getFile()`. Den første funksjonen, `countLines()`, bruker `os.Open()` på filen og `bufio.NewScanner()` og teller hver gang den gjør en `Scan()` på filen, da finner man antall linjer i filen. Funksjon nummer 2, MapMaker(), blir også kjørt på filen og bruker `os.ReadFile()`. Filen blir dermed konvertert til en string, og så splittet opp i en string slice med `strings.Split()`. Hver posisjon i slicen tilsvarer én karakter / rune i filen. Deretter legger vi runer vi finner i filen inn et map, og legger til 1 til int-verdien for en rune hver gang vi ser den en gang til. Eksempel: vi finner 'a' i filen, 'a' i mapet : [a:1] blir til [a:2]. I xMostUsed går vi gjennom (med en for-loop) alle runene vi har funnet, og hver gang vi finner en som har blitt skrevet flere ganger enn `highestCount`, setter vi `highestCount` lik den verdien. Når vi er ferdig printer vi ut 'highestCount' og tilsvarende rune. Runen blir deretter slettet fra mapet, og den for-loopen x ganger (x er satt til 5 i `main()`).  

Koden er i [filecount.go](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave2/filecount.go) i src/Oppgave2.  


Her er filen kjørt i cmd på text.txt:  

![Bilde av oppgave 2 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave2.png "Oppgave 2 resultat i cmd")  

Filen kan eventuelt bli kjørt i cmd direkte fra [filecount.bat](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave2/filecount.bat).

#### b)

[filecount.exe](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin/filecount.exe) filen ligger i [bin](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin)-mappen på git. Filstien for å finne text.txt er relativ, og programmet vil lete i [files](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/files)-mappen på git etter filen du vil bruke. For å kjøre .exe filen skriver man kun 'filecount text.txt'.

----

### Oppgave 3

#### a)

Se [addup.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/addup.go) for koden.  

![Bilde av oppgave 3a ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave3a.png "Oppgave 3a resultat i cmd")  

Filen kan eventuelt bli kjørt i cmd direkte fra [addup.bat](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave3/addup.bat).

#### b)

I denne oppgaven har vi to filer, [addtofile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/addtofile.go) og [sumfromfile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/sumfromfile.go), som skal kommunisere gjennom en fil (her kalt readwritenumbers.txt). For å løse oppgaven har vi laget et program [3b.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/3b_run/3b.go) i mappen 3b_run, som kaller på funksjoner i de andre programmene. De andre programmene ligger da i mappen notmain og har pakke notmain, slik at de kan bli importert til 3b.go. 3b.go koordinerer de andre programmene til å:
1. Lese inn to tall fra terminal
2. Skrive inn tallene i en fil
3. Lese tallene fra filen
4. Summere tallene
5. Skrive inn summen i filen
6. Lese summen fra filen
7. Skrive ut summen

Se [3b.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/3b_run/3b.go), [addtofile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/addtofile.go) og [sumfromfile.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/notmain/sumfromfile.go) for koden.   

![Bilde av oppgave 3b ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/oppgave3b.png "Oppgave 3b resultat i cmd")  

Filen kan eventuelt bli kjørt i cmd direkte fra [3b.bat](https://github.com/TobiasAlbert123/IS-105/blob/master/Oblig2/src/oppgave3/3b_run/3b.bat).

#### c)

##### Feilhåndtering i 3a):  
Når input blir scannet fra terminalen, blir en error variabel initialisert. Hvis det er en error, (for eksempel bokstaver i input i stedet for tall) printes en feilmelding ut, og programmet avsluttes. Siden programmet prøver å scanne en variabel av type int vil det bare gå videre om input er av type int.  

##### Eksempel:  
![3a errormelding](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/error3a.png "3a errormelding")

Alle err i dette programmet har en delay på 2 sekunder `(time.Sleep(2*time.Second)`) for å forhindre at programmet avslutter for tidlig når et SIGINT-signal blir motatt, som gjør at kun errormelding blir printet men ikke SIGINT-meldingen.  

##### Feilhåndtering i 3b):
Her blir en error initialisert når:
- Tall blir scannet fra terminal  
  * gir error om typen ikke er int  
- Filen blir laget  
- Filen blir skrevet til  
  * f.eks. om tall ikke er int eller ikke blir funnet  
- Filen blir lest  
  * gir error om filen ikke eksisterer eller om filnavn ikke er string  
- String blir omgjort til int (sctrconv.Atoi) etter filen blir lest  
  * hvis typen den konverterer ikke er et tall som string  

##### Eksempel:  
![3a errormelding](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/error3b.png "3a errormelding")

Alle err i dette programmet har en delay på 2 sekunder `(time.Sleep(2*time.Second)`) for å forhindre at programmet avslutter for tidlig når et SIGINT-signal blir motatt, som gjør at kun errormelding blir printet men ikke SIGINT-meldingen.




#### d)

SIGINT blir implementert ved å ha en go-funksjon som lytter etter et SIGINT-signal, og printer ut en avslutningsmelding og avslutter programmet når SIGINT-signalet blir sendt. Se `main()` og `stop()` i [addup.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/addup.go) og [3b.go](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/oppgave3/3b_run/3b.go).

Eksempler på SIGINT-signalet (CTRL+C) brukt i oppgave 3a og 3b:  

![Bilde av oppgave 3c1 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/sigint3a.png "Oppgave 3c resultat i cmd")  
_ _  
![Bilde av oppgave 3c2 ikke funnet](https://raw.githubusercontent.com/TobiasAlbert123/IS-105/master/Oblig2/images/sigint3b.png "Oppgave 3c resultat i cmd")

#### e)

Se [bin-mappen](https://github.com/TobiasAlbert123/IS-105/tree/master/Oblig2/src/bin). Disse filene ble laget ved å skrive '`go build <filnavn>`' i cmd i riktig directory.