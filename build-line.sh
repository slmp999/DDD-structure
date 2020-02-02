

#!/bin/bash
#IghdgRgijXFNgL9tE7bvcc37WvOUFmZvpHRlPDDGej1
#V7QQFIjZ1VEAKcRcTC8TlfDttv4S85rZUwyhetn0zHH
msg="message=$1:$2"
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Bearer V7QQFIjZ1VEAKcRcTC8TlfDttv4S85rZUwyhetn0zHH" \
    --data "$msg" \
    https://notify-api.line.me/api/notify