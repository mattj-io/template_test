package main

var test_var = `apiVersion: kudo.dev/v1beta1
blah = gcomm://mariadb-0,{{ range $i, $v := untilStep 1 (int .Count) 1 }},mariadb-{{ $i }}{{end}}
`


