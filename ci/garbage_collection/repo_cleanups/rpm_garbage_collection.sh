#!/bin/bash

set -e

. ../garbage_collection_misc.sh
. ../prune_branches.sh

time_limit=14   ## Days. Set this to give the "deadline". 
                ## All branches older than this a removed.

rpm_base_dir=/var/www/html/openvdc-repos

## Remove all directories whose branch (on git) no longer exists
## or which has not beenm pushed to within $time_limit days.
for directory in $(TIME_LIMIT=${time_limit} dirs_to_prune ${rpm_base_dir}); do
   remove_dir ${rpm_base_dir}/${directory}
done
 
## Now delete "old" (> ${time_limit} days) rpm's from the master directory

here=$PWD
cd ${rpm_base_dir}/master

nrepos=$(ls -1 . | wc -l)
if [[ ${nrepos} -lt 2 ]]; then
   echo "Something is wrong. The master directory contains one or less repos. Quitting."
   exit 1
fi
 
current=$(readlink current)
if [[ -z ${current} ]]; then
   echo "No 'current' symlink in master! "
   exit 1                # There is no "current" symlink. Don't remove anything!
fi
echo "'current' rpm repo is ${current}"

readlink current
current=${current##*\/}

cutoff_date=$(get_cutoff_date)

echo "Checking for stale rpm repos under master..."
for directory in $(ls -d 2*); do
   dr=${directory}
   rpmdate=${dr:0:8}     # yyyymmddgitxxxx is the rpm repo directory format

   if [[ "${dr}" = "${current}" ]]; then
     continue
   fi

   if [[ "${rpmdate}" < "${cutoff_date}" ]]; then  
     full_dir_name=${rpm_base_dir}/master/${dr}
     remove_dir ${full_dir_name}
   fi

done

exit 0   ## Explicit notice: We are done.
