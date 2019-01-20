help: ## ドキュメントのヘルプを表示する。
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## テストコードを実行する
	go test -v -cover ./...

bootstrap: ## 外部ツールをインストールする
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		go get $$t ; \
	done
	gometalinter --install --update

.PHONY: help test bootstrap
