package main

import (
	"embed"
)

//go:embed static/*
var staticEmbed embed.FS
