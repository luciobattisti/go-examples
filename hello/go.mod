module hello

go 1.20

replace examples.com/helper => ./helper

require examples.com/helper v0.0.0-00010101000000-000000000000 // indirect
