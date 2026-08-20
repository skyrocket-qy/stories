bench:
	cd benchmarks/fx-vs-wire && go test -bench=. -benchmem

bk:
	git add .
	git commit -mupdate
	git push