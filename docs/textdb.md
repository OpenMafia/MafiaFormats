# TextDB_xx.def
This file is a database of texts used within the game.

Per each text blob, we can use TextOffset to find a null-terminated string within the file.

## Specification

### Spec: FormatTextDB

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | Unknown |  |
| uint32 | NumBlobs |  |
| TextBlob | TextBlobs | N definitions of TextBlob;  |
### Spec: TextBlob

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | TextID |  |
| uint32 | TextOffset | Absolute position to a null-terminated string |
