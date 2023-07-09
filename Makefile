.PHONY: publish-npm publish-pypi publish

publish: publish-npm publish-pypi

publish-npm: build-ts
	cd client-ts && npm publish

publish-pypi: build-py
	cd client-python && poetry publish

.PHONY: test test-ts test-py test-go

test: test-ts test-py test-go

test-ts:
	cd client-ts && npm test

test-py:
	cd client-python && python -m unittest discover

test-go:
	cd go && go test ./...

.PHONY: build build-ts build-py

build: build-ts build-py

build-ts:
	cd client-ts && npm run build

build-py:
	cd client-python && poetry build

docs-dev:
	cd docs.thiggle.com && pnpm dev