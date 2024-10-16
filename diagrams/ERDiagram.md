# Mermaid ER Diagram

```mermaid
erDiagram

    x[user]
    b[Budget]
    i[Income]
    t[Transactions]

    x ||--|{ b : Creates
    b ||--|{ t : Tracks
    x ||--|{ i : Makes

```