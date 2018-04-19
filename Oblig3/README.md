# Oppgave 1
### a) Sett opp en webserver som lytter til port 8080.
### b) Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client".
### Strengen skal vises i nettleseren.  

# Oppgave 2 API kall og behandling av JSON og HTML templates
### a) Presenter data på webserveren slik at den er leselig for mennesker(Ren tekst, f.eks. "Sted: Jorenhholem, dato: 13.04.2018", ikke i curly brackets.) når klienter aksesserer stiene /1, /2, /3, /4, /5.

### Dataen som skal presenteres skal hentes fra fem ulike APIer, hvor alle returnerer data i JSON format. Dere velger selv hvilke datasett dere benytter. Det er denne dataen som skal presenteres på de ulike stiene på webserveren.

### Et eksempel på et API som returnerer JSON data; https://hotell.difi.no/api/json/stavanger/miljostasjoner



### b) Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates.



# Oppgave 3 UDP/TCP server og Internett standarder
### Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.
