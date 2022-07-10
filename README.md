# golang-clean-architecture

Этот репозиторий, артефакт 3 роликов про Чистую Архитектуру на примере языка Go.

- Часть 4 - https://youtu.be/B_GUqUO42cA (текущий код)
- Часть 3 - https://youtu.be/PqQyCFygiZg
- Часть 2 - https://youtu.be/s_Bou_mChKs
- Часть 1 - https://youtu.be/eVhIlhLl4e4

`internal/adapters/` - external layer, has no relation to business logic.
    - `api/` - contains API handlers and interfaces for books, users and other entities.
    - `db/` - contains DB methods for books, users and other entities.
`internal/domain` - business logic for the application.
    - `book/service.go` - contains implementation of use-cases (business logic), like books get/create/... Realization is in place of use (`adapters/api/books/handlers/interface.go`) 
    - `book/dto.go` - represents structures to create/update books. It is needed for `Service` interface and is linking chain between business logic and controller layers.
    - `book/model.go` - contains book model and business logic (like `Take()` method). 
    - `book/storage.go` - contains `Storage` interface for DB methods required for business logic.

`pkg/client` - contains needed clients (like database).
