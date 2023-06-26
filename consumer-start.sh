#restart mysql
sudo systemctl restart mysql

while IFS='=' read -r key value; do
  if [ "$key" = "RENEWAL_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-renewal@"thread_$i" restart
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"

while IFS='=' read -r key value; do
  if [ "$key" = "RETRY_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-retry@"thread_$i" restart
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"
