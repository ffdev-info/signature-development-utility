# CLI helpers

# Help
help:
 @just -l

# Run all pre-commit checks
all-checks:
 pre-commit run --all-files

# Setup pre-commit
setup:
 pre-commit install --install-hooks # uninstall: `pre-commit uninstall`
 # Setup golang checks.
 go install golang.org/x/tools/cmd/godoc@latest
 go install golang.org/x/tools/cmd/goimports@latest
 go install honnef.co/go/tools/cmd/staticcheck@latest

# Remove pre-commit
teardown:
 pre-commit uninstall

# Run go server
goserver:
 ./signature-development-utility -port 8001

# Run go linting
linting:
 - goimports -w .
 - go fmt ./...
 - go vet ./...
 - staticcheck ./...

# Fix imports
fix-imports:
  goimports -w .

# Run tests
test:
  go test ./...

# Run tests without cache
test-nocache:
  go test -count=1 ./...

# Docs
docs:
  godoc -http 0.0.0.0:8000

# Build a release snapshot
build-snapshot:
    goreleaser build --snapshot --single-target --clean -f .goreleaser.yml

# Build a tagged release
build-release:
    goreleaser release --skip=publish --clean -f .goreleaser.yml --skip=sign
