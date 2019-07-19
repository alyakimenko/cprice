# cprice

Cryptocurrency systray widget written in Golang.  
![demo](https://raw.githubusercontent.com/alyakimenko/cprice/master/assets/demo.png)

### Sentry support
To connect sentry error tracking service, you should to set the `SENTRY_DSN` env variable.
```bash
export SENTRY_DSN=<your sentry dsn>
```

### Building
```bash
make build
```

Compiled package will be inside the `bin/` folder