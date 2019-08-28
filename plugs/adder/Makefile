TARGETS = average mod simple

all: $(TARGETS)

.PHONY: $(TARGETS)
$(TARGETS):
	$(call build_go_plugin,$@)

define build_go_plugin
	target=$(1); \
	cd $$target; \
	go build -buildmode=plugin -o $$target.so
endef

.PHONY: clean
clean:
	@find . -type f -iname "*.so" -print -delete

.PHONY: list
list:
	@find . -type f -iname "*.so" -print
