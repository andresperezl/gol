test:
	go test -cover -covermode atomic `go list ./... | grep -v lolapi`
gen-spec:
	swagger generate client -f specs/api.yaml -t lolapi -A LeagueOfLegends
