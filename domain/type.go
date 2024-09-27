package domain

import "github.com/gorilla/websocket"

type CodeLanguageType string

var (
	Golang     CodeLanguageType = "golang"
	Javascript CodeLanguageType = "javascript"
	Java       CodeLanguageType = "java"
	Python     CodeLanguageType = "python"
	CPP        CodeLanguageType = "c++"
	C          CodeLanguageType = "c"
	Rust       CodeLanguageType = "rust"
	CSharp     CodeLanguageType = "c#"
	TypeScript CodeLanguageType = "typescript"
	PHP        CodeLanguageType = "php"
	Swift      CodeLanguageType = "swift"
	Kotlin     CodeLanguageType = "kotlin"
	Dart       CodeLanguageType = "dart"
	Ruby       CodeLanguageType = "ruby"
	Scala      CodeLanguageType = "scala"
	Racket     CodeLanguageType = "racket"
	Erlang     CodeLanguageType = "erlang"

	WsClients = make(map[uint]*websocket.Conn) // IDE id to conn

	GetDefaultFileName = map[CodeLanguageType]string{
		Golang:     "main.go",
		Javascript: "main.js",
		Java:       "Main.java",
		Python:     "main.py",
		CPP:        "main.cpp",
		C:          "main.c",
		Rust:       "main.rs",
		CSharp:     "main.cs",
		TypeScript: "main.ts",
		PHP:        "main.php",
		Swift:      "main.swift",
		Kotlin:     "main.kt",
		Dart:       "main.dart",
		Ruby:       "main.rb",
		Scala:      "main.scala",
		Racket:     "main.rkt",
		Erlang:     "main.erl",
	}
)
