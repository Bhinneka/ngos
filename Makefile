SOURCES := $(shell find . -name '*.go' -type f -not -path './vendor/*'  -not -path '*/mocks/*')

ngos: $(SOURCES)
	go build -o ngos github.com/Bhinneka/ngos/cmd/ngos