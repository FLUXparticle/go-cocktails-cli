# Cocktail-CLI

## Aufgaben

### Vorbereitung für die Remote-Aufgaben

Trage vor dem ersten Remote-Test deine Benutzer-ID in die zunächst leere
Konstante `fridgeUser` in `pkg/cocktails/repository/remote/RemoteRepository.go`
ein. Der Konstruktor des `RemoteRepository` löst bei leerer Benutzer-ID
eine Panic aus. Die CLI verwendet dieselbe ID später für den Kühlschrank.

### Aufgabe 1 – Milchmenge in Cocktails berechnen

Implementiere im `CocktailService` die Methode `MilkAmount() (int, error)`, die die gesamte Milchmenge
aller Cocktails berechnet, deren Name `Milk` enthält. Die Methode gibt die
Summe in Zentilitern und einen Fehler zurück.

1. Rufe über das Repository die Cocktail-Liste ab und filtere die Namen nach
   der Teilzeichenfolge `Milk`.
2. Lade die Details der gefundenen Cocktails anhand ihrer ID. Gehe die
   `instructions` durch und addiere alle `amount`-Werte, bei denen
   `ingredient.name` genau `Milch` ist.
3. Gib einen Fehler aus dem Repository an den Aufrufer weiter.

Schreibe im Service-Paket einen Go-Test in einer Datei mit der Endung
`_test.go`. Erzeuge den Service mit einem `RemoteRepository` und übergib ihm
dieses über `NewCocktailService`. Rufe im Test die neue Methode auf und prüfe
mit `testing`, dass sie ohne Fehler genau **38 cl** liefert. Gib bei einer
Abweichung den berechneten und den erwarteten Wert aus. Die Filterung und
Summierung gehören in die Service-Methode. Verwende im Test `remote.NewRemoteRepository()`. Der Test ruft den
öffentlichen Kursserver auf und benötigt eine Internetverbindung.

#### Endpunkte der Remote-API

