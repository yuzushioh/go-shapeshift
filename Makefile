CMDDIR="cmd/shapeshift"

build:
	cd $(CMDDIR) && go build

install:
	cd $(CMDDIR) && go install

clean:
	cd $(CMDDIR) && go clean
