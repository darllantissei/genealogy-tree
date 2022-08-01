### Gerar mock tests
Utilizar o framework de mock disponível em [gomock](https://github.com/golang/mock) \
Seguir as instruções disponíveis em [gomock-install](https://github.com/golang/mock#go-116) \
Gerar mocks:
```bash
GOPATH=$(go env GOPATH) ~/go/bin/mockgen -source=$(pwd)/<path_interface> -destination=$(pwd)/<path_mocks>
```

### Arquitetura utilizada no projeto - Hexagonal
![Scheme](project_architecture.jpg)