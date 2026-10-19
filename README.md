# gatorcli — learning repo

Simple Go CLI (learning project). Needs **Postgres** + **Go**.

```bash
go install https://github.com/pcybak/gatorcli
```

Set up `config` + DB schema, then:

```bash
go run .        # dev
gatorcli        # built binary
```

Commands: `login`, `register`, `reset`, `users`, `agg`, `addfeed`, `feeds`, `follow`, `following`, `unfollow`, `browse`.
