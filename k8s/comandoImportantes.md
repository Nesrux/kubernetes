# Comandos impotantes do kind/kubernetes

criar e deletar cluster pelo kind
```
kind create/delete clustes
```
criar cluster com configs especiais
```
kind create cluster --config kind.yaml
```


mostrar cluster disponiveis
```
kind get clusters
```
mostrar os nodes
```
kubectl get nodes
```

se tiver na raiz ou na pasta atual, só precisa do nome do arquivo, a flag
-f é fe file, e o apply aplica as modificações
```
kubectl apply -f pasta/arquivo.extensao
```
Mostra os Pods que estao rodando
```
kubectl get po/pod/pods
```

fazer um bind de portas da minha maquina para o kubernets/kind
```
kubectl port-forward pod/goserver <aporta que eu quero acessar>:<porta exportada>
```