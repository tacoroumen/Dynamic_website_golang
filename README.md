# Go Webserver Project

Dit is een Dynamic Website gebasseerd op Golang dit is gemaakt voor het Vakantiepark Fonteyn. In de Read me file wordt vertelt wat je moet doen als er een fout code komt. (dit kan altijd bewerkt worden.) 

## Aan de slag

### Vereisten

- Docker geïnstalleerd 
- Configuratiebestand (`config.json`) voor database- en mailinstellingen
    tijdens het maken van de Email config moet je een App wachtwoord maken (als men een google mail gebruikt). Zonder dat kan je geen email sturen.
- (link https://support.google.com/accounts/answer/185833?visit_id=638385731453045893-3032648351&p=InvalidSecondFactor&rd=1)

## Problemen

### Config.json problemen
 1. Fout bij het lezen van het configuratiebestand
    Oplossing:
      **1.1 Bestandsnaam of pad incorrect:** 
     Controleer of de opgegeven bestandsnaam en het pad naar het configuratiebestand correct zijz. Zorg ervoor dat het bestand daadwerkelijk bestaat op de gespecificeerde locatie.
    
    **1.2 Configuratiebestand bevat syntaxisfouten:** 
    Bekijk de inhoud van het configuratiebestand en zorg ervoor dat het geldige JSON-syntaxis heeft.
    Controleer op ontbrekende komma's, accolades of ongeldige tekens in het configuratiebestand.

    **1.3 Bestand beschadigd of niet toegankelijk:** 
    Controleer of het configuratiebestand niet beschadigd is en toegankelijk is voor het programma.

2. **fout bij het decoderen van JSON**
    Oplossing:
     **2.1 Quotes rondom sleutels en waarden:**
    Controleer of alle sleutels en waarden in het JSON-bestand tussen dubbele aanhalingstekens staan. JSON vereist dat sleutels en waarden tussen aanhalingstekens worden geplaatst.
    **2.2 JSON-validator gebruiken:**
    Maak gebruik van online JSON-validatietools om te controleren of het configuratiebestand een geldig JSON-formaat heeft.

3. **fout bij het laden van configuratie**
    Oplossing:
     **3.1 Bestandstoegangsrechten Controleren:**
     Zorg ervoor dat het bestand toegankelijk is voor het uitvoerbare bestand. Controleer de bestandsrechten en pas ze indien nodig aan.
     **3.2 Controleer Bestandsnaam en Pad:**
    Zorg ervoor dat de bestandsnaam en het pad naar het configuratiebestand correct zijn gespecificeerd.

### Fout met het laden van html
1. **error executing Forms**
   **1.1 Oplossing:Controleren op Template Bestaan:**
   Zorg ervoor dat het templatebestand reservering.gohtml daadwerkelijk bestaat op de verwachte locatie. De naam en het pad moeten overeenkomen met wat je verwacht in de filepath.Join.
   **1.2 Debugtools gebruiken:** 
   Maak gebruik van debugtools of print-statements om de waarde van variabelen op verschillende punten in je code te controleren en te zien waar de fout mogelijk optreedt.
2. **error executing template**
    Oplossing:
    **2.1 Controleren op lege templates:** 
    Zorg ervoor dat het template dat je probeert uit te voeren daadwerkelijk HTML-inhoud bevat. Als het leeg is of ongeldige HTML bevat, kan dit fouten veroorzaken    
    **2.2 Variabele details controleren:** 
    Controleer of de variabele details die je doorgeeft aan het template correct is gevuld en de verwachte structuur heeft.
    