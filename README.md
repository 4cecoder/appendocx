# appendocx

focuses on appending text to the content of the DOCX file. 


contains only the core functionality for reading a DOCX file, and appending text, and writing it back to a file.


### Build Program
```bash
go build -o docx-append main.go
```

### Cli Usage
```bash
./docx-append -i input.docx -o output.docx -t "Text to append"
```
