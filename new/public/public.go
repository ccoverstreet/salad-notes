package public

import "embed"

//go:embed all:build/*
var Content embed.FS
