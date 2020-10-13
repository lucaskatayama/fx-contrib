
SUBDIRS := $(dir $(wildcard */Makefile))

all: $(SUBDIRS)

test: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)


.PHONY: all $(SUBDIRS)
