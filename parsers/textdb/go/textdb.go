/* This file has been generated by go-specgen */
package main

type FormatTextDB struct {
	Unknown   uint32
	NumBlobs  uint32
	TextBlobs []TextBlob
}

type TextBlob struct {
	TextID     uint32
	TextOffset uint32 /* Absolute position to a null-terminated string */
}