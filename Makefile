SUBDIRS = plugs
EXEC    = example

all: $(SUBDIRS)
	go build -o $(EXEC)

clean: $(SUBDIRS)
	rm -f $(EXEC)

include mk/intermediate.mk
