#!/bin/bash

here="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
shard=$1

(cd "$here" && runuser postgres -c 'psql -v ON_ERROR_STOP=1 -A -t -q -d mas' <<EOD

set role mas;
set search_path to ${shard};

select refresh_views();
select refresh_polygons();

EOD
)
