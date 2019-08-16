# tcltm Makefile
PROJECT			= tcltm
PROJECTDIR		= $(CURDIR)
SCRIPTDIR		= $(CURDIR)/.build
SOURCEDIR		= $(CURDIR)/src
SOURCEFILES		= $(shell find $(SOURCEDIR) -name '*.tcl' -not -name '*test*')
TARGETDIR		= $(CURDIR)/target
VERSION			= $(shell $(SCRIPTDIR)/version)
COMMIT			= $(shell git rev-parse --short HEAD)

# Create Target Directory
$(TARGETDIR):
	@mkdir -p $(TARGETDIR)
	@chmod 777 $(TARGETDIR)

.PHONY: source
source: $(SOURCEFILES) | $(TARGETDIR)
	@sed '/^[[:blank:]]*#/d' $^ > $(TARGETDIR)/$(PROJECT).src
	@sed -i '/^[[:space:]]*$$/d' $(TARGETDIR)/$(PROJECT).src

.PHONY: build
build: source
	@echo "Building tcltm"
	@sed -e '/@SOURCE@/{r $(TARGETDIR)/tcltm.src' -e 'd' -e 'N' -e 'G}' $(SOURCEDIR)/tcltm.tmpl > $(TARGETDIR)/$(PROJECT)
	@sed -i -e '/@USAGE@/{r $(SOURCEDIR)/usage.inc' -e 'd' -e 'N' -e 'G}' $(TARGETDIR)/$(PROJECT)
	@sed -i -e 's/@VERSION@/$(VERSION)/' $(TARGETDIR)/$(PROJECT)
	@sed -i -e 's/@COMMIT@/$(COMMIT)/' $(TARGETDIR)/$(PROJECT)
	@chmod +x $(TARGETDIR)/$(PROJECT)

.PHONY: install
install: build
	@echo "Installing tcltm => ~/.local/bin"
	@cp $(TARGETDIR)/$(PROJECT)  ~/.local/bin

.PHONY: test
test: | $(SOURCEDIR)
	@tclsh $(SOURCEDIR)/test.tcl

.PHONY: clean
clean: | $(PROJECTDIR)
	@rm -rf $(TARGETDIR)
	@rm -f $(PROJECTDIR)/tcltm