Basis-URL ist [cocktails.fluxparticle.com](https://cocktails.fluxparticle.com).
Ergänze im `RemoteRepository` den gemeinsamen HTTP-Abruf sowie die
Methoden für Liste und Details. Dekodiere die JSON-Antworten in die
Cocktail-Modelle. Die folgenden Antworten zeigen die benötigten Felder.

`GET /api/cocktails`

```json
[
  { "id": 123, "name": "Beispiel-Milk" }
]
```

`GET /api/cocktails/123`

```json
{
  "id": 123,
  "name": "Beispiel-Milk",
  "instructions": [
    {
      "amount": 15,
      "ingredient": { "id": 8, "name": "Milch" }
    }
  ]
}
```

### Aufgabe 2 – Rezeptdetails parallel abrufen

Ergänze im `CocktailService` die Methode `MilkAmountParallel() (int, error)`.
Rufe wie in Aufgabe 1 zuerst die Cocktail-Liste ab und filtere nach `Milk`.
Lade dann die Rezeptdetails der drei gefundenen Cocktails parallel.

1. Starte für jeden Treffer mit `sync.WaitGroup.Go` eine Goroutine.
2. Sende die Milchmenge eines erfolgreich geladenen Rezepts über einen
   `chan int`. Sende bei einem fehlgeschlagenen Abruf `-1`.
3. Warte in einer separaten Goroutine mit `Wait` auf alle Abrufe und
   schließe danach den Kanal.
4. Lies den Kanal mit `range` bis zum Ende und addiere die Teilmengen.
   Überspringe `-1` beim Addieren und lies den Kanal weiter aus. Wenn
   mindestens ein Abruf fehlgeschlagen ist, gib danach einen Fehler zurück.

Ergänze einen zweiten Test. Injiziere wieder `remote.NewRemoteRepository()`,
rufe `MilkAmountParallel` auf und prüfe für den echten Kursserver **38 cl**.

### Aufgabe 3 – Parallele Abrufe bei einem Fehler abbrechen

Implementiere `MilkAmountErrGroup(group *errgroup.Group) (int, error)`.
Rufe die Cocktail-Liste ab und starte für jeden Namen mit `Milk` den
Detailabruf mit `group.Go`. Erzeuge im Test die Gruppe und ihren Kontext
mit `errgroup.WithContext`. Übergebe den Kontext beim Erzeugen des
Remote-Repositorys, damit dessen HTTP-Requests ihn verwenden.
Speichere die Teilmengen an
getrennten Positionen in einem Slice und addiere sie nach einem
erfolgreichen `Wait`. Bei einem Fehler gib den Fehler der Gruppe zurück.

Addiere in dieser Methode vor jedem Detailabruf `45` zur Cocktail-ID.
Aus den drei Treffern `23`, `24` und `25` werden so die IDs `68`,
`69` und `70`.

Teste den Abbruch gegen den öffentlichen Kursserver. Die gültigen IDs
`68` und `69` werden mit einer Sekunde Verzögerung beantwortet. Für die
ungültige ID `70` soll der Kursserver sofort `404` liefern. Prüfe, dass die Methode diesen Fehler zurückgibt. Beobachte die im
Testlauf angezeigte Dauer mit `go test -v ./pkg/cocktails/service -run
TestMilkAmountErrGroupCancelsRequests`. Sie enthält auch das Laden der
Cocktail-Liste.

### Aufgabe 4 – Cocktail-CLI für die Remote-API

Schreibe unter `cmd/cocktail` ein CLI-Programm. Lies in `main` mit
`flag` die Cocktail-ID über `-id` oder `-i` ein und übergib den Wert
an `run`. Verwende den `CocktailService` mit `RemoteRepository`.
Mit ID zeigt das Programm das Rezept. Ohne ID listet es IDs und Namen
aller Cocktails. Formatiere die Ausgabe im CLI-Programm.

**Bonus:** Ergänze `cmd/ingredients` nach demselben Muster. Zeige ohne
ID die Zutatenliste und mit ID den Namen der Zutat und die zugehörigen
Cocktails. Implementiere dazu im Remote-Repository die Abrufe für
`GET /api/ingredients` und `GET /api/ingredients/{id}`.

```sh
go run ./cmd/cocktail
go run ./cmd/cocktail -i 24
# Bonus
go run ./cmd/ingredients -id 1
```

### Aufgabe 5 – Kühlschrank-CLI für die Remote-API

Schreibe das Programm unter `cmd/fridge`. Verwende
`remote.NewRemoteRepository()` für den Zugriff auf den Kursserver
und übergib das Repository über das Interface `FridgeRepository` an
den `FridgeService`. Führe die Kühlschrankaktionen über diesen Service
aus und lies die Optionen mit dem Standardpaket `flag` ein.

Die Remote-Zutatenabfragen aus dem Bonus von Aufgabe 4 sind in diesem
Checkpoint vorgegeben. Ergänze die Remote-Zugriffe für den
Kühlschrankbestand und dessen Änderungen.

Mit `-display` zeigt das Programm den Kühlschrankinhalt. Mit `-add ID`
fügst du eine Zutat hinzu, mit `-remove ID` entfernst du sie und mit
`-possible` zeigst du die möglichen Cocktails. Die Flags lassen sich
kombinieren. Das Programm führt die gesetzten Aktionen in dieser
Reihenfolge aus: `-add`, `-remove`, `-display`, `-possible`. Der
`FridgeService` berechnet die möglichen Cocktails aus dem
Kühlschrankbestand und den Rezeptdetails des Katalogs. Verwende für
jedes Rezept die ID aus dem Zutatenobjekt der Anweisung
(`instruction.Ingredient.IngredientID`).

**Hinweis zu `-possible`:** Nutze die Zutaten im Kühlschrank, um die
Kandidaten einzugrenzen. Lade jedes Rezept höchstens einmal, auch wenn der
Cocktail bei mehreren Zutaten als Treffer vorkommt. Zur Sichtprüfung kannst
du einen Kühlschrank mit Buttermilch (31) und Fruchtsirup Preiselbeere (8)
verwenden. Pink Power (27) kommt in beiden Zutatenlisten vor und sollte
einmal als möglicher Cocktail erscheinen.

Gib Fehler bei
Kühlschrankzugriffen auf der Standardfehlerausgabe aus und beende das
Programm mit einem Exitcode ungleich null.

```sh
go run ./cmd/fridge -display
go run ./cmd/fridge -add 27 -possible
go run ./cmd/fridge -remove 27
```

### Aufgabe 6 – Cocktail-CLI auf SQLite umstellen

Erzeuge die Datenbank mit `go run ./cmd/initdb`. Implementiere im
`DatabaseRepository` den Listen-Helfer `listCocktails`. Seine Aufrufer
übergeben bereits die SQL-Anweisung und die Parameter. Führe die
Abfrage aus und lies die Ergebniszeilen als Cocktail-Liste ein.

Implementiere auch `GetCocktail` vollständig. Die benötigten
SQL-Anweisungen stehen im TODO. Übergib die angefragte ID als
SQL-Parameter und lies das Rezept samt Zutaten ein. Stelle
`cmd/cocktail` auf `database.NewDatabaseRepository()` um und
schließe das Repository am Ende von `run`. Flags und Konsolenausgabe
aus Aufgabe 4 bleiben gleich.

**Bonus:** Implementiere `listIngredients` und
`GetCocktailsByIngredient`. Die SQL-Anweisung für die Zutatenliste
steht bereits im Aufruf. Die Abfragen für die Cocktails zu einer
Zutat stehen im TODO. Stelle auch `cmd/ingredients` auf SQLite um.

```sh
go run ./cmd/initdb
go run ./cmd/cocktail
go run ./cmd/cocktail -i 24
# Bonus
go run ./cmd/ingredients -id 1
```

### Aufgabe 7 – Kühlschrank im Speicher und Race Detector

Das In-Memory-Repository und der Test sind vorgegeben. Im Test fügt
eine Goroutine dieselbe Zutat 10.000-mal hinzu, während eine zweite
sie 10.000-mal entfernt. Welche Änderung zuletzt ausgeführt wird,
hängt von der Reihenfolge der Goroutinen ab.

Führe den Test mit dem Race Detector aus.

```sh
go test -race -run '^TestFridgeConcurrentAccess$' ./pkg/cocktails/repository/memory
```

Vor der Lösung greifen beide Goroutinen ungeschützt schreibend auf
dieselbe Map zu. Der Race Detector meldet ein Data Race. Go kann
auch mit `concurrent map writes` abbrechen.

Ergänze einen `sync.RWMutex` im Repository. Schütze das Setzen des
Bool-Werts mit `Lock` und `Unlock`. Halte in `Ingredients()` beim
vorgegebenen `maps.Clone` den Leselock mit `RLock` und `RUnlock`.
Führe den Test anschließend erneut mit `-race` aus.

### Aufgabe 8 – HTTP-Server für den SQLite-Katalog

Das Startprogramm unter `cmd/server` ist vorgegeben. Es startet den
Server auf Port 8080 und verdrahtet den SQLite-Katalog mit dem
In-Memory-Kühlschrank. Die SQLite-Zutatenabfragen aus Aufgabe 6
sind im Checkpoint vorhanden.

Schreibe für jede Route eine eigene Handler-Methode und registriere
sie in `RegisterRoutes` beim `http.ServeMux`.

Nutze die vorgegebenen Hilfsfunktionen für Pfadparameter und
HTTP-Statuscodes. Ergänze in `writeJSON` die Ausgabe des
Antwortobjekts als JSON.

- `GET /api/cocktails` liefert die Cocktail-Liste.
- `GET /api/cocktails/{id}` liefert das Rezept zur angegebenen ID.
  Antworte bei einer ungültigen ID mit `400` und bei einer unbekannten
  ID mit `404`.

**Bonus**

- Ergänze `GET /api/ingredients` für die Zutatenliste und
  `GET /api/ingredients/{id}` für den Namen einer Zutat und die
  Cocktails, die sie enthalten.
- Ergänze `GET /api/search?query=Milk` für die Suche. Lies den
  Query-Parameter `query`, rufe `CocktailService.Search` auf und
  liefere die Treffer als JSON. Die Suche berücksichtigt die Namen
  von Cocktails und Zutaten.
- Ergänze die Kühlschrank-API im `FridgeHandler`. Registriere
  `GET /api/user/{user}/fridge` für den Bestand,
  `PATCH /api/user/{user}/fridge/{id}` für Änderungen und
  `GET /api/user/{user}/possible` für mögliche Cocktails.
  Rufe dafür die Methoden des `FridgeService` auf. Der lokale
  Server verwendet für alle Benutzer denselben In-Memory-Kühlschrank.

Teste die Cocktail-Liste mit `httptest.NewTestServer`. Registriere die
Routen und sende über den HTTP-Client einen GET-Request an die vollständige
URL `server.URL + "/api/cocktails"`. Prüfe den Status `200` und dass die
JSON-Antwort genau 69 Cocktails enthält.

```sh
go run ./cmd/initdb
go test ./pkg/cocktails/handler
go run ./cmd/server
```
