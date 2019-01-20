test: ## テストコードを実行する
	go test -v -cover ./...

bootstrap: ## 外部ツールをインストールする
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		go get $$t ; \
	done
	gometalinter --install --update

.PHONY: help test bootstrap
