
.PHONY = install

SRCDIR = pack
DSTFILE = icons.go

ICONS := $(shell find $(SRCDIR) -name '*.svg')

# Generate icons
generate: $(ICONS) #install
	@echo "Generating go-interpretable icons"
	for icon in ${ICONS} ; do \
		NORMALIZED_NAME="$$(basename $$icon .svg)"; \
		CAMEL_NAME="$$(echo $$NORMALIZED_NAME | sed -e 's/\b\(.\)/\u\1/g' -e 's/-//g')"; \
        $$HOME/go/bin/2goarray $$CAMEL_NAME icons < $$icon >> $$NORMALIZED_NAME.go; \
    done

# Get dependencies
install:
	go install github.com/cratonica/2goarray@latest