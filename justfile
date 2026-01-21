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

# Remove pre-commit
teardown:
 pre-commit uninstall

# Run go server
goserver:
 ./signature-development-utility -port 8001

# Run PHP server
phpserver:
 php -S localhost:8000
