package web

import "embed"

//go:embed all:build/*
var Web embed.FS
