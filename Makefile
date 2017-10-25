# Tcl Module Makefile

files = binary.tcl config.tcl markup.tcl license.tcl

tcltm:
	for f in $(files); do (cat $${f}; echo) >> tcltm.src; done
	sed -e '/#SOURCE#/{r tcltm.src' -e 'N' -e 'G}' tcltm.tcl > tcltm
	chmod +x tcltm
	rm tcltm.src

# Targets
all: tcltm

install:
	cp tcltm /usr/bin/tcltm

clean:
	if [ -f tcltm ]; then rm tcltm; fi