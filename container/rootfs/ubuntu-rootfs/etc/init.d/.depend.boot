TARGETS = mountkernfs.sh hostname.sh mountdevsubfs.sh procps urandom hwclock.sh mountnfs-bootclean.sh mountnfs.sh bootmisc.sh checkfs.sh checkroot.sh mountall.sh checkroot-bootclean.sh mountall-bootclean.sh
INTERACTIVE = checkfs.sh checkroot.sh
mountdevsubfs.sh: mountkernfs.sh
procps: mountkernfs.sh
urandom: hwclock.sh
hwclock.sh: mountdevsubfs.sh
mountnfs-bootclean.sh: mountnfs.sh
bootmisc.sh: mountnfs-bootclean.sh checkroot-bootclean.sh mountall-bootclean.sh
checkfs.sh: checkroot.sh
checkroot.sh: mountdevsubfs.sh hostname.sh hwclock.sh
mountall.sh: checkfs.sh checkroot-bootclean.sh
checkroot-bootclean.sh: checkroot.sh
mountall-bootclean.sh: mountall.sh
