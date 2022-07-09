# golang-clean-architecture

Этот репозиторий, артефакт 3 роликов про Чистую Архитектуру на примере языка Go.

- Часть 4 - https://youtu.be/B_GUqUO42cA (текущий код)
- Часть 3 - https://youtu.be/PqQyCFygiZg
- Часть 2 - https://youtu.be/s_Bou_mChKs
- Часть 1 - https://youtu.be/eVhIlhLl4e4

`internal/adapters` - external layer, has no relation to business logic.
`internal/domain/book/service.go` - contains `Service` interface which represents use-cases (business logic)

