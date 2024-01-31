# spack

*S*pace *Ch*eck: A tool to check the open/close status of hackspaces.

# Examples

To list the status of all hackspaces:
```sh
$ ./spack
/dev/tal                                          false
/usr/space                                        true
57North Hacklab                                   false
<name>space                                       true
[...]
```

To get a specific hackspace:
```sh
$ ./spack | grep Entropia
Entropia                                          true
```
