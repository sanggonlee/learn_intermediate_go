// Deprecated: use github.com/sanggonlee/learn_intermediate_go/01_modules/v2 instead.
module github.com/sanggonlee/learn_intermediate_go/01_modules

go 1.15

require (
	github.com/lib/pq v1.10.2 // indirect
	github.com/sanggonlee/pogo v1.0.0
)

exclude github.com/lib/pq v1.10.2

replace github.com/sanggonlee/pogo => ../../pogo

retract (
	v1.0.0 // Published accidentally.
	v1.0.1 // Contains retractions only.
)
