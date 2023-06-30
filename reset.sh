curl -X DELETE --user admin:admin http://ngxmp.com:15675/api/queues/%2f/Q_RENEWAL/contents
echo "\n"
sleep 1

curl -X DELETE --user admin:admin http://35.247.131.49:15675/api/queues/%2f/Q_RETRY/contents
echo "\n"
sleep 1