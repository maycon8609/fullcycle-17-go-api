### Iniciar docker
```bash
  docker-compose up -d
```

### Encerrar docker
```bash
  docker-compose down
```

### Criar tabelas manualmente
- Acessar bash do mysql
  ```bash
    docker-compose exec mysql bash
  ```
  ```bash
    mysql -Uroot -p imersao17
  ```
- sera solicitado password que e `root`.
- copiar script do aquivo db.sql e colar no bash para executar os mesmos.
- conferir se as tabelas foram criadas com sucesso.
  ```bash
    show tables
  ```

### Iniciar aplicação
```bash
  cd cmd/catalog
```
```bash
  go run main.go
```