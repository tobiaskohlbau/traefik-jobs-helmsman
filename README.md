# Howto

```SHELL
$ make modinfo
$ make
```

Inspect kube-apiserver calls:
```SHELL
$ kubectl --namespace=kube-system logs -f kube-apiserver-kind-control-plane | grep "start/v0.0.0"
```

Inspect ValidatingAdmissionController:
```SHELL
$ kubectl --namespace=logger logs -l app=logger -f
```


Cleanup:

```SHELL
$ make clean
```
